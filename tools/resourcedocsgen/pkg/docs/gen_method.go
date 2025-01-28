// Copyright 2024, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package docs

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/golang/glog"
	"github.com/pulumi/pulumi/pkg/v3/codegen/python"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/slice"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"github.com/pulumi/registry/tools/resourcedocsgen/pkg/docs/templates"
	"github.com/pulumi/registry/tools/resourcedocsgen/pkg/util/language"
)

type methodDocArgs struct {
	Title string

	ResourceName string

	DeprecationMessage string
	Comment            string
	ExamplesSection    examplesSection

	// MethodName is a map of the language and the method name in that language.
	MethodName map[language.Language]string
	// MethodArgs is map per language view of the parameters in the method.
	MethodArgs map[language.Language]string
	// MethodResult is a map per language property types that is returned as a result of calling a method.
	MethodResult map[language.Language]propertyType

	// InputProperties is a map per language and the corresponding slice of input properties accepted by the method.
	InputProperties map[language.Language][]property
	// OutputProperties is a map per language and the corresponding slice of output properties, which are properties of
	// the MethodResult type.
	OutputProperties map[language.Language][]property
}

func (mod *modContext) genMethods(r *schema.Resource) []methodDocArgs {
	methods := slice.Prealloc[methodDocArgs](len(r.Methods))
	for _, m := range r.Methods {
		methods = append(methods, mod.genMethod(r, m))
	}
	return methods
}

func (mod *modContext) genMethod(r *schema.Resource, m *schema.Method) methodDocArgs {
	dctx := mod.context
	f := m.Function
	inputProps, outputProps := make(map[language.Language][]property), make(map[language.Language][]property)
	for lang := range language.All() {
		if f.Inputs != nil {
			exclude := func(name string) bool {
				return name == "__self__"
			}
			props := mod.getPropertiesWithIDPrefixAndExclude(f.Inputs.Properties, lang, true, false, false,
				m.Name+"_arg_", exclude)
			if len(props) > 0 {
				inputProps[lang] = props
			}
		}
		if f.ReturnType != nil {
			if objectType, ok := f.ReturnType.(*schema.ObjectType); ok && objectType != nil {
				outputProps[lang] = mod.getPropertiesWithIDPrefixAndExclude(objectType.Properties, lang, false, false, false,
					m.Name+"_result_", nil)
			}
		}
	}

	// Generate the per-language map for the method name.
	methodNameMap := map[language.Language]string{}
	for lang := range language.All() {
		docHelper := dctx.getLanguageDocHelper(lang)
		methodNameMap[lang] = docHelper.GetMethodName(m)
	}

	defer glog.Flush()
	resourceSnippetLanguages := mod.context.getSupportedSnippetLanguages(r.IsOverlay, r.OverlaySupportedLanguages)
	methodSnippetLanguages := mod.context.getSupportedSnippetLanguages(f.IsOverlay, f.OverlaySupportedLanguages)
	if resourceSnippetLanguages != methodSnippetLanguages {
		// All code choosers are tied together. If the method doesn't support the same languages as the resource, we default
		// to the resource's supported languages for consistency.
		glog.Warningf(
			"Resource %s (%s) and method %s (%s) have different supported snippet languages. "+
				"Defaulting to the resources supported snippet languages.",
			r.Token, resourceSnippetLanguages, m.Name, methodSnippetLanguages,
		)
	}

	docInfo := dctx.decomposeDocstring(f.Comment, resourceSnippetLanguages)
	args := methodDocArgs{
		Title: strings.Title(m.Name), //nolint:staticcheck // Maintain usage of strings.Title for backwards compatibility.

		ResourceName: resourceName(r),

		MethodName:   methodNameMap,
		MethodArgs:   mod.genMethodArgs(r, m, methodNameMap),
		MethodResult: mod.getMethodResult(r, m),

		Comment:            docInfo.description,
		DeprecationMessage: f.DeprecationMessage,
		ExamplesSection: examplesSection{
			Examples:             docInfo.examples,
			LangChooserLanguages: resourceSnippetLanguages,
		},

		InputProperties:  inputProps,
		OutputProperties: outputProps,
	}

	return args
}

func (mod *modContext) genMethodTS(f *schema.Function, resourceName, methodName string,
	optionalArgs bool,
) []formalParam {
	argsType := fmt.Sprintf("%s.%sArgs", resourceName, title(methodName, language.NodeJS))

	var optionalFlag string
	if optionalArgs {
		optionalFlag = "?"
	}

	var params []formalParam
	if f.Inputs != nil {
		params = append(params, formalParam{
			Name:         "args",
			OptionalFlag: optionalFlag,
			Type: propertyType{
				Name: argsType,
			},
		})
	}
	return params
}

func (mod *modContext) genMethodGo(f *schema.Function, resourceName, methodName string,
	optionalArgs bool,
) []formalParam {
	argsType := fmt.Sprintf("%s%sArgs", resourceName, title(methodName, language.Go))

	params := []formalParam{
		{
			Name:         "ctx",
			OptionalFlag: "*",
			Type: propertyType{
				Name: "Context",
				Link: "https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context",
			},
		},
	}

	var optionalFlag string
	if optionalArgs {
		optionalFlag = "*"
	}

	if f.Inputs != nil {
		params = append(params, formalParam{
			Name:         "args",
			OptionalFlag: optionalFlag,
			Type: propertyType{
				Name: argsType,
			},
		})
	}
	return params
}

