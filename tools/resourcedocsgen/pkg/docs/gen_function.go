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
	"fmt"
	"strings"

	"github.com/golang/glog"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	go_gen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
	"github.com/pulumi/pulumi/pkg/v3/codegen/python"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/slice"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"github.com/pulumi/registry/tools/resourcedocsgen/pkg/util/language"
)

// functionDocArgs represents the args that a Function doc template needs.
type functionDocArgs struct {
	Header header

	Tool string
	// LangChooserLanguages is a comma-separated list of languages to pass to the language chooser shortcode. Use this to
	// customize the languages shown for a function. By default, the language chooser will show all languages supported by
	// Pulumi.
	//
	// Supported values are "typescript", "python", "go", "csharp", "java", "yaml"
	LangChooserLanguages string

	DeprecationMessage string
	Comment            string
	ExamplesSection    examplesSection

	// FunctionName is a map of the language and the function name in that language.
	FunctionName map[language.Language]string
	// FunctionArgs is map per language view of the parameters in the Function.
	FunctionArgs map[language.Language]string
	// FunctionResult is a map per language property types that is returned as a result of calling a Function.
	FunctionResult map[language.Language]propertyType

	// InputProperties is a map per language and the corresponding slice of input properties accepted by the Function.
	InputProperties map[language.Language][]property
	// InputProperties is a map per language and the corresponding slice of output properties, which are properties of the
	// FunctionResult type.
	OutputProperties map[language.Language][]property

	// NestedTypes is a slice of the nested types used in the input and output properties.
	NestedTypes []docNestedType

	PackageDetails packageDetails

	// Check if the function supports an `Output` version that is automatically lifted to accept `Input` values and return
	// an `Output` (per language).
	HasOutputVersion map[language.Language]bool

	// True if any of the entries in `HasOutputVersion` are true.
	AnyLanguageHasOutputVersion bool

	// Same as FunctionArgs, but specific to the Output version of the function.
	FunctionArgsOutputVersion map[language.Language]string

	// Same as FunctionResult, but specific to the Output version of the function. In languages like Go, `Output<Result>`
	// gets a dedicated nominal type to emulate generics, which will be passed in here.
	FunctionResultOutputVersion map[language.Language]propertyType
}

// getFunctionResourceInfo returns a map of per-language information about the resource being looked-up using a static
// "getter" function.
func (mod *modContext) getFunctionResourceInfo(
	f *schema.Function, outputVersion bool,
) map[language.Language]propertyType {
	dctx := mod.context
	resourceMap := make(map[language.Language]propertyType)

	var resultTypeName string
	for lang := range language.All() {
		docLangHelper := dctx.getLanguageDocHelper(lang)
		switch lang {
		case language.NodeJS:
			resultTypeName = docLangHelper.GetResourceFunctionResultName(mod.mod, f)
		case language.Go:
			resultTypeName = docLangHelper.GetResourceFunctionResultName(mod.mod, f)
			if outputVersion {
				resultTypeName = resultTypeName + "Output"
			}
		case language.CSharp:
			namespace := title(mod.pkg.Name(), language.CSharp)
			if ns, ok := dctx.csharpPkgInfo.Namespaces[mod.pkg.Name()]; ok {
				namespace = ns
			}
			resultTypeName = docLangHelper.GetResourceFunctionResultName(mod.mod, f)
			if mod.mod == "" {
				resultTypeName = fmt.Sprintf("Pulumi.%s.%s", namespace, resultTypeName)
			} else {
				resultTypeName = fmt.Sprintf("Pulumi.%s.%s.%s", namespace, title(mod.mod, language.CSharp), resultTypeName)
			}

		case language.Python:
			resultTypeName = docLangHelper.GetResourceFunctionResultName(mod.mod, f)
		case language.Java:
			resultTypeName = docLangHelper.GetResourceFunctionResultName(mod.mod, f)
		case language.YAML:
			resultTypeName = docLangHelper.GetResourceFunctionResultName(mod.mod, f)
		default:
			panic(fmt.Errorf("cannot generate function resource info for unhandled language %#v", lang))
		}

		parts := strings.Split(resultTypeName, ".")
		displayName := parts[len(parts)-1]
		resourceMap[lang] = propertyType{
			Name:        resultTypeName,
			DisplayName: displayName,
			Link:        "#result",
		}
	}

	return resourceMap
}

