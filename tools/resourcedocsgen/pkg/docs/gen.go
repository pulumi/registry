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

//go:generate go run bundler.go

package docs

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"html"
	"path"
	"sort"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v3/codegen/pcl"
	"github.com/pulumi/registry/tools/resourcedocsgen/pkg/docs/templates"
	"github.com/pulumi/registry/tools/resourcedocsgen/pkg/util/language"

	"github.com/golang/glog"

	"github.com/pulumi/pulumi-java/pkg/codegen/java"
	yaml "github.com/pulumi/pulumi-yaml/pkg/pulumiyaml/codegen"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/pulumi/pulumi/pkg/v3/codegen/dotnet"
	go_gen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
	"github.com/pulumi/pulumi/pkg/v3/codegen/nodejs"
	"github.com/pulumi/pulumi/pkg/v3/codegen/python"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/slice"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

// titleLookup maps a package name to a desired display name.
//
// NOTE: This can be removed when all Pulumi-managed packages have a DisplayName in their schema. See
// pulumi/pulumi#7813. This lookup table no longer needs to be updated for new providers and is considered stale.
func titleLookup(shortName string) (string, bool) {
	v, ok := map[string]string{
		"aiven":                                "Aiven",
		"akamai":                               "Akamai",
		"alicloud":                             "Alibaba Cloud",
		"auth0":                                "Auth0",
		"aws":                                  "AWS Classic",
		"awsx":                                 "AWSx (Pulumi Crosswalk for AWS)",
		"aws-apigateway":                       "AWS API Gateway",
		"aws-miniflux":                         "Miniflux",
		"aws-native":                           "AWS Native",
		"aws-quickstart-aurora-mysql":          "AWS QuickStart Aurora MySQL",
		"aws-quickstart-aurora-postgres":       "AWS QuickStart Aurora PostgreSQL",
		"aws-quickstart-redshift":              "AWS QuickStart Redshift",
		"aws-serverless":                       "AWS Serverless",
		"aws-quickstart-vpc":                   "AWS QuickStart VPC",
		"aws-s3-replicated-bucket":             "AWS S3 Replicated Bucket",
		"azure":                                "Azure Classic",
		"azure-justrun":                        "Azure Justrun",
		"azure-native":                         "Azure Native",
		"azure-quickstart-acr-geo-replication": "Azure QuickStart ACR Geo Replication",
		"azure-quickstart-aks":                 "Azure QuickStart AKS",
		"azure-quickstart-compute":             "Azure QuickStart Compute",
		"azure-quickstart-sql":                 "Azure QuickStart SQL",
		"azuread":                              "Azure Active Directory (Azure AD)",
		"azuredevops":                          "Azure DevOps",
		"azuresel":                             "Azure",
		"civo":                                 "Civo",
		"cloudamqp":                            "CloudAMQP",
		"cloudflare":                           "Cloudflare",
		"cloudinit":                            "cloud-init",
		"confluentcloud":                       "Confluent Cloud",
		"confluent":                            "Confluent Cloud (Deprecated)",
		"consul":                               "HashiCorp Consul",
		"coredns-helm":                         "CoreDNS (Helm)",
		"datadog":                              "Datadog",
		"digitalocean":                         "DigitalOcean",
		"dnsimple":                             "DNSimple",
		"docker":                               "Docker",
		"docker-buildkit":                      "Docker BuildKit",
		"eks":                                  "Amazon EKS",
		"equinix-metal":                        "Equinix Metal",
		"f5bigip":                              "f5 BIG-IP",
		"fastly":                               "Fastly",
		"gcp":                                  "Google Cloud (GCP) Classic",
		"gcp-global-cloudrun":                  "Google Global Cloud Run",
		"gcp-project-scaffold":                 "Google Project Scaffolding",
		"google-native":                        "Google Cloud Native",
		"github":                               "GitHub",
		"github-serverless-webhook":            "GitHub Serverless Webhook",
		"gitlab":                               "GitLab",
		"hcloud":                               "Hetzner Cloud",
		"istio-helm":                           "Istio (Helm)",
		"jaeger-helm":                          "Jaeger (Helm)",
		"kafka":                                "Kafka",
		"keycloak":                             "Keycloak",
		"kong":                                 "Kong",
		"kubernetes":                           "Kubernetes",
		"libvirt":                              "libvirt",
		"linode":                               "Linode",
		"mailgun":                              "Mailgun",
		"minio":                                "MinIO",
		"mongodbatlas":                         "MongoDB Atlas",
		"mysql":                                "MySQL",
		"newrelic":                             "New Relic",
		"kubernetes-ingress-nginx":             "NGINX Ingress Controller (Helm)",
		"kubernetes-coredns":                   "CoreDNS (Helm)",
		"kubernetes-cert-manager":              "Jetstack Cert Manager (Helm)",
		"nomad":                                "HashiCorp Nomad",
		"ns1":                                  "NS1",
		"okta":                                 "Okta",
		"openstack":                            "OpenStack",
		"opsgenie":                             "Opsgenie",
		"packet":                               "Packet",
		"pagerduty":                            "PagerDuty",
		"pulumi-std":                           "Pulumi Standard Library",
		"postgresql":                           "PostgreSQL",
		"prometheus-helm":                      "Prometheus (Helm)",
		"rabbitmq":                             "RabbitMQ",
		"rancher2":                             "Rancher2",
		"random":                               "random",
		"rke":                                  "Rancher Kubernetes Engine (RKE)",
		"run-my-darn-container":                "Run My Darn Container",
		"shipa":                                "Shipa",
		"signalfx":                             "SignalFx",
		"snowflake":                            "Snowflake",
		"splunk":                               "Splunk",
		"spotinst":                             "Spotinst",
		"sumologic":                            "Sumo Logic",
		"tls":                                  "TLS",
		"vault":                                "Vault",
		"venafi":                               "Venafi",
		"vsphere":                              "vSphere",
		"wavefront":                            "Wavefront",
		"yandex":                               "Yandex",
	}[shortName]
	return v, ok
}

// Property anchor tag separator, used in a property anchor tag id to separate the property and language (e.g.
// property_lang).
const propertyLangSeparator = "_"

type languageConstructorSyntax struct {
	resources map[string]string
	invokes   map[string]string
}

type constructorSyntaxData struct {
	typescript *languageConstructorSyntax
	python     *languageConstructorSyntax
	golang     *languageConstructorSyntax
	csharp     *languageConstructorSyntax
	java       *languageConstructorSyntax
	yaml       *languageConstructorSyntax
}

func emptyLanguageConstructorSyntax() *languageConstructorSyntax {
	return &languageConstructorSyntax{
		resources: map[string]string{},
		invokes:   map[string]string{},
	}
}

func newConstructorSyntaxData() *constructorSyntaxData {
	return &constructorSyntaxData{
		typescript: emptyLanguageConstructorSyntax(),
		python:     emptyLanguageConstructorSyntax(),
		golang:     emptyLanguageConstructorSyntax(),
		csharp:     emptyLanguageConstructorSyntax(),
		java:       emptyLanguageConstructorSyntax(),
		yaml:       emptyLanguageConstructorSyntax(),
	}
}

// Context provides ambient context for docs generation.
type Context struct {
	internalModMap map[string]*modContext

	docHelpers map[language.Language]codegen.DocLanguageHelper

	// The language-specific info objects for a certain package (provider).
	goPkgInfo     go_gen.GoPackageInfo
	csharpPkgInfo dotnet.CSharpPackageInfo
	nodePkgInfo   nodejs.NodePackageInfo
	pythonPkgInfo python.PackageInfo

	// Maps a *modContext, *schema.Resource, or *schema.Function to the link that was assigned to it.
	moduleConflictLinkMap map[interface{}]string

	constructorSyntaxData *constructorSyntaxData

	// Context should not be copied to prevent decentralization between fields, i.e. map mutations will persist between
	// copied but appending to arrays will be witnessed only by a single copy.
	_ noCopy
}

// modules is a map of a module name and information about it. This is crux of all API docs generation as the modContext
// carries information about the resources, functions, as well other modules within each module.
func (dctx *Context) modules() map[string]*modContext {
	return dctx.internalModMap
}

func (dctx *Context) setModules(modules map[string]*modContext) {
	m := map[string]*modContext{}
	for k, v := range modules {
		m[k] = v.withDocGenContext(dctx)
	}
	dctx.internalModMap = m
}

func NewContext(tool string, pkg *schema.Package) *Context {
	dctx := &Context{
		docHelpers: map[language.Language]codegen.DocLanguageHelper{
			language.CSharp: &dotnet.DocLanguageHelper{},
			language.Go:     &go_gen.DocLanguageHelper{},
			language.NodeJS: &nodejs.DocLanguageHelper{},
			language.Python: &python.DocLanguageHelper{},
			language.YAML:   &yaml.DocLanguageHelper{},
			language.Java:   &java.DocLanguageHelper{},
		},
		moduleConflictLinkMap: map[interface{}]string{},
		constructorSyntaxData: generateConstructorSyntaxData(pkg),
	}

	defer glog.Flush()

	// Generate the modules from the schema, and for every module run the generator functions to generate markdown files.
	dctx.setModules(dctx.generateModulesFromSchemaPackage(tool, pkg))

	return dctx
}

type typeDetails struct {
	inputType bool
}

// header represents the header of each resource markdown file.
type header struct {
	Title    string
	TitleTag string
	MetaDesc string
}

// property represents an input or an output property.
type property struct {
	// ID is the `id` attribute that will be attached to the DOM element containing the property.
	ID string
	// DisplayName is the property name with word-breaks.
	DisplayName        string
	Name               string
	Comment            string
	Types              []propertyType
	DeprecationMessage string
	Link               string

	IsRequired         bool
	IsInput            bool
	IsReplaceOnChanges bool
}

// enum represents an enum.
type enum struct {
	// ID is the `id` attribute attached to the DOM element containing the enum.
	ID string
	// DisplayName is the enum name with word-breaks.
	DisplayName string
	// Name is the language-specific enum name.
	Name               string
	Value              string
	Comment            string
	DeprecationMessage string
}

// docNestedType represents a complex type.
type docNestedType struct {
	Name       string
	Input      bool
	AnchorID   string
	Properties map[language.Language][]property
	EnumValues map[language.Language][]enum
}

