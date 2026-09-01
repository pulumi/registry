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
	"regexp"
	"strings"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/registry/tools/resourcedocsgen/pkg/util/language"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	providerPackage = "prov"
	codeFence       = "```"
)

var simpleProperties = map[string]schema.PropertySpec{
	"stringProp": {
		Description: "A string prop.",
		TypeSpec: schema.TypeSpec{
			Type: "string",
		},
	},
	"boolProp": {
		Description: "A bool prop.",
		TypeSpec: schema.TypeSpec{
			Type: "boolean",
		},
	},
}

// newTestPackageSpec returns a new fake package spec for a Provider used for testing.
func newTestPackageSpec() schema.PackageSpec {
	pythonMapCase := map[string]schema.RawMessage{
		"python": schema.RawMessage(`{"mapCase":false}`),
	}
	return schema.PackageSpec{
		Name:        providerPackage,
		Version:     "0.0.1",
		Description: "A fake provider package used for testing.",
		Meta: &schema.MetadataSpec{
			ModuleFormat: "(.*)(?:/[^/]*)",
		},
		Types: map[string]schema.ComplexTypeSpec{
			// Package-level types.
			"prov:/getPackageResourceOptions:getPackageResourceOptions": {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: "Options object for the package-level function getPackageResource.",
					Type:        "object",
					Properties:  simpleProperties,
				},
			},

			// Module-level types.
			"prov:module/getModuleResourceOptions:getModuleResourceOptions": {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: "Options object for the module-level function getModuleResource.",
					Type:        "object",
					Properties:  simpleProperties,
				},
			},
			"prov:module/ResourceOptions:ResourceOptions": {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: "The resource options object.",
					Type:        "object",
					Properties: map[string]schema.PropertySpec{
						"stringProp": {
							Description: "A string prop.",
							Language:    pythonMapCase,
							TypeSpec: schema.TypeSpec{
								Type: "string",
							},
						},
						"boolProp": {
							Description: "A bool prop.",
							Language:    pythonMapCase,
							TypeSpec: schema.TypeSpec{
								Type: "boolean",
							},
						},
						"recursiveType": {
							Description: "I am a recursive type.",
							Language:    pythonMapCase,
							TypeSpec: schema.TypeSpec{
								Ref: "#/types/prov:module/ResourceOptions:ResourceOptions",
							},
						},
					},
				},
			},
			"prov:module/ResourceOptions2:ResourceOptions2": {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: "The resource options object.",
					Type:        "object",
					Properties: map[string]schema.PropertySpec{
						"uniqueProp": {
							Description: "This is a property unique to this type.",
							Language:    pythonMapCase,
							TypeSpec: schema.TypeSpec{
								Type: "number",
							},
						},
					},
				},
			},
		},
		Provider: &schema.ResourceSpec{
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: fmt.Sprintf("The provider type for the %s package.", providerPackage),
				Type:        "object",
			},
			InputProperties: map[string]schema.PropertySpec{
				"stringProp": {
					Description: "A stringProp for the provider resource.",
					TypeSpec: schema.TypeSpec{
						Type: "string",
					},
				},
			},
		},
		Resources: map[string]schema.ResourceSpec{
			"prov:module2/resource2:Resource2": {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: `This is a module-level resource called Resource.
{{% examples %}}
## Example Usage

{{% example %}}
### Basic Example

` + codeFence + `typescript
					// Some TypeScript code.
` + codeFence + `
` + codeFence + `python
					# Some Python code.
` + codeFence + `
{{% /example %}}
{{% example %}}
### Custom Sub-Domain Example

` + codeFence + `typescript
					// Some typescript code
` + codeFence + `
` + codeFence + `python
					# Some Python code.
` + codeFence + `
{{% /example %}}
{{% /examples %}}

## Import

The import docs would be here

` + codeFence + `sh
$ pulumi import prov:module/resource:Resource test test
` + codeFence + `
`,
				},
				InputProperties: map[string]schema.PropertySpec{
					"integerProp": {
						Description: "This is integerProp's description.",
						TypeSpec: schema.TypeSpec{
							Type: "integer",
						},
					},
					"stringProp": {
						Description: "This is stringProp's description.",
						TypeSpec: schema.TypeSpec{
							Type: "string",
						},
					},
					"boolProp": {
						Description: "A bool prop.",
						TypeSpec: schema.TypeSpec{
							Type: "boolean",
						},
					},
					"optionsProp": {
						TypeSpec: schema.TypeSpec{
							Ref: "#/types/prov:module/ResourceOptions:ResourceOptions",
						},
					},
					"options2Prop": {
						TypeSpec: schema.TypeSpec{
							Ref: "#/types/prov:module/ResourceOptions2:ResourceOptions2",
						},
					},
					"recursiveType": {
						Description: "I am a recursive type.",
						TypeSpec: schema.TypeSpec{
							Ref: "#/types/prov:module/ResourceOptions:ResourceOptions",
						},
					},
				},
			},
			"prov:module/resource:Resource": {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: `This is a module-level resource called Resource.
{{% examples %}}
## Example Usage

{{% example %}}
### Basic Example

` + codeFence + `typescript
					// Some TypeScript code.
` + codeFence + `
` + codeFence + `python
					# Some Python code.
` + codeFence + `
{{% /example %}}
{{% example %}}
### Custom Sub-Domain Example

` + codeFence + `typescript
					// Some typescript code
` + codeFence + `
` + codeFence + `python
					# Some Python code.
` + codeFence + `
{{% /example %}}
{{% /examples %}}

## Import

The import docs would be here

` + codeFence + `sh
$ pulumi import prov:module/resource:Resource test test
` + codeFence + `
`,
				},
				InputProperties: map[string]schema.PropertySpec{
					"integerProp": {
						Description: "This is integerProp's description.",
						TypeSpec: schema.TypeSpec{
							Type: "integer",
						},
					},
					"stringProp": {
						Description: "This is stringProp's description.",
						TypeSpec: schema.TypeSpec{
							Type: "string",
						},
					},
					"boolProp": {
						Description: "A bool prop.",
						TypeSpec: schema.TypeSpec{
							Type: "boolean",
						},
					},
					"optionsProp": {
						TypeSpec: schema.TypeSpec{
							Ref: "#/types/prov:module/ResourceOptions:ResourceOptions",
						},
					},
					"options2Prop": {
						TypeSpec: schema.TypeSpec{
							Ref: "#/types/prov:module/ResourceOptions2:ResourceOptions2",
						},
					},
					"recursiveType": {
						Description: "I am a recursive type.",
						TypeSpec: schema.TypeSpec{
							Ref: "#/types/prov:module/ResourceOptions:ResourceOptions",
						},
					},
				},
			},
			"prov:/packageLevelResource:PackageLevelResource": {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: "This is a package-level resource.",
				},
				InputProperties: map[string]schema.PropertySpec{
					"prop": {
						Description: "An input property.",
						TypeSpec: schema.TypeSpec{
							Type: "string",
						},
					},
				},
			},
		},
		Functions: map[string]schema.FunctionSpec{
			// Package-level Functions.
			"prov:/getPackageResource:getPackageResource": {
				Description: "A package-level function.",
				Inputs: &schema.ObjectTypeSpec{
					Description: "Inputs for getPackageResource.",
					Type:        "object",
					Properties: map[string]schema.PropertySpec{
						"options": {
							TypeSpec: schema.TypeSpec{
								Ref: "#/types/prov:/getPackageResourceOptions:getPackageResourceOptions",
							},
						},
					},
				},
				Outputs: &schema.ObjectTypeSpec{
					Description: "Outputs for getPackageResource.",
					Properties:  simpleProperties,
					Type:        "object",
				},
			},

			// Module-level Functions.
			"prov:module/getModuleResource:getModuleResource": {
				Description: "A module-level function.",
				Inputs: &schema.ObjectTypeSpec{
					Description: "Inputs for getModuleResource.",
					Type:        "object",
					Properties: map[string]schema.PropertySpec{
						"options": {
							TypeSpec: schema.TypeSpec{
								Ref: "#/types/prov:module/getModuleResource:getModuleResource",
							},
						},
					},
				},
				Outputs: &schema.ObjectTypeSpec{
					Description: "Outputs for getModuleResource.",
					Properties:  simpleProperties,
					Type:        "object",
				},
			},
		},
	}
}

