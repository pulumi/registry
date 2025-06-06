
---
title: "getProviderActionOverviewStatus"
title_tag: "azure-native.appcomplianceautomation.getProviderActionOverviewStatus"
meta_desc: "Documentation for the azure-native.appcomplianceautomation.getProviderActionOverviewStatus function with examples, input properties, output properties, and supporting types."
layout: api
no_edit_this_page: true
---



<!-- WARNING: this file was generated by Pulumi Docs Generator. -->
<!-- Do not edit by hand unless you're certain you know what you are doing! -->

Get the resource overview status.

Azure REST API version: 2024-06-27.




## Using getProviderActionOverviewStatus {#using}

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
><span class="k">function </span>getProviderActionOverviewStatus<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">GetProviderActionOverviewStatusArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Promise&lt;<span class="nx"><a href="#result">GetProviderActionOverviewStatusResult</a></span>></span
><span class="k">
function </span>getProviderActionOverviewStatusOutput<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">GetProviderActionOverviewStatusOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Output&lt;<span class="nx"><a href="#result">GetProviderActionOverviewStatusResult</a></span>></span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="python">
<div class="highlight"><pre class="chroma"><code class="language-python" data-lang="python"
><span class="k">def </span>get_provider_action_overview_status<span class="p">(</span><span class="nx">type</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                                        <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>GetProviderActionOverviewStatusResult</span
><span class="k">
def </span>get_provider_action_overview_status_output<span class="p">(</span><span class="nx">type</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                                        <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>Output[GetProviderActionOverviewStatusResult]</span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="go">
<div class="highlight"><pre class="chroma"><code class="language-go" data-lang="go"
><span class="k">func </span>GetProviderActionOverviewStatus<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">GetProviderActionOverviewStatusArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) (*<span class="nx"><a href="#result">GetProviderActionOverviewStatusResult</a></span>, error)</span
><span class="k">
func </span>GetProviderActionOverviewStatusOutput<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">GetProviderActionOverviewStatusOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) GetProviderActionOverviewStatusResultOutput</span
></code></pre></div>

&gt; Note: This function is named `GetProviderActionOverviewStatus` in the Go SDK.

</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="csharp">
<div class="highlight"><pre class="chroma"><code class="language-csharp" data-lang="csharp"><span class="k">public static class </span><span class="nx">GetProviderActionOverviewStatus </span><span class="p">
{</span><span class="k">
    public static </span>Task&lt;<span class="nx"><a href="#result">GetProviderActionOverviewStatusResult</a></span>> <span class="p">InvokeAsync(</span><span class="nx">GetProviderActionOverviewStatusArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="k">
    public static </span>Output&lt;<span class="nx"><a href="#result">GetProviderActionOverviewStatusResult</a></span>> <span class="p">Invoke(</span><span class="nx">GetProviderActionOverviewStatusInvokeArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="p">
}</span></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="java">
<div class="highlight"><pre class="chroma"><code class="language-java" data-lang="java"><span class="k">public static CompletableFuture&lt;<span class="nx"><a href="#result">GetProviderActionOverviewStatusResult</a></span>> </span>getProviderActionOverviewStatus<span class="p">(</span><span class="nx">GetProviderActionOverviewStatusArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
<span class="k">public static Output&lt;<span class="nx"><a href="#result">GetProviderActionOverviewStatusResult</a></span>> </span>getProviderActionOverviewStatus<span class="p">(</span><span class="nx">GetProviderActionOverviewStatusArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
</code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="yaml">
<div class="highlight"><pre class="chroma"><code class="language-yaml" data-lang="yaml"><span class="k">fn::invoke:</span>
<span class="k">&nbsp;&nbsp;function:</span> azure-native:appcomplianceautomation:getProviderActionOverviewStatus
<span class="k">&nbsp;&nbsp;arguments:</span>
<span class="c">&nbsp;&nbsp;&nbsp;&nbsp;# arguments dictionary</span></code></pre></div>
</pulumi-choosable>
</div>



The following arguments are supported:


<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="type_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_csharp" style="color: inherit; text-decoration: inherit;">Type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The resource type.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="type_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_go" style="color: inherit; text-decoration: inherit;">Type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The resource type.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="type_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_java" style="color: inherit; text-decoration: inherit;">type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The resource type.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="type_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_nodejs" style="color: inherit; text-decoration: inherit;">type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The resource type.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="type_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_python" style="color: inherit; text-decoration: inherit;">type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The resource type.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="type_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_yaml" style="color: inherit; text-decoration: inherit;">type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The resource type.</dd></dl>
</pulumi-choosable>
</div>




## getProviderActionOverviewStatus Result {#result}

The following output properties are available:



<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="statuslist_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#statuslist_csharp" style="color: inherit; text-decoration: inherit;">Status<wbr>List</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#statusitemresponse">List&lt;Pulumi.<wbr>Azure<wbr>Native.<wbr>App<wbr>Compliance<wbr>Automation.<wbr>Outputs.<wbr>Status<wbr>Item<wbr>Response&gt;</a></span>
    </dt>
    <dd>List of different status items.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="statuslist_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#statuslist_go" style="color: inherit; text-decoration: inherit;">Status<wbr>List</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#statusitemresponse">[]Status<wbr>Item<wbr>Response</a></span>
    </dt>
    <dd>List of different status items.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="statuslist_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#statuslist_java" style="color: inherit; text-decoration: inherit;">status<wbr>List</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#statusitemresponse">List&lt;Status<wbr>Item<wbr>Response&gt;</a></span>
    </dt>
    <dd>List of different status items.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="statuslist_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#statuslist_nodejs" style="color: inherit; text-decoration: inherit;">status<wbr>List</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#statusitemresponse">Status<wbr>Item<wbr>Response[]</a></span>
    </dt>
    <dd>List of different status items.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="status_list_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#status_list_python" style="color: inherit; text-decoration: inherit;">status_<wbr>list</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#statusitemresponse">Sequence[Status<wbr>Item<wbr>Response]</a></span>
    </dt>
    <dd>List of different status items.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="statuslist_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#statuslist_yaml" style="color: inherit; text-decoration: inherit;">status<wbr>List</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#statusitemresponse">List&lt;Property Map&gt;</a></span>
    </dt>
    <dd>List of different status items.</dd></dl>
</pulumi-choosable>
</div>




## Supporting Types


<h4 id="statusitemresponse">Status<wbr>Item<wbr>Response</h4>



<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="statusname_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#statusname_csharp" style="color: inherit; text-decoration: inherit;">Status<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Status name - e.g. &quot;Active&quot;, &quot;Failed&quot;.</dd><dt class="property-optional"
            title="Optional">
        <span id="statusvalue_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#statusvalue_csharp" style="color: inherit; text-decoration: inherit;">Status<wbr>Value</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Status value. e.g. &quot;100&quot;, or &quot;100%&quot;.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="statusname_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#statusname_go" style="color: inherit; text-decoration: inherit;">Status<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Status name - e.g. &quot;Active&quot;, &quot;Failed&quot;.</dd><dt class="property-optional"
            title="Optional">
        <span id="statusvalue_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#statusvalue_go" style="color: inherit; text-decoration: inherit;">Status<wbr>Value</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Status value. e.g. &quot;100&quot;, or &quot;100%&quot;.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="statusname_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#statusname_java" style="color: inherit; text-decoration: inherit;">status<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Status name - e.g. &quot;Active&quot;, &quot;Failed&quot;.</dd><dt class="property-optional"
            title="Optional">
        <span id="statusvalue_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#statusvalue_java" style="color: inherit; text-decoration: inherit;">status<wbr>Value</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Status value. e.g. &quot;100&quot;, or &quot;100%&quot;.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="statusname_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#statusname_nodejs" style="color: inherit; text-decoration: inherit;">status<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Status name - e.g. &quot;Active&quot;, &quot;Failed&quot;.</dd><dt class="property-optional"
            title="Optional">
        <span id="statusvalue_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#statusvalue_nodejs" style="color: inherit; text-decoration: inherit;">status<wbr>Value</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Status value. e.g. &quot;100&quot;, or &quot;100%&quot;.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="status_name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#status_name_python" style="color: inherit; text-decoration: inherit;">status_<wbr>name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Status name - e.g. &quot;Active&quot;, &quot;Failed&quot;.</dd><dt class="property-optional"
            title="Optional">
        <span id="status_value_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#status_value_python" style="color: inherit; text-decoration: inherit;">status_<wbr>value</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Status value. e.g. &quot;100&quot;, or &quot;100%&quot;.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="statusname_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#statusname_yaml" style="color: inherit; text-decoration: inherit;">status<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Status name - e.g. &quot;Active&quot;, &quot;Failed&quot;.</dd><dt class="property-optional"
            title="Optional">
        <span id="statusvalue_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#statusvalue_yaml" style="color: inherit; text-decoration: inherit;">status<wbr>Value</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Status value. e.g. &quot;100&quot;, or &quot;100%&quot;.</dd></dl>
</pulumi-choosable>
</div>





<h2 id="package-details">Package Details</h2>
<dl class="package-details">
	<dt>Repository</dt>
	<dd><a href="https://github.com/pulumi/pulumi-azure-native">Azure Native pulumi/pulumi-azure-native</a></dd>
	<dt>License</dt>
	<dd>Apache-2.0</dd>
</dl>