// propertyType represents the type of a property.
type propertyType struct {
	DisplayName string
	// Name used in description list.
	DescriptionName string
	Name            string
	// Link can be a link to an anchor tag on the same page, or to another page/site.
	Link string
}

// paramSeparator is for data passed to the separator template.
type paramSeparator struct {
	Indent string
}

// formalParam represents the formal parameters of a constructor or a lookup function.
type formalParam struct {
	Name string
	Type propertyType

	// This is the language specific optional type indicator. For example, in nodejs this is the character "?" and in Go
	// it's "*".
	OptionalFlag string

	DefaultValue string

	// Comment is an optional description of the parameter.
	Comment string
}

type packageDetails struct {
	DisplayName    string
	Repository     string
	RepositoryName string
	License        string
	Notes          string
	Version        string
}

type resourceDocArgs struct {
	Header header

	Tool string

	// LangChooserLanguages is a comma-separated list of languages to pass to the language chooser shortcode. Use this to
	// customize the languages shown for a resource. By default, the language chooser will show all languages supported by
	// Pulumi for all resources. Supported values are "typescript", "python", "go", "csharp", "java", "yaml"
	LangChooserLanguages string

	// CreationExampleSyntax is a map from language to the rendered HTML for the creation example syntax where the key is
	// the language name and the value is a piece of code that shows how to create a new instance of the resource with
	// default placeholder values.
	CreationExampleSyntax map[language.Language]string

	// Comment represents the introductory resource comment.
	Comment            string
	ExamplesSection    examplesSection
	DeprecationMessage string

	// Import
	ImportDocs string

	// ConstructorParams is a map from language to the rendered HTML for the constructor's arguments.
	ConstructorParams map[language.Language]renderedConstructorParam
	// ConstructorParamsTyped is the typed set of parameters for the constructor, in order.
	ConstructorParamsTyped map[language.Language]formalConstructParams
	// ConstructorResource is the resource that is being constructed or is the result of a constructor-like function.
	ConstructorResource map[language.Language]propertyType
	// ArgsRequired is a flag indicating if the args param is required when creating a new resource.
	ArgsRequired bool

	// InputProperties is a map per language and a corresponding slice of input properties accepted as args while creating
	// a new resource.
	InputProperties map[language.Language][]property
	// OutputProperties is a map per language and a corresponding slice of output properties returned when a new instance
	// of the resource is created.
	OutputProperties map[language.Language][]property

	// LookupParams is a map of the param string to be rendered per language for looking-up a resource.
	LookupParams map[language.Language]string
	// StateInputs is a map per language and the corresponding slice of state input properties required while looking-up
	// an existing resource.
	StateInputs map[language.Language][]property
	// StateParam is the type name of the state param, if any.
	StateParam string

	// NestedTypes is a slice of the nested types used in the input and output properties.
	NestedTypes []docNestedType
	// Maximum number of nested types to show.
	MaxNestedTypes int

	// A list of methods associated with the resource.
	Methods []methodDocArgs

	PackageDetails packageDetails
}

// typeUsage represents a nested type's usage.
type typeUsage struct {
	Input  bool
	Output bool
}

// nestedTypeUsageInfo is a type-alias for a map of Pulumi type-tokens and whether or not the type is used as an input
// and/or output properties.
type nestedTypeUsageInfo map[string]typeUsage

func (ss nestedTypeUsageInfo) add(s string, input bool) {
	if v, ok := ss[s]; ok {
		if input {
			v.Input = true
		} else {
			v.Output = true
		}
		ss[s] = v
		return
	}

	ss[s] = typeUsage{
		Input:  input,
		Output: !input,
	}
}

// contains returns true if the token already exists and matches the input or output flag of the token.
func (ss nestedTypeUsageInfo) contains(token string, input bool) bool {
	a, ok := ss[token]
	if !ok {
		return false
	}

	if input && a.Input {
		return true
	} else if !input && a.Output {
		return true
	}
	return false
}

type modContext struct {
	pkg         schema.PackageReference
	mod         string
	moduleName  map[language.Language]string
	inputTypes  []*schema.ObjectType
	resources   []*schema.Resource
	functions   []*schema.Function
	typeDetails map[*schema.ObjectType]*typeDetails
	children    []*modContext
	tool        string
	context     *Context
}

func (mod *modContext) withDocGenContext(dctx *Context) *modContext {
	if mod == nil {
		return nil
	}
	newctx := *mod
	newctx.context = dctx
	children := slice.Prealloc[*modContext](len(newctx.children))
	for _, c := range newctx.children {
		children = append(children, c.withDocGenContext(dctx))
	}
	newctx.children = children
	return &newctx
}

// getSupportedLanguages returns the list of supported languages based on the overlay configuration. If `isOverlay` is
// false or `overlaySupportedLanguages` is empty/nil, it returns the default list of supported languages. Otherwise, it
// filters the `overlaySupportedLanguages` to ensure that they are a subset of the default supported languages.
func (dctx *Context) getSupportedLanguages(isOverlay bool, overlaySupportedLanguages []string) []string {
	contextSupportedLanguages := make([]string, 0, 6 /* the number of supported languages */)
	for l := range language.All() {
		contextSupportedLanguages = append(contextSupportedLanguages, l.String())
	}
	if !isOverlay || len(overlaySupportedLanguages) == 0 {
		return contextSupportedLanguages
	}
	var supportedLanguages []string
	allLanguages := codegen.NewStringSet(contextSupportedLanguages...)
	for _, lang := range overlaySupportedLanguages {
		if allLanguages.Has(lang) {
			supportedLanguages = append(supportedLanguages, lang)
		}
	}

	return supportedLanguages
}

// getSupportedSnippetLanguages returns a comma separated string containing the supported snippet languages for the
// given type based on the overlay configuration. If the type is not an overlay or if there are no overlay supported
// languages, all languages are supported. Internally this calls the getSupportedLanguages function to retrieve the
// supported languages and replaces "nodejs" with "typescript" because snippet languages expect "typescript" instead of
// "nodejs".
func (dctx *Context) getSupportedSnippetLanguages(isOverlay bool, overlaySupportedLanguages []string) string {
	supportedLanguages := dctx.getSupportedLanguages(isOverlay, overlaySupportedLanguages)
	supportedSnippetLanguages := make([]string, len(supportedLanguages))
	for idx, lang := range supportedLanguages {
		if lang == "nodejs" {
			lang = "typescript"
		}
		supportedSnippetLanguages[idx] = lang
	}

	return strings.Join(supportedSnippetLanguages, ",")
}

func resourceName(r *schema.Resource) string {
	if r.IsProvider {
		return "Provider"
	}

	//nolint:staticcheck
	return strings.Title(tokenToName(r.Token))
}

func (dctx *Context) getLanguageDocHelper(lang language.Language) codegen.DocLanguageHelper {
	if h, ok := dctx.docHelpers[lang]; ok {
		return h
	}
	panic(fmt.Errorf("could not find a doc lang helper for %s", lang))
}

type propertyCharacteristics struct {
	// input is a flag indicating if the property is an input type.
	input bool
}

func (mod *modContext) details(t *schema.ObjectType) *typeDetails {
	details, ok := mod.typeDetails[t]
	if !ok {
		details = &typeDetails{}
		if mod.typeDetails == nil {
			mod.typeDetails = map[*schema.ObjectType]*typeDetails{}
		}
		mod.typeDetails[t] = details
	}
	return details
}

// getLanguageModuleName transforms the current module's name to a language-specific name using the language info, if
// any, for the current package.
func (mod *modContext) getLanguageModuleName(lang language.Language) string {
	if mod.moduleName == nil {
		dctx := mod.context
		name := mod.mod
		withOverride := func(defaultName string, overrides map[string]string) string {
			if override, ok := overrides[name]; ok {
				return override
			}
			return defaultName
		}

		mod.moduleName = map[language.Language]string{
			language.CSharp: withOverride(name, dctx.csharpPkgInfo.Namespaces),
			// Go module names use lowercase.
			language.Go:     withOverride(strings.ToLower(name), dctx.goPkgInfo.ModuleToPackage),
			language.Java:   name,
			language.NodeJS: withOverride(name, dctx.nodePkgInfo.ModuleToPackage),
			language.Python: withOverride(name, dctx.pythonPkgInfo.ModuleNameOverrides),
			language.YAML:   name,
		}
	}
	return mod.moduleName[lang]
}

