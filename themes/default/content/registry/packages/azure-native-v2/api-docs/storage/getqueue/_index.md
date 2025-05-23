
---
title: "getQueue"
title_tag: "azure-native.storage.getQueue"
meta_desc: "Documentation for the azure-native.storage.getQueue function with examples, input properties, output properties, and supporting types."
layout: api
no_edit_this_page: true
---



<!-- WARNING: this file was generated by Pulumi Docs Generator. -->
<!-- Do not edit by hand unless you're certain you know what you are doing! -->

Gets the queue with the specified queue name, under the specified account if it exists.

Azure REST API version: 2024-01-01.

Other available API versions: 2022-09-01.

These versions are not included in the SDK but you can generate them as a local SDK package; see the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.




## Using getQueue {#using}

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
><span class="k">function </span>getQueue<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">GetQueueArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Promise&lt;<span class="nx"><a href="#result">GetQueueResult</a></span>></span
><span class="k">
function </span>getQueueOutput<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">GetQueueOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Output&lt;<span class="nx"><a href="#result">GetQueueResult</a></span>></span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="python">
<div class="highlight"><pre class="chroma"><code class="language-python" data-lang="python"
><span class="k">def </span>get_queue<span class="p">(</span><span class="nx">account_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
              <span class="nx">queue_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
              <span class="nx">resource_group_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
              <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>GetQueueResult</span
><span class="k">
def </span>get_queue_output<span class="p">(</span><span class="nx">account_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
              <span class="nx">queue_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
              <span class="nx">resource_group_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
              <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>Output[GetQueueResult]</span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="go">
<div class="highlight"><pre class="chroma"><code class="language-go" data-lang="go"
><span class="k">func </span>LookupQueue<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">LookupQueueArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) (*<span class="nx"><a href="#result">LookupQueueResult</a></span>, error)</span
><span class="k">
func </span>LookupQueueOutput<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">LookupQueueOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) LookupQueueResultOutput</span
></code></pre></div>

&gt; Note: This function is named `LookupQueue` in the Go SDK.

</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="csharp">
<div class="highlight"><pre class="chroma"><code class="language-csharp" data-lang="csharp"><span class="k">public static class </span><span class="nx">GetQueue </span><span class="p">
{</span><span class="k">
    public static </span>Task&lt;<span class="nx"><a href="#result">GetQueueResult</a></span>> <span class="p">InvokeAsync(</span><span class="nx">GetQueueArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="k">
    public static </span>Output&lt;<span class="nx"><a href="#result">GetQueueResult</a></span>> <span class="p">Invoke(</span><span class="nx">GetQueueInvokeArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="p">
}</span></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="java">
<div class="highlight"><pre class="chroma"><code class="language-java" data-lang="java"><span class="k">public static CompletableFuture&lt;<span class="nx"><a href="#result">GetQueueResult</a></span>> </span>getQueue<span class="p">(</span><span class="nx">GetQueueArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
<span class="k">public static Output&lt;<span class="nx"><a href="#result">GetQueueResult</a></span>> </span>getQueue<span class="p">(</span><span class="nx">GetQueueArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
</code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="yaml">
<div class="highlight"><pre class="chroma"><code class="language-yaml" data-lang="yaml"><span class="k">fn::invoke:</span>
<span class="k">&nbsp;&nbsp;function:</span> azure-native:storage:getQueue
<span class="k">&nbsp;&nbsp;arguments:</span>
<span class="c">&nbsp;&nbsp;&nbsp;&nbsp;# arguments dictionary</span></code></pre></div>
</pulumi-choosable>
</div>



The following arguments are supported:


<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="accountname_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#accountname_csharp" style="color: inherit; text-decoration: inherit;">Account<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="queuename_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#queuename_csharp" style="color: inherit; text-decoration: inherit;">Queue<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>A queue name must be unique within a storage account and must be between 3 and 63 characters.The name must comprise of lowercase alphanumeric and dash(-) characters only, it should begin and end with an alphanumeric character and it cannot have two consecutive dash(-) characters.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_csharp" style="color: inherit; text-decoration: inherit;">Resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the resource group within the user's subscription. The name is case insensitive.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="accountname_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#accountname_go" style="color: inherit; text-decoration: inherit;">Account<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="queuename_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#queuename_go" style="color: inherit; text-decoration: inherit;">Queue<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>A queue name must be unique within a storage account and must be between 3 and 63 characters.The name must comprise of lowercase alphanumeric and dash(-) characters only, it should begin and end with an alphanumeric character and it cannot have two consecutive dash(-) characters.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_go" style="color: inherit; text-decoration: inherit;">Resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the resource group within the user's subscription. The name is case insensitive.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="accountname_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#accountname_java" style="color: inherit; text-decoration: inherit;">account<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="queuename_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#queuename_java" style="color: inherit; text-decoration: inherit;">queue<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>A queue name must be unique within a storage account and must be between 3 and 63 characters.The name must comprise of lowercase alphanumeric and dash(-) characters only, it should begin and end with an alphanumeric character and it cannot have two consecutive dash(-) characters.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_java" style="color: inherit; text-decoration: inherit;">resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The name of the resource group within the user's subscription. The name is case insensitive.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="accountname_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#accountname_nodejs" style="color: inherit; text-decoration: inherit;">account<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="queuename_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#queuename_nodejs" style="color: inherit; text-decoration: inherit;">queue<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>A queue name must be unique within a storage account and must be between 3 and 63 characters.The name must comprise of lowercase alphanumeric and dash(-) characters only, it should begin and end with an alphanumeric character and it cannot have two consecutive dash(-) characters.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_nodejs" style="color: inherit; text-decoration: inherit;">resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the resource group within the user's subscription. The name is case insensitive.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="account_name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#account_name_python" style="color: inherit; text-decoration: inherit;">account_<wbr>name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="queue_name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#queue_name_python" style="color: inherit; text-decoration: inherit;">queue_<wbr>name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>A queue name must be unique within a storage account and must be between 3 and 63 characters.The name must comprise of lowercase alphanumeric and dash(-) characters only, it should begin and end with an alphanumeric character and it cannot have two consecutive dash(-) characters.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resource_group_name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resource_group_name_python" style="color: inherit; text-decoration: inherit;">resource_<wbr>group_<wbr>name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The name of the resource group within the user's subscription. The name is case insensitive.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="accountname_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#accountname_yaml" style="color: inherit; text-decoration: inherit;">account<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="queuename_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#queuename_yaml" style="color: inherit; text-decoration: inherit;">queue<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>A queue name must be unique within a storage account and must be between 3 and 63 characters.The name must comprise of lowercase alphanumeric and dash(-) characters only, it should begin and end with an alphanumeric character and it cannot have two consecutive dash(-) characters.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_yaml" style="color: inherit; text-decoration: inherit;">resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The name of the resource group within the user's subscription. The name is case insensitive.</dd></dl>
</pulumi-choosable>
</div>




## getQueue Result {#result}

The following output properties are available:



<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="approximatemessagecount_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#approximatemessagecount_csharp" style="color: inherit; text-decoration: inherit;">Approximate<wbr>Message<wbr>Count</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">int</span>
    </dt>
    <dd>Integer indicating an approximate number of messages in the queue. This number is not lower than the actual number of messages in the queue, but could be higher.</dd><dt class="property-"
            title="">
        <span id="azureapiversion_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#azureapiversion_csharp" style="color: inherit; text-decoration: inherit;">Azure<wbr>Api<wbr>Version</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The Azure API version of the resource.</dd><dt class="property-"
            title="">
        <span id="id_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_csharp" style="color: inherit; text-decoration: inherit;">Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}</dd><dt class="property-"
            title="">
        <span id="name_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_csharp" style="color: inherit; text-decoration: inherit;">Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the resource</dd><dt class="property-"
            title="">
        <span id="type_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_csharp" style="color: inherit; text-decoration: inherit;">Type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The type of the resource. E.g. &quot;Microsoft.Compute/virtualMachines&quot; or &quot;Microsoft.Storage/storageAccounts&quot;</dd><dt class="property-"
            title="">
        <span id="metadata_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#metadata_csharp" style="color: inherit; text-decoration: inherit;">Metadata</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Dictionary&lt;string, string&gt;</span>
    </dt>
    <dd>A name-value pair that represents queue metadata.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="approximatemessagecount_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#approximatemessagecount_go" style="color: inherit; text-decoration: inherit;">Approximate<wbr>Message<wbr>Count</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">int</span>
    </dt>
    <dd>Integer indicating an approximate number of messages in the queue. This number is not lower than the actual number of messages in the queue, but could be higher.</dd><dt class="property-"
            title="">
        <span id="azureapiversion_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#azureapiversion_go" style="color: inherit; text-decoration: inherit;">Azure<wbr>Api<wbr>Version</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The Azure API version of the resource.</dd><dt class="property-"
            title="">
        <span id="id_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_go" style="color: inherit; text-decoration: inherit;">Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}</dd><dt class="property-"
            title="">
        <span id="name_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_go" style="color: inherit; text-decoration: inherit;">Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the resource</dd><dt class="property-"
            title="">
        <span id="type_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_go" style="color: inherit; text-decoration: inherit;">Type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The type of the resource. E.g. &quot;Microsoft.Compute/virtualMachines&quot; or &quot;Microsoft.Storage/storageAccounts&quot;</dd><dt class="property-"
            title="">
        <span id="metadata_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#metadata_go" style="color: inherit; text-decoration: inherit;">Metadata</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">map[string]string</span>
    </dt>
    <dd>A name-value pair that represents queue metadata.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="approximatemessagecount_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#approximatemessagecount_java" style="color: inherit; text-decoration: inherit;">approximate<wbr>Message<wbr>Count</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Integer</span>
    </dt>
    <dd>Integer indicating an approximate number of messages in the queue. This number is not lower than the actual number of messages in the queue, but could be higher.</dd><dt class="property-"
            title="">
        <span id="azureapiversion_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#azureapiversion_java" style="color: inherit; text-decoration: inherit;">azure<wbr>Api<wbr>Version</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The Azure API version of the resource.</dd><dt class="property-"
            title="">
        <span id="id_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_java" style="color: inherit; text-decoration: inherit;">id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}</dd><dt class="property-"
            title="">
        <span id="name_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_java" style="color: inherit; text-decoration: inherit;">name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The name of the resource</dd><dt class="property-"
            title="">
        <span id="type_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_java" style="color: inherit; text-decoration: inherit;">type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The type of the resource. E.g. &quot;Microsoft.Compute/virtualMachines&quot; or &quot;Microsoft.Storage/storageAccounts&quot;</dd><dt class="property-"
            title="">
        <span id="metadata_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#metadata_java" style="color: inherit; text-decoration: inherit;">metadata</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Map&lt;String,String&gt;</span>
    </dt>
    <dd>A name-value pair that represents queue metadata.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="approximatemessagecount_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#approximatemessagecount_nodejs" style="color: inherit; text-decoration: inherit;">approximate<wbr>Message<wbr>Count</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">number</span>
    </dt>
    <dd>Integer indicating an approximate number of messages in the queue. This number is not lower than the actual number of messages in the queue, but could be higher.</dd><dt class="property-"
            title="">
        <span id="azureapiversion_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#azureapiversion_nodejs" style="color: inherit; text-decoration: inherit;">azure<wbr>Api<wbr>Version</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The Azure API version of the resource.</dd><dt class="property-"
            title="">
        <span id="id_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_nodejs" style="color: inherit; text-decoration: inherit;">id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}</dd><dt class="property-"
            title="">
        <span id="name_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_nodejs" style="color: inherit; text-decoration: inherit;">name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the resource</dd><dt class="property-"
            title="">
        <span id="type_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_nodejs" style="color: inherit; text-decoration: inherit;">type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The type of the resource. E.g. &quot;Microsoft.Compute/virtualMachines&quot; or &quot;Microsoft.Storage/storageAccounts&quot;</dd><dt class="property-"
            title="">
        <span id="metadata_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#metadata_nodejs" style="color: inherit; text-decoration: inherit;">metadata</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">{[key: string]: string}</span>
    </dt>
    <dd>A name-value pair that represents queue metadata.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="approximate_message_count_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#approximate_message_count_python" style="color: inherit; text-decoration: inherit;">approximate_<wbr>message_<wbr>count</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">int</span>
    </dt>
    <dd>Integer indicating an approximate number of messages in the queue. This number is not lower than the actual number of messages in the queue, but could be higher.</dd><dt class="property-"
            title="">
        <span id="azure_api_version_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#azure_api_version_python" style="color: inherit; text-decoration: inherit;">azure_<wbr>api_<wbr>version</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The Azure API version of the resource.</dd><dt class="property-"
            title="">
        <span id="id_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_python" style="color: inherit; text-decoration: inherit;">id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}</dd><dt class="property-"
            title="">
        <span id="name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_python" style="color: inherit; text-decoration: inherit;">name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The name of the resource</dd><dt class="property-"
            title="">
        <span id="type_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_python" style="color: inherit; text-decoration: inherit;">type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The type of the resource. E.g. &quot;Microsoft.Compute/virtualMachines&quot; or &quot;Microsoft.Storage/storageAccounts&quot;</dd><dt class="property-"
            title="">
        <span id="metadata_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#metadata_python" style="color: inherit; text-decoration: inherit;">metadata</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Mapping[str, str]</span>
    </dt>
    <dd>A name-value pair that represents queue metadata.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="approximatemessagecount_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#approximatemessagecount_yaml" style="color: inherit; text-decoration: inherit;">approximate<wbr>Message<wbr>Count</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Number</span>
    </dt>
    <dd>Integer indicating an approximate number of messages in the queue. This number is not lower than the actual number of messages in the queue, but could be higher.</dd><dt class="property-"
            title="">
        <span id="azureapiversion_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#azureapiversion_yaml" style="color: inherit; text-decoration: inherit;">azure<wbr>Api<wbr>Version</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The Azure API version of the resource.</dd><dt class="property-"
            title="">
        <span id="id_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_yaml" style="color: inherit; text-decoration: inherit;">id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}</dd><dt class="property-"
            title="">
        <span id="name_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_yaml" style="color: inherit; text-decoration: inherit;">name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The name of the resource</dd><dt class="property-"
            title="">
        <span id="type_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_yaml" style="color: inherit; text-decoration: inherit;">type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The type of the resource. E.g. &quot;Microsoft.Compute/virtualMachines&quot; or &quot;Microsoft.Storage/storageAccounts&quot;</dd><dt class="property-"
            title="">
        <span id="metadata_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#metadata_yaml" style="color: inherit; text-decoration: inherit;">metadata</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Map&lt;String&gt;</span>
    </dt>
    <dd>A name-value pair that represents queue metadata.</dd></dl>
</pulumi-choosable>
</div>





<h2 id="package-details">Package Details</h2>
<dl class="package-details">
	<dt>Repository</dt>
	<dd><a href="https://github.com/pulumi/pulumi-azure-native">Azure Native pulumi/pulumi-azure-native</a></dd>
	<dt>License</dt>
	<dd>Apache-2.0</dd>
</dl>

