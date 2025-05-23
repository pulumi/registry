
---
title: "getAccessControlRecord"
title_tag: "azure-native.storsimple.getAccessControlRecord"
meta_desc: "Documentation for the azure-native.storsimple.getAccessControlRecord function with examples, input properties, output properties, and supporting types."
layout: api
no_edit_this_page: true
---



<!-- WARNING: this file was generated by Pulumi Docs Generator. -->
<!-- Do not edit by hand unless you're certain you know what you are doing! -->

Returns the properties of the specified access control record name.

Azure REST API version: 2017-06-01.




## Using getAccessControlRecord {#using}

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
><span class="k">function </span>getAccessControlRecord<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">GetAccessControlRecordArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Promise&lt;<span class="nx"><a href="#result">GetAccessControlRecordResult</a></span>></span
><span class="k">
function </span>getAccessControlRecordOutput<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">GetAccessControlRecordOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Output&lt;<span class="nx"><a href="#result">GetAccessControlRecordResult</a></span>></span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="python">
<div class="highlight"><pre class="chroma"><code class="language-python" data-lang="python"
><span class="k">def </span>get_access_control_record<span class="p">(</span><span class="nx">access_control_record_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                              <span class="nx">manager_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                              <span class="nx">resource_group_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                              <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>GetAccessControlRecordResult</span
><span class="k">
def </span>get_access_control_record_output<span class="p">(</span><span class="nx">access_control_record_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                              <span class="nx">manager_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                              <span class="nx">resource_group_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                              <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>Output[GetAccessControlRecordResult]</span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="go">
<div class="highlight"><pre class="chroma"><code class="language-go" data-lang="go"
><span class="k">func </span>LookupAccessControlRecord<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">LookupAccessControlRecordArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) (*<span class="nx"><a href="#result">LookupAccessControlRecordResult</a></span>, error)</span
><span class="k">
func </span>LookupAccessControlRecordOutput<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">LookupAccessControlRecordOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) LookupAccessControlRecordResultOutput</span
></code></pre></div>

&gt; Note: This function is named `LookupAccessControlRecord` in the Go SDK.

</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="csharp">
<div class="highlight"><pre class="chroma"><code class="language-csharp" data-lang="csharp"><span class="k">public static class </span><span class="nx">GetAccessControlRecord </span><span class="p">
{</span><span class="k">
    public static </span>Task&lt;<span class="nx"><a href="#result">GetAccessControlRecordResult</a></span>> <span class="p">InvokeAsync(</span><span class="nx">GetAccessControlRecordArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="k">
    public static </span>Output&lt;<span class="nx"><a href="#result">GetAccessControlRecordResult</a></span>> <span class="p">Invoke(</span><span class="nx">GetAccessControlRecordInvokeArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="p">
}</span></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="java">
<div class="highlight"><pre class="chroma"><code class="language-java" data-lang="java"><span class="k">public static CompletableFuture&lt;<span class="nx"><a href="#result">GetAccessControlRecordResult</a></span>> </span>getAccessControlRecord<span class="p">(</span><span class="nx">GetAccessControlRecordArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
<span class="k">public static Output&lt;<span class="nx"><a href="#result">GetAccessControlRecordResult</a></span>> </span>getAccessControlRecord<span class="p">(</span><span class="nx">GetAccessControlRecordArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
</code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="yaml">
<div class="highlight"><pre class="chroma"><code class="language-yaml" data-lang="yaml"><span class="k">fn::invoke:</span>
<span class="k">&nbsp;&nbsp;function:</span> azure-native:storsimple:getAccessControlRecord
<span class="k">&nbsp;&nbsp;arguments:</span>
<span class="c">&nbsp;&nbsp;&nbsp;&nbsp;# arguments dictionary</span></code></pre></div>
</pulumi-choosable>
</div>



The following arguments are supported:


<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="accesscontrolrecordname_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#accesscontrolrecordname_csharp" style="color: inherit; text-decoration: inherit;">Access<wbr>Control<wbr>Record<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Name of access control record to be fetched.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="managername_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#managername_csharp" style="color: inherit; text-decoration: inherit;">Manager<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The manager name</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_csharp" style="color: inherit; text-decoration: inherit;">Resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The resource group name</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="accesscontrolrecordname_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#accesscontrolrecordname_go" style="color: inherit; text-decoration: inherit;">Access<wbr>Control<wbr>Record<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Name of access control record to be fetched.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="managername_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#managername_go" style="color: inherit; text-decoration: inherit;">Manager<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The manager name</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_go" style="color: inherit; text-decoration: inherit;">Resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The resource group name</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="accesscontrolrecordname_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#accesscontrolrecordname_java" style="color: inherit; text-decoration: inherit;">access<wbr>Control<wbr>Record<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Name of access control record to be fetched.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="managername_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#managername_java" style="color: inherit; text-decoration: inherit;">manager<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The manager name</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_java" style="color: inherit; text-decoration: inherit;">resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The resource group name</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="accesscontrolrecordname_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#accesscontrolrecordname_nodejs" style="color: inherit; text-decoration: inherit;">access<wbr>Control<wbr>Record<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Name of access control record to be fetched.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="managername_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#managername_nodejs" style="color: inherit; text-decoration: inherit;">manager<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The manager name</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_nodejs" style="color: inherit; text-decoration: inherit;">resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The resource group name</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="access_control_record_name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#access_control_record_name_python" style="color: inherit; text-decoration: inherit;">access_<wbr>control_<wbr>record_<wbr>name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Name of access control record to be fetched.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="manager_name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#manager_name_python" style="color: inherit; text-decoration: inherit;">manager_<wbr>name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The manager name</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resource_group_name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resource_group_name_python" style="color: inherit; text-decoration: inherit;">resource_<wbr>group_<wbr>name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The resource group name</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="accesscontrolrecordname_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#accesscontrolrecordname_yaml" style="color: inherit; text-decoration: inherit;">access<wbr>Control<wbr>Record<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Name of access control record to be fetched.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="managername_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#managername_yaml" style="color: inherit; text-decoration: inherit;">manager<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The manager name</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_yaml" style="color: inherit; text-decoration: inherit;">resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The resource group name</dd></dl>
</pulumi-choosable>
</div>




## getAccessControlRecord Result {#result}

The following output properties are available:



<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-"
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
    <dd>The path ID that uniquely identifies the object.</dd><dt class="property-"
            title="">
        <span id="initiatorname_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#initiatorname_csharp" style="color: inherit; text-decoration: inherit;">Initiator<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The iSCSI initiator name (IQN).</dd><dt class="property-"
            title="">
        <span id="name_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_csharp" style="color: inherit; text-decoration: inherit;">Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the object.</dd><dt class="property-"
            title="">
        <span id="type_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_csharp" style="color: inherit; text-decoration: inherit;">Type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The hierarchical type of the object.</dd><dt class="property-"
            title="">
        <span id="volumecount_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#volumecount_csharp" style="color: inherit; text-decoration: inherit;">Volume<wbr>Count</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">int</span>
    </dt>
    <dd>The number of volumes using the access control record.</dd><dt class="property-"
            title="">
        <span id="kind_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#kind_csharp" style="color: inherit; text-decoration: inherit;">Kind</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The Kind of the object. Currently only Series8000 is supported</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-"
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
    <dd>The path ID that uniquely identifies the object.</dd><dt class="property-"
            title="">
        <span id="initiatorname_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#initiatorname_go" style="color: inherit; text-decoration: inherit;">Initiator<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The iSCSI initiator name (IQN).</dd><dt class="property-"
            title="">
        <span id="name_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_go" style="color: inherit; text-decoration: inherit;">Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the object.</dd><dt class="property-"
            title="">
        <span id="type_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_go" style="color: inherit; text-decoration: inherit;">Type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The hierarchical type of the object.</dd><dt class="property-"
            title="">
        <span id="volumecount_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#volumecount_go" style="color: inherit; text-decoration: inherit;">Volume<wbr>Count</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">int</span>
    </dt>
    <dd>The number of volumes using the access control record.</dd><dt class="property-"
            title="">
        <span id="kind_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#kind_go" style="color: inherit; text-decoration: inherit;">Kind</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The Kind of the object. Currently only Series8000 is supported</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-"
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
    <dd>The path ID that uniquely identifies the object.</dd><dt class="property-"
            title="">
        <span id="initiatorname_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#initiatorname_java" style="color: inherit; text-decoration: inherit;">initiator<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The iSCSI initiator name (IQN).</dd><dt class="property-"
            title="">
        <span id="name_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_java" style="color: inherit; text-decoration: inherit;">name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The name of the object.</dd><dt class="property-"
            title="">
        <span id="type_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_java" style="color: inherit; text-decoration: inherit;">type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The hierarchical type of the object.</dd><dt class="property-"
            title="">
        <span id="volumecount_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#volumecount_java" style="color: inherit; text-decoration: inherit;">volume<wbr>Count</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Integer</span>
    </dt>
    <dd>The number of volumes using the access control record.</dd><dt class="property-"
            title="">
        <span id="kind_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#kind_java" style="color: inherit; text-decoration: inherit;">kind</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The Kind of the object. Currently only Series8000 is supported</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-"
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
    <dd>The path ID that uniquely identifies the object.</dd><dt class="property-"
            title="">
        <span id="initiatorname_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#initiatorname_nodejs" style="color: inherit; text-decoration: inherit;">initiator<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The iSCSI initiator name (IQN).</dd><dt class="property-"
            title="">
        <span id="name_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_nodejs" style="color: inherit; text-decoration: inherit;">name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the object.</dd><dt class="property-"
            title="">
        <span id="type_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_nodejs" style="color: inherit; text-decoration: inherit;">type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The hierarchical type of the object.</dd><dt class="property-"
            title="">
        <span id="volumecount_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#volumecount_nodejs" style="color: inherit; text-decoration: inherit;">volume<wbr>Count</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">number</span>
    </dt>
    <dd>The number of volumes using the access control record.</dd><dt class="property-"
            title="">
        <span id="kind_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#kind_nodejs" style="color: inherit; text-decoration: inherit;">kind</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The Kind of the object. Currently only Series8000 is supported</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-"
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
    <dd>The path ID that uniquely identifies the object.</dd><dt class="property-"
            title="">
        <span id="initiator_name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#initiator_name_python" style="color: inherit; text-decoration: inherit;">initiator_<wbr>name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The iSCSI initiator name (IQN).</dd><dt class="property-"
            title="">
        <span id="name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_python" style="color: inherit; text-decoration: inherit;">name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The name of the object.</dd><dt class="property-"
            title="">
        <span id="type_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_python" style="color: inherit; text-decoration: inherit;">type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The hierarchical type of the object.</dd><dt class="property-"
            title="">
        <span id="volume_count_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#volume_count_python" style="color: inherit; text-decoration: inherit;">volume_<wbr>count</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">int</span>
    </dt>
    <dd>The number of volumes using the access control record.</dd><dt class="property-"
            title="">
        <span id="kind_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#kind_python" style="color: inherit; text-decoration: inherit;">kind</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The Kind of the object. Currently only Series8000 is supported</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-"
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
    <dd>The path ID that uniquely identifies the object.</dd><dt class="property-"
            title="">
        <span id="initiatorname_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#initiatorname_yaml" style="color: inherit; text-decoration: inherit;">initiator<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The iSCSI initiator name (IQN).</dd><dt class="property-"
            title="">
        <span id="name_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_yaml" style="color: inherit; text-decoration: inherit;">name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The name of the object.</dd><dt class="property-"
            title="">
        <span id="type_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_yaml" style="color: inherit; text-decoration: inherit;">type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The hierarchical type of the object.</dd><dt class="property-"
            title="">
        <span id="volumecount_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#volumecount_yaml" style="color: inherit; text-decoration: inherit;">volume<wbr>Count</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Number</span>
    </dt>
    <dd>The number of volumes using the access control record.</dd><dt class="property-"
            title="">
        <span id="kind_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#kind_yaml" style="color: inherit; text-decoration: inherit;">kind</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The Kind of the object. Currently only Series8000 is supported</dd></dl>
</pulumi-choosable>
</div>





<h2 id="package-details">Package Details</h2>
<dl class="package-details">
	<dt>Repository</dt>
	<dd><a href="https://github.com/pulumi/pulumi-azure-native">Azure Native pulumi/pulumi-azure-native</a></dd>
	<dt>License</dt>
	<dd>Apache-2.0</dd>
</dl>