// cleanTypeString removes any namespaces from the generated type string for all languages. The result of this function
// should be used display purposes only.
func (mod *modContext) cleanTypeString(
	t schema.Type, langTypeString string, lang language.Language, modName string, isInput bool,
) string {
	switch lang {
	case language.Go:
		langTypeString = cleanOptionalIdentifier(langTypeString, lang)
		parts := strings.Split(langTypeString, ".")
		return parts[len(parts)-1]
	}

	cleanCSharpName := func(pkgName, objModName string) string {
		// C# types can be wrapped in enumerable types such as List<> or Dictionary<>, so we have to only replace the
		// namespace between the < and the > characters.
		qualifier := "Inputs"
		if !isInput {
			qualifier = "Outputs"
		}

		var csharpNS string
		// This type could be at the package-level, so it won't have a module name.
		if objModName != "" {
			csharpNS = fmt.Sprintf("Pulumi.%s.%s.%s.",
				title(pkgName, language.CSharp),
				title(objModName, language.CSharp),
				qualifier)
		} else {
			csharpNS = fmt.Sprintf("Pulumi.%s.%s.", title(pkgName, language.CSharp), qualifier)
		}
		return strings.ReplaceAll(langTypeString, csharpNS, "")
	}

	cleanNodeJSName := func(objModName string) string {
		// The nodejs codegen currently doesn't use the ModuleToPackage override available in the k8s package's schema. So
		// we'll manually strip some known module names for k8s.
		//
		// TODO[pulumi/pulumi#4325]: Remove this block once the nodejs code gen is able to use the package name overrides
		// for modules.
		if isKubernetesPackage(mod.pkg) {
			langTypeString = strings.ReplaceAll(langTypeString, "k8s.io.", "")
			langTypeString = strings.ReplaceAll(langTypeString, "apiserver.", "")
			langTypeString = strings.ReplaceAll(langTypeString, "rbac.authorization.v1.", "")
			langTypeString = strings.ReplaceAll(langTypeString, "rbac.authorization.v1alpha1.", "")
			langTypeString = strings.ReplaceAll(langTypeString, "rbac.authorization.v1beta1.", "")
		}
		objModName = strings.ReplaceAll(objModName, "/", ".") + "."
		return strings.ReplaceAll(langTypeString, objModName, "")
	}

	cleanPythonName := func(objModName string) string {
		// SDK codegen aliases imported top-level inputs as _root_inputs and _root_enums, e.g.
		//
		// from .. import _inputs as _root_inputs
		//
		// This is our implementation detail, so we shouldn't show this prefix in end-user documentation. Remove them here.
		// Top-level arg types will be displayed without any module prefix.
		objModName = strings.Replace(objModName, "root_enums.", "", 1)
		objModName = strings.Replace(objModName, "root_inputs.", "", 1)

		// SDK codegen aliases imported modules with underscored names, e.g.
		//
		// from ... import meta as _meta
		//
		// Therefore, type references for Python all have _ before module names. This is our implementation detail, so we
		// shouldn't show those underscores in end-user documentation. Remove them here. We need to keep underscores inside
		// module names though, as in 'pulumi_random'.
		return removeLeadingUnderscores(objModName)
	}

	switch t := t.(type) {
	case *schema.ObjectType:
		// Strip "Args" suffixes from display names for everything but Python inputs.
		if lang != language.Python || !isInput {
			name := tokenToName(t.Token)
			nameWithArgs := name + "Args"

			// If the langTypeString looks like it's a concatenation of its name and "Args", strip out the "Args".
			if strings.Contains(langTypeString, nameWithArgs) {
				langTypeString = strings.ReplaceAll(langTypeString, nameWithArgs, name)
			}
		}
	}

	switch t := t.(type) {
	case *schema.ArrayType:
		if schema.IsPrimitiveType(t.ElementType) {
			break
		}
		return mod.cleanTypeString(t.ElementType, langTypeString, lang, modName, isInput)
	case *schema.UnionType:
		for _, e := range t.ElementTypes {
			if schema.IsPrimitiveType(e) {
				continue
			}
			return mod.cleanTypeString(e, langTypeString, lang, modName, isInput)
		}
	case *schema.ObjectType:
		objTypeModName := mod.pkg.TokenToModule(t.Token)
		if objTypeModName != mod.mod {
			modName = mod.getLanguageModuleName(lang)
		}
	}

	switch lang {
	case language.NodeJS:
		return cleanNodeJSName(modName)
	case language.CSharp:
		return cleanCSharpName(mod.pkg.Name(), modName)
	case language.Python:
		return cleanPythonName(langTypeString)
	default:
		return strings.ReplaceAll(langTypeString, modName, "")
	}
}

// typeString returns a property type suitable for docs with its display name and the anchor link to a type if the type
// of the property is an array or an object.
func (mod *modContext) typeString(
	t schema.Type,
	lang language.Language,
	characteristics propertyCharacteristics,
	insertWordBreaks bool,
) propertyType {
	t = codegen.PlainType(t)

	docLanguageHelper := mod.context.getLanguageDocHelper(lang)
	modName := mod.getLanguageModuleName(lang)
	langTypeString := docLanguageHelper.GetTypeName(mod.pkg, t, characteristics.input, mod.mod)

	if optional, ok := t.(*schema.OptionalType); ok {
		t = optional.ElementType
	}

	// If the type is an object type, let's also wrap it with a link to the supporting type on the same page using an
	// anchor tag.
	var href string
	switch t := t.(type) {
	case *schema.ArrayType:
		elementLangType := mod.typeString(t.ElementType, lang, characteristics, false)
		href = elementLangType.Link
	case *schema.ObjectType:
		tokenName := tokenToName(t.Token)
		// Links to anchor tags on the same page must be lower-cased.
		href = "#" + strings.ToLower(tokenName)
	case *schema.EnumType:
		tokenName := tokenToName(t.Token)
		// Links to anchor tags on the same page must be lower-cased.
		href = "#" + strings.ToLower(tokenName)
	case *schema.UnionType:
		var elements []string
		for _, e := range t.ElementTypes {
			elementLangType := mod.typeString(e, lang, characteristics, false)
			elements = append(elements, elementLangType.DisplayName)
		}
		langTypeString = strings.Join(elements, " | ")
	}

	// Strip the namespace/module prefix for the type's display name.
	displayName := langTypeString
	if !schema.IsPrimitiveType(t) {
		// TODO[https://github.com/pulumi/registry/issues/7512]: This shouldn't be necessary.
		displayName = mod.cleanTypeString(t, langTypeString, lang, modName, characteristics.input)
	}

	displayName = cleanOptionalIdentifier(displayName, lang)
	langTypeString = cleanOptionalIdentifier(langTypeString, lang)

	// Name and DisplayName should be html-escaped to avoid throwing off rendering for template types in languages like
	// csharp, Java etc. If word-breaks need to be inserted, then the type string should be html-escaped first.
	displayName = html.EscapeString(displayName)
	if insertWordBreaks {
		displayName = wbr(displayName)
	}

	return propertyType{
		Name:        html.EscapeString(langTypeString),
		DisplayName: displayName,
		Link:        href,
	}
}

// cleanOptionalIdentifier removes the type identifier (i.e. "?" in "string?").
func cleanOptionalIdentifier(s string, lang language.Language) string {
	switch lang {
	case language.NodeJS, language.CSharp:
		return strings.TrimSuffix(s, "?")
	case language.Go:
		return strings.TrimPrefix(s, "*")
	case language.Python:
		if strings.HasPrefix(s, "Optional[") && strings.HasSuffix(s, "]") {
			s = strings.TrimPrefix(s, "Optional[")
			s = strings.TrimSuffix(s, "]")
			return s
		}
	}
	return s
}

// Resources typically take the same set of parameters to their constructors, and these are the default
// comments/descriptions for them.
const (
	ctorNameArgComment = "The unique name of the resource."
	ctorArgsArgComment = "The arguments to resource properties."
	ctorOptsArgComment = "Bag of options to control resource's behavior."
)

func (mod *modContext) genConstructorTS(r *schema.Resource, argsOptional bool) []formalParam {
	name := resourceName(r)
	docLangHelper := mod.context.getLanguageDocHelper(language.NodeJS)

	var argsType string

	optsType := "CustomResourceOptions"
	if (isKubernetesPackage(mod.pkg) && mod.isComponentResource()) || r.IsComponent {
		optsType = "ComponentResourceOptions"
	}

	// The args type for k8s package differs from the rest depending on whether we are dealing with overlay resources or
	// regular k8s resources.
	if isKubernetesPackage(mod.pkg) {
		if mod.isKubernetesOverlayModule() {
			if name == "CustomResource" {
				argsType = name + "Args"
			} else {
				argsType = name + "Opts"
			}
		} else {
			// The non-schema-based k8s codegen does not apply a suffix to the input types.
			argsType = name
		}
	} else {
		argsType = name + "Args"
	}

	argsFlag := ""
	if argsOptional {
		argsFlag = "?"
	}

	def, err := mod.pkg.Definition()
	contract.AssertNoErrorf(err, "failed to get definition for package %q", mod.pkg.Name())

	return []formalParam{
		{
			Name: "name",
			Type: propertyType{
				Name: "string",
			},
			Comment: ctorNameArgComment,
		},
		{
			Name:         "args",
			OptionalFlag: argsFlag,
			Type: propertyType{
				Name: argsType,
				Link: "#inputs",
			},
			Comment: ctorArgsArgComment,
		},
		{
			Name:         "opts",
			OptionalFlag: "?",
			Type: propertyType{
				Name: optsType,
				Link: docLangHelper.GetDocLinkForPulumiType(def, optsType),
			},
			Comment: ctorOptsArgComment,
		},
	}
}

func (mod *modContext) genConstructorGo(r *schema.Resource, argsOptional bool) []formalParam {
	name := resourceName(r)
	argsType := name + "Args"
	argsFlag := ""
	if argsOptional {
		argsFlag = "*"
	}

	docLangHelper := mod.context.getLanguageDocHelper(language.Go)

	def, err := mod.pkg.Definition()
	contract.AssertNoErrorf(err, "failed to get definition for package %q", mod.pkg.Name())

	return []formalParam{
		{
			Name:         "ctx",
			OptionalFlag: "*",
			Type: propertyType{
				Name: "Context",
				Link: docLangHelper.GetDocLinkForPulumiType(def, "Context"),
			},
			Comment: "Context object for the current deployment.",
		},
		{
			Name: "name",
			Type: propertyType{
				Name: "string",
			},
			Comment: ctorNameArgComment,
		},
		{
			Name:         "args",
			OptionalFlag: argsFlag,
			Type: propertyType{
				Name: argsType,
				Link: "#inputs",
			},
			Comment: ctorArgsArgComment,
		},
		{
			Name:         "opts",
			OptionalFlag: "...",
			Type: propertyType{
				Name: "ResourceOption",
				Link: docLangHelper.GetDocLinkForPulumiType(def, "ResourceOption"),
			},
			Comment: ctorOptsArgComment,
		},
	}
}

func (mod *modContext) genConstructorCS(r *schema.Resource, argsOptional bool) []formalParam {
	name := resourceName(r)

	optsType := "CustomResourceOptions"
	if (isKubernetesPackage(mod.pkg) && mod.isComponentResource()) || r.IsComponent {
		optsType = "ComponentResourceOptions"
	}

	var argsFlag string
	var argsDefault string
	if argsOptional {
		// If the number of required input properties was zero, we can make the args object optional.
		argsDefault = " = null"
		argsFlag = "?"
	}

	docLangHelper := mod.context.getLanguageDocHelper(language.CSharp)

	def, err := mod.pkg.Definition()
	contract.AssertNoErrorf(err, "failed to get definition for package %q", mod.pkg.Name())

	return []formalParam{
		{
			Name: "name",
			Type: propertyType{
				Name: "string",
			},
			Comment: ctorNameArgComment,
		},
		{
			Name:         "args",
			OptionalFlag: argsFlag,
			DefaultValue: argsDefault,
			Type: propertyType{
				Name: name + "Args",
				Link: "#inputs",
			},
			Comment: ctorArgsArgComment,
		},
		{
			Name:         "opts",
			OptionalFlag: "?",
			DefaultValue: " = null",
			Type: propertyType{
				Name: optsType,
				Link: docLangHelper.GetDocLinkForPulumiType(def, "Pulumi."+optsType),
			},
			Comment: ctorOptsArgComment,
		},
	}
}