func getResourceFromModule(resource string, mod *modContext) *schema.Resource {
	for _, r := range mod.resources {
		if resourceName(r) != resource {
			continue
		}
		return r
	}
	return nil
}

func getFunctionFromModule(function string, mod *modContext) *schema.Function {
	for _, f := range mod.functions {
		if tokenToName(f.Token) != function {
			continue
		}
		return f
	}
	return nil
}

func TestFunctionHeaders(t *testing.T) {
	t.Parallel()

	testPackageSpec := newTestPackageSpec()

	schemaPkg, err := schema.ImportSpec(testPackageSpec, nil, schema.NewNullLoader(), schema.ValidationOptions{
		AllowDanglingReferences: true,
	})
	assert.NoError(t, err, "importing spec")

	dctx := NewContext("test", schemaPkg)

	tests := []struct {
		ExpectedTitleTag string
		FunctionName     string
		ModuleName       string
		ExpectedMetaDesc string
	}{
		{
			FunctionName: "getPackageResource",
			// Empty string indicates the package-level root module.
			ModuleName:       "",
			ExpectedTitleTag: "prov.getPackageResource",
			// The schema description is too short to stand alone as a meta description, so
			// this falls back to the generic template.
			ExpectedMetaDesc: "Use prov.getPackageResource with Pulumi. Full API reference with input" +
				" and output properties and examples in TypeScript, Python, Go, C#, Java, and YAML.",
		},
		{
			FunctionName:     "getModuleResource",
			ModuleName:       "module",
			ExpectedTitleTag: "prov.module.getModuleResource",
			ExpectedMetaDesc: "Use prov.module.getModuleResource with Pulumi. Full API reference with" +
				" input and output properties and examples in TypeScript, Python, Go, C#, Java, and YAML.",
		},
	}

	modules := dctx.generateModulesFromSchemaPackage("test", schemaPkg)
	for _, test := range tests {
		test := test
		t.Run(test.FunctionName, func(t *testing.T) {
			t.Parallel()

			mod, ok := modules[test.ModuleName]
			if !ok {
				t.Fatalf("could not find the module %s in modules map", test.ModuleName)
			}

			f := getFunctionFromModule(test.FunctionName, mod)
			if f == nil {
				t.Fatalf("could not find %s in modules", test.FunctionName)
			}
			h := mod.genFunctionHeader(f)
			assert.Equal(t, test.ExpectedTitleTag, h.TitleTag)
			assert.Equal(t, test.ExpectedMetaDesc, h.MetaDesc)
		})
	}
}

