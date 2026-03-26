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
	"regexp"
	"strings"

	"github.com/pulumi/registry/tools/resourcedocsgen/pkg/util/language"
)

// stripHTML removes HTML tags from a string, leaving just the text content.
var htmlTagRe = regexp.MustCompile(`<[^>]*>`)

func stripHTML(s string) string {
	// Replace <br> and <br/> with newlines
	s = strings.ReplaceAll(s, "<br>", "\n")
	s = strings.ReplaceAll(s, "<br/>", "\n")
	s = strings.ReplaceAll(s, "<br />", "\n")
	// Strip all remaining HTML tags
	s = htmlTagRe.ReplaceAllString(s, "")
	// Decode common HTML entities
	s = strings.ReplaceAll(s, "&amp;", "&")
	s = strings.ReplaceAll(s, "&lt;", "<")
	s = strings.ReplaceAll(s, "&gt;", ">")
	s = strings.ReplaceAll(s, "&quot;", "\"")
	s = strings.ReplaceAll(s, "&#39;", "'")
	s = strings.ReplaceAll(s, "&nbsp;", " ")
	return s
}

// propertyTypeName returns a display name for a propertyType, preferring DisplayName, then Name.
func propertyTypeName(pt propertyType) string {
	if pt.DisplayName != "" {
		return stripHTML(pt.DisplayName)
	}
	if pt.DescriptionName != "" {
		return stripHTML(pt.DescriptionName)
	}
	return stripHTML(pt.Name)
}

// stripHTMLFromProperties returns a copy of the properties with HTML stripped from text fields.
func stripHTMLFromProperties(props []property) []property {
	out := make([]property, len(props))
	for i, p := range props {
		p.Comment = stripHTML(p.Comment)
		p.DeprecationMessage = stripHTML(p.DeprecationMessage)
		p.DisplayName = stripHTML(p.DisplayName)
		for j, t := range p.Types {
			p.Types[j] = propertyType{
				DisplayName:     stripHTML(t.DisplayName),
				DescriptionName: stripHTML(t.DescriptionName),
				Name:            stripHTML(t.Name),
			}
		}
		out[i] = p
	}
	return out
}

// cliResourceDocArgs is a single-language view of resourceDocArgs for terminal-friendly rendering.
type cliResourceDocArgs struct {
	Title              string
	Comment            string
	DeprecationMessage string
	Language           string // e.g. "typescript", "python", "go"

	Examples []cliExample

	// Constructor
	ConstructorSyntax string // rendered code for the chosen language
	ConstructorParams []formalParam

	InputProperties  []property
	OutputProperties []property

	// State lookup
	StateInputs []property

	// Methods
	Methods []cliMethodDocArgs

	// Nested/supporting types
	NestedTypes    []cliNestedType
	MaxNestedTypes int

	// Import
	ImportDocs string

	PackageDetails packageDetails
}

// cliMethodDocArgs is a single-language view of methodDocArgs.
type cliMethodDocArgs struct {
	Title              string
	Comment            string
	DeprecationMessage string
	Syntax             string // rendered method signature
	InputProperties    []property
	OutputProperties   []property
	Examples           []cliExample
}

// cliFunctionDocArgs is a single-language view of functionDocArgs for terminal-friendly rendering.
type cliFunctionDocArgs struct {
	Title              string
	Comment            string
	DeprecationMessage string
	Language           string

	Examples []cliExample

	FunctionName   string
	FunctionSyntax string // rendered function signature
	FunctionResult string // result type name

	InputProperties  []property
	OutputProperties []property

	NestedTypes []cliNestedType

	PackageDetails packageDetails
}

// cliExample is a single-language example.
type cliExample struct {
	Title   string
	Snippet string
}

// cliNestedType is a single-language nested type.
type cliNestedType struct {
	Name        string
	Description string
	Properties  []property
}

// cliLanguageTag returns the template language tag for a given language.
// NodeJS maps to "typescript", others use their string representation.
func cliLanguageTag(lang language.Language) string {
	if lang == language.NodeJS {
		return "typescript"
	}
	return lang.String()
}

