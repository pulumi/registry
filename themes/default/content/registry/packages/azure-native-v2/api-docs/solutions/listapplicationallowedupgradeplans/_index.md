
---
title: "listApplicationAllowedUpgradePlans"
title_tag: "azure-native.solutions.listApplicationAllowedUpgradePlans"
meta_desc: "Documentation for the azure-native.solutions.listApplicationAllowedUpgradePlans function with examples, input properties, output properties, and supporting types."
layout: api
no_edit_this_page: true
---



<!-- WARNING: this file was generated by Pulumi Docs Generator. -->
<!-- Do not edit by hand unless you're certain you know what you are doing! -->

List allowed upgrade plans for application.

Azure REST API version: 2021-07-01.

Other available API versions: 2023-12-01-preview.

These versions are not included in the SDK but you can generate them as a local SDK package; see the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.




## Using listApplicationAllowedUpgradePlans {#using}

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
><span class="k">function </span>listApplicationAllowedUpgradePlans<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">ListApplicationAllowedUpgradePlansArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Promise&lt;<span class="nx"><a href="#result">ListApplicationAllowedUpgradePlansResult</a></span>></span
><span class="k">
function </span>listApplicationAllowedUpgradePlansOutput<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">ListApplicationAllowedUpgradePlansOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Output&lt;<span class="nx"><a href="#result">ListApplicationAllowedUpgradePlansResult</a></span>></span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="python">
<div class="highlight"><pre class="chroma"><code class="language-python" data-lang="python"
><span class="k">def </span>list_application_allowed_upgrade_plans<span class="p">(</span><span class="nx">application_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                                           <span class="nx">resource_group_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                                           <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>ListApplicationAllowedUpgradePlansResult</span
><span class="k">
def </span>list_application_allowed_upgrade_plans_output<span class="p">(</span><span class="nx">application_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                                           <span class="nx">resource_group_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                                           <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>Output[ListApplicationAllowedUpgradePlansResult]</span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="go">
<div class="highlight"><pre class="chroma"><code class="language-go" data-lang="go"
><span class="k">func </span>ListApplicationAllowedUpgradePlans<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">ListApplicationAllowedUpgradePlansArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) (*<span class="nx"><a href="#result">ListApplicationAllowedUpgradePlansResult</a></span>, error)</span
><span class="k">
func </span>ListApplicationAllowedUpgradePlansOutput<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">ListApplicationAllowedUpgradePlansOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) ListApplicationAllowedUpgradePlansResultOutput</span
></code></pre></div>

&gt; Note: This function is named `ListApplicationAllowedUpgradePlans` in the Go SDK.

</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="csharp">
<div class="highlight"><pre class="chroma"><code class="language-csharp" data-lang="csharp"><span class="k">public static class </span><span class="nx">ListApplicationAllowedUpgradePlans </span><span class="p">
{</span><span class="k">
    public static </span>Task&lt;<span class="nx"><a href="#result">ListApplicationAllowedUpgradePlansResult</a></span>> <span class="p">InvokeAsync(</span><span class="nx">ListApplicationAllowedUpgradePlansArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="k">
    public static </span>Output&lt;<span class="nx"><a href="#result">ListApplicationAllowedUpgradePlansResult</a></span>> <span class="p">Invoke(</span><span class="nx">ListApplicationAllowedUpgradePlansInvokeArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="p">
}</span></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="java">
<div class="highlight"><pre class="chroma"><code class="language-java" data-lang="java"><span class="k">public static CompletableFuture&lt;<span class="nx"><a href="#result">ListApplicationAllowedUpgradePlansResult</a></span>> </span>listApplicationAllowedUpgradePlans<span class="p">(</span><span class="nx">ListApplicationAllowedUpgradePlansArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
<span class="k">public static Output&lt;<span class="nx"><a href="#result">ListApplicationAllowedUpgradePlansResult</a></span>> </span>listApplicationAllowedUpgradePlans<span class="p">(</span><span class="nx">ListApplicationAllowedUpgradePlansArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
</code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="yaml">
<div class="highlight"><pre class="chroma"><code class="language-yaml" data-lang="yaml"><span class="k">fn::invoke:</span>
<span class="k">&nbsp;&nbsp;function:</span> azure-native:solutions:listApplicationAllowedUpgradePlans
<span class="k">&nbsp;&nbsp;arguments:</span>
<span class="c">&nbsp;&nbsp;&nbsp;&nbsp;# arguments dictionary</span></code></pre></div>
</pulumi-choosable>
</div>



The following arguments are supported:


<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="applicationname_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#applicationname_csharp" style="color: inherit; text-decoration: inherit;">Application<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the managed application.</dd><dt class="property-required property-replacement"
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
        <span id="applicationname_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#applicationname_go" style="color: inherit; text-decoration: inherit;">Application<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the managed application.</dd><dt class="property-required property-replacement"
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
        <span id="applicationname_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#applicationname_java" style="color: inherit; text-decoration: inherit;">application<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The name of the managed application.</dd><dt class="property-required property-replacement"
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
        <span id="applicationname_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#applicationname_nodejs" style="color: inherit; text-decoration: inherit;">application<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the managed application.</dd><dt class="property-required property-replacement"
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
        <span id="application_name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#application_name_python" style="color: inherit; text-decoration: inherit;">application_<wbr>name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The name of the managed application.</dd><dt class="property-required property-replacement"
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
        <span id="applicationname_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#applicationname_yaml" style="color: inherit; text-decoration: inherit;">application<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The name of the managed application.</dd><dt class="property-required property-replacement"
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




## listApplicationAllowedUpgradePlans Result {#result}

The following output properties are available:



<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="value_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#value_csharp" style="color: inherit; text-decoration: inherit;">Value</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#planresponse">List&lt;Pulumi.<wbr>Azure<wbr>Native.<wbr>Solutions.<wbr>Outputs.<wbr>Plan<wbr>Response&gt;</a></span>
    </dt>
    <dd>The array of plans.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="value_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#value_go" style="color: inherit; text-decoration: inherit;">Value</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#planresponse">[]Plan<wbr>Response</a></span>
    </dt>
    <dd>The array of plans.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="value_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#value_java" style="color: inherit; text-decoration: inherit;">value</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#planresponse">List&lt;Plan<wbr>Response&gt;</a></span>
    </dt>
    <dd>The array of plans.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="value_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#value_nodejs" style="color: inherit; text-decoration: inherit;">value</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#planresponse">Plan<wbr>Response[]</a></span>
    </dt>
    <dd>The array of plans.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="value_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#value_python" style="color: inherit; text-decoration: inherit;">value</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#planresponse">Sequence[Plan<wbr>Response]</a></span>
    </dt>
    <dd>The array of plans.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="value_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#value_yaml" style="color: inherit; text-decoration: inherit;">value</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#planresponse">List&lt;Property Map&gt;</a></span>
    </dt>
    <dd>The array of plans.</dd></dl>
</pulumi-choosable>
</div>




## Supporting Types


<h4 id="planresponse">Plan<wbr>Response</h4>



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
    <dd>The plan name.</dd><dt class="property-required"
            title="Required">
        <span id="product_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#product_csharp" style="color: inherit; text-decoration: inherit;">Product</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The product code.</dd><dt class="property-required"
            title="Required">
        <span id="publisher_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#publisher_csharp" style="color: inherit; text-decoration: inherit;">Publisher</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The publisher ID.</dd><dt class="property-required"
            title="Required">
        <span id="version_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#version_csharp" style="color: inherit; text-decoration: inherit;">Version</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The plan's version.</dd><dt class="property-optional"
            title="Optional">
        <span id="promotioncode_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#promotioncode_csharp" style="color: inherit; text-decoration: inherit;">Promotion<wbr>Code</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The promotion code.</dd></dl>
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
    <dd>The plan name.</dd><dt class="property-required"
            title="Required">
        <span id="product_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#product_go" style="color: inherit; text-decoration: inherit;">Product</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The product code.</dd><dt class="property-required"
            title="Required">
        <span id="publisher_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#publisher_go" style="color: inherit; text-decoration: inherit;">Publisher</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The publisher ID.</dd><dt class="property-required"
            title="Required">
        <span id="version_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#version_go" style="color: inherit; text-decoration: inherit;">Version</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The plan's version.</dd><dt class="property-optional"
            title="Optional">
        <span id="promotioncode_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#promotioncode_go" style="color: inherit; text-decoration: inherit;">Promotion<wbr>Code</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The promotion code.</dd></dl>
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
    <dd>The plan name.</dd><dt class="property-required"
            title="Required">
        <span id="product_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#product_java" style="color: inherit; text-decoration: inherit;">product</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The product code.</dd><dt class="property-required"
            title="Required">
        <span id="publisher_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#publisher_java" style="color: inherit; text-decoration: inherit;">publisher</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The publisher ID.</dd><dt class="property-required"
            title="Required">
        <span id="version_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#version_java" style="color: inherit; text-decoration: inherit;">version</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The plan's version.</dd><dt class="property-optional"
            title="Optional">
        <span id="promotioncode_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#promotioncode_java" style="color: inherit; text-decoration: inherit;">promotion<wbr>Code</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The promotion code.</dd></dl>
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
    <dd>The plan name.</dd><dt class="property-required"
            title="Required">
        <span id="product_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#product_nodejs" style="color: inherit; text-decoration: inherit;">product</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The product code.</dd><dt class="property-required"
            title="Required">
        <span id="publisher_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#publisher_nodejs" style="color: inherit; text-decoration: inherit;">publisher</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The publisher ID.</dd><dt class="property-required"
            title="Required">
        <span id="version_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#version_nodejs" style="color: inherit; text-decoration: inherit;">version</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The plan's version.</dd><dt class="property-optional"
            title="Optional">
        <span id="promotioncode_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#promotioncode_nodejs" style="color: inherit; text-decoration: inherit;">promotion<wbr>Code</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The promotion code.</dd></dl>
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
    <dd>The plan name.</dd><dt class="property-required"
            title="Required">
        <span id="product_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#product_python" style="color: inherit; text-decoration: inherit;">product</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The product code.</dd><dt class="property-required"
            title="Required">
        <span id="publisher_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#publisher_python" style="color: inherit; text-decoration: inherit;">publisher</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The publisher ID.</dd><dt class="property-required"
            title="Required">
        <span id="version_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#version_python" style="color: inherit; text-decoration: inherit;">version</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The plan's version.</dd><dt class="property-optional"
            title="Optional">
        <span id="promotion_code_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#promotion_code_python" style="color: inherit; text-decoration: inherit;">promotion_<wbr>code</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The promotion code.</dd></dl>
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
    <dd>The plan name.</dd><dt class="property-required"
            title="Required">
        <span id="product_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#product_yaml" style="color: inherit; text-decoration: inherit;">product</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The product code.</dd><dt class="property-required"
            title="Required">
        <span id="publisher_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#publisher_yaml" style="color: inherit; text-decoration: inherit;">publisher</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The publisher ID.</dd><dt class="property-required"
            title="Required">
        <span id="version_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#version_yaml" style="color: inherit; text-decoration: inherit;">version</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The plan's version.</dd><dt class="property-optional"
            title="Optional">
        <span id="promotioncode_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#promotioncode_yaml" style="color: inherit; text-decoration: inherit;">promotion<wbr>Code</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The promotion code.</dd></dl>
</pulumi-choosable>
</div>





<h2 id="package-details">Package Details</h2>
<dl class="package-details">
	<dt>Repository</dt>
	<dd><a href="https://github.com/pulumi/pulumi-azure-native">Azure Native pulumi/pulumi-azure-native</a></dd>
	<dt>License</dt>
	<dd>Apache-2.0</dd>
</dl>