func TestResourceDocHeader(t *testing.T) {
	t.Parallel()

	testPackageSpec := newTestPackageSpec()

	schemaPkg, err := schema.ImportSpec(testPackageSpec, nil, schema.NewNullLoader(), schema.ValidationOptions{
		AllowDanglingReferences: true,
	})
	assert.NoError(t, err, "importing spec")

	dctx := NewContext("test", schemaPkg)

	tests := []struct {
		Name             string
		ExpectedTitleTag string
		ResourceName     string
		ModuleName       string
		ExpectedMetaDesc string
	}{
		{
			Name:         "PackageLevelResourceHeader",
			ResourceName: "PackageLevelResource",
			// Empty string indicates the package-level root module.
			ModuleName:       "",
			ExpectedTitleTag: "prov.PackageLevelResource",
			// The schema description is too short to stand alone as a meta description, so
			// this falls back to the generic template.
			ExpectedMetaDesc: "Create and manage prov.PackageLevelResource with Pulumi. Full API reference" +
				" with input and output properties, lookup functions, and examples in TypeScript, Python," +
				" Go, C#, Java, and YAML.",
		},
		{
			Name:             "ModuleLevelResourceHeader",
			ResourceName:     "Resource",
			ModuleName:       "module",
			ExpectedTitleTag: "prov.module.Resource",
			// The schema description's opening summary sentence is used directly, with the
			// following {{% examples %}} shortcode block dropped.
			ExpectedMetaDesc: "This is a module-level resource called Resource.",
		},
	}

	modules := dctx.generateModulesFromSchemaPackage("test", schemaPkg)
	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			mod, ok := modules[test.ModuleName]
			if !ok {
				t.Fatalf("could not find the module %s in modules map", test.ModuleName)
			}

			r := getResourceFromModule(test.ResourceName, mod)
			if r == nil {
				t.Fatalf("could not find %s in modules", test.ResourceName)
			}
			h := mod.genResourceHeader(r)
			assert.Equal(t, test.ExpectedTitleTag, h.TitleTag)
			assert.Equal(t, test.ExpectedMetaDesc, h.MetaDesc)
		})
	}
}

