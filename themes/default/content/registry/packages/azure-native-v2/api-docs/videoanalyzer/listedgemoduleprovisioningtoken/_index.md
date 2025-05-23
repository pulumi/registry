
---
title: "listEdgeModuleProvisioningToken"
title_tag: "azure-native.videoanalyzer.listEdgeModuleProvisioningToken"
meta_desc: "Documentation for the azure-native.videoanalyzer.listEdgeModuleProvisioningToken function with examples, input properties, output properties, and supporting types."
layout: api
no_edit_this_page: true
---



<!-- WARNING: this file was generated by Pulumi Docs Generator. -->
<!-- Do not edit by hand unless you're certain you know what you are doing! -->

Creates a new provisioning token. A provisioning token allows for a single instance of Azure Video analyzer IoT edge module to be initialized and authorized to the cloud account. The provisioning token itself is short lived and it is only used for the initial handshake between IoT edge module and the cloud. After the initial handshake, the IoT edge module will agree on a set of authentication keys which will be auto-rotated as long as the module is able to periodically connect to the cloud. A new provisioning token can be generated for the same IoT edge module in case the module state lost or reset.

Azure REST API version: 2021-11-01-preview.




## Using listEdgeModuleProvisioningToken {#using}

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
><span class="k">function </span>listEdgeModuleProvisioningToken<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">ListEdgeModuleProvisioningTokenArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Promise&lt;<span class="nx"><a href="#result">ListEdgeModuleProvisioningTokenResult</a></span>></span
><span class="k">
function </span>listEdgeModuleProvisioningTokenOutput<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">ListEdgeModuleProvisioningTokenOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Output&lt;<span class="nx"><a href="#result">ListEdgeModuleProvisioningTokenResult</a></span>></span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="python">
<div class="highlight"><pre class="chroma"><code class="language-python" data-lang="python"
><span class="k">def </span>list_edge_module_provisioning_token<span class="p">(</span><span class="nx">account_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                                        <span class="nx">edge_module_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                                        <span class="nx">expiration_date</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                                        <span class="nx">resource_group_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                                        <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>ListEdgeModuleProvisioningTokenResult</span
><span class="k">
def </span>list_edge_module_provisioning_token_output<span class="p">(</span><span class="nx">account_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                                        <span class="nx">edge_module_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                                        <span class="nx">expiration_date</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                                        <span class="nx">resource_group_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                                        <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>Output[ListEdgeModuleProvisioningTokenResult]</span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="go">
<div class="highlight"><pre class="chroma"><code class="language-go" data-lang="go"
><span class="k">func </span>ListEdgeModuleProvisioningToken<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">ListEdgeModuleProvisioningTokenArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) (*<span class="nx"><a href="#result">ListEdgeModuleProvisioningTokenResult</a></span>, error)</span
><span class="k">
func </span>ListEdgeModuleProvisioningTokenOutput<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">ListEdgeModuleProvisioningTokenOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) ListEdgeModuleProvisioningTokenResultOutput</span
></code></pre></div>

&gt; Note: This function is named `ListEdgeModuleProvisioningToken` in the Go SDK.

</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="csharp">
<div class="highlight"><pre class="chroma"><code class="language-csharp" data-lang="csharp"><span class="k">public static class </span><span class="nx">ListEdgeModuleProvisioningToken </span><span class="p">
{</span><span class="k">
    public static </span>Task&lt;<span class="nx"><a href="#result">ListEdgeModuleProvisioningTokenResult</a></span>> <span class="p">InvokeAsync(</span><span class="nx">ListEdgeModuleProvisioningTokenArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="k">
    public static </span>Output&lt;<span class="nx"><a href="#result">ListEdgeModuleProvisioningTokenResult</a></span>> <span class="p">Invoke(</span><span class="nx">ListEdgeModuleProvisioningTokenInvokeArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="p">
}</span></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="java">
<div class="highlight"><pre class="chroma"><code class="language-java" data-lang="java"><span class="k">public static CompletableFuture&lt;<span class="nx"><a href="#result">ListEdgeModuleProvisioningTokenResult</a></span>> </span>listEdgeModuleProvisioningToken<span class="p">(</span><span class="nx">ListEdgeModuleProvisioningTokenArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
<span class="k">public static Output&lt;<span class="nx"><a href="#result">ListEdgeModuleProvisioningTokenResult</a></span>> </span>listEdgeModuleProvisioningToken<span class="p">(</span><span class="nx">ListEdgeModuleProvisioningTokenArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
</code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="yaml">
<div class="highlight"><pre class="chroma"><code class="language-yaml" data-lang="yaml"><span class="k">fn::invoke:</span>
<span class="k">&nbsp;&nbsp;function:</span> azure-native:videoanalyzer:listEdgeModuleProvisioningToken
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
    <dd>The Azure Video Analyzer account name.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="edgemodulename_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#edgemodulename_csharp" style="color: inherit; text-decoration: inherit;">Edge<wbr>Module<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The Edge Module name.</dd><dt class="property-required"
            title="Required">
        <span id="expirationdate_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#expirationdate_csharp" style="color: inherit; text-decoration: inherit;">Expiration<wbr>Date</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The desired expiration date of the registration token. The Azure Video Analyzer IoT edge module must be initialized and connected to the Internet prior to the token expiration date.</dd><dt class="property-required property-replacement"
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
        <span id="accountname_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#accountname_go" style="color: inherit; text-decoration: inherit;">Account<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The Azure Video Analyzer account name.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="edgemodulename_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#edgemodulename_go" style="color: inherit; text-decoration: inherit;">Edge<wbr>Module<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The Edge Module name.</dd><dt class="property-required"
            title="Required">
        <span id="expirationdate_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#expirationdate_go" style="color: inherit; text-decoration: inherit;">Expiration<wbr>Date</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The desired expiration date of the registration token. The Azure Video Analyzer IoT edge module must be initialized and connected to the Internet prior to the token expiration date.</dd><dt class="property-required property-replacement"
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
        <span id="accountname_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#accountname_java" style="color: inherit; text-decoration: inherit;">account<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The Azure Video Analyzer account name.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="edgemodulename_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#edgemodulename_java" style="color: inherit; text-decoration: inherit;">edge<wbr>Module<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The Edge Module name.</dd><dt class="property-required"
            title="Required">
        <span id="expirationdate_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#expirationdate_java" style="color: inherit; text-decoration: inherit;">expiration<wbr>Date</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The desired expiration date of the registration token. The Azure Video Analyzer IoT edge module must be initialized and connected to the Internet prior to the token expiration date.</dd><dt class="property-required property-replacement"
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
        <span id="accountname_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#accountname_nodejs" style="color: inherit; text-decoration: inherit;">account<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The Azure Video Analyzer account name.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="edgemodulename_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#edgemodulename_nodejs" style="color: inherit; text-decoration: inherit;">edge<wbr>Module<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The Edge Module name.</dd><dt class="property-required"
            title="Required">
        <span id="expirationdate_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#expirationdate_nodejs" style="color: inherit; text-decoration: inherit;">expiration<wbr>Date</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The desired expiration date of the registration token. The Azure Video Analyzer IoT edge module must be initialized and connected to the Internet prior to the token expiration date.</dd><dt class="property-required property-replacement"
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
        <span id="account_name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#account_name_python" style="color: inherit; text-decoration: inherit;">account_<wbr>name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The Azure Video Analyzer account name.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="edge_module_name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#edge_module_name_python" style="color: inherit; text-decoration: inherit;">edge_<wbr>module_<wbr>name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The Edge Module name.</dd><dt class="property-required"
            title="Required">
        <span id="expiration_date_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#expiration_date_python" style="color: inherit; text-decoration: inherit;">expiration_<wbr>date</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The desired expiration date of the registration token. The Azure Video Analyzer IoT edge module must be initialized and connected to the Internet prior to the token expiration date.</dd><dt class="property-required property-replacement"
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
        <span id="accountname_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#accountname_yaml" style="color: inherit; text-decoration: inherit;">account<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The Azure Video Analyzer account name.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="edgemodulename_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#edgemodulename_yaml" style="color: inherit; text-decoration: inherit;">edge<wbr>Module<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The Edge Module name.</dd><dt class="property-required"
            title="Required">
        <span id="expirationdate_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#expirationdate_yaml" style="color: inherit; text-decoration: inherit;">expiration<wbr>Date</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The desired expiration date of the registration token. The Azure Video Analyzer IoT edge module must be initialized and connected to the Internet prior to the token expiration date.</dd><dt class="property-required property-replacement"
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




## listEdgeModuleProvisioningToken Result {#result}

The following output properties are available:



<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="expirationdate_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#expirationdate_csharp" style="color: inherit; text-decoration: inherit;">Expiration<wbr>Date</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The expiration date of the registration token. The Azure Video Analyzer IoT edge module must be initialized and connected to the Internet prior to the token expiration date.</dd><dt class="property-"
            title="">
        <span id="token_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#token_csharp" style="color: inherit; text-decoration: inherit;">Token</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The token blob to be provided to the Azure Video Analyzer IoT edge module through the Azure IoT Edge module twin properties.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="expirationdate_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#expirationdate_go" style="color: inherit; text-decoration: inherit;">Expiration<wbr>Date</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The expiration date of the registration token. The Azure Video Analyzer IoT edge module must be initialized and connected to the Internet prior to the token expiration date.</dd><dt class="property-"
            title="">
        <span id="token_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#token_go" style="color: inherit; text-decoration: inherit;">Token</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The token blob to be provided to the Azure Video Analyzer IoT edge module through the Azure IoT Edge module twin properties.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="expirationdate_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#expirationdate_java" style="color: inherit; text-decoration: inherit;">expiration<wbr>Date</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The expiration date of the registration token. The Azure Video Analyzer IoT edge module must be initialized and connected to the Internet prior to the token expiration date.</dd><dt class="property-"
            title="">
        <span id="token_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#token_java" style="color: inherit; text-decoration: inherit;">token</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The token blob to be provided to the Azure Video Analyzer IoT edge module through the Azure IoT Edge module twin properties.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="expirationdate_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#expirationdate_nodejs" style="color: inherit; text-decoration: inherit;">expiration<wbr>Date</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The expiration date of the registration token. The Azure Video Analyzer IoT edge module must be initialized and connected to the Internet prior to the token expiration date.</dd><dt class="property-"
            title="">
        <span id="token_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#token_nodejs" style="color: inherit; text-decoration: inherit;">token</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The token blob to be provided to the Azure Video Analyzer IoT edge module through the Azure IoT Edge module twin properties.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="expiration_date_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#expiration_date_python" style="color: inherit; text-decoration: inherit;">expiration_<wbr>date</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The expiration date of the registration token. The Azure Video Analyzer IoT edge module must be initialized and connected to the Internet prior to the token expiration date.</dd><dt class="property-"
            title="">
        <span id="token_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#token_python" style="color: inherit; text-decoration: inherit;">token</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The token blob to be provided to the Azure Video Analyzer IoT edge module through the Azure IoT Edge module twin properties.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="expirationdate_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#expirationdate_yaml" style="color: inherit; text-decoration: inherit;">expiration<wbr>Date</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The expiration date of the registration token. The Azure Video Analyzer IoT edge module must be initialized and connected to the Internet prior to the token expiration date.</dd><dt class="property-"
            title="">
        <span id="token_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#token_yaml" style="color: inherit; text-decoration: inherit;">token</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The token blob to be provided to the Azure Video Analyzer IoT edge module through the Azure IoT Edge module twin properties.</dd></dl>
</pulumi-choosable>
</div>





<h2 id="package-details">Package Details</h2>
<dl class="package-details">
	<dt>Repository</dt>
	<dd><a href="https://github.com/pulumi/pulumi-azure-native">Azure Native pulumi/pulumi-azure-native</a></dd>
	<dt>License</dt>
	<dd>Apache-2.0</dd>
</dl>

