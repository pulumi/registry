
---
title: "listDnsForwardingRulesetByVirtualNetwork"
title_tag: "azure-native.dnsresolver.listDnsForwardingRulesetByVirtualNetwork"
meta_desc: "Documentation for the azure-native.dnsresolver.listDnsForwardingRulesetByVirtualNetwork function with examples, input properties, output properties, and supporting types."
layout: api
no_edit_this_page: true
---



<!-- WARNING: this file was generated by Pulumi Docs Generator. -->
<!-- Do not edit by hand unless you're certain you know what you are doing! -->

Lists DNS forwarding ruleset resource IDs attached to a virtual network.

Azure REST API version: 2023-07-01-preview.

Other available API versions: 2020-04-01-preview, 2022-07-01.

These versions are not included in the SDK but you can generate them as a local SDK package; see the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.




## Using listDnsForwardingRulesetByVirtualNetwork {#using}

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
><span class="k">function </span>listDnsForwardingRulesetByVirtualNetwork<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">ListDnsForwardingRulesetByVirtualNetworkArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Promise&lt;<span class="nx"><a href="#result">ListDnsForwardingRulesetByVirtualNetworkResult</a></span>></span
><span class="k">
function </span>listDnsForwardingRulesetByVirtualNetworkOutput<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">ListDnsForwardingRulesetByVirtualNetworkOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Output&lt;<span class="nx"><a href="#result">ListDnsForwardingRulesetByVirtualNetworkResult</a></span>></span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="python">
<div class="highlight"><pre class="chroma"><code class="language-python" data-lang="python"
><span class="k">def </span>list_dns_forwarding_ruleset_by_virtual_network<span class="p">(</span><span class="nx">resource_group_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                                                   <span class="nx">top</span><span class="p">:</span> <span class="nx">Optional[int]</span> = None<span class="p">,</span>
                                                   <span class="nx">virtual_network_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                                                   <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>ListDnsForwardingRulesetByVirtualNetworkResult</span
><span class="k">
def </span>list_dns_forwarding_ruleset_by_virtual_network_output<span class="p">(</span><span class="nx">resource_group_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                                                   <span class="nx">top</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[int]]</span> = None<span class="p">,</span>
                                                   <span class="nx">virtual_network_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                                                   <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>Output[ListDnsForwardingRulesetByVirtualNetworkResult]</span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="go">
<div class="highlight"><pre class="chroma"><code class="language-go" data-lang="go"
><span class="k">func </span>ListDnsForwardingRulesetByVirtualNetwork<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">ListDnsForwardingRulesetByVirtualNetworkArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) (*<span class="nx"><a href="#result">ListDnsForwardingRulesetByVirtualNetworkResult</a></span>, error)</span
><span class="k">
func </span>ListDnsForwardingRulesetByVirtualNetworkOutput<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">ListDnsForwardingRulesetByVirtualNetworkOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) ListDnsForwardingRulesetByVirtualNetworkResultOutput</span
></code></pre></div>

&gt; Note: This function is named `ListDnsForwardingRulesetByVirtualNetwork` in the Go SDK.

</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="csharp">
<div class="highlight"><pre class="chroma"><code class="language-csharp" data-lang="csharp"><span class="k">public static class </span><span class="nx">ListDnsForwardingRulesetByVirtualNetwork </span><span class="p">
{</span><span class="k">
    public static </span>Task&lt;<span class="nx"><a href="#result">ListDnsForwardingRulesetByVirtualNetworkResult</a></span>> <span class="p">InvokeAsync(</span><span class="nx">ListDnsForwardingRulesetByVirtualNetworkArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="k">
    public static </span>Output&lt;<span class="nx"><a href="#result">ListDnsForwardingRulesetByVirtualNetworkResult</a></span>> <span class="p">Invoke(</span><span class="nx">ListDnsForwardingRulesetByVirtualNetworkInvokeArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="p">
}</span></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="java">
<div class="highlight"><pre class="chroma"><code class="language-java" data-lang="java"><span class="k">public static CompletableFuture&lt;<span class="nx"><a href="#result">ListDnsForwardingRulesetByVirtualNetworkResult</a></span>> </span>listDnsForwardingRulesetByVirtualNetwork<span class="p">(</span><span class="nx">ListDnsForwardingRulesetByVirtualNetworkArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
<span class="k">public static Output&lt;<span class="nx"><a href="#result">ListDnsForwardingRulesetByVirtualNetworkResult</a></span>> </span>listDnsForwardingRulesetByVirtualNetwork<span class="p">(</span><span class="nx">ListDnsForwardingRulesetByVirtualNetworkArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
</code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="yaml">
<div class="highlight"><pre class="chroma"><code class="language-yaml" data-lang="yaml"><span class="k">fn::invoke:</span>
<span class="k">&nbsp;&nbsp;function:</span> azure-native:dnsresolver:listDnsForwardingRulesetByVirtualNetwork
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
    <dd>The name of the resource group. The name is case insensitive.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="virtualnetworkname_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#virtualnetworkname_csharp" style="color: inherit; text-decoration: inherit;">Virtual<wbr>Network<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the virtual network.</dd><dt class="property-optional"
            title="Optional">
        <span id="top_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#top_csharp" style="color: inherit; text-decoration: inherit;">Top</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">int</span>
    </dt>
    <dd>The maximum number of results to return. If not specified, returns up to 100 results.</dd></dl>
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
    <dd>The name of the resource group. The name is case insensitive.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="virtualnetworkname_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#virtualnetworkname_go" style="color: inherit; text-decoration: inherit;">Virtual<wbr>Network<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the virtual network.</dd><dt class="property-optional"
            title="Optional">
        <span id="top_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#top_go" style="color: inherit; text-decoration: inherit;">Top</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">int</span>
    </dt>
    <dd>The maximum number of results to return. If not specified, returns up to 100 results.</dd></dl>
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
    <dd>The name of the resource group. The name is case insensitive.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="virtualnetworkname_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#virtualnetworkname_java" style="color: inherit; text-decoration: inherit;">virtual<wbr>Network<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The name of the virtual network.</dd><dt class="property-optional"
            title="Optional">
        <span id="top_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#top_java" style="color: inherit; text-decoration: inherit;">top</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Integer</span>
    </dt>
    <dd>The maximum number of results to return. If not specified, returns up to 100 results.</dd></dl>
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
    <dd>The name of the resource group. The name is case insensitive.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="virtualnetworkname_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#virtualnetworkname_nodejs" style="color: inherit; text-decoration: inherit;">virtual<wbr>Network<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the virtual network.</dd><dt class="property-optional"
            title="Optional">
        <span id="top_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#top_nodejs" style="color: inherit; text-decoration: inherit;">top</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">number</span>
    </dt>
    <dd>The maximum number of results to return. If not specified, returns up to 100 results.</dd></dl>
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
    <dd>The name of the resource group. The name is case insensitive.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="virtual_network_name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#virtual_network_name_python" style="color: inherit; text-decoration: inherit;">virtual_<wbr>network_<wbr>name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The name of the virtual network.</dd><dt class="property-optional"
            title="Optional">
        <span id="top_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#top_python" style="color: inherit; text-decoration: inherit;">top</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">int</span>
    </dt>
    <dd>The maximum number of results to return. If not specified, returns up to 100 results.</dd></dl>
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
    <dd>The name of the resource group. The name is case insensitive.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="virtualnetworkname_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#virtualnetworkname_yaml" style="color: inherit; text-decoration: inherit;">virtual<wbr>Network<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The name of the virtual network.</dd><dt class="property-optional"
            title="Optional">
        <span id="top_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#top_yaml" style="color: inherit; text-decoration: inherit;">top</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Number</span>
    </dt>
    <dd>The maximum number of results to return. If not specified, returns up to 100 results.</dd></dl>
</pulumi-choosable>
</div>




## listDnsForwardingRulesetByVirtualNetwork Result {#result}

The following output properties are available:



<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="nextlink_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#nextlink_csharp" style="color: inherit; text-decoration: inherit;">Next<wbr>Link</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The continuation token for the next page of results.</dd><dt class="property-"
            title="">
        <span id="value_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#value_csharp" style="color: inherit; text-decoration: inherit;">Value</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#virtualnetworkdnsforwardingrulesetresponse">List&lt;Pulumi.<wbr>Azure<wbr>Native.<wbr>Dns<wbr>Resolver.<wbr>Outputs.<wbr>Virtual<wbr>Network<wbr>Dns<wbr>Forwarding<wbr>Ruleset<wbr>Response&gt;</a></span>
    </dt>
    <dd>Enumeration of the Virtual Network DNS Forwarding Ruleset.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="nextlink_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#nextlink_go" style="color: inherit; text-decoration: inherit;">Next<wbr>Link</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The continuation token for the next page of results.</dd><dt class="property-"
            title="">
        <span id="value_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#value_go" style="color: inherit; text-decoration: inherit;">Value</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#virtualnetworkdnsforwardingrulesetresponse">[]Virtual<wbr>Network<wbr>Dns<wbr>Forwarding<wbr>Ruleset<wbr>Response</a></span>
    </dt>
    <dd>Enumeration of the Virtual Network DNS Forwarding Ruleset.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="nextlink_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#nextlink_java" style="color: inherit; text-decoration: inherit;">next<wbr>Link</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The continuation token for the next page of results.</dd><dt class="property-"
            title="">
        <span id="value_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#value_java" style="color: inherit; text-decoration: inherit;">value</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#virtualnetworkdnsforwardingrulesetresponse">List&lt;Virtual<wbr>Network<wbr>Dns<wbr>Forwarding<wbr>Ruleset<wbr>Response&gt;</a></span>
    </dt>
    <dd>Enumeration of the Virtual Network DNS Forwarding Ruleset.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="nextlink_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#nextlink_nodejs" style="color: inherit; text-decoration: inherit;">next<wbr>Link</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The continuation token for the next page of results.</dd><dt class="property-"
            title="">
        <span id="value_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#value_nodejs" style="color: inherit; text-decoration: inherit;">value</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#virtualnetworkdnsforwardingrulesetresponse">Virtual<wbr>Network<wbr>Dns<wbr>Forwarding<wbr>Ruleset<wbr>Response[]</a></span>
    </dt>
    <dd>Enumeration of the Virtual Network DNS Forwarding Ruleset.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="next_link_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#next_link_python" style="color: inherit; text-decoration: inherit;">next_<wbr>link</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The continuation token for the next page of results.</dd><dt class="property-"
            title="">
        <span id="value_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#value_python" style="color: inherit; text-decoration: inherit;">value</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#virtualnetworkdnsforwardingrulesetresponse">Sequence[Virtual<wbr>Network<wbr>Dns<wbr>Forwarding<wbr>Ruleset<wbr>Response]</a></span>
    </dt>
    <dd>Enumeration of the Virtual Network DNS Forwarding Ruleset.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="nextlink_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#nextlink_yaml" style="color: inherit; text-decoration: inherit;">next<wbr>Link</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The continuation token for the next page of results.</dd><dt class="property-"
            title="">
        <span id="value_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#value_yaml" style="color: inherit; text-decoration: inherit;">value</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#virtualnetworkdnsforwardingrulesetresponse">List&lt;Property Map&gt;</a></span>
    </dt>
    <dd>Enumeration of the Virtual Network DNS Forwarding Ruleset.</dd></dl>
</pulumi-choosable>
</div>




## Supporting Types


<h4 id="subresourceresponse">Sub<wbr>Resource<wbr>Response</h4>



<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="id_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_csharp" style="color: inherit; text-decoration: inherit;">Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Resource ID.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="id_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_go" style="color: inherit; text-decoration: inherit;">Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Resource ID.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="id_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_java" style="color: inherit; text-decoration: inherit;">id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Resource ID.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="id_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_nodejs" style="color: inherit; text-decoration: inherit;">id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Resource ID.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="id_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_python" style="color: inherit; text-decoration: inherit;">id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Resource ID.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="id_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_yaml" style="color: inherit; text-decoration: inherit;">id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Resource ID.</dd></dl>
</pulumi-choosable>
</div>

<h4 id="virtualnetworkdnsforwardingrulesetresponse">Virtual<wbr>Network<wbr>Dns<wbr>Forwarding<wbr>Ruleset<wbr>Response</h4>



<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="id_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_csharp" style="color: inherit; text-decoration: inherit;">Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>DNS Forwarding Ruleset Resource ID.</dd><dt class="property-optional"
            title="Optional">
        <span id="virtualnetworklink_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#virtualnetworklink_csharp" style="color: inherit; text-decoration: inherit;">Virtual<wbr>Network<wbr>Link</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#subresourceresponse">Pulumi.<wbr>Azure<wbr>Native.<wbr>Dns<wbr>Resolver.<wbr>Inputs.<wbr>Sub<wbr>Resource<wbr>Response</a></span>
    </dt>
    <dd>The reference to the virtual network link.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="id_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_go" style="color: inherit; text-decoration: inherit;">Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>DNS Forwarding Ruleset Resource ID.</dd><dt class="property-optional"
            title="Optional">
        <span id="virtualnetworklink_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#virtualnetworklink_go" style="color: inherit; text-decoration: inherit;">Virtual<wbr>Network<wbr>Link</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#subresourceresponse">Sub<wbr>Resource<wbr>Response</a></span>
    </dt>
    <dd>The reference to the virtual network link.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="id_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_java" style="color: inherit; text-decoration: inherit;">id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>DNS Forwarding Ruleset Resource ID.</dd><dt class="property-optional"
            title="Optional">
        <span id="virtualnetworklink_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#virtualnetworklink_java" style="color: inherit; text-decoration: inherit;">virtual<wbr>Network<wbr>Link</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#subresourceresponse">Sub<wbr>Resource<wbr>Response</a></span>
    </dt>
    <dd>The reference to the virtual network link.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="id_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_nodejs" style="color: inherit; text-decoration: inherit;">id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>DNS Forwarding Ruleset Resource ID.</dd><dt class="property-optional"
            title="Optional">
        <span id="virtualnetworklink_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#virtualnetworklink_nodejs" style="color: inherit; text-decoration: inherit;">virtual<wbr>Network<wbr>Link</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#subresourceresponse">Sub<wbr>Resource<wbr>Response</a></span>
    </dt>
    <dd>The reference to the virtual network link.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="id_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_python" style="color: inherit; text-decoration: inherit;">id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>DNS Forwarding Ruleset Resource ID.</dd><dt class="property-optional"
            title="Optional">
        <span id="virtual_network_link_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#virtual_network_link_python" style="color: inherit; text-decoration: inherit;">virtual_<wbr>network_<wbr>link</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#subresourceresponse">Sub<wbr>Resource<wbr>Response</a></span>
    </dt>
    <dd>The reference to the virtual network link.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="id_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_yaml" style="color: inherit; text-decoration: inherit;">id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>DNS Forwarding Ruleset Resource ID.</dd><dt class="property-optional"
            title="Optional">
        <span id="virtualnetworklink_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#virtualnetworklink_yaml" style="color: inherit; text-decoration: inherit;">virtual<wbr>Network<wbr>Link</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#subresourceresponse">Property Map</a></span>
    </dt>
    <dd>The reference to the virtual network link.</dd></dl>
</pulumi-choosable>
</div>





<h2 id="package-details">Package Details</h2>
<dl class="package-details">
	<dt>Repository</dt>
	<dd><a href="https://github.com/pulumi/pulumi-azure-native">Azure Native pulumi/pulumi-azure-native</a></dd>
	<dt>License</dt>
	<dd>Apache-2.0</dd>
</dl>