func TestExamplesProcessing(t *testing.T) {
	t.Parallel()

	testPackageSpec := newTestPackageSpec()

	schemaPkg, err := schema.ImportSpec(testPackageSpec, nil, schema.NewNullLoader(), schema.ValidationOptions{
		AllowDanglingReferences: true,
	})
	assert.NoError(t, err, "importing spec")
	dctx := NewContext("test", schemaPkg)

	description := testPackageSpec.Resources["prov:module/resource:Resource"].Description
	docInfo := dctx.decomposeDocstring(description, dctx.getSupportedSnippetLanguages(false, nil))
	examplesSection := docInfo.examples
	importSection := docInfo.importDetails

	assert.NotEmpty(t, importSection)

	// The resource under test has two examples and both have TS and Python examples.
	assert.Equal(t, 2, len(examplesSection))
	assert.Equal(t, "### Basic Example", examplesSection[0].Title)
	assert.Equal(t, "### Custom Sub-Domain Example", examplesSection[1].Title)
	expectedLangSnippets := []language.Language{language.NodeJS, language.Python}
	otherLangSnippets := []language.Language{language.CSharp, language.Go}
	for _, e := range examplesSection {
		for _, lang := range expectedLangSnippets {
			_, ok := e.Snippets[lang]
			assert.True(t, ok, "Could not find %s snippet", lang)
		}
		for _, lang := range otherLangSnippets {
			snippet, ok := e.Snippets[lang]
			assert.True(t, ok, "Expected to find default placeholders for other languages")
			assert.Contains(t, "Example coming soon!", snippet)
		}
	}
}

