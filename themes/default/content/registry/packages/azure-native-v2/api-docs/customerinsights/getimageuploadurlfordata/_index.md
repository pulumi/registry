
---
title: "getImageUploadUrlForData"
title_tag: "azure-native.customerinsights.getImageUploadUrlForData"
meta_desc: "Documentation for the azure-native.customerinsights.getImageUploadUrlForData function with examples, input properties, output properties, and supporting types."
layout: api
no_edit_this_page: true
---



<!-- WARNING: this file was generated by Pulumi Docs Generator. -->
<!-- Do not edit by hand unless you're certain you know what you are doing! -->

Gets data image upload URL.

Azure REST API version: 2017-04-26.




## Using getImageUploadUrlForData {#using}

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
><span class="k">function </span>getImageUploadUrlForData<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">GetImageUploadUrlForDataArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Promise&lt;<span class="nx"><a href="#result">GetImageUploadUrlForDataResult</a></span>></span
><span class="k">
function </span>getImageUploadUrlForDataOutput<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">GetImageUploadUrlForDataOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Output&lt;<span class="nx"><a href="#result">GetImageUploadUrlForDataResult</a></span>></span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="python">
<div class="highlight"><pre class="chroma"><code class="language-python" data-lang="python"
><span class="k">def </span>get_image_upload_url_for_data<span class="p">(</span><span class="nx">entity_type</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                                  <span class="nx">entity_type_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                                  <span class="nx">hub_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                                  <span class="nx">relative_path</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                                  <span class="nx">resource_group_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                                  <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>GetImageUploadUrlForDataResult</span
><span class="k">
def </span>get_image_upload_url_for_data_output<span class="p">(</span><span class="nx">entity_type</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                                  <span class="nx">entity_type_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                                  <span class="nx">hub_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                                  <span class="nx">relative_path</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                                  <span class="nx">resource_group_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                                  <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>Output[GetImageUploadUrlForDataResult]</span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="go">
<div class="highlight"><pre class="chroma"><code class="language-go" data-lang="go"
><span class="k">func </span>GetImageUploadUrlForData<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">GetImageUploadUrlForDataArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) (*<span class="nx"><a href="#result">GetImageUploadUrlForDataResult</a></span>, error)</span
><span class="k">
func </span>GetImageUploadUrlForDataOutput<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">GetImageUploadUrlForDataOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) GetImageUploadUrlForDataResultOutput</span
></code></pre></div>

&gt; Note: This function is named `GetImageUploadUrlForData` in the Go SDK.

</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="csharp">
<div class="highlight"><pre class="chroma"><code class="language-csharp" data-lang="csharp"><span class="k">public static class </span><span class="nx">GetImageUploadUrlForData </span><span class="p">
{</span><span class="k">
    public static </span>Task&lt;<span class="nx"><a href="#result">GetImageUploadUrlForDataResult</a></span>> <span class="p">InvokeAsync(</span><span class="nx">GetImageUploadUrlForDataArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="k">
    public static </span>Output&lt;<span class="nx"><a href="#result">GetImageUploadUrlForDataResult</a></span>> <span class="p">Invoke(</span><span class="nx">GetImageUploadUrlForDataInvokeArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="p">
}</span></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="java">
<div class="highlight"><pre class="chroma"><code class="language-java" data-lang="java"><span class="k">public static CompletableFuture&lt;<span class="nx"><a href="#result">GetImageUploadUrlForDataResult</a></span>> </span>getImageUploadUrlForData<span class="p">(</span><span class="nx">GetImageUploadUrlForDataArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
<span class="k">public static Output&lt;<span class="nx"><a href="#result">GetImageUploadUrlForDataResult</a></span>> </span>getImageUploadUrlForData<span class="p">(</span><span class="nx">GetImageUploadUrlForDataArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
</code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="yaml">
<div class="highlight"><pre class="chroma"><code class="language-yaml" data-lang="yaml"><span class="k">fn::invoke:</span>
<span class="k">&nbsp;&nbsp;function:</span> azure-native:customerinsights:getImageUploadUrlForData
<span class="k">&nbsp;&nbsp;arguments:</span>
<span class="c">&nbsp;&nbsp;&nbsp;&nbsp;# arguments dictionary</span></code></pre></div>
</pulumi-choosable>
</div>



The following arguments are supported:


<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="hubname_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#hubname_csharp" style="color: inherit; text-decoration: inherit;">Hub<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the hub.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_csharp" style="color: inherit; text-decoration: inherit;">Resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the resource group.</dd><dt class="property-optional"
            title="Optional">
        <span id="entitytype_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#entitytype_csharp" style="color: inherit; text-decoration: inherit;">Entity<wbr>Type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Type of entity. Can be Profile or Interaction.</dd><dt class="property-optional"
            title="Optional">
        <span id="entitytypename_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#entitytypename_csharp" style="color: inherit; text-decoration: inherit;">Entity<wbr>Type<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Name of the entity type.</dd><dt class="property-optional"
            title="Optional">
        <span id="relativepath_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#relativepath_csharp" style="color: inherit; text-decoration: inherit;">Relative<wbr>Path</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Relative path of the image.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="hubname_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#hubname_go" style="color: inherit; text-decoration: inherit;">Hub<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the hub.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_go" style="color: inherit; text-decoration: inherit;">Resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the resource group.</dd><dt class="property-optional"
            title="Optional">
        <span id="entitytype_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#entitytype_go" style="color: inherit; text-decoration: inherit;">Entity<wbr>Type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Type of entity. Can be Profile or Interaction.</dd><dt class="property-optional"
            title="Optional">
        <span id="entitytypename_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#entitytypename_go" style="color: inherit; text-decoration: inherit;">Entity<wbr>Type<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Name of the entity type.</dd><dt class="property-optional"
            title="Optional">
        <span id="relativepath_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#relativepath_go" style="color: inherit; text-decoration: inherit;">Relative<wbr>Path</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Relative path of the image.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="hubname_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#hubname_java" style="color: inherit; text-decoration: inherit;">hub<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The name of the hub.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_java" style="color: inherit; text-decoration: inherit;">resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The name of the resource group.</dd><dt class="property-optional"
            title="Optional">
        <span id="entitytype_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#entitytype_java" style="color: inherit; text-decoration: inherit;">entity<wbr>Type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Type of entity. Can be Profile or Interaction.</dd><dt class="property-optional"
            title="Optional">
        <span id="entitytypename_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#entitytypename_java" style="color: inherit; text-decoration: inherit;">entity<wbr>Type<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Name of the entity type.</dd><dt class="property-optional"
            title="Optional">
        <span id="relativepath_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#relativepath_java" style="color: inherit; text-decoration: inherit;">relative<wbr>Path</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Relative path of the image.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="hubname_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#hubname_nodejs" style="color: inherit; text-decoration: inherit;">hub<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the hub.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_nodejs" style="color: inherit; text-decoration: inherit;">resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the resource group.</dd><dt class="property-optional"
            title="Optional">
        <span id="entitytype_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#entitytype_nodejs" style="color: inherit; text-decoration: inherit;">entity<wbr>Type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Type of entity. Can be Profile or Interaction.</dd><dt class="property-optional"
            title="Optional">
        <span id="entitytypename_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#entitytypename_nodejs" style="color: inherit; text-decoration: inherit;">entity<wbr>Type<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Name of the entity type.</dd><dt class="property-optional"
            title="Optional">
        <span id="relativepath_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#relativepath_nodejs" style="color: inherit; text-decoration: inherit;">relative<wbr>Path</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Relative path of the image.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="hub_name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#hub_name_python" style="color: inherit; text-decoration: inherit;">hub_<wbr>name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The name of the hub.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resource_group_name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resource_group_name_python" style="color: inherit; text-decoration: inherit;">resource_<wbr>group_<wbr>name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The name of the resource group.</dd><dt class="property-optional"
            title="Optional">
        <span id="entity_type_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#entity_type_python" style="color: inherit; text-decoration: inherit;">entity_<wbr>type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Type of entity. Can be Profile or Interaction.</dd><dt class="property-optional"
            title="Optional">
        <span id="entity_type_name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#entity_type_name_python" style="color: inherit; text-decoration: inherit;">entity_<wbr>type_<wbr>name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Name of the entity type.</dd><dt class="property-optional"
            title="Optional">
        <span id="relative_path_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#relative_path_python" style="color: inherit; text-decoration: inherit;">relative_<wbr>path</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Relative path of the image.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="hubname_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#hubname_yaml" style="color: inherit; text-decoration: inherit;">hub<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The name of the hub.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_yaml" style="color: inherit; text-decoration: inherit;">resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The name of the resource group.</dd><dt class="property-optional"
            title="Optional">
        <span id="entitytype_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#entitytype_yaml" style="color: inherit; text-decoration: inherit;">entity<wbr>Type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Type of entity. Can be Profile or Interaction.</dd><dt class="property-optional"
            title="Optional">
        <span id="entitytypename_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#entitytypename_yaml" style="color: inherit; text-decoration: inherit;">entity<wbr>Type<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Name of the entity type.</dd><dt class="property-optional"
            title="Optional">
        <span id="relativepath_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#relativepath_yaml" style="color: inherit; text-decoration: inherit;">relative<wbr>Path</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Relative path of the image.</dd></dl>
</pulumi-choosable>
</div>




## getImageUploadUrlForData Result {#result}

The following output properties are available:



<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="contenturl_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#contenturl_csharp" style="color: inherit; text-decoration: inherit;">Content<wbr>Url</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Content URL for the image blob.</dd><dt class="property-"
            title="">
        <span id="imageexists_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#imageexists_csharp" style="color: inherit; text-decoration: inherit;">Image<wbr>Exists</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">bool</span>
    </dt>
    <dd>Whether image exists already.</dd><dt class="property-"
            title="">
        <span id="relativepath_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#relativepath_csharp" style="color: inherit; text-decoration: inherit;">Relative<wbr>Path</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Relative path of the image.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="contenturl_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#contenturl_go" style="color: inherit; text-decoration: inherit;">Content<wbr>Url</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Content URL for the image blob.</dd><dt class="property-"
            title="">
        <span id="imageexists_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#imageexists_go" style="color: inherit; text-decoration: inherit;">Image<wbr>Exists</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">bool</span>
    </dt>
    <dd>Whether image exists already.</dd><dt class="property-"
            title="">
        <span id="relativepath_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#relativepath_go" style="color: inherit; text-decoration: inherit;">Relative<wbr>Path</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Relative path of the image.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="contenturl_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#contenturl_java" style="color: inherit; text-decoration: inherit;">content<wbr>Url</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Content URL for the image blob.</dd><dt class="property-"
            title="">
        <span id="imageexists_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#imageexists_java" style="color: inherit; text-decoration: inherit;">image<wbr>Exists</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Boolean</span>
    </dt>
    <dd>Whether image exists already.</dd><dt class="property-"
            title="">
        <span id="relativepath_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#relativepath_java" style="color: inherit; text-decoration: inherit;">relative<wbr>Path</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Relative path of the image.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="contenturl_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#contenturl_nodejs" style="color: inherit; text-decoration: inherit;">content<wbr>Url</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Content URL for the image blob.</dd><dt class="property-"
            title="">
        <span id="imageexists_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#imageexists_nodejs" style="color: inherit; text-decoration: inherit;">image<wbr>Exists</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">boolean</span>
    </dt>
    <dd>Whether image exists already.</dd><dt class="property-"
            title="">
        <span id="relativepath_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#relativepath_nodejs" style="color: inherit; text-decoration: inherit;">relative<wbr>Path</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Relative path of the image.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="content_url_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#content_url_python" style="color: inherit; text-decoration: inherit;">content_<wbr>url</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Content URL for the image blob.</dd><dt class="property-"
            title="">
        <span id="image_exists_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#image_exists_python" style="color: inherit; text-decoration: inherit;">image_<wbr>exists</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">bool</span>
    </dt>
    <dd>Whether image exists already.</dd><dt class="property-"
            title="">
        <span id="relative_path_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#relative_path_python" style="color: inherit; text-decoration: inherit;">relative_<wbr>path</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Relative path of the image.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="contenturl_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#contenturl_yaml" style="color: inherit; text-decoration: inherit;">content<wbr>Url</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Content URL for the image blob.</dd><dt class="property-"
            title="">
        <span id="imageexists_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#imageexists_yaml" style="color: inherit; text-decoration: inherit;">image<wbr>Exists</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Boolean</span>
    </dt>
    <dd>Whether image exists already.</dd><dt class="property-"
            title="">
        <span id="relativepath_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#relativepath_yaml" style="color: inherit; text-decoration: inherit;">relative<wbr>Path</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Relative path of the image.</dd></dl>
</pulumi-choosable>
</div>





<h2 id="package-details">Package Details</h2>
<dl class="package-details">
	<dt>Repository</dt>
	<dd><a href="https://github.com/pulumi/pulumi-azure-native">Azure Native pulumi/pulumi-azure-native</a></dd>
	<dt>License</dt>
	<dd>Apache-2.0</dd>
</dl>