func (mod *modContext) genConstructorYaml() []formalParam {
	return []formalParam{
		{
			Name:    "properties",
			Comment: ctorArgsArgComment,
		},
		{
			Name:    "options",
			Comment: ctorOptsArgComment,
		},
	}
}

func (mod *modContext) genConstructorJava(r *schema.Resource, argsOverload bool) []formalParam {
	name := resourceName(r)

	optsType := "CustomResourceOptions"
	if (isKubernetesPackage(mod.pkg) && mod.isComponentResource()) || r.IsComponent {
		optsType = "ComponentResourceOptions"
	}

	docLangHelper := mod.context.getLanguageDocHelper(language.Java)

	def, err := mod.pkg.Definition()
	contract.AssertNoErrorf(err, "failed to get definition for package %q", mod.pkg.Name())

	result := []formalParam{
		{
			Name: "name",
			Type: propertyType{
				Name: "String",
			},
			Comment: ctorNameArgComment,
		},
		{
			Name: "args",
			Type: propertyType{
				Name: name + "Args",
				Link: "#inputs",
			},
			Comment: ctorArgsArgComment,
		},
	}
	if !argsOverload {
		result = append(result, formalParam{
			Name:         "options",
			OptionalFlag: "@Nullable",
			Type: propertyType{
				Name: optsType,
				Link: docLangHelper.GetDocLinkForPulumiType(def, optsType),
			},
			Comment: ctorOptsArgComment,
		})
	}
	return result
}

func (mod *modContext) genConstructorPython(r *schema.Resource, argsOptional, argsOverload bool) []formalParam {
	docLanguageHelper := mod.context.getLanguageDocHelper(language.Python)
	isK8sOverlayMod := mod.isKubernetesOverlayModule()
	isDockerImageResource := mod.pkg.Name() == "docker" && resourceName(r) == "Image"

	// Kubernetes overlay resources use a different ordering of formal params in Python.
	if isK8sOverlayMod && r.IsOverlay {
		return getKubernetesOverlayPythonFormalParams(mod.mod)
	} else if isDockerImageResource {
		return getDockerImagePythonFormalParams()
	}

	// We perform at least three appends before iterating over input types.
	params := slice.Prealloc[formalParam](3 + len(r.InputProperties))

	params = append(params, formalParam{
		Name: "resource_name",
		Type: propertyType{
			Name: "str",
		},
		Comment: ctorNameArgComment,
	})

	if argsOverload {
		// Determine whether we need to use the alternate args class name (e.g. `<Resource>InitArgs` instead of
		// `<Resource>Args`) due to an input type with the same name as the resource in the same module.
		resName := resourceName(r)
		resArgsName := resName + "Args"
		for _, inputType := range mod.inputTypes {
			//nolint:staticcheck
			inputTypeName := strings.Title(tokenToName(inputType.Token))
			if resName == inputTypeName {
				resArgsName = resName + "InitArgs"
			}
		}

		optionalFlag, defaultVal, descriptionName := "", "", resArgsName
		typeName := descriptionName
		if argsOptional {
			optionalFlag, defaultVal, typeName = "optional", " = None", fmt.Sprintf("Optional[%s]", typeName)
		}
		params = append(params, formalParam{
			Name:         "args",
			OptionalFlag: optionalFlag,
			DefaultValue: defaultVal,
			Type: propertyType{
				Name:            typeName,
				DescriptionName: descriptionName,
				Link:            "#inputs",
			},
			Comment: ctorArgsArgComment,
		})
	}

	params = append(params, formalParam{
		Name:         "opts",
		OptionalFlag: "optional",
		DefaultValue: " = None",
		Type: propertyType{
			Name:            "Optional[ResourceOptions]",
			DescriptionName: "ResourceOptions",
			Link:            "/docs/reference/pkg/python/pulumi/#pulumi.ResourceOptions",
		},
		Comment: ctorOptsArgComment,
	})

	if argsOverload {
		return params
	}

	for _, p := range r.InputProperties {
		// If the property defines a const value, then skip it. For example, in k8s, `apiVersion` and `kind` are often
		// hard-coded in the SDK and are not really user-provided input properties.
		if p.ConstValue != nil {
			continue
		}
		typ := docLanguageHelper.GetTypeName(
			mod.pkg,
			codegen.PlainType(codegen.OptionalType(p)),
			true, /*input*/
			mod.mod,
		)
		params = append(params, formalParam{
			Name:         python.InitParamName(p.Name),
			DefaultValue: " = None",
			Type: propertyType{
				Name: typ,
			},
		})
	}
	return params
}

func (mod *modContext) genNestedTypes(member interface{}, resourceType, isProvider bool) []docNestedType {
	dctx := mod.context
	tokens := nestedTypeUsageInfo{}
	// Collect all of the types for this "member" as a map of resource names and if it appears in an input object and/or
	// output object.
	mod.getTypes(member, tokens)

	sortedTokens := slice.Prealloc[string](len(tokens))
	for token := range tokens {
		sortedTokens = append(sortedTokens, token)
	}
	sort.Strings(sortedTokens)

	var typs []docNestedType
	for _, token := range sortedTokens {
		for iter := mod.pkg.Types().Range(); iter.Next(); {
			t, err := iter.Type()
			contract.AssertNoErrorf(err, "error iterating types")
			switch typ := t.(type) {
			case *schema.ObjectType:
				if typ.Token != token || len(typ.Properties) == 0 || typ.IsInputShape() {
					continue
				}

				// Create a map to hold the per-language properties of this object.
				props := make(map[language.Language][]property)
				for lang := range language.All() {
					props[lang] = mod.getProperties(typ.Properties, lang, true, true, isProvider)
				}

				//nolint:staticcheck
				name := strings.Title(tokenToName(typ.Token))
				typs = append(typs, docNestedType{
					Name:       wbr(name),
					AnchorID:   strings.ToLower(name),
					Properties: props,
				})
			case *schema.EnumType:
				if typ.Token != token || len(typ.Elements) == 0 {
					continue
				}
				//nolint:staticcheck
				name := strings.Title(tokenToName(typ.Token))

				enums := make(map[language.Language][]enum)
				for lang := range language.All() {
					docLangHelper := dctx.getLanguageDocHelper(lang)

					var langEnumValues []enum
					for _, e := range typ.Elements {
						enumName, err := docLangHelper.GetEnumName(e, name)
						if err != nil {
							panic(err)
						}
						enumID := strings.ToLower(name + propertyLangSeparator + lang.String())
						langEnumValues = append(langEnumValues, enum{
							ID:                 enumID,
							DisplayName:        wbr(enumName),
							Name:               enumName,
							Value:              fmt.Sprintf("%v", e.Value),
							Comment:            e.Comment,
							DeprecationMessage: e.DeprecationMessage,
						})
					}
					enums[lang] = langEnumValues
				}

				typs = append(typs, docNestedType{
					Name:       wbr(name),
					AnchorID:   strings.ToLower(name),
					EnumValues: enums,
				})
			}
		}
	}

	sort.Slice(typs, func(i, j int) bool {
		return typs[i].Name < typs[j].Name
	})

	return typs
}

// getProperties returns a slice of properties that can be rendered for docs for the provided slice of properties in the
// schema.
func (mod *modContext) getProperties(
	properties []*schema.Property, lang language.Language, input, nested, isProvider bool,
) []property {
	return mod.getPropertiesWithIDPrefixAndExclude(properties, lang, input, nested, isProvider, "", nil)
}

func (mod *modContext) getPropertiesWithIDPrefixAndExclude(
	properties []*schema.Property, lang language.Language, input, nested, isProvider bool,
	idPrefix string, exclude func(name string) bool,
) []property {
	dctx := mod.context
	if len(properties) == 0 {
		return nil
	}
	docProperties := slice.Prealloc[property](len(properties))
	for _, prop := range properties {
		if prop == nil {
			continue
		}

		if exclude != nil && exclude(prop.Name) {
			continue
		}

		// If the property has a const value, then don't show it as an input property. Even though it is a valid property,
		// it is used by the language code gen to generate the appropriate defaults for it. These cannot be overridden by
		// users.
		if prop.ConstValue != nil {
			continue
		}

		characteristics := propertyCharacteristics{input: input}

		langDocHelper := dctx.getLanguageDocHelper(lang)
		name, err := langDocHelper.GetPropertyName(prop)
		if err != nil {
			panic(err)
		}
		propLangName := name

		propID := idPrefix + strings.ToLower(propLangName+propertyLangSeparator+lang.String())

		propTypes := make([]propertyType, 0)
		if typ, isUnion := codegen.UnwrapType(prop.Type).(*schema.UnionType); isUnion {
			for _, elementType := range typ.ElementTypes {
				propTypes = append(propTypes, mod.typeString(elementType, lang, characteristics, true))
			}
		} else {
			propTypes = append(propTypes, mod.typeString(prop.Type, lang, characteristics, true))
		}

		comment := prop.Comment
		link := "#" + propID

		// Check if type is defined in a package external to the current package. If it is external, update comment to
		// indicate to user that type is defined in another package and link there.
		if isExternalType(codegen.UnwrapType(prop.Type), mod.pkg) {
			packageName := tokenToPackageName(fmt.Sprintf("%v", codegen.UnwrapType(prop.Type)))
			extPkgLink := "/registry/packages/" + packageName
			comment += fmt.Sprintf(
				"\nThis type is defined in the [%s](%s) package.",
				getPackageDisplayName(packageName),
				extPkgLink,
			)
		}

		// Default values for Provider inputs correspond to environment variables, so add that info to the docs.
		if isProvider && input && prop.DefaultValue != nil && len(prop.DefaultValue.Environment) > 0 {
			var suffix string
			if len(prop.DefaultValue.Environment) > 1 {
				suffix = "s"
			}
			comment += fmt.Sprintf(" It can also be sourced from the following environment variable%s: ", suffix)
			for i, v := range prop.DefaultValue.Environment {
				comment += fmt.Sprintf("`%s`", v)
				if i != len(prop.DefaultValue.Environment)-1 {
					comment += ", "
				}
			}
		}

		docProperties = append(docProperties, property{
			ID:                 propID,
			DisplayName:        wbr(propLangName),
			Name:               propLangName,
			Comment:            comment,
			DeprecationMessage: prop.DeprecationMessage,
			IsRequired:         prop.IsRequired(),
			IsInput:            input,
			// We indicate that a property will replace if either
			// a) we will force the replace at the engine level
			// b) we are told that the provider will require a replace
			IsReplaceOnChanges: prop.ReplaceOnChanges || prop.WillReplaceOnChanges,
			Link:               link,
			Types:              propTypes,
		})
	}

	// Sort required props to move them to the top of the properties list, then by name.
	sort.SliceStable(docProperties, func(i, j int) bool {
		pi, pj := docProperties[i], docProperties[j]
		switch {
		case pi.IsRequired != pj.IsRequired:
			return pi.IsRequired && !pj.IsRequired
		default:
			return pi.Name < pj.Name
		}
	})

	return docProperties
}