func (mod *modContext) genMethodCS(
	f *schema.Function,
	resourceName string,
	methodName string,
	optionalArgs bool,
) []formalParam {
	argsType := fmt.Sprintf("%s.%sArgs", resourceName, title(methodName, language.CSharp))
	var params []formalParam
	if f.Inputs != nil {
		var optionalFlag string
		if optionalArgs {
			optionalFlag = "?"
		}
		params = append(params, formalParam{
			Name:         "args",
			OptionalFlag: optionalFlag,
			DefaultValue: "",
			Type: propertyType{
				Name: argsType,
			},
		})
	}
	return params
}

func (mod *modContext) genMethodPython(f *schema.Function) []formalParam {
	dctx := mod.context
	docLanguageHelper := dctx.getLanguageDocHelper(language.Python)
	var params []formalParam

	params = append(params, formalParam{
		Name: "self",
	})

	if f.Inputs != nil {
		// Filter out the __self__ argument from the inputs.
		args := slice.Prealloc[*schema.Property](len(f.Inputs.InputShape.Properties) - 1)
		for _, arg := range f.Inputs.InputShape.Properties {
			if arg.Name == "__self__" {
				continue
			}
			args = append(args, arg)
		}
		// Sort required args first.
		sort.Slice(args, func(i, j int) bool {
			pi, pj := args[i], args[j]
			switch {
			case pi.IsRequired() != pj.IsRequired():
				return pi.IsRequired() && !pj.IsRequired()
			default:
				return pi.Name < pj.Name
			}
		})
		for _, arg := range args {
			def, err := mod.pkg.Definition()
			contract.AssertNoErrorf(err, "failed to get definition for package %q", mod.pkg.Name())
			typ := docLanguageHelper.GetLanguageTypeString(def, mod.mod, arg.Type, true /*input*/)
			var defaultValue string
			if !arg.IsRequired() {
				defaultValue = " = None"
			}
			params = append(params, formalParam{
				Name:         python.PyName(arg.Name),
				DefaultValue: defaultValue,
				Type: propertyType{
					Name: typ,
				},
			})
		}
	}
	return params
}

// genMethodArgs generates the arguments string for a given method that can be rendered directly into a template. An
// empty string indicates no args.
func (mod *modContext) genMethodArgs(r *schema.Resource, m *schema.Method,
	methodNameMap map[language.Language]string,
) map[language.Language]string {
	dctx := mod.context
	f := m.Function

	functionParams := make(map[language.Language]string)
	for lang := range language.All() {
		var (
			paramTemplate templates.Template
			params        []formalParam
		)
		b := &bytes.Buffer{}

		paramSeparatorTemplate := templates.ParamSeparator
		ps := paramSeparator{}

		var hasArgs bool
		optionalArgs := true
		if f.Inputs != nil {
			for _, arg := range f.Inputs.InputShape.Properties {
				if arg.Name == "__self__" {
					continue
				}
				hasArgs = true
				if arg.IsRequired() {
					optionalArgs = false
				}
			}
		}

		if !hasArgs {
			functionParams[lang] = ""
			continue
		}

		switch lang {
		case language.NodeJS:
			params = mod.genMethodTS(f, resourceName(r), methodNameMap[language.NodeJS], optionalArgs)
			paramTemplate = templates.TsFormalParam
		case language.Go:
			params = mod.genMethodGo(f, resourceName(r), methodNameMap[language.Go], optionalArgs)
			paramTemplate = templates.GoFormalParam
		case language.CSharp:
			params = mod.genMethodCS(f, resourceName(r), methodNameMap[language.CSharp], optionalArgs)
			paramTemplate = templates.CSharpFormalParam
		case language.Python:
			params = mod.genMethodPython(f)
			paramTemplate = templates.PyFormalParam
			paramSeparatorTemplate = templates.PyParamSeparator

			docHelper := dctx.getLanguageDocHelper(lang)
			methodName := docHelper.GetMethodName(m)
			ps = paramSeparator{Indent: strings.Repeat(" ", len("def (")+len(methodName))}
		}

		n := len(params)
		if n == 0 {
			functionParams[lang] = ""
			continue
		}

		for i, p := range params {
			if err := paramTemplate(b, p); err != nil {
				panic(err)
			}
			if i != n-1 {
				if err := paramSeparatorTemplate(b, ps); err != nil {
					panic(err)
				}
			}
		}
		functionParams[lang] = b.String()
	}
	return functionParams
}

// getMethodResult returns a map of per-language information about the method result. An empty propertyType.Name
// indicates no result.
func (mod *modContext) getMethodResult(r *schema.Resource, m *schema.Method) map[language.Language]propertyType {
	resourceMap := make(map[language.Language]propertyType)

	dctx := mod.context

	var resultTypeName string
	for lang := range language.All() {
		if m.Function.ReturnType != nil {
			def, err := mod.pkg.Definition()
			contract.AssertNoErrorf(err, "failed to get definition for package %q", mod.pkg.Name())
			resultTypeName = dctx.getLanguageDocHelper(lang).
				GetMethodResultName(def, mod.mod, r, m)
		}
		resourceMap[lang] = propertyType{
			Name: resultTypeName,
			//nolint:staticcheck // Maintain usage of strings.Title for backwards compatibility.
			Link: fmt.Sprintf("#method_%s_result", strings.Title(m.Name)),
		}
	}

	return resourceMap
}