func TestDecomposeDocstring(t *testing.T) {
	t.Parallel()
	awsVpcDocs := "Provides a VPC resource.\n" +
		"\n" +
		"{{% examples %}}\n" +
		"## Example Usage\n" +
		"{{% example %}}\n" +
		"\n" +
		"Basic usage:\n" +
		"\n" +
		"```typescript\n" +
		"Basic usage: typescript\n" +
		"```\n" +
		"```python\n" +
		"Basic usage: python\n" +
		"```\n" +
		"```csharp\n" +
		"Basic usage: csharp\n" +
		"```\n" +
		"```go\n" +
		"Basic usage: go\n" +
		"```\n" +
		"```java\n" +
		"Basic usage: java\n" +
		"```\n" +
		"```yaml\n" +
		"Basic usage: yaml\n" +
		"```\n" +
		"\n" +
		"Basic usage with tags:\n" +
		"\n" +
		"```typescript\n" +
		"Basic usage with tags: typescript\n" +
		"```\n" +
		"```python\n" +
		"Basic usage with tags: python\n" +
		"```\n" +
		"```csharp\n" +
		"Basic usage with tags: csharp\n" +
		"```\n" +
		"```go\n" +
		"Basic usage with tags: go\n" +
		"```\n" +
		"```java\n" +
		"Basic usage with tags: java\n" +
		"```\n" +
		"```yaml\n" +
		"Basic usage with tags: yaml\n" +
		"```\n" +
		"\n" +
		"VPC with CIDR from AWS IPAM:\n" +
		"\n" +
		"```typescript\n" +
		"VPC with CIDR from AWS IPAM: typescript\n" +
		"```\n" +
		"```python\n" +
		"VPC with CIDR from AWS IPAM: python\n" +
		"```\n" +
		"```csharp\n" +
		"VPC with CIDR from AWS IPAM: csharp\n" +
		"```\n" +
		"```java\n" +
		"VPC with CIDR from AWS IPAM: java\n" +
		"```\n" +
		"```yaml\n" +
		"VPC with CIDR from AWS IPAM: yaml\n" +
		"```\n" +
		"{{% /example %}}\n" +
		"{{% /examples %}}\n" +
		"\n" +
		"## Import\n" +
		"\n" +
		"VPCs can be imported using the `vpc id`, e.g.,\n" +
		"\n" +
		"```sh\n" +
		" $ pulumi import aws:ec2/vpc:Vpc test_vpc vpc-a01106c2\n" +
		"```\n" +
		"\n" +
		" "

	testPackageSpec := newTestPackageSpec()
	schemaPkg, err := schema.ImportSpec(testPackageSpec, nil, schema.NewNullLoader(), schema.ValidationOptions{
		AllowDanglingReferences: true,
	})
	assert.NoError(t, err, "importing spec")
	dctx := NewContext("test", schemaPkg)

	info := dctx.decomposeDocstring(awsVpcDocs, dctx.getSupportedSnippetLanguages(false, nil))
	assert.Equal(t, docInfo{
		description: "Provides a VPC resource.\n",
		examples: []exampleSection{
			{
				Title: "Basic usage",
				Snippets: map[language.Language]string{
					language.CSharp: "```csharp\nBasic usage: csharp\n```",
					language.Go:     "```go\nBasic usage: go\n```",
					language.Java:   "```java\nBasic usage: java\n```",
					language.Python: "```python\nBasic usage: python\n```",
					language.NodeJS: "\n```typescript\nBasic usage: typescript\n```",
					language.YAML:   "```yaml\nBasic usage: yaml\n```",
					language.HCL:    "Example coming soon!",
				},
			},
			{
				Title: "Basic usage with tags",
				Snippets: map[language.Language]string{
					language.CSharp: "```csharp\nBasic usage with tags: csharp\n```",
					language.Go:     "```go\nBasic usage with tags: go\n```",
					language.Java:   "```java\nBasic usage with tags: java\n```",
					language.Python: "```python\nBasic usage with tags: python\n```",
					language.NodeJS: "\n```typescript\nBasic usage with tags: typescript\n```",
					language.YAML:   "```yaml\nBasic usage with tags: yaml\n```",
					language.HCL:    "Example coming soon!",
				},
			},
			{
				Title: "VPC with CIDR from AWS IPAM",
				Snippets: map[language.Language]string{
					language.CSharp: "```csharp\nVPC with CIDR from AWS IPAM: csharp\n```",
					language.Go:     "Example coming soon!",
					language.Java:   "```java\nVPC with CIDR from AWS IPAM: java\n```",
					language.Python: "```python\nVPC with CIDR from AWS IPAM: python\n```",
					language.NodeJS: "\n```typescript\nVPC with CIDR from AWS IPAM: typescript\n```",
					language.YAML:   "```yaml\nVPC with CIDR from AWS IPAM: yaml\n```",
					language.HCL:    "Example coming soon!",
				},
			},
		},
		importDetails: "\n\nVPCs can be imported using the `vpc id`, e.g.,\n\n" +
			"```sh\n $ pulumi import aws:ec2/vpc:Vpc test_vpc vpc-a01106c2\n```",
	},
		info)
}

