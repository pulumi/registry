
---
title: "getInferencePoolStatus"
title_tag: "azure-native.machinelearningservices.getInferencePoolStatus"
meta_desc: "Documentation for the azure-native.machinelearningservices.getInferencePoolStatus function with examples, input properties, output properties, and supporting types."
layout: api
no_edit_this_page: true
---



<!-- WARNING: this file was generated by Pulumi Docs Generator. -->
<!-- Do not edit by hand unless you're certain you know what you are doing! -->

Azure REST API version: 2024-01-01-preview.

Other available API versions: 2023-08-01-preview.

These versions are not included in the SDK but you can generate them as a local SDK package; see the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.




## Using getInferencePoolStatus {#using}

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
><span class="k">function </span>getInferencePoolStatus<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">GetInferencePoolStatusArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Promise&lt;<span class="nx"><a href="#result">GetInferencePoolStatusResult</a></span>></span
><span class="k">
function </span>getInferencePoolStatusOutput<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">GetInferencePoolStatusOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Output&lt;<span class="nx"><a href="#result">GetInferencePoolStatusResult</a></span>></span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="python">
<div class="highlight"><pre class="chroma"><code class="language-python" data-lang="python"
><span class="k">def </span>get_inference_pool_status<span class="p">(</span><span class="nx">inference_pool_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                              <span class="nx">resource_group_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                              <span class="nx">workspace_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                              <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>GetInferencePoolStatusResult</span
><span class="k">
def </span>get_inference_pool_status_output<span class="p">(</span><span class="nx">inference_pool_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                              <span class="nx">resource_group_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                              <span class="nx">workspace_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                              <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>Output[GetInferencePoolStatusResult]</span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="go">
<div class="highlight"><pre class="chroma"><code class="language-go" data-lang="go"
><span class="k">func </span>GetInferencePoolStatus<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">GetInferencePoolStatusArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) (*<span class="nx"><a href="#result">GetInferencePoolStatusResult</a></span>, error)</span
><span class="k">
func </span>GetInferencePoolStatusOutput<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">GetInferencePoolStatusOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) GetInferencePoolStatusResultOutput</span
></code></pre></div>

&gt; Note: This function is named `GetInferencePoolStatus` in the Go SDK.

</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="csharp">
<div class="highlight"><pre class="chroma"><code class="language-csharp" data-lang="csharp"><span class="k">public static class </span><span class="nx">GetInferencePoolStatus </span><span class="p">
{</span><span class="k">
    public static </span>Task&lt;<span class="nx"><a href="#result">GetInferencePoolStatusResult</a></span>> <span class="p">InvokeAsync(</span><span class="nx">GetInferencePoolStatusArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="k">
    public static </span>Output&lt;<span class="nx"><a href="#result">GetInferencePoolStatusResult</a></span>> <span class="p">Invoke(</span><span class="nx">GetInferencePoolStatusInvokeArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="p">
}</span></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="java">
<div class="highlight"><pre class="chroma"><code class="language-java" data-lang="java"><span class="k">public static CompletableFuture&lt;<span class="nx"><a href="#result">GetInferencePoolStatusResult</a></span>> </span>getInferencePoolStatus<span class="p">(</span><span class="nx">GetInferencePoolStatusArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
<span class="k">public static Output&lt;<span class="nx"><a href="#result">GetInferencePoolStatusResult</a></span>> </span>getInferencePoolStatus<span class="p">(</span><span class="nx">GetInferencePoolStatusArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
</code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="yaml">
<div class="highlight"><pre class="chroma"><code class="language-yaml" data-lang="yaml"><span class="k">fn::invoke:</span>
<span class="k">&nbsp;&nbsp;function:</span> azure-native:machinelearningservices:getInferencePoolStatus
<span class="k">&nbsp;&nbsp;arguments:</span>
<span class="c">&nbsp;&nbsp;&nbsp;&nbsp;# arguments dictionary</span></code></pre></div>
</pulumi-choosable>
</div>



The following arguments are supported:


<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="inferencepoolname_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#inferencepoolname_csharp" style="color: inherit; text-decoration: inherit;">Inference<wbr>Pool<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Name of InferencePool</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_csharp" style="color: inherit; text-decoration: inherit;">Resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the resource group. The name is case insensitive.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="workspacename_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#workspacename_csharp" style="color: inherit; text-decoration: inherit;">Workspace<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Name of Azure Machine Learning workspace.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="inferencepoolname_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#inferencepoolname_go" style="color: inherit; text-decoration: inherit;">Inference<wbr>Pool<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Name of InferencePool</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_go" style="color: inherit; text-decoration: inherit;">Resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the resource group. The name is case insensitive.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="workspacename_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#workspacename_go" style="color: inherit; text-decoration: inherit;">Workspace<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Name of Azure Machine Learning workspace.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="inferencepoolname_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#inferencepoolname_java" style="color: inherit; text-decoration: inherit;">inference<wbr>Pool<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Name of InferencePool</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_java" style="color: inherit; text-decoration: inherit;">resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The name of the resource group. The name is case insensitive.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="workspacename_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#workspacename_java" style="color: inherit; text-decoration: inherit;">workspace<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Name of Azure Machine Learning workspace.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="inferencepoolname_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#inferencepoolname_nodejs" style="color: inherit; text-decoration: inherit;">inference<wbr>Pool<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Name of InferencePool</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_nodejs" style="color: inherit; text-decoration: inherit;">resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the resource group. The name is case insensitive.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="workspacename_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#workspacename_nodejs" style="color: inherit; text-decoration: inherit;">workspace<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Name of Azure Machine Learning workspace.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="inference_pool_name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#inference_pool_name_python" style="color: inherit; text-decoration: inherit;">inference_<wbr>pool_<wbr>name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Name of InferencePool</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resource_group_name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resource_group_name_python" style="color: inherit; text-decoration: inherit;">resource_<wbr>group_<wbr>name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The name of the resource group. The name is case insensitive.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="workspace_name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#workspace_name_python" style="color: inherit; text-decoration: inherit;">workspace_<wbr>name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Name of Azure Machine Learning workspace.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="inferencepoolname_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#inferencepoolname_yaml" style="color: inherit; text-decoration: inherit;">inference<wbr>Pool<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Name of InferencePool</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_yaml" style="color: inherit; text-decoration: inherit;">resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The name of the resource group. The name is case insensitive.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="workspacename_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#workspacename_yaml" style="color: inherit; text-decoration: inherit;">workspace<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Name of Azure Machine Learning workspace.</dd></dl>
</pulumi-choosable>
</div>




## getInferencePoolStatus Result {#result}

The following output properties are available:



<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="actualcapacity_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#actualcapacity_csharp" style="color: inherit; text-decoration: inherit;">Actual<wbr>Capacity</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">int</span>
    </dt>
    <dd>Gets or sets the actual number of instances in the pool.</dd><dt class="property-"
            title="">
        <span id="groupcount_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#groupcount_csharp" style="color: inherit; text-decoration: inherit;">Group<wbr>Count</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">int</span>
    </dt>
    <dd>Gets or sets the actual number of groups in the pool.</dd><dt class="property-"
            title="">
        <span id="requestedcapacity_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#requestedcapacity_csharp" style="color: inherit; text-decoration: inherit;">Requested<wbr>Capacity</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">int</span>
    </dt>
    <dd>Gets or sets the requested number of instances for the pool.</dd><dt class="property-"
            title="">
        <span id="reservedcapacity_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#reservedcapacity_csharp" style="color: inherit; text-decoration: inherit;">Reserved<wbr>Capacity</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">int</span>
    </dt>
    <dd>Gets or sets the number of instances in the pool reserved by the system.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="actualcapacity_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#actualcapacity_go" style="color: inherit; text-decoration: inherit;">Actual<wbr>Capacity</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">int</span>
    </dt>
    <dd>Gets or sets the actual number of instances in the pool.</dd><dt class="property-"
            title="">
        <span id="groupcount_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#groupcount_go" style="color: inherit; text-decoration: inherit;">Group<wbr>Count</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">int</span>
    </dt>
    <dd>Gets or sets the actual number of groups in the pool.</dd><dt class="property-"
            title="">
        <span id="requestedcapacity_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#requestedcapacity_go" style="color: inherit; text-decoration: inherit;">Requested<wbr>Capacity</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">int</span>
    </dt>
    <dd>Gets or sets the requested number of instances for the pool.</dd><dt class="property-"
            title="">
        <span id="reservedcapacity_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#reservedcapacity_go" style="color: inherit; text-decoration: inherit;">Reserved<wbr>Capacity</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">int</span>
    </dt>
    <dd>Gets or sets the number of instances in the pool reserved by the system.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="actualcapacity_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#actualcapacity_java" style="color: inherit; text-decoration: inherit;">actual<wbr>Capacity</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Integer</span>
    </dt>
    <dd>Gets or sets the actual number of instances in the pool.</dd><dt class="property-"
            title="">
        <span id="groupcount_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#groupcount_java" style="color: inherit; text-decoration: inherit;">group<wbr>Count</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Integer</span>
    </dt>
    <dd>Gets or sets the actual number of groups in the pool.</dd><dt class="property-"
            title="">
        <span id="requestedcapacity_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#requestedcapacity_java" style="color: inherit; text-decoration: inherit;">requested<wbr>Capacity</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Integer</span>
    </dt>
    <dd>Gets or sets the requested number of instances for the pool.</dd><dt class="property-"
            title="">
        <span id="reservedcapacity_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#reservedcapacity_java" style="color: inherit; text-decoration: inherit;">reserved<wbr>Capacity</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Integer</span>
    </dt>
    <dd>Gets or sets the number of instances in the pool reserved by the system.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="actualcapacity_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#actualcapacity_nodejs" style="color: inherit; text-decoration: inherit;">actual<wbr>Capacity</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">number</span>
    </dt>
    <dd>Gets or sets the actual number of instances in the pool.</dd><dt class="property-"
            title="">
        <span id="groupcount_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#groupcount_nodejs" style="color: inherit; text-decoration: inherit;">group<wbr>Count</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">number</span>
    </dt>
    <dd>Gets or sets the actual number of groups in the pool.</dd><dt class="property-"
            title="">
        <span id="requestedcapacity_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#requestedcapacity_nodejs" style="color: inherit; text-decoration: inherit;">requested<wbr>Capacity</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">number</span>
    </dt>
    <dd>Gets or sets the requested number of instances for the pool.</dd><dt class="property-"
            title="">
        <span id="reservedcapacity_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#reservedcapacity_nodejs" style="color: inherit; text-decoration: inherit;">reserved<wbr>Capacity</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">number</span>
    </dt>
    <dd>Gets or sets the number of instances in the pool reserved by the system.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="actual_capacity_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#actual_capacity_python" style="color: inherit; text-decoration: inherit;">actual_<wbr>capacity</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">int</span>
    </dt>
    <dd>Gets or sets the actual number of instances in the pool.</dd><dt class="property-"
            title="">
        <span id="group_count_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#group_count_python" style="color: inherit; text-decoration: inherit;">group_<wbr>count</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">int</span>
    </dt>
    <dd>Gets or sets the actual number of groups in the pool.</dd><dt class="property-"
            title="">
        <span id="requested_capacity_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#requested_capacity_python" style="color: inherit; text-decoration: inherit;">requested_<wbr>capacity</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">int</span>
    </dt>
    <dd>Gets or sets the requested number of instances for the pool.</dd><dt class="property-"
            title="">
        <span id="reserved_capacity_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#reserved_capacity_python" style="color: inherit; text-decoration: inherit;">reserved_<wbr>capacity</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">int</span>
    </dt>
    <dd>Gets or sets the number of instances in the pool reserved by the system.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="actualcapacity_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#actualcapacity_yaml" style="color: inherit; text-decoration: inherit;">actual<wbr>Capacity</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Number</span>
    </dt>
    <dd>Gets or sets the actual number of instances in the pool.</dd><dt class="property-"
            title="">
        <span id="groupcount_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#groupcount_yaml" style="color: inherit; text-decoration: inherit;">group<wbr>Count</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Number</span>
    </dt>
    <dd>Gets or sets the actual number of groups in the pool.</dd><dt class="property-"
            title="">
        <span id="requestedcapacity_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#requestedcapacity_yaml" style="color: inherit; text-decoration: inherit;">requested<wbr>Capacity</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Number</span>
    </dt>
    <dd>Gets or sets the requested number of instances for the pool.</dd><dt class="property-"
            title="">
        <span id="reservedcapacity_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#reservedcapacity_yaml" style="color: inherit; text-decoration: inherit;">reserved<wbr>Capacity</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Number</span>
    </dt>
    <dd>Gets or sets the number of instances in the pool reserved by the system.</dd></dl>
</pulumi-choosable>
</div>





<h2 id="package-details">Package Details</h2>
<dl class="package-details">
	<dt>Repository</dt>
	<dd><a href="https://github.com/pulumi/pulumi-azure-native">Azure Native pulumi/pulumi-azure-native</a></dd>
	<dt>License</dt>
	<dd>Apache-2.0</dd>
</dl>

