
---
title: "getClientConfig"
title_tag: "azure-native.authorization.getClientConfig"
meta_desc: "Documentation for the azure-native.authorization.getClientConfig function with examples, input properties, output properties, and supporting types."
layout: api
no_edit_this_page: true
---



<!-- WARNING: this file was generated by Pulumi Docs Generator. -->
<!-- Do not edit by hand unless you're certain you know what you are doing! -->

Use this function to access the current configuration of the native Azure provider.




## Using getClientConfig {#using}

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
><span class="k">function </span>getClientConfig<span class="p">(</span><span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Promise&lt;<span class="nx"><a href="#result">GetClientConfigResult</a></span>></span
><span class="k">
function </span>getClientConfigOutput<span class="p">(</span><span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Output&lt;<span class="nx"><a href="#result">GetClientConfigResult</a></span>></span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="python">
<div class="highlight"><pre class="chroma"><code class="language-python" data-lang="python"
><span class="k">def </span>get_client_config<span class="p">(</span><span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>GetClientConfigResult</span
><span class="k">
def </span>get_client_config_output<span class="p">(</span><span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>Output[GetClientConfigResult]</span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="go">
<div class="highlight"><pre class="chroma"><code class="language-go" data-lang="go"
><span class="k">func </span>GetClientConfig<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) (*<span class="nx"><a href="#result">GetClientConfigResult</a></span>, error)</span
><span class="k">
func </span>GetClientConfigOutput<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) GetClientConfigResultOutput</span
></code></pre></div>

&gt; Note: This function is named `GetClientConfig` in the Go SDK.

</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="csharp">
<div class="highlight"><pre class="chroma"><code class="language-csharp" data-lang="csharp"><span class="k">public static class </span><span class="nx">GetClientConfig </span><span class="p">
{</span><span class="k">
    public static </span>Task&lt;<span class="nx"><a href="#result">GetClientConfigResult</a></span>> <span class="p">InvokeAsync(</span><span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="k">
    public static </span>Output&lt;<span class="nx"><a href="#result">GetClientConfigResult</a></span>> <span class="p">Invoke(</span><span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="p">
}</span></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="java">
<div class="highlight"><pre class="chroma"><code class="language-java" data-lang="java"><span class="k">public static CompletableFuture&lt;<span class="nx"><a href="#result">GetClientConfigResult</a></span>> </span>getClientConfig<span class="p">(</span><span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
<span class="k">public static Output&lt;<span class="nx"><a href="#result">GetClientConfigResult</a></span>> </span>getClientConfig<span class="p">(</span><span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
</code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="yaml">
<div class="highlight"><pre class="chroma"><code class="language-yaml" data-lang="yaml"><span class="k">fn::invoke:</span>
<span class="k">&nbsp;&nbsp;function:</span> azure-native:authorization:getClientConfig
<span class="k">&nbsp;&nbsp;arguments:</span>
<span class="c">&nbsp;&nbsp;&nbsp;&nbsp;# arguments dictionary</span></code></pre></div>
</pulumi-choosable>
</div>




## getClientConfig Result {#result}

The following output properties are available:



<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="clientid_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#clientid_csharp" style="color: inherit; text-decoration: inherit;">Client<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Azure Client ID (Application Object ID).</dd><dt class="property-"
            title="">
        <span id="objectid_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#objectid_csharp" style="color: inherit; text-decoration: inherit;">Object<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Azure Object ID of the current user or service principal.</dd><dt class="property-"
            title="">
        <span id="subscriptionid_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#subscriptionid_csharp" style="color: inherit; text-decoration: inherit;">Subscription<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Azure Subscription ID</dd><dt class="property-"
            title="">
        <span id="tenantid_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#tenantid_csharp" style="color: inherit; text-decoration: inherit;">Tenant<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Azure Tenant ID</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="clientid_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#clientid_go" style="color: inherit; text-decoration: inherit;">Client<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Azure Client ID (Application Object ID).</dd><dt class="property-"
            title="">
        <span id="objectid_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#objectid_go" style="color: inherit; text-decoration: inherit;">Object<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Azure Object ID of the current user or service principal.</dd><dt class="property-"
            title="">
        <span id="subscriptionid_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#subscriptionid_go" style="color: inherit; text-decoration: inherit;">Subscription<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Azure Subscription ID</dd><dt class="property-"
            title="">
        <span id="tenantid_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#tenantid_go" style="color: inherit; text-decoration: inherit;">Tenant<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Azure Tenant ID</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="clientid_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#clientid_java" style="color: inherit; text-decoration: inherit;">client<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Azure Client ID (Application Object ID).</dd><dt class="property-"
            title="">
        <span id="objectid_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#objectid_java" style="color: inherit; text-decoration: inherit;">object<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Azure Object ID of the current user or service principal.</dd><dt class="property-"
            title="">
        <span id="subscriptionid_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#subscriptionid_java" style="color: inherit; text-decoration: inherit;">subscription<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Azure Subscription ID</dd><dt class="property-"
            title="">
        <span id="tenantid_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#tenantid_java" style="color: inherit; text-decoration: inherit;">tenant<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Azure Tenant ID</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="clientid_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#clientid_nodejs" style="color: inherit; text-decoration: inherit;">client<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Azure Client ID (Application Object ID).</dd><dt class="property-"
            title="">
        <span id="objectid_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#objectid_nodejs" style="color: inherit; text-decoration: inherit;">object<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Azure Object ID of the current user or service principal.</dd><dt class="property-"
            title="">
        <span id="subscriptionid_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#subscriptionid_nodejs" style="color: inherit; text-decoration: inherit;">subscription<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Azure Subscription ID</dd><dt class="property-"
            title="">
        <span id="tenantid_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#tenantid_nodejs" style="color: inherit; text-decoration: inherit;">tenant<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Azure Tenant ID</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="client_id_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#client_id_python" style="color: inherit; text-decoration: inherit;">client_<wbr>id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Azure Client ID (Application Object ID).</dd><dt class="property-"
            title="">
        <span id="object_id_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#object_id_python" style="color: inherit; text-decoration: inherit;">object_<wbr>id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Azure Object ID of the current user or service principal.</dd><dt class="property-"
            title="">
        <span id="subscription_id_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#subscription_id_python" style="color: inherit; text-decoration: inherit;">subscription_<wbr>id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Azure Subscription ID</dd><dt class="property-"
            title="">
        <span id="tenant_id_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#tenant_id_python" style="color: inherit; text-decoration: inherit;">tenant_<wbr>id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Azure Tenant ID</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="clientid_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#clientid_yaml" style="color: inherit; text-decoration: inherit;">client<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Azure Client ID (Application Object ID).</dd><dt class="property-"
            title="">
        <span id="objectid_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#objectid_yaml" style="color: inherit; text-decoration: inherit;">object<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Azure Object ID of the current user or service principal.</dd><dt class="property-"
            title="">
        <span id="subscriptionid_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#subscriptionid_yaml" style="color: inherit; text-decoration: inherit;">subscription<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Azure Subscription ID</dd><dt class="property-"
            title="">
        <span id="tenantid_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#tenantid_yaml" style="color: inherit; text-decoration: inherit;">tenant<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Azure Tenant ID</dd></dl>
</pulumi-choosable>
</div>





<h2 id="package-details">Package Details</h2>
<dl class="package-details">
	<dt>Repository</dt>
	<dd><a href="https://github.com/pulumi/pulumi-azure-native">Azure Native pulumi/pulumi-azure-native</a></dd>
	<dt>License</dt>
	<dd>Apache-2.0</dd>
</dl>

