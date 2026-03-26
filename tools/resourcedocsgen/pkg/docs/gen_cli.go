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

// cliLangEntry holds per-language data for a single language within a CLI doc.
type cliLangEntry struct {
	Tag        string // "typescript", "python", "go", "csharp", "java", "yaml"
	Properties []property
}

// cliResourceDocArgs holds all-language data for terminal-friendly resource rendering.
type cliResourceDocArgs struct {
	Title              string
	Comment            string
	DeprecationMessage string

	Examples []cliExample

	// Constructor — one entry per language
	Constructors []cliConstructorEntry

	// Properties — one entry per language
	InputProperties  []cliLangEntry
	OutputProperties []cliLangEntry

	// State lookup — one entry per language
	StateInputs []cliLangEntry

	// Methods
	Methods []cliMethodDocArgs

	// Nested/supporting types — one entry per language
	NestedTypes    []cliNestedType
	MaxNestedTypes int

	// Import
	ImportDocs string

	PackageDetails packageDetails
}

// cliConstructorEntry holds constructor info for a single language.
type cliConstructorEntry struct {
	Tag    string
	Syntax string
	Params []formalParam
}

// cliMethodDocArgs holds method info with per-language data.
type cliMethodDocArgs struct {
	Title              string
	Comment            string
	DeprecationMessage string
	// Syntax per language
	Syntaxes         []cliMethodSyntax
	InputProperties  []cliLangEntry
	OutputProperties []cliLangEntry
	Examples         []cliExample
}

// cliMethodSyntax holds a method signature for one language.
type cliMethodSyntax struct {
	Tag    string
	Syntax string
}

// cliFunctionDocArgs holds all-language data for terminal-friendly function rendering.
type cliFunctionDocArgs struct {
	Title              string
	Comment            string
	DeprecationMessage string

	Examples []cliExample

	// Function signature per language
	Signatures []cliMethodSyntax

	InputProperties  []cliLangEntry
	OutputProperties []cliLangEntry

	NestedTypes []cliNestedType

	PackageDetails packageDetails
}

// cliExample is a single-language example.
type cliExample struct {
	Title   string
	Snippet string
	Tag     string // language tag for the code fence
}

// cliNestedType is a nested type with per-language properties.
type cliNestedType struct {
	Name        string
	Description string
	Properties  []cliLangEntry
}

// cliLanguageTag returns the template language tag for a given language.
// NodeJS maps to "typescript", others use their string representation.
func cliLanguageTag(lang language.Language) string {
	if lang == language.NodeJS {
		return "typescript"
	}
	return lang.String()
}