// toCLIResourceArgs extracts a single-language view from resourceDocArgs.
func toCLIResourceArgs(args resourceDocArgs, lang language.Language) cliResourceDocArgs {
	langTag := cliLanguageTag(lang)

	cli := cliResourceDocArgs{
		Title:              args.Header.Title,
		Comment:            stripHTML(args.Comment),
		DeprecationMessage: stripHTML(args.DeprecationMessage),
		Language:           langTag,
		ImportDocs:         stripHTML(args.ImportDocs),
		MaxNestedTypes:     args.MaxNestedTypes,
		PackageDetails:     args.PackageDetails,
	}

	// Constructor syntax
	if cs, ok := args.CreationExampleSyntax[lang]; ok {
		cli.ConstructorSyntax = stripHTML(cs)
	}

	// Constructor params
	if fp, ok := args.ConstructorParamsTyped[lang]; ok {
		cli.ConstructorParams = fp.Params
	}

	// Properties
	if props, ok := args.InputProperties[lang]; ok {
		cli.InputProperties = stripHTMLFromProperties(props)
	}
	if props, ok := args.OutputProperties[lang]; ok {
		cli.OutputProperties = stripHTMLFromProperties(props)
	}

	// State inputs
	if props, ok := args.StateInputs[lang]; ok {
		cli.StateInputs = stripHTMLFromProperties(props)
	}

	// Examples
	cli.Examples = extractCLIExamples(args.ExamplesSection, lang)

	// Methods
	for _, m := range args.Methods {
		cm := cliMethodDocArgs{
			Title:              m.Title,
			Comment:            stripHTML(m.Comment),
			DeprecationMessage: stripHTML(m.DeprecationMessage),
			Examples:           extractCLIExamples(m.ExamplesSection, lang),
		}
		if name, ok := m.MethodName[lang]; ok {
			if mArgs, ok := m.MethodArgs[lang]; ok {
				cleanArgs := stripHTML(mArgs)
				if result, ok := m.MethodResult[lang]; ok && result.Name != "" {
					cm.Syntax = name + "(" + cleanArgs + ") -> " + propertyTypeName(result)
				} else {
					cm.Syntax = name + "(" + cleanArgs + ")"
				}
			}
		}
		if props, ok := m.InputProperties[lang]; ok {
			cm.InputProperties = stripHTMLFromProperties(props)
		}
		if props, ok := m.OutputProperties[lang]; ok {
			cm.OutputProperties = stripHTMLFromProperties(props)
		}
		cli.Methods = append(cli.Methods, cm)
	}

	// Nested types
	for i, nt := range args.NestedTypes {
		if i >= args.MaxNestedTypes {
			break
		}
		cnt := cliNestedType{
			Name:        stripHTML(nt.Name),
			Description: stripHTML(nt.Description),
		}
		if props, ok := nt.Properties[lang]; ok {
			cnt.Properties = stripHTMLFromProperties(props)
		}
		cli.NestedTypes = append(cli.NestedTypes, cnt)
	}

	return cli
}

// toCLIFunctionArgs extracts a single-language view from functionDocArgs.
func toCLIFunctionArgs(args functionDocArgs, lang language.Language) cliFunctionDocArgs {
	langTag := cliLanguageTag(lang)

	cli := cliFunctionDocArgs{
		Title:              args.Header.Title,
		Comment:            stripHTML(args.Comment),
		DeprecationMessage: stripHTML(args.DeprecationMessage),
		Language:           langTag,
		PackageDetails:     args.PackageDetails,
	}

	if name, ok := args.FunctionName[lang]; ok {
		cli.FunctionName = name
	}
	if fArgs, ok := args.FunctionArgs[lang]; ok {
		cleanArgs := stripHTML(fArgs)
		if result, ok := args.FunctionResult[lang]; ok && result.Name != "" {
			cli.FunctionSyntax = cli.FunctionName + "(" + cleanArgs + ") -> " + propertyTypeName(result)
		} else {
			cli.FunctionSyntax = cli.FunctionName + "(" + cleanArgs + ")"
		}
		if result, ok := args.FunctionResult[lang]; ok {
			cli.FunctionResult = propertyTypeName(result)
		}
	}

	if props, ok := args.InputProperties[lang]; ok {
		cli.InputProperties = stripHTMLFromProperties(props)
	}
	if props, ok := args.OutputProperties[lang]; ok {
		cli.OutputProperties = stripHTMLFromProperties(props)
	}

	cli.Examples = extractCLIExamples(args.ExamplesSection, lang)

	for _, nt := range args.NestedTypes {
		cnt := cliNestedType{
			Name:        stripHTML(nt.Name),
			Description: stripHTML(nt.Description),
		}
		if props, ok := nt.Properties[lang]; ok {
			cnt.Properties = stripHTMLFromProperties(props)
		}
		cli.NestedTypes = append(cli.NestedTypes, cnt)
	}

	return cli
}

// extractCLIExamples pulls single-language examples from an examplesSection.
func extractCLIExamples(es examplesSection, lang language.Language) []cliExample {
	var examples []cliExample
	for _, ex := range es.Examples {
		snippet, ok := ex.Snippets[lang]
		if !ok || snippet == "" || snippet == defaultMissingExampleSnippetPlaceholder {
			continue
		}
		examples = append(examples, cliExample{
			Title:   ex.Title,
			Snippet: snippet,
		})
	}
	return examples
}