func TestGenOverlayResource(t *testing.T) {
	t.Parallel()
	getSchemaPackage := func() *schema.Package {
		testPackageSpec := newTestPackageSpec()
		testPackageSpec.Resources["prov:module/overlayResource:OverlayResource"] = schema.ResourceSpec{
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "This is a module-level resource called OverlayResource.",
				IsOverlay:   true,
			},
			InputProperties: map[string]schema.PropertySpec{
				"prop": {
					Description: "An input property.",
					TypeSpec: schema.TypeSpec{
						Type: "string",
					},
				},
			},
		}

		//nolint:lll
		testPackageSpec.Resources["prov:module/overlayResourceConstrainedLanguages:OverlayResourceConstrainedLanguages"] = schema.ResourceSpec{
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description:               "This is a module-level resource called OverlayResourceConstrainedLanguages.",
				IsOverlay:                 true,
				OverlaySupportedLanguages: []string{"python", "go", "nodejs"},
			},
			InputProperties: map[string]schema.PropertySpec{
				"prop": {
					Description: "An input property.",
					TypeSpec: schema.TypeSpec{
						Type: "string",
					},
				},
			},
		}

		//nolint:lll
		testPackageSpec.Resources["prov:module/overlayResourceWrongLanguage:OverlayResourceWrongLanguage"] = schema.ResourceSpec{
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description:               "This is a module-level resource called OverlayResourceWrongLanguage.",
				IsOverlay:                 true,
				OverlaySupportedLanguages: []string{"python", "go", "nodejs", "smalltalk"},
			},
			InputProperties: map[string]schema.PropertySpec{
				"prop": {
					Description: "An input property.",
					TypeSpec: schema.TypeSpec{
						Type: "string",
					},
				},
			},
		}

		schemaPkg, err := schema.ImportSpec(testPackageSpec, nil, schema.NewNullLoader(), schema.ValidationOptions{
			AllowDanglingReferences: true,
		})
		assert.NoError(t, err, "importing spec")
		return schemaPkg
	}

	tests := []struct {
		ResourceName                 string
		ExpectedLangChooserLanguages string
	}{
		{
			// regular resource, should support all languages (i.e. ExpectedLangChooserLanguages should be empty)
			ResourceName:                 "Resource",
			ExpectedLangChooserLanguages: "csharp,go,typescript,python,yaml,java,hcl",
		},
		{
			// regular overlay resource, should support all languages (i.e. ExpectedLangChooserLanguages should be empty)
			ResourceName:                 "OverlayResource",
			ExpectedLangChooserLanguages: "csharp,go,typescript,python,yaml,java,hcl",
		},
		{
			// overlay resource with a constrained list of supported languages should support only the languages specified in
			// OverlaySupportedLanguages
			ResourceName:                 "OverlayResourceConstrainedLanguages",
			ExpectedLangChooserLanguages: "python,go,typescript",
		},
		{
			// overlay resource with a wrong language in OverlaySupportedLanguages should filter out the wrong language
			ResourceName:                 "OverlayResourceWrongLanguage",
			ExpectedLangChooserLanguages: "python,go,typescript",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.ResourceName, func(t *testing.T) {
			t.Parallel()
			schemaPkg := getSchemaPackage()
			dctx := NewContext("test", schemaPkg)
			modules := dctx.generateModulesFromSchemaPackage("test", schemaPkg)
			mod, ok := modules["module"]
			if !ok {
				t.Fatalf("could not find the module 'module' in modules map")
			}

			r := getResourceFromModule(test.ResourceName, mod)
			if r == nil {
				t.Fatalf("could not find %s in modules", test.ResourceName)
			}
			resourceDocs := mod.genResource(r)
			assert.Equal(t, test.ExpectedLangChooserLanguages, resourceDocs.LangChooserLanguages)
		})
	}
}