func getDockerImagePythonFormalParams() []formalParam {
	return []formalParam{
		{
			Name: "image_name",
		},
		{
			Name: "build",
		},
		{
			Name:         "local_image_name",
			DefaultValue: "=None",
		},
		{
			Name:         "registry",
			DefaultValue: "=None",
		},
		{
			Name:         "skip_push",
			DefaultValue: "=None",
		},
		{
			Name:         "opts",
			DefaultValue: "=None",
		},
	}
}

type renderedConstructorParam struct {
	// The constructor parameter, rendered to HTML.
	Param string
	// An optional arg parameter.
	ArgParam string
}

type formalConstructParams struct {
	Params []formalParam

	ArgParams []formalParam
}

// Returns the rendered HTML for the resource's constructor, as well as the specific arguments.
func (mod *modContext) genConstructors(
	r *schema.Resource,
	allOptionalInputs bool,
) (map[language.Language]renderedConstructorParam, map[language.Language]formalConstructParams) {
	renderedParams := make(map[language.Language]renderedConstructorParam)
	formalParams := make(map[language.Language]formalConstructParams)

	pyIdent := strings.Repeat(" ", len("def (")+len(resourceName(r)))

	for lang := range language.All() {
		switch lang {
		case language.NodeJS:
			params := mod.genConstructorTS(r, allOptionalInputs)
			formalParams[lang] = formalConstructParams{Params: params}
			renderedParams[lang] = renderedConstructorParam{Param: renderParams(lang, params, pyIdent)}
		case language.Go:
			params := mod.genConstructorGo(r, allOptionalInputs)
			formalParams[lang] = formalConstructParams{Params: params}
			renderedParams[lang] = renderedConstructorParam{Param: renderParams(lang, params, pyIdent)}
		case language.CSharp:
			params := mod.genConstructorCS(r, allOptionalInputs)
			formalParams[lang] = formalConstructParams{Params: params}
			renderedParams[lang] = renderedConstructorParam{Param: renderParams(lang, params, pyIdent)}
		case language.Java:
			params := mod.genConstructorJava(r, false)
			argsOverloadParams := mod.genConstructorJava(r, true)
			formalParams[lang] = formalConstructParams{
				Params:    params,
				ArgParams: argsOverloadParams,
			}
			renderedParams[lang] = renderedConstructorParam{
				Param:    renderParams(lang, params, pyIdent),
				ArgParam: renderParams(lang, argsOverloadParams, pyIdent),
			}
		case language.Python:
			params := mod.genConstructorPython(r, allOptionalInputs, false)
			argsOverloadParams := mod.genConstructorPython(r, allOptionalInputs, true)
			formalParams[lang] = formalConstructParams{
				Params:    params,
				ArgParams: argsOverloadParams,
			}
			renderedParams[lang] = renderedConstructorParam{
				Param:    renderParams(lang, params, pyIdent),
				ArgParam: renderParams(lang, argsOverloadParams, pyIdent),
			}
		case language.YAML:
			formalParams[lang] = formalConstructParams{Params: mod.genConstructorYaml()}
			renderedParams[lang] = renderedConstructorParam{}
		}
	}

	return renderedParams, formalParams
}

// getConstructorResourceInfo returns a map of per-language information about the resource being constructed.
func (mod *modContext) getConstructorResourceInfo(resourceTypeName, tok string) map[language.Language]propertyType {
	dctx := mod.context
	docLangHelper := dctx.getLanguageDocHelper(language.YAML)
	resourceMap := make(map[language.Language]propertyType)
	resourceDisplayName := resourceTypeName

	for lang := range language.All() {
		// Use the module to package lookup to transform the module name to its normalized package name.
		modName := mod.getLanguageModuleName(lang)
		// Reset the type name back to the display name.
		resourceTypeName = resourceDisplayName

		switch lang {
		case language.NodeJS, language.Go, language.Python, language.Java:
			// Intentionally left blank.
		case language.CSharp:
			namespace := title(mod.pkg.Name(), language.CSharp)
			if ns, ok := dctx.csharpPkgInfo.Namespaces[mod.pkg.Name()]; ok {
				namespace = ns
			}
			if mod.mod == "" {
				resourceTypeName = fmt.Sprintf("Pulumi.%s.%s", namespace, resourceTypeName)
				break
			}

			resourceTypeName = fmt.Sprintf("Pulumi.%s.%s.%s", namespace, modName, resourceTypeName)
		case language.YAML:
			resourceMap[lang] = propertyType{
				Name:        resourceTypeName,
				DisplayName: docLangHelper.GetTypeName(mod.pkg, &schema.ResourceType{Token: tok}, false, mod.mod),
			}
			continue
		default:
			panic(fmt.Errorf("cannot generate constructor info for unhandled language %q", lang))
		}

		parts := strings.Split(resourceTypeName, ".")
		displayName := parts[len(parts)-1]

		resourceMap[lang] = propertyType{
			Name:        resourceDisplayName,
			DisplayName: displayName,
		}
	}

	return resourceMap
}

func (mod *modContext) getTSLookupParams(r *schema.Resource, stateParam string) []formalParam {
	dctx := mod.context
	docLangHelper := dctx.getLanguageDocHelper(language.NodeJS)
	def, err := mod.pkg.Definition()
	contract.AssertNoErrorf(err, "failed to get definition for package %q", mod.pkg.Name())

	return []formalParam{
		{
			Name: "name",

			Type: propertyType{
				Name: "string",
			},
		},
		{
			Name: "id",
			Type: propertyType{
				Name: "Input<ID>",
				Link: docLangHelper.GetDocLinkForPulumiType(def, "ID"),
			},
		},
		{
			Name:         "state",
			OptionalFlag: "?",
			Type: propertyType{
				Name: stateParam,
			},
		},
		{
			Name:         "opts",
			OptionalFlag: "?",
			Type: propertyType{
				Name: "CustomResourceOptions",
				Link: docLangHelper.GetDocLinkForPulumiType(def, "CustomResourceOptions"),
			},
		},
	}
}

func (mod *modContext) getGoLookupParams(r *schema.Resource, stateParam string) []formalParam {
	dctx := mod.context
	docLangHelper := dctx.getLanguageDocHelper(language.Go)

	def, err := mod.pkg.Definition()
	contract.AssertNoErrorf(err, "failed to get definition for package %q", mod.pkg.Name())

	return []formalParam{
		{
			Name:         "ctx",
			OptionalFlag: "*",
			Type: propertyType{
				Name: "Context",
				Link: docLangHelper.GetDocLinkForPulumiType(def, "Context"),
			},
		},
		{
			Name: "name",
			Type: propertyType{
				Name: "string",
			},
		},
		{
			Name: "id",
			Type: propertyType{
				Name: "IDInput",
				Link: docLangHelper.GetDocLinkForPulumiType(def, "IDInput"),
			},
		},
		{
			Name:         "state",
			OptionalFlag: "*",
			Type: propertyType{
				Name: stateParam,
			},
		},
		{
			Name:         "opts",
			OptionalFlag: "...",
			Type: propertyType{
				Name: "ResourceOption",
				Link: docLangHelper.GetDocLinkForPulumiType(def, "ResourceOption"),
			},
		},
	}
}

func (mod *modContext) getCSLookupParams(r *schema.Resource, stateParam string) []formalParam {
	dctx := mod.context
	docLangHelper := dctx.getLanguageDocHelper(language.CSharp)

	def, err := mod.pkg.Definition()
	contract.AssertNoErrorf(err, "failed to get definition for package %q", mod.pkg.Name())

	return []formalParam{
		{
			Name: "name",
			Type: propertyType{
				Name: "string",
			},
		},
		{
			Name: "id",
			Type: propertyType{
				Name: "Input<string>",
				Link: docLangHelper.GetDocLinkForPulumiType(def, "Pulumi.Input"),
			},
		},
		{
			Name:         "state",
			OptionalFlag: "?",
			Type: propertyType{
				Name: stateParam,
			},
		},
		{
			Name:         "opts",
			OptionalFlag: "?",
			DefaultValue: " = null",
			Type: propertyType{
				Name: "CustomResourceOptions",
				Link: docLangHelper.GetDocLinkForPulumiType(def, "Pulumi.CustomResourceOptions"),
			},
		},
	}
}

func (mod *modContext) getJavaLookupParams(r *schema.Resource, stateParam string) []formalParam {
	dctx := mod.context
	docLangHelper := dctx.getLanguageDocHelper(language.Java)
	def, err := mod.pkg.Definition()
	contract.AssertNoErrorf(err, "failed to get definition for package %q", mod.pkg.Name())

	return []formalParam{
		{
			Name: "name",
			Type: propertyType{
				Name: "String",
			},
		},
		{
			Name: "id",
			Type: propertyType{
				Name: "Output<String>",
				Link: docLangHelper.GetDocLinkForPulumiType(def, "Output"),
			},
		},
		{
			Name: "state",
			Type: propertyType{
				Name: stateParam,
			},
		},
		{
			Name: "options",
			Type: propertyType{
				Name: "CustomResourceOptions",
				Link: docLangHelper.GetDocLinkForPulumiType(def, "CustomResourceOptions"),
			},
		},
	}
}

