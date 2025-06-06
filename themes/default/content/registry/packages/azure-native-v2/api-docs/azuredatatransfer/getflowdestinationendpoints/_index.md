
---
title: "getFlowDestinationEndpoints"
title_tag: "azure-native.azuredatatransfer.getFlowDestinationEndpoints"
meta_desc: "Documentation for the azure-native.azuredatatransfer.getFlowDestinationEndpoints function with examples, input properties, output properties, and supporting types."
layout: api
no_edit_this_page: true
---



<!-- WARNING: this file was generated by Pulumi Docs Generator. -->
<!-- Do not edit by hand unless you're certain you know what you are doing! -->

Get the destination endpoints for the specified flow and stream ID.

Azure REST API version: 2024-09-27.




## Using getFlowDestinationEndpoints {#using}

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
><span class="k">function </span>getFlowDestinationEndpoints<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">GetFlowDestinationEndpointsArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Promise&lt;<span class="nx"><a href="#result">GetFlowDestinationEndpointsResult</a></span>></span
><span class="k">
function </span>getFlowDestinationEndpointsOutput<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">GetFlowDestinationEndpointsOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Output&lt;<span class="nx"><a href="#result">GetFlowDestinationEndpointsResult</a></span>></span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="python">
<div class="highlight"><pre class="chroma"><code class="language-python" data-lang="python"
><span class="k">def </span>get_flow_destination_endpoints<span class="p">(</span><span class="nx">connection_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                                   <span class="nx">flow_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                                   <span class="nx">resource_group_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                                   <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>GetFlowDestinationEndpointsResult</span
><span class="k">
def </span>get_flow_destination_endpoints_output<span class="p">(</span><span class="nx">connection_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                                   <span class="nx">flow_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                                   <span class="nx">resource_group_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                                   <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>Output[GetFlowDestinationEndpointsResult]</span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="go">
<div class="highlight"><pre class="chroma"><code class="language-go" data-lang="go"
><span class="k">func </span>GetFlowDestinationEndpoints<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">GetFlowDestinationEndpointsArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) (*<span class="nx"><a href="#result">GetFlowDestinationEndpointsResult</a></span>, error)</span
><span class="k">
func </span>GetFlowDestinationEndpointsOutput<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">GetFlowDestinationEndpointsOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) GetFlowDestinationEndpointsResultOutput</span
></code></pre></div>

&gt; Note: This function is named `GetFlowDestinationEndpoints` in the Go SDK.

</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="csharp">
<div class="highlight"><pre class="chroma"><code class="language-csharp" data-lang="csharp"><span class="k">public static class </span><span class="nx">GetFlowDestinationEndpoints </span><span class="p">
{</span><span class="k">
    public static </span>Task&lt;<span class="nx"><a href="#result">GetFlowDestinationEndpointsResult</a></span>> <span class="p">InvokeAsync(</span><span class="nx">GetFlowDestinationEndpointsArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="k">
    public static </span>Output&lt;<span class="nx"><a href="#result">GetFlowDestinationEndpointsResult</a></span>> <span class="p">Invoke(</span><span class="nx">GetFlowDestinationEndpointsInvokeArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="p">
}</span></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="java">
<div class="highlight"><pre class="chroma"><code class="language-java" data-lang="java"><span class="k">public static CompletableFuture&lt;<span class="nx"><a href="#result">GetFlowDestinationEndpointsResult</a></span>> </span>getFlowDestinationEndpoints<span class="p">(</span><span class="nx">GetFlowDestinationEndpointsArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
<span class="k">public static Output&lt;<span class="nx"><a href="#result">GetFlowDestinationEndpointsResult</a></span>> </span>getFlowDestinationEndpoints<span class="p">(</span><span class="nx">GetFlowDestinationEndpointsArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
</code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="yaml">
<div class="highlight"><pre class="chroma"><code class="language-yaml" data-lang="yaml"><span class="k">fn::invoke:</span>
<span class="k">&nbsp;&nbsp;function:</span> azure-native:azuredatatransfer:getFlowDestinationEndpoints
<span class="k">&nbsp;&nbsp;arguments:</span>
<span class="c">&nbsp;&nbsp;&nbsp;&nbsp;# arguments dictionary</span></code></pre></div>
</pulumi-choosable>
</div>



The following arguments are supported:


<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="connectionname_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#connectionname_csharp" style="color: inherit; text-decoration: inherit;">Connection<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name for the connection that is to be requested.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="flowname_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#flowname_csharp" style="color: inherit; text-decoration: inherit;">Flow<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name for the flow that is to be onboarded.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_csharp" style="color: inherit; text-decoration: inherit;">Resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the resource group. The name is case insensitive.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="connectionname_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#connectionname_go" style="color: inherit; text-decoration: inherit;">Connection<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name for the connection that is to be requested.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="flowname_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#flowname_go" style="color: inherit; text-decoration: inherit;">Flow<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name for the flow that is to be onboarded.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_go" style="color: inherit; text-decoration: inherit;">Resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the resource group. The name is case insensitive.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="connectionname_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#connectionname_java" style="color: inherit; text-decoration: inherit;">connection<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The name for the connection that is to be requested.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="flowname_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#flowname_java" style="color: inherit; text-decoration: inherit;">flow<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The name for the flow that is to be onboarded.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_java" style="color: inherit; text-decoration: inherit;">resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The name of the resource group. The name is case insensitive.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="connectionname_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#connectionname_nodejs" style="color: inherit; text-decoration: inherit;">connection<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name for the connection that is to be requested.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="flowname_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#flowname_nodejs" style="color: inherit; text-decoration: inherit;">flow<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name for the flow that is to be onboarded.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_nodejs" style="color: inherit; text-decoration: inherit;">resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the resource group. The name is case insensitive.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="connection_name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#connection_name_python" style="color: inherit; text-decoration: inherit;">connection_<wbr>name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The name for the connection that is to be requested.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="flow_name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#flow_name_python" style="color: inherit; text-decoration: inherit;">flow_<wbr>name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The name for the flow that is to be onboarded.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resource_group_name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resource_group_name_python" style="color: inherit; text-decoration: inherit;">resource_<wbr>group_<wbr>name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The name of the resource group. The name is case insensitive.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="connectionname_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#connectionname_yaml" style="color: inherit; text-decoration: inherit;">connection<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The name for the connection that is to be requested.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="flowname_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#flowname_yaml" style="color: inherit; text-decoration: inherit;">flow<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The name for the flow that is to be onboarded.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_yaml" style="color: inherit; text-decoration: inherit;">resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The name of the resource group. The name is case insensitive.</dd></dl>
</pulumi-choosable>
</div>




## getFlowDestinationEndpoints Result {#result}

The following output properties are available:



<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="endpoints_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#endpoints_csharp" style="color: inherit; text-decoration: inherit;">Endpoints</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">List&lt;string&gt;</span>
    </dt>
    <dd>The destination endpoints for the flow stream</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="endpoints_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#endpoints_go" style="color: inherit; text-decoration: inherit;">Endpoints</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">[]string</span>
    </dt>
    <dd>The destination endpoints for the flow stream</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="endpoints_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#endpoints_java" style="color: inherit; text-decoration: inherit;">endpoints</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">List&lt;String&gt;</span>
    </dt>
    <dd>The destination endpoints for the flow stream</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="endpoints_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#endpoints_nodejs" style="color: inherit; text-decoration: inherit;">endpoints</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string[]</span>
    </dt>
    <dd>The destination endpoints for the flow stream</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="endpoints_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#endpoints_python" style="color: inherit; text-decoration: inherit;">endpoints</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Sequence[str]</span>
    </dt>
    <dd>The destination endpoints for the flow stream</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="endpoints_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#endpoints_yaml" style="color: inherit; text-decoration: inherit;">endpoints</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">List&lt;String&gt;</span>
    </dt>
    <dd>The destination endpoints for the flow stream</dd></dl>
</pulumi-choosable>
</div>





<h2 id="package-details">Package Details</h2>
<dl class="package-details">
	<dt>Repository</dt>
	<dd><a href="https://github.com/pulumi/pulumi-azure-native">Azure Native pulumi/pulumi-azure-native</a></dd>
	<dt>License</dt>
	<dd>Apache-2.0</dd>
</dl>

