
---
title: "getWorkspaceCollection"
title_tag: "azure-native.powerbi.getWorkspaceCollection"
meta_desc: "Documentation for the azure-native.powerbi.getWorkspaceCollection function with examples, input properties, output properties, and supporting types."
layout: api
no_edit_this_page: true
---



<!-- WARNING: this file was generated by Pulumi Docs Generator. -->
<!-- Do not edit by hand unless you're certain you know what you are doing! -->

Retrieves an existing Power BI Workspace Collection.

Azure REST API version: 2016-01-29.




## Using getWorkspaceCollection {#using}

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
><span class="k">function </span>getWorkspaceCollection<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">GetWorkspaceCollectionArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Promise&lt;<span class="nx"><a href="#result">GetWorkspaceCollectionResult</a></span>></span
><span class="k">
function </span>getWorkspaceCollectionOutput<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">GetWorkspaceCollectionOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Output&lt;<span class="nx"><a href="#result">GetWorkspaceCollectionResult</a></span>></span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="python">
<div class="highlight"><pre class="chroma"><code class="language-python" data-lang="python"
><span class="k">def </span>get_workspace_collection<span class="p">(</span><span class="nx">resource_group_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                             <span class="nx">workspace_collection_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                             <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>GetWorkspaceCollectionResult</span
><span class="k">
def </span>get_workspace_collection_output<span class="p">(</span><span class="nx">resource_group_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                             <span class="nx">workspace_collection_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                             <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>Output[GetWorkspaceCollectionResult]</span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="go">
<div class="highlight"><pre class="chroma"><code class="language-go" data-lang="go"
><span class="k">func </span>LookupWorkspaceCollection<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">LookupWorkspaceCollectionArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) (*<span class="nx"><a href="#result">LookupWorkspaceCollectionResult</a></span>, error)</span
><span class="k">
func </span>LookupWorkspaceCollectionOutput<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">LookupWorkspaceCollectionOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) LookupWorkspaceCollectionResultOutput</span
></code></pre></div>

&gt; Note: This function is named `LookupWorkspaceCollection` in the Go SDK.

</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="csharp">
<div class="highlight"><pre class="chroma"><code class="language-csharp" data-lang="csharp"><span class="k">public static class </span><span class="nx">GetWorkspaceCollection </span><span class="p">
{</span><span class="k">
    public static </span>Task&lt;<span class="nx"><a href="#result">GetWorkspaceCollectionResult</a></span>> <span class="p">InvokeAsync(</span><span class="nx">GetWorkspaceCollectionArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="k">
    public static </span>Output&lt;<span class="nx"><a href="#result">GetWorkspaceCollectionResult</a></span>> <span class="p">Invoke(</span><span class="nx">GetWorkspaceCollectionInvokeArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="p">
}</span></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="java">
<div class="highlight"><pre class="chroma"><code class="language-java" data-lang="java"><span class="k">public static CompletableFuture&lt;<span class="nx"><a href="#result">GetWorkspaceCollectionResult</a></span>> </span>getWorkspaceCollection<span class="p">(</span><span class="nx">GetWorkspaceCollectionArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
<span class="k">public static Output&lt;<span class="nx"><a href="#result">GetWorkspaceCollectionResult</a></span>> </span>getWorkspaceCollection<span class="p">(</span><span class="nx">GetWorkspaceCollectionArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
</code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="yaml">
<div class="highlight"><pre class="chroma"><code class="language-yaml" data-lang="yaml"><span class="k">fn::invoke:</span>
<span class="k">&nbsp;&nbsp;function:</span> azure-native:powerbi:getWorkspaceCollection
<span class="k">&nbsp;&nbsp;arguments:</span>
<span class="c">&nbsp;&nbsp;&nbsp;&nbsp;# arguments dictionary</span></code></pre></div>
</pulumi-choosable>
</div>



The following arguments are supported:


<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_csharp" style="color: inherit; text-decoration: inherit;">Resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Azure resource group</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="workspacecollectionname_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#workspacecollectionname_csharp" style="color: inherit; text-decoration: inherit;">Workspace<wbr>Collection<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Power BI Embedded Workspace Collection name</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_go" style="color: inherit; text-decoration: inherit;">Resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Azure resource group</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="workspacecollectionname_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#workspacecollectionname_go" style="color: inherit; text-decoration: inherit;">Workspace<wbr>Collection<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Power BI Embedded Workspace Collection name</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_java" style="color: inherit; text-decoration: inherit;">resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Azure resource group</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="workspacecollectionname_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#workspacecollectionname_java" style="color: inherit; text-decoration: inherit;">workspace<wbr>Collection<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Power BI Embedded Workspace Collection name</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_nodejs" style="color: inherit; text-decoration: inherit;">resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Azure resource group</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="workspacecollectionname_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#workspacecollectionname_nodejs" style="color: inherit; text-decoration: inherit;">workspace<wbr>Collection<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Power BI Embedded Workspace Collection name</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="resource_group_name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resource_group_name_python" style="color: inherit; text-decoration: inherit;">resource_<wbr>group_<wbr>name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Azure resource group</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="workspace_collection_name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#workspace_collection_name_python" style="color: inherit; text-decoration: inherit;">workspace_<wbr>collection_<wbr>name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Power BI Embedded Workspace Collection name</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_yaml" style="color: inherit; text-decoration: inherit;">resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Azure resource group</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="workspacecollectionname_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#workspacecollectionname_yaml" style="color: inherit; text-decoration: inherit;">workspace<wbr>Collection<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Power BI Embedded Workspace Collection name</dd></dl>
</pulumi-choosable>
</div>




## getWorkspaceCollection Result {#result}

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
        <span id="properties_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#properties_csharp" style="color: inherit; text-decoration: inherit;">Properties</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">object</span>
    </dt>
    <dd>Properties</dd><dt class="property-"
            title="">
        <span id="id_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_csharp" style="color: inherit; text-decoration: inherit;">Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Resource id</dd><dt class="property-"
            title="">
        <span id="location_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#location_csharp" style="color: inherit; text-decoration: inherit;">Location</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Azure location</dd><dt class="property-"
            title="">
        <span id="name_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_csharp" style="color: inherit; text-decoration: inherit;">Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Workspace collection name</dd><dt class="property-"
            title="">
        <span id="sku_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#sku_csharp" style="color: inherit; text-decoration: inherit;">Sku</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#azureskuresponse">Pulumi.<wbr>Azure<wbr>Native.<wbr>Power<wbr>BI.<wbr>Outputs.<wbr>Azure<wbr>Sku<wbr>Response</a></span>
    </dt>
    <dd></dd><dt class="property-"
            title="">
        <span id="tags_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#tags_csharp" style="color: inherit; text-decoration: inherit;">Tags</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Dictionary&lt;string, string&gt;</span>
    </dt>
    <dd></dd><dt class="property-"
            title="">
        <span id="type_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_csharp" style="color: inherit; text-decoration: inherit;">Type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Resource type</dd></dl>
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
        <span id="properties_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#properties_go" style="color: inherit; text-decoration: inherit;">Properties</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">interface{}</span>
    </dt>
    <dd>Properties</dd><dt class="property-"
            title="">
        <span id="id_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_go" style="color: inherit; text-decoration: inherit;">Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Resource id</dd><dt class="property-"
            title="">
        <span id="location_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#location_go" style="color: inherit; text-decoration: inherit;">Location</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Azure location</dd><dt class="property-"
            title="">
        <span id="name_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_go" style="color: inherit; text-decoration: inherit;">Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Workspace collection name</dd><dt class="property-"
            title="">
        <span id="sku_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#sku_go" style="color: inherit; text-decoration: inherit;">Sku</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#azureskuresponse">Azure<wbr>Sku<wbr>Response</a></span>
    </dt>
    <dd></dd><dt class="property-"
            title="">
        <span id="tags_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#tags_go" style="color: inherit; text-decoration: inherit;">Tags</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">map[string]string</span>
    </dt>
    <dd></dd><dt class="property-"
            title="">
        <span id="type_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_go" style="color: inherit; text-decoration: inherit;">Type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Resource type</dd></dl>
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
        <span id="properties_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#properties_java" style="color: inherit; text-decoration: inherit;">properties</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Object</span>
    </dt>
    <dd>Properties</dd><dt class="property-"
            title="">
        <span id="id_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_java" style="color: inherit; text-decoration: inherit;">id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Resource id</dd><dt class="property-"
            title="">
        <span id="location_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#location_java" style="color: inherit; text-decoration: inherit;">location</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Azure location</dd><dt class="property-"
            title="">
        <span id="name_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_java" style="color: inherit; text-decoration: inherit;">name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Workspace collection name</dd><dt class="property-"
            title="">
        <span id="sku_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#sku_java" style="color: inherit; text-decoration: inherit;">sku</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#azureskuresponse">Azure<wbr>Sku<wbr>Response</a></span>
    </dt>
    <dd></dd><dt class="property-"
            title="">
        <span id="tags_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#tags_java" style="color: inherit; text-decoration: inherit;">tags</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Map&lt;String,String&gt;</span>
    </dt>
    <dd></dd><dt class="property-"
            title="">
        <span id="type_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_java" style="color: inherit; text-decoration: inherit;">type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Resource type</dd></dl>
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
        <span id="properties_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#properties_nodejs" style="color: inherit; text-decoration: inherit;">properties</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">any</span>
    </dt>
    <dd>Properties</dd><dt class="property-"
            title="">
        <span id="id_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_nodejs" style="color: inherit; text-decoration: inherit;">id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Resource id</dd><dt class="property-"
            title="">
        <span id="location_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#location_nodejs" style="color: inherit; text-decoration: inherit;">location</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Azure location</dd><dt class="property-"
            title="">
        <span id="name_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_nodejs" style="color: inherit; text-decoration: inherit;">name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Workspace collection name</dd><dt class="property-"
            title="">
        <span id="sku_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#sku_nodejs" style="color: inherit; text-decoration: inherit;">sku</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#azureskuresponse">Azure<wbr>Sku<wbr>Response</a></span>
    </dt>
    <dd></dd><dt class="property-"
            title="">
        <span id="tags_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#tags_nodejs" style="color: inherit; text-decoration: inherit;">tags</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">{[key: string]: string}</span>
    </dt>
    <dd></dd><dt class="property-"
            title="">
        <span id="type_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_nodejs" style="color: inherit; text-decoration: inherit;">type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Resource type</dd></dl>
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
        <span id="properties_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#properties_python" style="color: inherit; text-decoration: inherit;">properties</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Any</span>
    </dt>
    <dd>Properties</dd><dt class="property-"
            title="">
        <span id="id_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_python" style="color: inherit; text-decoration: inherit;">id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Resource id</dd><dt class="property-"
            title="">
        <span id="location_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#location_python" style="color: inherit; text-decoration: inherit;">location</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Azure location</dd><dt class="property-"
            title="">
        <span id="name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_python" style="color: inherit; text-decoration: inherit;">name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Workspace collection name</dd><dt class="property-"
            title="">
        <span id="sku_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#sku_python" style="color: inherit; text-decoration: inherit;">sku</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#azureskuresponse">Azure<wbr>Sku<wbr>Response</a></span>
    </dt>
    <dd></dd><dt class="property-"
            title="">
        <span id="tags_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#tags_python" style="color: inherit; text-decoration: inherit;">tags</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Mapping[str, str]</span>
    </dt>
    <dd></dd><dt class="property-"
            title="">
        <span id="type_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_python" style="color: inherit; text-decoration: inherit;">type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Resource type</dd></dl>
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
        <span id="properties_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#properties_yaml" style="color: inherit; text-decoration: inherit;">properties</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Any</span>
    </dt>
    <dd>Properties</dd><dt class="property-"
            title="">
        <span id="id_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_yaml" style="color: inherit; text-decoration: inherit;">id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Resource id</dd><dt class="property-"
            title="">
        <span id="location_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#location_yaml" style="color: inherit; text-decoration: inherit;">location</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Azure location</dd><dt class="property-"
            title="">
        <span id="name_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_yaml" style="color: inherit; text-decoration: inherit;">name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Workspace collection name</dd><dt class="property-"
            title="">
        <span id="sku_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#sku_yaml" style="color: inherit; text-decoration: inherit;">sku</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#azureskuresponse">Property Map</a></span>
    </dt>
    <dd></dd><dt class="property-"
            title="">
        <span id="tags_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#tags_yaml" style="color: inherit; text-decoration: inherit;">tags</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Map&lt;String&gt;</span>
    </dt>
    <dd></dd><dt class="property-"
            title="">
        <span id="type_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_yaml" style="color: inherit; text-decoration: inherit;">type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Resource type</dd></dl>
</pulumi-choosable>
</div>




## Supporting Types


<h4 id="azureskuresponse">Azure<wbr>Sku<wbr>Response</h4>



<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="name_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_csharp" style="color: inherit; text-decoration: inherit;">Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>SKU name</dd><dt class="property-required"
            title="Required">
        <span id="tier_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#tier_csharp" style="color: inherit; text-decoration: inherit;">Tier</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>SKU tier</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="name_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_go" style="color: inherit; text-decoration: inherit;">Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>SKU name</dd><dt class="property-required"
            title="Required">
        <span id="tier_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#tier_go" style="color: inherit; text-decoration: inherit;">Tier</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>SKU tier</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="name_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_java" style="color: inherit; text-decoration: inherit;">name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>SKU name</dd><dt class="property-required"
            title="Required">
        <span id="tier_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#tier_java" style="color: inherit; text-decoration: inherit;">tier</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>SKU tier</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="name_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_nodejs" style="color: inherit; text-decoration: inherit;">name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>SKU name</dd><dt class="property-required"
            title="Required">
        <span id="tier_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#tier_nodejs" style="color: inherit; text-decoration: inherit;">tier</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>SKU tier</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_python" style="color: inherit; text-decoration: inherit;">name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>SKU name</dd><dt class="property-required"
            title="Required">
        <span id="tier_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#tier_python" style="color: inherit; text-decoration: inherit;">tier</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>SKU tier</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="name_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_yaml" style="color: inherit; text-decoration: inherit;">name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>SKU name</dd><dt class="property-required"
            title="Required">
        <span id="tier_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#tier_yaml" style="color: inherit; text-decoration: inherit;">tier</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>SKU tier</dd></dl>
</pulumi-choosable>
</div>





<h2 id="package-details">Package Details</h2>
<dl class="package-details">
	<dt>Repository</dt>
	<dd><a href="https://github.com/pulumi/pulumi-azure-native">Azure Native pulumi/pulumi-azure-native</a></dd>
	<dt>License</dt>
	<dd>Apache-2.0</dd>
</dl>