func (mod *modContext) getPythonLookupParams(r *schema.Resource, stateParam string) []formalParam {
	dctx := mod.context
	// The input properties for a resource needs to be exploded as individual constructor params.
	docLanguageHelper := dctx.getLanguageDocHelper(language.Python)
	params := slice.Prealloc[formalParam](len(r.StateInputs.Properties))
	for _, p := range r.StateInputs.Properties {
		typ := docLanguageHelper.GetTypeName(
			mod.pkg,
			codegen.PlainType(codegen.OptionalType(p)),
			true, /*input*/
			mod.mod,
		)
		params = append(params, formalParam{
			Name:         python.PyName(p.Name),
			DefaultValue: " = None",
			Type: propertyType{
				Name: typ,
			},
		})
	}
	return params
}

// genLookupParams generates a map of per-language way of rendering the formal parameters of the lookup function used to
// lookup an existing resource.
func (mod *modContext) genLookupParams(r *schema.Resource, stateParam string) map[language.Language]string {
	lookupParams := make(map[language.Language]string)
	if r.StateInputs == nil {
		return lookupParams
	}

	pyIndent := strings.Repeat(" ", len("def get("))

	for lang := range language.All() {
		params := mod.getLookupParams(lang, r, stateParam)
		lookupParams[lang] = renderParams(lang, params, pyIndent)
	}
	return lookupParams
}

func (mod *modContext) getLookupParams(lang language.Language, r *schema.Resource, stateParam string) []formalParam {
	switch lang {
	case language.NodeJS:
		return mod.getTSLookupParams(r, stateParam)
	case language.Go:
		return mod.getGoLookupParams(r, stateParam)
	case language.CSharp:
		return mod.getCSLookupParams(r, stateParam)
	case language.Java:
		return mod.getJavaLookupParams(r, stateParam)
	case language.Python:
		return mod.getPythonLookupParams(r, stateParam)
	case language.YAML, language.Java:
		return nil // We don't render lookup parameters for Java and YAML
	default:
		glog.Fatalf("Unknown language %#v", lang)
		return nil
	}
}

// filterOutputProperties removes the input properties from the output properties list (since input props are implicitly
// output props), returning only "output" props.
func filterOutputProperties(inputProps []*schema.Property, props []*schema.Property) []*schema.Property {
	var outputProps []*schema.Property
	inputMap := make(map[string]bool, len(inputProps))
	for _, p := range inputProps {
		inputMap[p.Name] = true
	}
	for _, p := range props {
		if _, found := inputMap[p.Name]; !found {
			outputProps = append(outputProps, p)
		}
	}
	return outputProps
}

func (mod *modContext) genResourceHeader(r *schema.Resource) header {
	resourceName := resourceName(r)
	var metaDescription string
	var titleTag string
	if mod.mod == "" {
		metaDescription = fmt.Sprintf("Documentation for the %s.%s resource "+
			"with examples, input properties, output properties, "+
			"lookup functions, and supporting types.", mod.pkg.Name(), resourceName)
		titleTag = fmt.Sprintf("%s.%s", mod.pkg.Name(), resourceName)
	} else {
		metaDescription = fmt.Sprintf("Documentation for the %s.%s.%s resource "+
			"with examples, input properties, output properties, "+
			"lookup functions, and supporting types.", mod.pkg.Name(), mod.mod, resourceName)
		titleTag = fmt.Sprintf("%s.%s.%s", mod.pkg.Name(), mod.mod, resourceName)
	}

	return header{
		Title:    resourceName,
		TitleTag: titleTag,
		MetaDesc: metaDescription,
	}
}

// genResource is the entrypoint for generating a doc for a resource from its Pulumi schema.
func (mod *modContext) genResource(r *schema.Resource) resourceDocArgs {
	dctx := mod.context
	// Create a resource module file into which all of this resource's types will go.
	name := resourceName(r)

	inputProps := make(map[language.Language][]property)
	outputProps := make(map[language.Language][]property)
	stateInputs := make(map[language.Language][]property)

	var filteredOutputProps []*schema.Property
	// Provider resources do not have output properties, so there won't be anything to filter.
	if !r.IsProvider {
		filteredOutputProps = filterOutputProperties(r.InputProperties, r.Properties)
	}

	// All custom resources have an implicit `id` output property, that we must inject into the docs.
	if !r.IsComponent {
		filteredOutputProps = append(filteredOutputProps, &schema.Property{
			Name:    "id",
			Comment: "The provider-assigned unique ID for this managed resource.",
			Type:    schema.StringType,
		})
	}

	for lang := range language.All() {
		inputProps[lang] = mod.getProperties(r.InputProperties, lang, true, false, r.IsProvider)
		outputProps[lang] = mod.getProperties(filteredOutputProps, lang, false, false, r.IsProvider)
		if r.IsProvider {
			continue
		}
		if r.StateInputs != nil {
			stateProps := mod.getProperties(r.StateInputs.Properties, lang, true, false, r.IsProvider)
			for i := 0; i < len(stateProps); i++ {
				id := "state_" + stateProps[i].ID
				stateProps[i].ID = id
				stateProps[i].Link = "#" + id
			}
			stateInputs[lang] = stateProps
		}
	}

	allOptionalInputs := true
	for _, prop := range r.InputProperties {
		// If at least one prop is required, then break.
		if prop.IsRequired() {
			allOptionalInputs = false
			break
		}
	}

	def, err := mod.pkg.Definition()
	contract.AssertNoErrorf(err, "failed to get definition for package %s", mod.pkg.Name())
	packageDetails := packageDetails{
		DisplayName:    getPackageDisplayName(def.Name),
		Repository:     def.Repository,
		RepositoryName: getRepositoryName(def.Repository),
		License:        def.License,
		Notes:          def.Attribution,
	}

	renderedCtorParams, typedCtorParams := mod.genConstructors(r, allOptionalInputs)
	stateParam := name + "State"
	creationExampleSyntax := map[language.Language]string{}
	if !r.IsProvider && err == nil {
		if example, found := dctx.constructorSyntaxData.typescript.resources[r.Token]; found {
			if strings.Contains(example, "notImplemented") || strings.Contains(example, "PANIC") {
				example = ""
			}
			creationExampleSyntax[language.NodeJS] = example
		}
		if example, found := dctx.constructorSyntaxData.python.resources[r.Token]; found {
			if strings.Contains(example, "not_implemented") || strings.Contains(example, "PANIC") {
				example = ""
			}
			creationExampleSyntax[language.Python] = example
		}
		if example, found := dctx.constructorSyntaxData.csharp.resources[r.Token]; found {
			if strings.Contains(example, "NotImplemented") || strings.Contains(example, "PANIC") {
				example = ""
			}
			creationExampleSyntax[language.CSharp] = example
		}
		if example, found := dctx.constructorSyntaxData.golang.resources[r.Token]; found {
			modified := strings.ReplaceAll(example, "_, err =", "example, err :=")
			modified = strings.ReplaceAll(modified, "_, err :=", "example, err :=")
			if strings.Contains(modified, "notImplemented") || strings.Contains(modified, "PANIC") {
				modified = ""
			}
			creationExampleSyntax[language.Go] = modified
		}
		if example, found := dctx.constructorSyntaxData.java.resources[r.Token]; found {
			creationExampleSyntax[language.Java] = example
		}
		if example, found := dctx.constructorSyntaxData.yaml.resources[collapseYAMLToken(r.Token)]; found {
			creationExampleSyntax[language.YAML] = example
		}
	}

	// Set a cap on how many auxiliary types the docs will include. We do that to limit the maximum size of a doc page.
	// The default limit is rather high - we tried setting it to 200 for everyone but got pushback from third-party
	// consumers who have legitimate use cases with more that 200 types. Therefore, we currently apply a smaller limit of
	// 200 to packages that we know have some bloat in their types (AWS and AWS Native). See
	// https://github.com/pulumi/pulumi/issues/15507#issuecomment-2064361317
	//
	// Schema Tools will print a warning every time a new resources gets an increase in number of types that is beyond
	// 200. This should help us catch new instances and make decisions whether to include all types or add a limit. The
	// default is including all types up to 1000.
	maxNestedTypes := 1000
	switch mod.pkg.Name() {
	case "aws", "aws-native":
		maxNestedTypes = 200
	}

	supportedSnippetLanguages := mod.context.getSupportedSnippetLanguages(r.IsOverlay, r.OverlaySupportedLanguages)
	docInfo := dctx.decomposeDocstring(r.Comment, supportedSnippetLanguages)
	data := resourceDocArgs{
		Header: mod.genResourceHeader(r),

		Tool: mod.tool,

		Comment:            docInfo.description,
		DeprecationMessage: r.DeprecationMessage,
		ExamplesSection: examplesSection{
			Examples:             docInfo.examples,
			LangChooserLanguages: supportedSnippetLanguages,
		},
		ImportDocs:             docInfo.importDetails,
		CreationExampleSyntax:  creationExampleSyntax,
		ConstructorParams:      renderedCtorParams,
		ConstructorParamsTyped: typedCtorParams,

		ConstructorResource: mod.getConstructorResourceInfo(name, r.Token),
		ArgsRequired:        !allOptionalInputs,

		InputProperties:  inputProps,
		OutputProperties: outputProps,
		LookupParams:     mod.genLookupParams(r, stateParam),
		StateInputs:      stateInputs,
		StateParam:       stateParam,
		NestedTypes:      mod.genNestedTypes(r, true /*resourceType*/, r.IsProvider),
		MaxNestedTypes:   maxNestedTypes,

		Methods: mod.genMethods(r),

		PackageDetails:       packageDetails,
		LangChooserLanguages: supportedSnippetLanguages,
	}

	return data
}

func (mod *modContext) getNestedTypes(t schema.Type, types nestedTypeUsageInfo, input bool) {
	switch t := t.(type) {
	case *schema.InputType:
		mod.getNestedTypes(t.ElementType, types, input)
	case *schema.OptionalType:
		mod.getNestedTypes(t.ElementType, types, input)
	case *schema.ArrayType:
		mod.getNestedTypes(t.ElementType, types, input)
	case *schema.MapType:
		mod.getNestedTypes(t.ElementType, types, input)
	case *schema.ObjectType:
		if types.contains(t.Token, input) {
			break
		}

		types.add(t.Token, input)
		for _, p := range t.Properties {
			mod.getNestedTypes(p.Type, types, input)
		}
	case *schema.EnumType:
		types.add(t.Token, false)
	case *schema.UnionType:
		for _, e := range t.ElementTypes {
			mod.getNestedTypes(e, types, input)
		}
	}
}