func (mod *modContext) genFunctionParamsTS(f *schema.Function, funcName string, outputVersion bool) []formalParam {
	dctx := mod.context

	argsTypeSuffix := "Args"
	if outputVersion {
		argsTypeSuffix = "OutputArgs"
	}

	argsType := title(fmt.Sprintf("%s%s", funcName, argsTypeSuffix), language.NodeJS)

	docLangHelper := dctx.getLanguageDocHelper(language.NodeJS)
	var params []formalParam
	if f.Inputs != nil {
		params = append(params, formalParam{
			Name:         "args",
			OptionalFlag: "",
			Type: propertyType{
				Name: argsType,
			},
		})
	}
	def, err := mod.pkg.Definition()
	contract.AssertNoErrorf(err, "failed to get definition for package %q", mod.pkg.Name())
	params = append(params, formalParam{
		Name:         "opts",
		OptionalFlag: "?",
		Type: propertyType{
			Name: "InvokeOptions",
			Link: docLangHelper.GetDocLinkForPulumiType(def, "InvokeOptions"),
		},
	})

	return params
}

func (mod *modContext) genFunctionParamsGo(f *schema.Function, funcName string, outputVersion bool) []formalParam {
	argsTypeSuffix := "Args"
	if outputVersion {
		argsTypeSuffix = "OutputArgs"
	}

	argsType := fmt.Sprintf("%s%s", funcName, argsTypeSuffix)

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

	if f.Inputs != nil {
		params = append(params, formalParam{
			Name:         "args",
			OptionalFlag: "*",
			Type: propertyType{
				Name: argsType,
			},
		})
	}

	params = append(params, formalParam{
		Name:         "opts",
		OptionalFlag: "...",
		Type: propertyType{
			Name: "InvokeOption",
			Link: "https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption",
		},
	})
	return params
}

func (mod *modContext) genFunctionParamsCS(f *schema.Function, funcName string, outputVersion bool) []formalParam {
	dctx := mod.context

	argsTypeSuffix := "Args"
	if outputVersion {
		argsTypeSuffix = "InvokeArgs"
	}

	argsType := funcName + argsTypeSuffix
	docLangHelper := dctx.getLanguageDocHelper(language.CSharp)
	var params []formalParam
	if f.Inputs != nil {
		params = append(params, formalParam{
			Name:         "args",
			OptionalFlag: "",
			DefaultValue: "",
			Type: propertyType{
				Name: argsType,
			},
		})
	}

	def, err := mod.pkg.Definition()
	contract.AssertNoErrorf(err, "failed to get definition for package %q", mod.pkg.Name())
	params = append(params, formalParam{
		Name:         "opts",
		OptionalFlag: "?",
		DefaultValue: " = null",
		Type: propertyType{
			Name: "InvokeOptions",
			Link: docLangHelper.GetDocLinkForPulumiType(def, "Pulumi.InvokeOptions"),
		},
	})
	return params
}

func (mod *modContext) genFunctionParamsJava(f *schema.Function, funcName string) []formalParam {
	dctx := mod.context

	// java uses the same args type for both the regular and output versions of the function
	argsTypeSuffix := "Args"

	argsType := title(funcName+argsTypeSuffix, language.Java)
	docLangHelper := dctx.getLanguageDocHelper(language.Java)
	var params []formalParam
	if f.Inputs != nil {
		params = append(params, formalParam{
			Name:         "args",
			OptionalFlag: "",
			DefaultValue: "",
			Type: propertyType{
				Name: argsType,
			},
		})
	}
	def, err := mod.pkg.Definition()
	contract.AssertNoErrorf(err, "failed to get definition for package %q", mod.pkg.Name())

	params = append(params, formalParam{
		Name:         "options",
		OptionalFlag: "@Nullable",
		Type: propertyType{
			Name: "InvokeOptions",
			Link: docLangHelper.GetDocLinkForPulumiType(def, "InvokeOptions"),
		},
	})
	return params
}

