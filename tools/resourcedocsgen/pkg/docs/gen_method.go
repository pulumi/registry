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
			typ := docLanguageHelper.GetTypeName(mod.pkg, arg.Type, true /*input*/, mod.mod)
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

	pyDocHelper := dctx.getLanguageDocHelper(language.Python)
	pyIdent := strings.Repeat(" ", len("def (")+len(pyDocHelper.GetMethodName(m)))

	optionalArgs := true
	hasArgs := false
	if f.Inputs != nil {
		for _, arg := range f.Inputs.InputShape.Properties {
			if arg.Name == "__self__" {
				continue
			}
			hasArgs = true
			if arg.IsRequired() {
				optionalArgs = false
				break
			}
		}
	}

	functionParams := make(map[language.Language]string)
	for lang := range language.All() {
		if !hasArgs {
			functionParams[lang] = ""
			continue
		}
		params := mod.genMethodParams(lang, f, resourceName(r), methodNameMap[lang], optionalArgs)
		functionParams[lang] = renderParams(lang, params, pyIdent)
	}
	return functionParams
}

func (mod *modContext) genMethodParams(
	lang language.Language,
	f *schema.Function,
	resourceName string,
	methodName string,
	optionalArgs bool,
) []formalParam {
	switch lang {
	case language.NodeJS:
		return mod.genMethodTS(f, resourceName, methodName, optionalArgs)
	case language.Go:
		return mod.genMethodGo(f, resourceName, methodName, optionalArgs)
	case language.CSharp:
		return mod.genMethodCS(f, resourceName, methodName, optionalArgs)
	case language.Python:
		return mod.genMethodPython(f)
	case language.Java, language.YAML:
		// Java and YAML don't have method documentation.
		return nil
	default:
		glog.Fatalf("Unknown language %#v", lang)
		return nil
	}
}

func renderParams(lang language.Language, params []formalParam, pyIndent string) string {
	var param templates.Template
	var ps paramSeparator
	separator := templates.ParamSeparator

	switch lang {
	case language.CSharp:
		param = templates.CSharpFormalParam
	case language.Go:
		param = templates.GoFormalParam
	case language.Java:
		param = templates.JavaFormalParam
	case language.NodeJS:
		param = templates.TsFormalParam
	case language.Python:
		param = templates.PyFormalParam
		separator = templates.PyParamSeparator
		ps = paramSeparator{Indent: pyIndent}
	case language.YAML:
		return "" // YAML doesn't use rendered parameters
	default:
		glog.Fatalf("Unknown language %#v", lang)
	}

	b := &bytes.Buffer{}
	n := len(params)
	for i, p := range params {
		if err := param(b, p); err != nil {
			panic(err)
		}
		if i != n-1 {
			if err := separator(b, ps); err != nil {
				panic(err)
			}
		}
	}

	return b.String()
}

// getMethodResult returns a map of per-language information about the method result. An empty propertyType.Name
// indicates no result.
func (mod *modContext) getMethodResult(r *schema.Resource, m *schema.Method) map[language.Language]propertyType {
	resourceMap := make(map[language.Language]propertyType)

	dctx := mod.context

	var resultTypeName string
	for lang := range language.All() {
		if m.Function.ReturnType != nil {
			resultTypeName = dctx.getLanguageDocHelper(lang).
				GetMethodResultName(mod.pkg, mod.mod, r, m)
		}
		resourceMap[lang] = propertyType{
			Name: resultTypeName,
			//nolint:staticcheck // Maintain usage of strings.Title for backwards compatibility.
			Link: fmt.Sprintf("#method_%s_result", strings.Title(m.Name)),
		}
	}

	return resourceMap
}