func (mod *modContext) getTypes(member interface{}, types nestedTypeUsageInfo) {
	glog.V(3).Infoln("getting nested types for module", mod.mod)

	switch t := member.(type) {
	case *schema.ObjectType:
		for _, p := range t.Properties {
			mod.getNestedTypes(p.Type, types, false)
		}
	case *schema.Resource:
		for _, p := range t.Properties {
			mod.getNestedTypes(p.Type, types, false)
		}
		for _, p := range t.InputProperties {
			mod.getNestedTypes(p.Type, types, true)
		}
		for _, m := range t.Methods {
			mod.getTypes(m.Function, types)
		}
	case *schema.Function:
		if t.Inputs != nil && !t.MultiArgumentInputs {
			mod.getNestedTypes(t.Inputs, types, true)
		}

		if t.ReturnType != nil {
			if objectType, ok := t.ReturnType.(*schema.ObjectType); ok && objectType != nil {
				mod.getNestedTypes(objectType, types, false)
			}
		}
	}
}

// getModuleFileName returns the file name to use for a module.
func (mod *modContext) getModuleFileName() string {
	dctx := mod.context
	if !isKubernetesPackage(mod.pkg) {
		return mod.mod
	}

	// For k8s packages, use the Go-language info to get the file name for the module.
	if override, ok := dctx.goPkgInfo.ModuleToPackage[mod.mod]; ok {
		return override
	}
	return mod.mod
}

// moduleConflictResolver holds module-level information for resolving naming conflicts. It shares information with the
// top-level Context to ensure the same name is used across modules that reference each other.
type moduleConflictResolver struct {
	dctx *Context
	seen map[string]struct{}
}

func (dctx *Context) newModuleConflictResolver() moduleConflictResolver {
	return moduleConflictResolver{
		dctx: dctx,
		seen: map[string]struct{}{},
	}
}

// getSafeName returns a documentation name for an item that is unique within the module. If the item has already been
// resolved by any module, the previously-resolved name is returned.
func (r *moduleConflictResolver) getSafeName(name string, item interface{}) string {
	if safeName, ok := r.dctx.moduleConflictLinkMap[item]; ok {
		return safeName
	}

	var prefixes []string
	switch item.(type) {
	case *schema.Resource:
		prefixes = []string{"", "res-"}
	case *schema.Function:
		prefixes = []string{"", "fn-"}
	case *modContext:
		prefixes = []string{"", "mod-"}
	default:
		prefixes = []string{""}
	}
	for _, prefix := range prefixes {
		candidate := prefix + name
		if _, exists := r.seen[candidate]; exists {
			continue
		}
		r.seen[candidate] = struct{}{}
		r.dctx.moduleConflictLinkMap[item] = candidate
		return candidate
	}

	glog.Error("skipping unresolvable duplicate file name: ", name)
	return ""
}

func (mod *modContext) gen(fs codegen.Fs) error {
	glog.V(4).Infoln("genIndex for", mod.mod)

	modName := mod.getModuleFileName()
	conflictResolver := mod.context.newModuleConflictResolver()

	def, err := mod.pkg.Definition()
	contract.AssertNoErrorf(err, "failed to get definition for package %q", mod.pkg.Name())

	modTitle := modName
	if modTitle == "" {
		// An empty string indicates that this is the root module.
		if def.DisplayName != "" {
			modTitle = def.DisplayName
		} else {
			modTitle = getPackageDisplayName(mod.pkg.Name())
		}
	}

	// addFileTemplated executes template tmpl with data, and adds a file $dirName/_index.md with the result.
	addFileTemplated := func(dirName string, tmpl templates.Template, data interface{}) error {
		var buff bytes.Buffer
		if err := tmpl(&buff, data); err != nil {
			return err
		}
		p := path.Join(modName, dirName, "_index.md")
		fs.Add(p, buff.Bytes())
		return nil
	}

	// If there are submodules, list them.
	modules := slice.Prealloc[indexEntry](len(mod.children))
	for _, mod := range mod.children {
		modName := mod.getModuleFileName()
		displayName := modFilenameToDisplayName(modName)
		safeName := conflictResolver.getSafeName(displayName, mod)
		if safeName == "" {
			continue // unresolved conflict
		}
		modules = append(modules, indexEntry{
			Link:        getModuleLink(safeName),
			DisplayName: displayName,
		})
	}
	sortIndexEntries(modules)

	// If there are resources in the root, list them.
	resources := slice.Prealloc[indexEntry](len(mod.resources))
	for _, r := range mod.resources {
		title := resourceName(r)
		link := getResourceLink(title)
		link = conflictResolver.getSafeName(link, r)
		if link == "" {
			continue // unresolved conflict
		}

		data := mod.genResource(r)
		if err := addFileTemplated(link, templates.Resource, data); err != nil {
			return err
		}

		resources = append(resources, indexEntry{
			Link:        link + "/",
			DisplayName: title,
		})
	}
	sortIndexEntries(resources)

	// If there are functions in the root, list them.
	functions := slice.Prealloc[indexEntry](len(mod.functions))
	for _, f := range mod.functions {
		name := tokenToName(f.Token)
		link := getFunctionLink(name)
		link = conflictResolver.getSafeName(link, f)
		if link == "" {
			continue // unresolved conflict
		}

		data := mod.genFunction(f)
		if err := addFileTemplated(link, templates.Function, data); err != nil {
			return err
		}

		functions = append(functions, indexEntry{
			Link: link + "/",
			//nolint:staticcheck
			DisplayName: strings.Title(name),
		})
	}
	sortIndexEntries(functions)

	version := ""
	if mod.pkg.Version() != nil {
		version = mod.pkg.Version().String()
	}

	packageDetails := packageDetails{
		DisplayName:    getPackageDisplayName(def.Name),
		Repository:     def.Repository,
		RepositoryName: getRepositoryName(def.Repository),
		License:        def.License,
		Notes:          def.Attribution,
		Version:        version,
	}

	var modTitleTag string
	var packageDescription string
	// The same index.tmpl template is used for both top level package and module pages, if modules not present, assume
	// top level package index page when formatting title tags otherwise, if contains modules, assume modules top level
	// page when generating title tags.
	if len(modules) > 0 {
		modTitleTag = getPackageDisplayName(modTitle) + " Package"
	} else {
		modTitleTag = fmt.Sprintf("%s.%s", mod.pkg.Name(), modTitle)
		packageDescription = fmt.Sprintf("Explore the resources and functions of the %s.%s module.",
			mod.pkg.Name(), modTitle)
	}

	// Generate the index file.
	idxData := indexData{
		Tool:               mod.tool,
		PackageDescription: packageDescription,
		Title:              modTitle,
		TitleTag:           modTitleTag,
		Resources:          resources,
		Functions:          functions,
		Modules:            modules,
		PackageDetails:     packageDetails,
	}

	// If this is the root module, write out the package description.
	if mod.mod == "" {
		idxData.PackageDescription = mod.pkg.Description()
	}

	return addFileTemplated("", templates.Index, idxData)
}

// indexEntry represents an individual entry on an index page.
type indexEntry struct {
	Link        string
	DisplayName string
}

// indexData represents the index file data to be rendered as _index.md.
type indexData struct {
	Tool string

	Title              string
	TitleTag           string
	PackageDescription string

	Functions      []indexEntry
	Resources      []indexEntry
	Modules        []indexEntry
	PackageDetails packageDetails
}

func sortIndexEntries(entries []indexEntry) {
	if len(entries) == 0 {
		return
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].DisplayName < entries[j].DisplayName
	})
}

// getPackageDisplayName uses the title lookup map to look for a display name for the given title.
func getPackageDisplayName(title string) string {
	// If title not found in titleLookup map, default back to title given.
	if val, ok := titleLookup(title); ok {
		return val
	}
	return title
}

// getRepositoryName returns the repository name based on the repository's URL.
func getRepositoryName(repoURL string) string {
	return strings.TrimPrefix(repoURL, "https://github.com/")
}

func (dctx *Context) getMod(
	pkg schema.PackageReference,
	token string,
	tokenPkg schema.PackageReference,
	modules map[string]*modContext,
	tool string,
	add bool,
) *modContext {
	modName := pkg.TokenToModule(token)
	return dctx.getModByModName(pkg, tokens.ModuleName(modName), tokenPkg, modules, tool, add)
}

func (dctx *Context) getModByModName(
	pkg schema.PackageReference,
	modName tokens.ModuleName,
	tokenPkg schema.PackageReference,
	modules map[string]*modContext,
	tool string,
	add bool,
) *modContext {
	mod, ok := modules[string(modName)]
	if !ok {
		mod = &modContext{
			pkg:     pkg,
			mod:     string(modName),
			tool:    tool,
			context: dctx,
		}

		if modName != "" && codegen.PkgEquals(tokenPkg, pkg) {
			parentName, hasParent := parentModule(modName)
			if add && hasParent {
				parent := dctx.getModByModName(pkg, parentName, tokenPkg, modules, tool, add)
				parent.children = append(parent.children, mod)
			}
		}

		// Save the module only if we're adding and it's for the current package. This way, modules for external packages
		// are not saved.
		if add && tokenPkg == pkg {
			modules[string(modName)] = mod
		}
	}
	return mod
}