func (mod *modContext) genFunctionParamsPython(f *schema.Function, _ string, outputVersion bool) []formalParam {
	dctx := mod.context
	docLanguageHelper := dctx.getLanguageDocHelper(language.Python)
	var params []formalParam

	// Some functions don't have any inputs other than the InvokeOptions. For example, the `get_billing_service_account`
	// function.
	if f.Inputs != nil {
		inputs := f.Inputs
		if outputVersion {
			inputs = inputs.InputShape
		}

		params = slice.Prealloc[formalParam](len(inputs.Properties))
		for _, prop := range inputs.Properties {
			var schemaType schema.Type
			if outputVersion {
				schemaType = codegen.OptionalType(prop)
			} else {
				schemaType = codegen.PlainType(codegen.OptionalType(prop))
			}

			typ := docLanguageHelper.GetTypeName(mod.pkg,
				schemaType, true /*input*/, mod.mod)
			params = append(params, formalParam{
				Name:         python.PyName(prop.Name),
				DefaultValue: " = None",
				Type: propertyType{
					Name: typ,
				},
			})
		}
	} else {
		params = slice.Prealloc[formalParam](1)
	}

	params = append(params, formalParam{
		Name:         "opts",
		DefaultValue: " = None",
		Type: propertyType{
			Name: "Optional[InvokeOptions]",
			Link: "/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions",
		},
	})

	return params
}

// genFunctionArgs generates the arguments string for a given Function that can be rendered directly into a template.
func (mod *modContext) genFunctionArgs(
	f *schema.Function,
	funcNameMap map[language.Language]string,
	outputVersion bool,
) map[language.Language]string {
	pyDocHelper := mod.context.getLanguageDocHelper(language.Python)
	pyIdent := strings.Repeat(" ", len("def (")+len(pyDocHelper.GetFunctionName(f)))

	functionParams := make(map[language.Language]string)
	for lang := range language.All() {
		params := mod.genFunctionParams(lang, f, funcNameMap, outputVersion)
		if len(params) == 0 {
			continue
		}
		functionParams[lang] = renderParams(lang, params, pyIdent)
	}
	return functionParams
}

func (mod *modContext) genFunctionParams(
	lang language.Language, f *schema.Function,
	funcNameMap map[language.Language]string, outputVersion bool,
) []formalParam {
	funcName := funcNameMap[lang]
	switch lang {
	case language.NodeJS:
		return mod.genFunctionParamsTS(f, funcName, outputVersion)
	case language.Go:
		return mod.genFunctionParamsGo(f, funcName, outputVersion)
	case language.CSharp:
		return mod.genFunctionParamsCS(f, funcName, outputVersion)
	case language.Java:
		return mod.genFunctionParamsJava(f, funcName)
	case language.Python:
		return mod.genFunctionParamsPython(f, funcName, outputVersion)
	case language.YAML: // Left blank
		return nil
	default:
		glog.Fatalf("Unknown %#v", lang)
		return nil
	}
}

func (mod *modContext) genFunctionHeader(f *schema.Function) header {
	funcName := tokenToName(f.Token)
	var baseDescription string
	var titleTag string
	if mod.mod == "" {
		baseDescription = fmt.Sprintf("Documentation for the %s.%s function "+
			"with examples, input properties, output properties, "+
			"and supporting types.", mod.pkg.Name(), funcName)
		titleTag = fmt.Sprintf("%s.%s", mod.pkg.Name(), funcName)
	} else {
		baseDescription = fmt.Sprintf("Documentation for the %s.%s.%s function "+
			"with examples, input properties, output properties, "+
			"and supporting types.", mod.pkg.Name(), mod.mod, funcName)
		titleTag = fmt.Sprintf("%s.%s.%s", mod.pkg.Name(), mod.mod, funcName)
	}

	return header{
		Title:    funcName,
		TitleTag: titleTag,
		MetaDesc: baseDescription,
	}
}

