
---
title: "listOpenIdConnectProviderSecrets"
title_tag: "azure-native.apimanagement.listOpenIdConnectProviderSecrets"
meta_desc: "Documentation for the azure-native.apimanagement.listOpenIdConnectProviderSecrets function with examples, input properties, output properties, and supporting types."
layout: api
no_edit_this_page: true
---



<!-- WARNING: this file was generated by Pulumi Docs Generator. -->
<!-- Do not edit by hand unless you're certain you know what you are doing! -->

Gets the client secret details of the OpenID Connect Provider.

Azure REST API version: 2022-09-01-preview.

Other available API versions: 2022-08-01, 2024-06-01-preview.

These versions are not included in the SDK but you can generate them as a local SDK package; see the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.




## Using listOpenIdConnectProviderSecrets {#using}

Two invocation forms are available. The direct form accepts plain
arguments and either blocks until the result value is available, or
returns a Promise-wrapped result. The output form accepts
Input-wrapped arguments and returns an Output-wrapped result.

<div>
<pulumi-chooser type="language" options="csharp,go,typescript,python,yaml,java"></pulumi-chooser>
</div>


<div>
<pulumi-choosable type="language" values="javascript,typescript">
<div class="highlight"
><pre class="chroma"><code class="language-typescript" data-lang="typescript"
><span class="k">function </span>listOpenIdConnectProviderSecrets<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">ListOpenIdConnectProviderSecretsArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Promise&lt;<span class="nx"><a href="#result">ListOpenIdConnectProviderSecretsResult</a></span>></span
><span class="k">
function </span>listOpenIdConnectProviderSecretsOutput<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">ListOpenIdConnectProviderSecretsOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Output&lt;<span class="nx"><a href="#result">ListOpenIdConnectProviderSecretsResult</a></span>></span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="python">
<div class="highlight"><pre class="chroma"><code class="language-python" data-lang="python"
><span class="k">def </span>list_open_id_connect_provider_secrets<span class="p">(</span><span class="nx">opid</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                                          <span class="nx">resource_group_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                                          <span class="nx">service_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                                          <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>ListOpenIdConnectProviderSecretsResult</span
><span class="k">
def </span>list_open_id_connect_provider_secrets_output<span class="p">(</span><span class="nx">opid</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                                          <span class="nx">resource_group_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                                          <span class="nx">service_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                                          <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>Output[ListOpenIdConnectProviderSecretsResult]</span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="go">
<div class="highlight"><pre class="chroma"><code class="language-go" data-lang="go"
><span class="k">func </span>ListOpenIdConnectProviderSecrets<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">ListOpenIdConnectProviderSecretsArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) (*<span class="nx"><a href="#result">ListOpenIdConnectProviderSecretsResult</a></span>, error)</span
><span class="k">
func </span>ListOpenIdConnectProviderSecretsOutput<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">ListOpenIdConnectProviderSecretsOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) ListOpenIdConnectProviderSecretsResultOutput</span
></code></pre></div>

&gt; Note: This function is named `ListOpenIdConnectProviderSecrets` in the Go SDK.

</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="csharp">
<div class="highlight"><pre class="chroma"><code class="language-csharp" data-lang="csharp"><span class="k">public static class </span><span class="nx">ListOpenIdConnectProviderSecrets </span><span class="p">
{</span><span class="k">
    public static </span>Task&lt;<span class="nx"><a href="#result">ListOpenIdConnectProviderSecretsResult</a></span>> <span class="p">InvokeAsync(</span><span class="nx">ListOpenIdConnectProviderSecretsArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="k">
    public static </span>Output&lt;<span class="nx"><a href="#result">ListOpenIdConnectProviderSecretsResult</a></span>> <span class="p">Invoke(</span><span class="nx">ListOpenIdConnectProviderSecretsInvokeArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="p">
}</span></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="java">
<div class="highlight"><pre class="chroma"><code class="language-java" data-lang="java"><span class="k">public static CompletableFuture&lt;<span class="nx"><a href="#result">ListOpenIdConnectProviderSecretsResult</a></span>> </span>listOpenIdConnectProviderSecrets<span class="p">(</span><span class="nx">ListOpenIdConnectProviderSecretsArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
<span class="k">public static Output&lt;<span class="nx"><a href="#result">ListOpenIdConnectProviderSecretsResult</a></span>> </span>listOpenIdConnectProviderSecrets<span class="p">(</span><span class="nx">ListOpenIdConnectProviderSecretsArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
</code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="yaml">
<div class="highlight"><pre class="chroma"><code class="language-yaml" data-lang="yaml"><span class="k">fn::invoke:</span>
<span class="k">&nbsp;&nbsp;function:</span> azure-native:apimanagement:listOpenIdConnectProviderSecrets
<span class="k">&nbsp;&nbsp;arguments:</span>
<span class="c">&nbsp;&nbsp;&nbsp;&nbsp;# arguments dictionary</span></code></pre></div>
</pulumi-choosable>
</div>



The following arguments are supported:


<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="opid_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#opid_csharp" style="color: inherit; text-decoration: inherit;">Opid</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Identifier of the OpenID Connect Provider.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_csharp" style="color: inherit; text-decoration: inherit;">Resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the resource group. The name is case insensitive.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="servicename_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#servicename_csharp" style="color: inherit; text-decoration: inherit;">Service<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the API Management service.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="opid_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#opid_go" style="color: inherit; text-decoration: inherit;">Opid</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Identifier of the OpenID Connect Provider.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_go" style="color: inherit; text-decoration: inherit;">Resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the resource group. The name is case insensitive.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="servicename_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#servicename_go" style="color: inherit; text-decoration: inherit;">Service<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the API Management service.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="opid_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#opid_java" style="color: inherit; text-decoration: inherit;">opid</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Identifier of the OpenID Connect Provider.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_java" style="color: inherit; text-decoration: inherit;">resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The name of the resource group. The name is case insensitive.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="servicename_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#servicename_java" style="color: inherit; text-decoration: inherit;">service<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The name of the API Management service.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="opid_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#opid_nodejs" style="color: inherit; text-decoration: inherit;">opid</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Identifier of the OpenID Connect Provider.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_nodejs" style="color: inherit; text-decoration: inherit;">resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the resource group. The name is case insensitive.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="servicename_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#servicename_nodejs" style="color: inherit; text-decoration: inherit;">service<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the API Management service.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="opid_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#opid_python" style="color: inherit; text-decoration: inherit;">opid</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Identifier of the OpenID Connect Provider.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resource_group_name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resource_group_name_python" style="color: inherit; text-decoration: inherit;">resource_<wbr>group_<wbr>name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The name of the resource group. The name is case insensitive.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="service_name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#service_name_python" style="color: inherit; text-decoration: inherit;">service_<wbr>name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The name of the API Management service.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="opid_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#opid_yaml" style="color: inherit; text-decoration: inherit;">opid</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Identifier of the OpenID Connect Provider.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_yaml" style="color: inherit; text-decoration: inherit;">resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The name of the resource group. The name is case insensitive.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="servicename_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#servicename_yaml" style="color: inherit; text-decoration: inherit;">service<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The name of the API Management service.</dd></dl>
</pulumi-choosable>
</div>




## listOpenIdConnectProviderSecrets Result {#result}

The following output properties are available:



<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="clientsecret_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#clientsecret_csharp" style="color: inherit; text-decoration: inherit;">Client<wbr>Secret</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Client or app secret used in IdentityProviders, Aad, OpenID or OAuth.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="clientsecret_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#clientsecret_go" style="color: inherit; text-decoration: inherit;">Client<wbr>Secret</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Client or app secret used in IdentityProviders, Aad, OpenID or OAuth.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="clientsecret_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#clientsecret_java" style="color: inherit; text-decoration: inherit;">client<wbr>Secret</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Client or app secret used in IdentityProviders, Aad, OpenID or OAuth.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="clientsecret_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#clientsecret_nodejs" style="color: inherit; text-decoration: inherit;">client<wbr>Secret</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Client or app secret used in IdentityProviders, Aad, OpenID or OAuth.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="client_secret_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#client_secret_python" style="color: inherit; text-decoration: inherit;">client_<wbr>secret</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Client or app secret used in IdentityProviders, Aad, OpenID or OAuth.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="clientsecret_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#clientsecret_yaml" style="color: inherit; text-decoration: inherit;">client<wbr>Secret</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Client or app secret used in IdentityProviders, Aad, OpenID or OAuth.</dd></dl>
</pulumi-choosable>
</div>





<h2 id="package-details">Package Details</h2>
<dl class="package-details">
	<dt>Repository</dt>
	<dd><a href="https://github.com/pulumi/pulumi-azure-native">Azure Native pulumi/pulumi-azure-native</a></dd>
	<dt>License</dt>
	<dd>Apache-2.0</dd>
</dl>

