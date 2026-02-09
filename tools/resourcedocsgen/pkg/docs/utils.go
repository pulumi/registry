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
	"strings"
	"unicode"

	dotnet "github.com/pulumi/pulumi-dotnet/pulumi-language-dotnet/v3/codegen"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	go_gen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/slice"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"github.com/pulumi/registry/tools/resourcedocsgen/pkg/util/language"
)

// isDotNetTypeNameBoundary returns true if the prev and next runes represent a boundary between two parts of a .Net/C#
// type name. In C#/.Net, type names are written as Dot.Separated.PascalCase.Strings, so we are looking for a period
// followed by an uppercase character.
func isDotNetTypeNameBoundary(prev rune, next rune) bool {
	return prev == rune('.') && unicode.IsUpper(next)
}

// isPythonTypeNameBoundary returns true if the prev and next runes represent a boundary between two parts of a Python
// type name. In Python, names are snake_cased, so we are looking for an underscore followed by a lowercase character.
func isPythonTypeNameBoundary(prev rune, next rune) bool {
	return (prev == rune('_') && unicode.IsLower(next))
}

// wbr inserts an HTML <wbr> tag whenever a casing change is detected in a string. The <wbr> HTML element represents a
// word break opportunity—a position within text where the browser may optionally break a line, though its line-breaking
// rules would not otherwise create a break at that location. Using wbr on names when generating documentation allows us
// to help the browser wrap long names in a more readable way.
//
// Examples:
//
//	wbr("fooBar") == "foo<wbr>Bar"
//	wbr("fooBarBaz") == "foo<wbr>Bar<wbr>Baz"
func wbr(s string) string {
	runes := slice.Prealloc[rune](len(s))
	var prev rune
	for i, r := range s {
		if i != 0 &&
			// For TS, JS and Go, property names are camelCase and types are PascalCase.
			((unicode.IsLower(prev) && unicode.IsUpper(r)) ||
				isDotNetTypeNameBoundary(prev, r) ||
				isPythonTypeNameBoundary(prev, r)) {
			runes = append(runes, []rune("<wbr>")...)
		}
		runes = append(runes, r)
		prev = r
	}
	return string(runes)
}

// tokenToName returns the resource name from a Pulumi token.
//
// Examples:
//
//	tokenToName("aws:index:Foo") == "Foo"
func tokenToName(tok string) string {
	components := strings.Split(tok, ":")
	contract.Assertf(len(components) == 3, "malformed token %v", tok)
	return components[2]
}

// tokenToPackageName returns the package name from a Pulumi token.
//
// Examples:
//
//	tokenToPackageName("aws:index:Foo") == "aws"
func tokenToPackageName(tok string) string {
	components := strings.Split(tok, ":")
	contract.Assertf(len(components) == 3, "malformed token %v", tok)
	return components[0]
}

func title(s string, lang language.Language) string {
	switch lang {
	case language.Go:
		return go_gen.Title(s)
	case language.CSharp:
		return dotnet.Title(s)
	default:
		//nolint:staticcheck
		return strings.Title(s)
	}
}

func modFilenameToDisplayName(name string) string {
	parts := strings.Split(name, "/")
	return parts[len(parts)-1]
}

func getModuleLink(name string) string {
	return strings.ToLower(name) + "/"
}

func getResourceLink(name string) string {
	link := strings.ToLower(name)
	// Handle URL generation for resources named `index`. We prepend a double underscore here, since a link of
	// .../<module>/index has trouble resolving and returns a 404 in the browser, likely due to `index` being some sort of
	// reserved keyword.
	if link == "index" {
		return "--" + link
	}
	return link
}

func getFunctionLink(name string) string {
	return strings.ToLower(name)
}

// isExternalType retusn true if and only if the given type is external to (i.e. not defined within) the given package.
func isExternalType(t schema.Type, pkg schema.PackageReference) (isExternal bool) {
	switch typ := t.(type) {
	case *schema.ObjectType:
		return typ.PackageReference != nil && !codegen.PkgEquals(typ.PackageReference, pkg)
	case *schema.ResourceType:
		return typ.Resource != nil && pkg != nil && !codegen.PkgEquals(typ.Resource.PackageReference, pkg)
	case *schema.EnumType:
		return pkg != nil && !codegen.PkgEquals(typ.PackageReference, pkg)
	}
	return
}

// Iterate character by character and remove underscores if that underscore
// is at the very front of an identifier, follows a special character, and is not a delimeter
// within an identifier.
func removeLeadingUnderscores(s string) string {
	var sb strings.Builder
	lastChar := ' '
	for _, ch := range s {
		if ch != '_' || (unicode.IsLetter(lastChar)) {
			sb.WriteRune(ch)
		}
		lastChar = ch
	}
	return sb.String()
}

// noCopy tells go vet that any type that contains this type should not be copied.
//
// https://github.com/golang/go/issues/8005#issuecomment-190753527
type noCopy struct{}

// Lock is a type hint for go vet. It need not be called.
func (*noCopy) Lock() {}

// Convert a markdown language tag to a [language.Language] if there is a corresponding
// language. If no language exists, then _, false will be returned.
func convertMarkdownLanguage(s string) (language.Language, bool) {
	switch s {
	case "csharp":
		return language.CSharp, true
	case "go":
		return language.Go, true
	case "python":
		return language.Python, true
	case "typescript":
		return language.NodeJS, true
	case "yaml":
		return language.YAML, true
	case "java":
		return language.Java, true
	default:
		return language.Language{}, false
	}
}