func (mod *modContext) genFunctionOutputVersionMap(f *schema.Function) map[language.Language]bool {
	result := map[language.Language]bool{}
	for lang := range language.All() {
		hasOutputVersion := f.NeedsOutputVersion()
		switch lang {
		case language.Go:
			hasOutputVersion = go_gen.NeedsGoOutputVersion(f)
		case language.YAML:
			hasOutputVersion = false
		}
		result[lang] = hasOutputVersion
	}
	return result
}

// genFunction is the main entrypoint for generating docs for a Function. Returns args type that can be used to execute
// the `function.tmpl` doc template.
func (mod *modContext) genFunction(f *schema.Function) functionDocArgs {
	dctx := mod.context
	inputProps := make(map[language.Language][]property)
	outputProps := make(map[language.Language][]property)
	for lang := range language.All() {
		if f.Inputs != nil {
			inputProps[lang] = mod.getProperties(f.Inputs.Properties, lang, true, false, false)
		}
		if f.ReturnType != nil {
			if objectObject, ok := f.ReturnType.(*schema.ObjectType); ok {
				outputProps[lang] = mod.getProperties(objectObject.Properties,
					lang, false, false, false)
			}
		}
	}

	nestedTypes := mod.genNestedTypes(f, false /*resourceType*/, false /*isProvider*/)

	// Generate the per-language map for the function name.
	funcNameMap := map[language.Language]string{}
	for lang := range language.All() {
		docHelper := dctx.getLanguageDocHelper(lang)
		funcNameMap[lang] = docHelper.GetFunctionName(f)
	}

	def, err := mod.pkg.Definition()
	contract.AssertNoErrorf(err, "failed to get definition for package %q", mod.pkg.Name())

	packageDetails := packageDetails{
		DisplayName:    getPackageDisplayName(def.Name),
		Repository:     def.Repository,
		RepositoryName: getRepositoryName(def.Repository),
		License:        def.License,
		Notes:          def.Attribution,
	}

	supportedSnippetLanguages := mod.context.getSupportedSnippetLanguages(f.IsOverlay, f.OverlaySupportedLanguages)
	docInfo := dctx.decomposeDocstring(f.Comment, supportedSnippetLanguages)
	args := functionDocArgs{
		Header: mod.genFunctionHeader(f),

		Tool:                 mod.tool,
		LangChooserLanguages: supportedSnippetLanguages,

		FunctionName:   funcNameMap,
		FunctionArgs:   mod.genFunctionArgs(f, funcNameMap, false /*outputVersion*/),
		FunctionResult: mod.getFunctionResourceInfo(f, false /*outputVersion*/),

		Comment:            docInfo.description,
		DeprecationMessage: f.DeprecationMessage,
		ExamplesSection: examplesSection{
			Examples:             docInfo.examples,
			LangChooserLanguages: supportedSnippetLanguages,
		},

		InputProperties:  inputProps,
		OutputProperties: outputProps,

		NestedTypes: nestedTypes,

		PackageDetails: packageDetails,
	}

	args.HasOutputVersion = mod.genFunctionOutputVersionMap(f)

	for _, hasOutputVersion := range args.HasOutputVersion {
		if hasOutputVersion {
			args.AnyLanguageHasOutputVersion = true
			continue
		}
	}

	if f.NeedsOutputVersion() {
		args.FunctionArgsOutputVersion = mod.genFunctionArgs(f, funcNameMap, true /*outputVersion*/)
		args.FunctionResultOutputVersion = mod.getFunctionResourceInfo(f, true /*outputVersion*/)
	}

	return args
}