func TestGenOverlayFunction(t *testing.T) {
	t.Parallel()
	getSchemaPackage := func() *schema.Package {
		testPackageSpec := newTestPackageSpec()
		testPackageSpec.Functions["prov:module/overlayFunction:overlayFunction"] = schema.FunctionSpec{
			Description: "A module-level function.",
			IsOverlay:   true,
			Inputs: &schema.ObjectTypeSpec{
				Description: "Inputs for getModuleResource.",
				Type:        "object",
				Properties: map[string]schema.PropertySpec{
					"options": {
						TypeSpec: schema.TypeSpec{
							Ref: "#/types/prov:module/getModuleResource:getModuleResource",
						},
					},
				},
			},
			Outputs: &schema.ObjectTypeSpec{
				Description: "Outputs for getModuleResource.",
				Properties:  simpleProperties,
				Type:        "object",
			},
		}

		//nolint:lll
		testPackageSpec.Functions["prov:module/overlayFunctionConstrainedLanguages:overlayFunctionConstrainedLanguages"] = schema.FunctionSpec{
			Description:               "A module-level function.",
			IsOverlay:                 true,
			OverlaySupportedLanguages: []string{"python", "go", "nodejs"},
			Inputs: &schema.ObjectTypeSpec{
				Description: "Inputs for getModuleResource.",
				Type:        "object",
				Properties: map[string]schema.PropertySpec{
					"options": {
						TypeSpec: schema.TypeSpec{
							Ref: "#/types/prov:module/getModuleResource:getModuleResource",
						},
					},
				},
			},
			Outputs: &schema.ObjectTypeSpec{
				Description: "Outputs for getModuleResource.",
				Properties:  simpleProperties,
				Type:        "object",
			},
		}

		//nolint:lll
		testPackageSpec.Functions["prov:module/overlayFunctionWrongLanguage:overlayFunctionWrongLanguage"] = schema.FunctionSpec{
			Description:               "A module-level function.",
			IsOverlay:                 true,
			OverlaySupportedLanguages: []string{"python", "go", "nodejs", "smalltalk"},
			Inputs: &schema.ObjectTypeSpec{
				Description: "Inputs for getModuleResource.",
				Type:        "object",
				Properties: map[string]schema.PropertySpec{
					"options": {
						TypeSpec: schema.TypeSpec{
							Ref: "#/types/prov:module/getModuleResource:getModuleResource",
						},
					},
				},
			},
			Outputs: &schema.ObjectTypeSpec{
				Description: "Outputs for getModuleResource.",
				Properties:  simpleProperties,
				Type:        "object",
			},
		}

		schemaPkg, err := schema.ImportSpec(testPackageSpec, nil, schema.NewNullLoader(), schema.ValidationOptions{
			AllowDanglingReferences: true,
		})
		assert.NoError(t, err, "importing spec")
		return schemaPkg
	}

	tests := []struct {
		FunctionName                 string
		ExpectedLangChooserLanguages string
	}{
		{
			// regular function, should support all languages
			FunctionName:                 "getModuleResource",
			ExpectedLangChooserLanguages: "csharp,go,typescript,python,yaml,java,hcl",
		},
		{
			// regular overlay function, should support all languages
			FunctionName:                 "overlayFunction",
			ExpectedLangChooserLanguages: "csharp,go,typescript,python,yaml,java,hcl",
		},
		{
			// overlay function with a constrained list of supported languages should support only the languages specified in
			// OverlaySupportedLanguages
			FunctionName:                 "overlayFunctionConstrainedLanguages",
			ExpectedLangChooserLanguages: "python,go,typescript",
		},
		{
			// overlay function with a wrong language in OverlaySupportedLanguages should filter out the wrong language
			FunctionName:                 "overlayFunctionWrongLanguage",
			ExpectedLangChooserLanguages: "python,go,typescript",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.FunctionName, func(t *testing.T) {
			t.Parallel()
			schemaPkg := getSchemaPackage()
			dctx := NewContext("test", schemaPkg)
			modules := dctx.generateModulesFromSchemaPackage("test", schemaPkg)
			mod, ok := modules["module"]
			if !ok {
				t.Fatalf("could not find the module 'module' in modules map")
			}

			f := getFunctionFromModule(test.FunctionName, mod)
			if f == nil {
				t.Fatalf("could not find %s in modules", test.FunctionName)
			}
			resourceDocs := mod.genFunction(f)
			assert.Equal(t, test.ExpectedLangChooserLanguages, resourceDocs.LangChooserLanguages)
		})
	}
}

func TestHasOversizedCreationExample(t *testing.T) {
	t.Parallel()

	small := strings.Repeat("x", maxCreationExampleSyntaxBytes)
	tooBig := strings.Repeat("x", maxCreationExampleSyntaxBytes+1)

	assert.False(t, hasOversizedCreationExample(nil))
	assert.False(t, hasOversizedCreationExample(map[language.Language]string{
		language.Go:     small,
		language.Python: small,
	}))
	assert.True(t, hasOversizedCreationExample(map[language.Language]string{
		language.Go:     small,
		language.Python: tooBig,
	}))
}

func TestOversizedCreationExampleKeepsCLIConstructorSection(t *testing.T) {
	t.Parallel()

	spec := newTestPackageSpec()
	inputs := map[string]schema.PropertySpec{}
	for i := range 800 {
		name := fmt.Sprintf("configurationPropertyWithALongName%03d", i)
		inputs[name] = schema.PropertySpec{TypeSpec: schema.TypeSpec{Type: "string"}}
	}
	spec.Resources["prov:module/oversizedResource:OversizedResource"] = schema.ResourceSpec{
		ObjectTypeSpec:  schema.ObjectTypeSpec{Description: "Resource whose constructor example exceeds the size limit."},
		InputProperties: inputs,
	}

	schemaPkg, err := schema.ImportSpec(spec, nil, schema.NewNullLoader(), schema.ValidationOptions{
		AllowDanglingReferences: true,
	})
	require.NoError(t, err, "importing spec")

	bundle, err := NewContext("test", schemaPkg).GenerateCLIPackage()
	require.NoError(t, err, "generating CLI package")

	var content string
	for key, entry := range bundle.Resources {
		if strings.Contains(strings.ToLower(key), "oversizedresource") {
			content = entry.Content
			break
		}
	}
	require.NotEmpty(t, content, "expected an entry for the oversized resource")

	assert.Regexp(t, regexp.MustCompile("(?m)^"+regexp.QuoteMeta(codeFence)+`\w*\n\n`+regexp.QuoteMeta(codeFence)+"$"),
		content, "constructor syntax should be blank")
	assert.Contains(t, content, "## Create OversizedResource Resource")
	assert.Contains(t, content, "### Parameters")
}