// toCLIResourceArgs builds all-language CLI resource args from the Hugo resource args.
func toCLIResourceArgs(args resourceDocArgs) cliResourceDocArgs {
	cli := cliResourceDocArgs{
		Title:              args.Header.Title,
		Comment:            stripHTML(args.Comment),
		DeprecationMessage: stripHTML(args.DeprecationMessage),
		ImportDocs:         stripHTML(args.ImportDocs),
		MaxNestedTypes:     args.MaxNestedTypes,
		PackageDetails:     args.PackageDetails,
	}

	// Constructors, properties, state inputs — per language
	for lang := range language.All() {
		tag := cliLanguageTag(lang)

		if cs, ok := args.CreationExampleSyntax[lang]; ok {
			entry := cliConstructorEntry{Tag: tag, Syntax: stripHTML(cs)}
			if fp, ok := args.ConstructorParamsTyped[lang]; ok {
				entry.Params = fp.Params
			}
			cli.Constructors = append(cli.Constructors, entry)
		}

		if props, ok := args.InputProperties[lang]; ok {
			cli.InputProperties = append(cli.InputProperties, cliLangEntry{
				Tag: tag, Properties: stripHTMLFromProperties(props),
			})
		}
		if props, ok := args.OutputProperties[lang]; ok {
			cli.OutputProperties = append(cli.OutputProperties, cliLangEntry{
				Tag: tag, Properties: stripHTMLFromProperties(props),
			})
		}
		if props, ok := args.StateInputs[lang]; ok {
			cli.StateInputs = append(cli.StateInputs, cliLangEntry{
				Tag: tag, Properties: stripHTMLFromProperties(props),
			})
		}
	}

	// Examples — collect all languages
	cli.Examples = extractAllCLIExamples(args.ExamplesSection)

	// Methods
	for _, m := range args.Methods {
		cm := cliMethodDocArgs{
			Title:              m.Title,
			Comment:            stripHTML(m.Comment),
			DeprecationMessage: stripHTML(m.DeprecationMessage),
			Examples:           extractAllCLIExamples(m.ExamplesSection),
		}
		for lang := range language.All() {
			tag := cliLanguageTag(lang)
			if name, ok := m.MethodName[lang]; ok {
				if mArgs, ok := m.MethodArgs[lang]; ok {
					cleanArgs := stripHTML(mArgs)
					syntax := name + "(" + cleanArgs + ")"
					if result, ok := m.MethodResult[lang]; ok && result.Name != "" {
						syntax += " -> " + propertyTypeName(result)
					}
					cm.Syntaxes = append(cm.Syntaxes, cliMethodSyntax{Tag: tag, Syntax: syntax})
				}
			}
			if props, ok := m.InputProperties[lang]; ok {
				cm.InputProperties = append(cm.InputProperties, cliLangEntry{
					Tag: tag, Properties: stripHTMLFromProperties(props),
				})
			}
			if props, ok := m.OutputProperties[lang]; ok {
				cm.OutputProperties = append(cm.OutputProperties, cliLangEntry{
					Tag: tag, Properties: stripHTMLFromProperties(props),
				})
			}
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
		for lang := range language.All() {
			tag := cliLanguageTag(lang)
			if props, ok := nt.Properties[lang]; ok {
				cnt.Properties = append(cnt.Properties, cliLangEntry{
					Tag: tag, Properties: stripHTMLFromProperties(props),
				})
			}
		}
		cli.NestedTypes = append(cli.NestedTypes, cnt)
	}

	return cli
}

// toCLIFunctionArgs builds all-language CLI function args from the Hugo function args.
func toCLIFunctionArgs(args functionDocArgs) cliFunctionDocArgs {
	cli := cliFunctionDocArgs{
		Title:              args.Header.Title,
		Comment:            stripHTML(args.Comment),
		DeprecationMessage: stripHTML(args.DeprecationMessage),
		PackageDetails:     args.PackageDetails,
	}

	for lang := range language.All() {
		tag := cliLanguageTag(lang)

		if name, ok := args.FunctionName[lang]; ok {
			if fArgs, ok := args.FunctionArgs[lang]; ok {
				cleanArgs := stripHTML(fArgs)
				syntax := name + "(" + cleanArgs + ")"
				if result, ok := args.FunctionResult[lang]; ok && result.Name != "" {
					syntax += " -> " + propertyTypeName(result)
				}
				cli.Signatures = append(cli.Signatures, cliMethodSyntax{Tag: tag, Syntax: syntax})
			}
		}

		if props, ok := args.InputProperties[lang]; ok {
			cli.InputProperties = append(cli.InputProperties, cliLangEntry{
				Tag: tag, Properties: stripHTMLFromProperties(props),
			})
		}
		if props, ok := args.OutputProperties[lang]; ok {
			cli.OutputProperties = append(cli.OutputProperties, cliLangEntry{
				Tag: tag, Properties: stripHTMLFromProperties(props),
			})
		}
	}

	cli.Examples = extractAllCLIExamples(args.ExamplesSection)

	for _, nt := range args.NestedTypes {
		cnt := cliNestedType{
			Name:        stripHTML(nt.Name),
			Description: stripHTML(nt.Description),
		}
		for lang := range language.All() {
			tag := cliLanguageTag(lang)
			if props, ok := nt.Properties[lang]; ok {
				cnt.Properties = append(cnt.Properties, cliLangEntry{
					Tag: tag, Properties: stripHTMLFromProperties(props),
				})
			}
		}
		cli.NestedTypes = append(cli.NestedTypes, cnt)
	}

	return cli
}

// CLIDocsBundle is the JSON structure for bundled CLI docs per package.
// It contains all resource and function markdown docs keyed by module/name path.
type CLIDocsBundle struct {
	Version        int               `json:"version"`
	Package        string            `json:"package"`
	PackageVersion string            `json:"packageVersion"`
	Resources      map[string]string `json:"resources"`
	Functions      map[string]string `json:"functions"`
}

// extractAllCLIExamples collects examples for all languages.
func extractAllCLIExamples(es examplesSection) []cliExample {
	var examples []cliExample
	for _, ex := range es.Examples {
		for lang := range language.All() {
			snippet, ok := ex.Snippets[lang]
			if !ok || snippet == "" || snippet == defaultMissingExampleSnippetPlaceholder {
				continue
			}
			examples = append(examples, cliExample{
				Title:   ex.Title,
				Snippet: snippet,
				Tag:     cliLanguageTag(lang),
			})
		}
	}
	return examples
}