func (dctx *Context) generateModulesFromSchemaPackage(tool string, pkg *schema.Package) map[string]*modContext {
	// Group resources, types, and functions into modules.
	modules := map[string]*modContext{}

	// Decode language-specific info.
	if err := pkg.ImportLanguages(map[string]schema.Language{
		"go":     go_gen.Importer,
		"python": python.Importer,
		"csharp": dotnet.Importer,
		"nodejs": nodejs.Importer,
	}); err != nil {
		panic(err)
	}
	dctx.goPkgInfo, _ = pkg.Language["go"].(go_gen.GoPackageInfo)
	dctx.csharpPkgInfo, _ = pkg.Language["csharp"].(dotnet.CSharpPackageInfo)
	dctx.nodePkgInfo, _ = pkg.Language["nodejs"].(nodejs.NodePackageInfo)
	dctx.pythonPkgInfo, _ = pkg.Language["python"].(python.PackageInfo)

	goLangHelper := dctx.getLanguageDocHelper(language.Go).(*go_gen.DocLanguageHelper)
	// Generate the Go package map info now, so we can use that to get the type string names later.
	goLangHelper.GeneratePackagesMap(pkg.Reference(), tool, dctx.goPkgInfo)

	visitObjects := func(r *schema.Resource) {
		visitObjectTypes(r.InputProperties, func(t schema.Type) {
			switch T := t.(type) {
			case *schema.ObjectType:
				dctx.getMod(pkg.Reference(), T.Token, T.PackageReference, modules, tool, true).details(T).inputType = true
			}
		})
		if r.StateInputs != nil {
			visitObjectTypes(r.StateInputs.Properties, func(t schema.Type) {
				switch T := t.(type) {
				case *schema.ObjectType:
					dctx.getMod(pkg.Reference(), T.Token, T.PackageReference, modules, tool, true).details(T).inputType = true
				}
			})
		}
	}

	scanResource := func(r *schema.Resource) {
		mod := dctx.getMod(pkg.Reference(), r.Token, r.PackageReference, modules, tool, true)
		mod.resources = append(mod.resources, r)
		visitObjects(r)
	}

	scanK8SResource := func(r *schema.Resource) {
		mod := getKubernetesMod(pkg, r.Token, modules, tool)
		mod.resources = append(mod.resources, r)
		visitObjects(r)
	}

	glog.V(3).Infoln("scanning resources")
	if isKubernetesPackage(pkg.Reference()) {
		scanK8SResource(pkg.Provider)
		for _, r := range pkg.Resources {
			scanK8SResource(r)
		}
	} else {
		scanResource(pkg.Provider)
		for _, r := range pkg.Resources {
			scanResource(r)
		}
	}
	glog.V(3).Infoln("done scanning resources")

	for _, f := range pkg.Functions {
		if !f.IsMethod {
			mod := dctx.getMod(pkg.Reference(), f.Token, f.PackageReference, modules, tool, true)
			mod.functions = append(mod.functions, f)
		}
	}

	// Find nested types.
	for _, t := range pkg.Types {
		switch typ := t.(type) {
		case *schema.ObjectType:
			mod := dctx.getMod(pkg.Reference(), typ.Token, typ.PackageReference, modules, tool, false)
			if mod.details(typ).inputType {
				mod.inputTypes = append(mod.inputTypes, typ)
			}
		}
	}

	return modules
}

// generateModules generates constructor syntax examples for all resources of the input the package in the given input
// languages. We first generate a PCL program from the resource definitions, then for each language we call
// `GenerateProgram` to generate a language program that has all resources with their inputs being default values. After
// that, we extract the resource declarations from each program into a structured map.
//
// The reason we generate a full program a schema with all the resources is that if we did every resource separately, it
// would take too long to generate the programs and the rest of the docs. That is why we batch generate them and split
// the resource declarations by comment delimiters.
func generateConstructorSyntaxData(pkg *schema.Package) *constructorSyntaxData {
	loader := NewStaticSchemaLoader(pkg)
	constructorGenerator := &constructorSyntaxGenerator{
		indentSize:             0,
		requiredPropertiesOnly: false,
	}

	schemaProgram := constructorGenerator.generateAll(pkg, generateAllOptions{
		includeResources: true,
		includeFunctions: false,
		resourcesToSkip: []string{
			// Skipping problematic resources which block example generation
			"aws:wafv2/ruleGroup:RuleGroup",
			"aws:wafv2/webAcl:WebAcl",
		},
	})

	packagesToSkip := map[string]bool{"aws-native": true}

	type ProgramGenerator = func(program *pcl.Program) (files map[string][]byte, diags hcl.Diagnostics, err error)

	constructorSyntax := newConstructorSyntaxData()
	for lang := range language.All() {
		if packagesToSkip[pkg.Name] {
			continue
		}

		boundProgram, err := constructorGenerator.bindProgram(loader, schemaProgram)
		if err != nil {
			continue
		}

		safeExtract := func(generator ProgramGenerator) (files map[string][]byte, diags hcl.Diagnostics, err error) {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("panic: %v", r)
					files = map[string][]byte{}
				}
			}()
			return generator(boundProgram)
		}

		switch lang {
		case language.NodeJS:
			files, diags, err := safeExtract(nodejs.GenerateProgram)
			if !diags.HasErrors() && err == nil {
				program := string(files["index.ts"])
				constructorSyntax.typescript = extractConstructorSyntaxExamples(
					program, /* program */
					"",      /* indentation to trim */
					"//",    /* comment prefix */
					func(line string) bool { return strings.HasSuffix(line, "});") })
			}
		case language.Python:
			files, diags, err := safeExtract(python.GenerateProgram)
			if !diags.HasErrors() && err == nil {
				program := string(files["__main__.py"])
				constructorSyntax.python = extractConstructorSyntaxExamples(
					program, /* program */
					"",      /* indentation to trim */
					"#",     /* comment prefix */
					func(line string) bool { return strings.HasSuffix(line, ")") })
			}
		case language.CSharp:
			files, diags, err := safeExtract(dotnet.GenerateProgram)
			if !diags.HasErrors() && err == nil {
				exampleEnd := func(line string) bool {
					return strings.HasSuffix(line, "});") || strings.HasSuffix(line, `");`)
				}
				program := string(files["Program.cs"])
				constructorSyntax.csharp = extractConstructorSyntaxExamples(
					program, /* program */
					"    ",  /* indentation to trim */
					"//",    /* comment prefix */
					exampleEnd)
			}
		case language.Go:
			files, diags, err := safeExtract(go_gen.GenerateProgram)
			if !diags.HasErrors() && err == nil {
				var program string
				content := files["main.go"]
				if formatted, err := format.Source(content); err == nil {
					program = string(formatted)
				} else {
					program = string(content)
				}

				constructorSyntax.golang = extractConstructorSyntaxExamples(
					program, /* program */
					"\t\t",  /* indentation to trim */
					"//",    /* comment prefix */
					func(line string) bool { return strings.HasSuffix(strings.TrimSpace(line), ")") })
			}
		case language.YAML:
			files, diags, err := safeExtract(yaml.GenerateProgram)
			if !diags.HasErrors() && err == nil {
				program := string(files["Main.yaml"])
				constructorSyntax.yaml = extractConstructorSyntaxExamplesFromYAML(program)
			}
		case language.Java:
			files, diags, err := safeExtract(java.GenerateProgram)
			if !diags.HasErrors() && err == nil {
				program := string(files["MyStack.java"])
				constructorSyntax.java = extractConstructorSyntaxExamples(
					program,    /* program */
					"        ", /* indentation to trim */
					"//",       /* comment prefix */
					func(line string) bool { return strings.HasSuffix(line, ");") })
			}
		}
	}

	return constructorSyntax
}

// collapseYAMLToken converts an exact token to a token more suitable for display.
//
// Examples:
//
//	fizz:index/buzz:Buzz => fizz:Buzz
//	fizz:mode/buzz:Buzz  => fizz:mode:Buzz
//	foo:index:Bar        => foo:Bar
//	foo::Bar             => foo:Bar
//	fizz:mod:buzz        => fizz:mod:buzz
func collapseYAMLToken(token string) string {
	tokenParts := strings.Split(token, ":")

	if len(tokenParts) == 3 {
		title := func(s string) string {
			r := []rune(s)
			if len(r) == 0 {
				return ""
			}
			return strings.ToTitle(string(r[0])) + string(r[1:])
		}
		if mod := strings.Split(tokenParts[1], "/"); len(mod) == 2 && title(mod[1]) == tokenParts[2] {
			// aws:s3/bucket:Bucket => aws:s3:Bucket
			// We handle the case foo:index/bar:Bar => foo:index:Bar
			tokenParts = []string{tokenParts[0], mod[0], tokenParts[2]}
		}

		if tokenParts[1] == "index" || tokenParts[1] == "" {
			// foo:index:Bar => foo:Bar
			// or
			// foo::Bar => foo:Bar
			tokenParts = []string{tokenParts[0], tokenParts[2]}
		}
	}

	return strings.Join(tokenParts, ":")
}

// GeneratePackage generates docs for each resource given the Pulumi schema. The returned map contains the filename with
// path as the key and the contents as its value.
func (dctx *Context) GeneratePackage() (map[string][]byte, error) {
	if dctx.modules() == nil {
		return nil, errors.New("must call Initialize before generating the docs package")
	}

	defer glog.Flush()

	glog.V(3).Infoln("generating package docs now...")
	files := codegen.Fs{}
	modules := []string{}
	modMap := dctx.modules()
	for k := range modMap {
		modules = append(modules, k)
	}
	sort.Strings(modules)
	for _, mod := range modules {
		if err := modMap[mod].gen(files); err != nil {
			return nil, err
		}
	}

	return files, nil
}

const (
	// "" indicates the top-most module.
	topMostModule tokens.ModuleName = ""
)

// GeneratePackageTree returns a navigable structure starting from the top-most module.
func (dctx *Context) GeneratePackageTree() ([]PackageTreeItem, error) {
	if dctx.modules() == nil {
		return nil, errors.New("must call Initialize before generating the docs package")
	}

	defer glog.Flush()

	var packageTree []PackageTreeItem
	if rootMod, ok := dctx.modules()[string(topMostModule)]; ok {
		tree, err := generatePackageTree(*rootMod)
		if err != nil {
			glog.Errorf("Error generating the package tree for package: %v", err)
		}

		packageTree = tree
	} else {
		glog.Error("A root module entry was not found for the package. Cannot generate the package tree...")
	}

	return packageTree, nil
}

func visitObjectTypes(properties []*schema.Property, visitor func(t schema.Type)) {
	codegen.VisitTypeClosure(properties, func(t schema.Type) {
		switch st := t.(type) {
		case *schema.EnumType, *schema.ObjectType, *schema.ResourceType:
			visitor(st)
		}
	})
}

// Returns the parent module, if available. For top-level modules the parent is a special topMostModule. For
// topMostModule itself there is no parent and this function returns false.
func parentModule(modName tokens.ModuleName) (tokens.ModuleName, bool) {
	switch {
	case modName == topMostModule:
		return "", false
	case strings.Contains(string(modName), tokens.QNameDelimiter):
		return tokens.ModuleName(tokens.QName(modName).Namespace()), true
	default:
		return topMostModule, true
	}
}