// TestFunctionInvokeOptionsTypes pins the options type rendered on each of a function's signatures. The Output version
// accepts `InvokeOutputOptions` (which adds `dependsOn`), but how that surfaces differs per language: TypeScript and
// Python swap the type outright, Go keeps its variadic `pulumi.InvokeOption`, and C# and Java gain a second overload.
func TestFunctionInvokeOptionsTypes(t *testing.T) {
	t.Parallel()

	schemaPkg, err := schema.ImportSpec(newTestPackageSpec(), nil, schema.NewNullLoader(), schema.ValidationOptions{
		AllowDanglingReferences: true,
	})
	require.NoError(t, err, "importing spec")

	dctx := NewContext("test", schemaPkg)
	mod, ok := dctx.generateModulesFromSchemaPackage("test", schemaPkg)["module"]
	require.True(t, ok, "could not find the module 'module' in modules map")

	f := getFunctionFromModule("getModuleResource", mod)
	require.NotNil(t, f, "could not find getModuleResource in module")
	require.True(t, f.NeedsOutputVersion(), "test function is expected to have an Output version")

	args := mod.genFunction(f)

	// optsType returns everything after the `opts` parameter name, which for the languages that put the type after the
	// name is the type itself plus any default value.
	optsType := func(rendered string) string {
		_, opts, found := strings.Cut(stripHTML(rendered), "opts")
		require.True(t, found, "no opts parameter in %q", rendered)
		return strings.TrimSpace(opts)
	}

	t.Run("direct form", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, "?: InvokeOptions", optsType(args.FunctionArgs[language.NodeJS]))
		assert.Equal(t, ": Optional[InvokeOptions] = None", optsType(args.FunctionArgs[language.Python]))
		assert.Equal(t, "...InvokeOption", optsType(args.FunctionArgs[language.Go]))
		// C# and Java put the type ahead of the parameter name.
		assert.Contains(t, stripHTML(args.FunctionArgs[language.CSharp]), "InvokeOptions? opts = null")
		assert.Contains(t, stripHTML(args.FunctionArgs[language.Java]), "InvokeOptions options")
	})

	t.Run("output form", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, "?: InvokeOutputOptions", optsType(args.FunctionArgsOutputVersion[language.NodeJS]))
		assert.Equal(t, ": Optional[InvokeOutputOptions] = None", optsType(args.FunctionArgsOutputVersion[language.Python]))
		// Go's Output version keeps the variadic `pulumi.InvokeOption`; it builds the InvokeOutputOptions internally.
		assert.Equal(t, "...InvokeOption", optsType(args.FunctionArgsOutputVersion[language.Go]))
		// C# and Java keep the InvokeOptions overload and add a second one below.
		assert.Contains(t, stripHTML(args.FunctionArgsOutputVersion[language.CSharp]), "InvokeOptions? opts = null")
		assert.Contains(t, stripHTML(args.FunctionArgsOutputVersion[language.Java]), "InvokeOptions options")
	})

	t.Run("output form InvokeOutputOptions overload", func(t *testing.T) {
		t.Parallel()
		assert.Contains(t, stripHTML(args.FunctionArgsOutputOptions[language.CSharp]), "InvokeOutputOptions opts")
		assert.Contains(t, stripHTML(args.FunctionArgsOutputOptions[language.Java]), "InvokeOutputOptions options")
		// Only C# and Java render this as a distinct overload.
		assert.NotContains(t, args.FunctionArgsOutputOptions, language.NodeJS)
		assert.NotContains(t, args.FunctionArgsOutputOptions, language.Python)
		assert.NotContains(t, args.FunctionArgsOutputOptions, language.Go)
	})
}
