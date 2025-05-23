
---
title: "listWebAppHostKeysSlot"
title_tag: "azure-native.web.listWebAppHostKeysSlot"
meta_desc: "Documentation for the azure-native.web.listWebAppHostKeysSlot function with examples, input properties, output properties, and supporting types."
layout: api
no_edit_this_page: true
---



<!-- WARNING: this file was generated by Pulumi Docs Generator. -->
<!-- Do not edit by hand unless you're certain you know what you are doing! -->

Description for Get host secrets for a function app.

Azure REST API version: 2024-04-01.

Other available API versions: 2020-10-01, 2022-09-01.

These versions are not included in the SDK but you can generate them as a local SDK package; see the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.




## Using listWebAppHostKeysSlot {#using}

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
><span class="k">function </span>listWebAppHostKeysSlot<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">ListWebAppHostKeysSlotArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Promise&lt;<span class="nx"><a href="#result">ListWebAppHostKeysSlotResult</a></span>></span
><span class="k">
function </span>listWebAppHostKeysSlotOutput<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">ListWebAppHostKeysSlotOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Output&lt;<span class="nx"><a href="#result">ListWebAppHostKeysSlotResult</a></span>></span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="python">
<div class="highlight"><pre class="chroma"><code class="language-python" data-lang="python"
><span class="k">def </span>list_web_app_host_keys_slot<span class="p">(</span><span class="nx">name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                                <span class="nx">resource_group_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                                <span class="nx">slot</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                                <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>ListWebAppHostKeysSlotResult</span
><span class="k">
def </span>list_web_app_host_keys_slot_output<span class="p">(</span><span class="nx">name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                                <span class="nx">resource_group_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                                <span class="nx">slot</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                                <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>Output[ListWebAppHostKeysSlotResult]</span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="go">
<div class="highlight"><pre class="chroma"><code class="language-go" data-lang="go"
><span class="k">func </span>ListWebAppHostKeysSlot<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">ListWebAppHostKeysSlotArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) (*<span class="nx"><a href="#result">ListWebAppHostKeysSlotResult</a></span>, error)</span
><span class="k">
func </span>ListWebAppHostKeysSlotOutput<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">ListWebAppHostKeysSlotOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) ListWebAppHostKeysSlotResultOutput</span
></code></pre></div>

&gt; Note: This function is named `ListWebAppHostKeysSlot` in the Go SDK.

</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="csharp">
<div class="highlight"><pre class="chroma"><code class="language-csharp" data-lang="csharp"><span class="k">public static class </span><span class="nx">ListWebAppHostKeysSlot </span><span class="p">
{</span><span class="k">
    public static </span>Task&lt;<span class="nx"><a href="#result">ListWebAppHostKeysSlotResult</a></span>> <span class="p">InvokeAsync(</span><span class="nx">ListWebAppHostKeysSlotArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="k">
    public static </span>Output&lt;<span class="nx"><a href="#result">ListWebAppHostKeysSlotResult</a></span>> <span class="p">Invoke(</span><span class="nx">ListWebAppHostKeysSlotInvokeArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="p">
}</span></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="java">
<div class="highlight"><pre class="chroma"><code class="language-java" data-lang="java"><span class="k">public static CompletableFuture&lt;<span class="nx"><a href="#result">ListWebAppHostKeysSlotResult</a></span>> </span>listWebAppHostKeysSlot<span class="p">(</span><span class="nx">ListWebAppHostKeysSlotArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
<span class="k">public static Output&lt;<span class="nx"><a href="#result">ListWebAppHostKeysSlotResult</a></span>> </span>listWebAppHostKeysSlot<span class="p">(</span><span class="nx">ListWebAppHostKeysSlotArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
</code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="yaml">
<div class="highlight"><pre class="chroma"><code class="language-yaml" data-lang="yaml"><span class="k">fn::invoke:</span>
<span class="k">&nbsp;&nbsp;function:</span> azure-native:web:listWebAppHostKeysSlot
<span class="k">&nbsp;&nbsp;arguments:</span>
<span class="c">&nbsp;&nbsp;&nbsp;&nbsp;# arguments dictionary</span></code></pre></div>
</pulumi-choosable>
</div>



The following arguments are supported:


<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="name_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_csharp" style="color: inherit; text-decoration: inherit;">Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Site name.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_csharp" style="color: inherit; text-decoration: inherit;">Resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Name of the resource group to which the resource belongs.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="slot_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#slot_csharp" style="color: inherit; text-decoration: inherit;">Slot</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Name of the deployment slot.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="name_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_go" style="color: inherit; text-decoration: inherit;">Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Site name.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_go" style="color: inherit; text-decoration: inherit;">Resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Name of the resource group to which the resource belongs.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="slot_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#slot_go" style="color: inherit; text-decoration: inherit;">Slot</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Name of the deployment slot.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="name_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_java" style="color: inherit; text-decoration: inherit;">name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Site name.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_java" style="color: inherit; text-decoration: inherit;">resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Name of the resource group to which the resource belongs.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="slot_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#slot_java" style="color: inherit; text-decoration: inherit;">slot</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Name of the deployment slot.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="name_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_nodejs" style="color: inherit; text-decoration: inherit;">name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Site name.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_nodejs" style="color: inherit; text-decoration: inherit;">resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Name of the resource group to which the resource belongs.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="slot_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#slot_nodejs" style="color: inherit; text-decoration: inherit;">slot</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Name of the deployment slot.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_python" style="color: inherit; text-decoration: inherit;">name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Site name.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resource_group_name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resource_group_name_python" style="color: inherit; text-decoration: inherit;">resource_<wbr>group_<wbr>name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Name of the resource group to which the resource belongs.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="slot_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#slot_python" style="color: inherit; text-decoration: inherit;">slot</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Name of the deployment slot.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="name_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_yaml" style="color: inherit; text-decoration: inherit;">name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Site name.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_yaml" style="color: inherit; text-decoration: inherit;">resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Name of the resource group to which the resource belongs.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="slot_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#slot_yaml" style="color: inherit; text-decoration: inherit;">slot</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Name of the deployment slot.</dd></dl>
</pulumi-choosable>
</div>




## listWebAppHostKeysSlot Result {#result}

The following output properties are available:



<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="functionkeys_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#functionkeys_csharp" style="color: inherit; text-decoration: inherit;">Function<wbr>Keys</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Dictionary&lt;string, string&gt;</span>
    </dt>
    <dd>Host level function keys.</dd><dt class="property-"
            title="">
        <span id="masterkey_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#masterkey_csharp" style="color: inherit; text-decoration: inherit;">Master<wbr>Key</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Secret key.</dd><dt class="property-"
            title="">
        <span id="systemkeys_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#systemkeys_csharp" style="color: inherit; text-decoration: inherit;">System<wbr>Keys</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Dictionary&lt;string, string&gt;</span>
    </dt>
    <dd>System keys.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="functionkeys_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#functionkeys_go" style="color: inherit; text-decoration: inherit;">Function<wbr>Keys</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">map[string]string</span>
    </dt>
    <dd>Host level function keys.</dd><dt class="property-"
            title="">
        <span id="masterkey_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#masterkey_go" style="color: inherit; text-decoration: inherit;">Master<wbr>Key</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Secret key.</dd><dt class="property-"
            title="">
        <span id="systemkeys_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#systemkeys_go" style="color: inherit; text-decoration: inherit;">System<wbr>Keys</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">map[string]string</span>
    </dt>
    <dd>System keys.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="functionkeys_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#functionkeys_java" style="color: inherit; text-decoration: inherit;">function<wbr>Keys</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Map&lt;String,String&gt;</span>
    </dt>
    <dd>Host level function keys.</dd><dt class="property-"
            title="">
        <span id="masterkey_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#masterkey_java" style="color: inherit; text-decoration: inherit;">master<wbr>Key</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Secret key.</dd><dt class="property-"
            title="">
        <span id="systemkeys_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#systemkeys_java" style="color: inherit; text-decoration: inherit;">system<wbr>Keys</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Map&lt;String,String&gt;</span>
    </dt>
    <dd>System keys.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="functionkeys_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#functionkeys_nodejs" style="color: inherit; text-decoration: inherit;">function<wbr>Keys</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">{[key: string]: string}</span>
    </dt>
    <dd>Host level function keys.</dd><dt class="property-"
            title="">
        <span id="masterkey_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#masterkey_nodejs" style="color: inherit; text-decoration: inherit;">master<wbr>Key</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Secret key.</dd><dt class="property-"
            title="">
        <span id="systemkeys_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#systemkeys_nodejs" style="color: inherit; text-decoration: inherit;">system<wbr>Keys</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">{[key: string]: string}</span>
    </dt>
    <dd>System keys.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="function_keys_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#function_keys_python" style="color: inherit; text-decoration: inherit;">function_<wbr>keys</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Mapping[str, str]</span>
    </dt>
    <dd>Host level function keys.</dd><dt class="property-"
            title="">
        <span id="master_key_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#master_key_python" style="color: inherit; text-decoration: inherit;">master_<wbr>key</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Secret key.</dd><dt class="property-"
            title="">
        <span id="system_keys_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#system_keys_python" style="color: inherit; text-decoration: inherit;">system_<wbr>keys</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Mapping[str, str]</span>
    </dt>
    <dd>System keys.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="functionkeys_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#functionkeys_yaml" style="color: inherit; text-decoration: inherit;">function<wbr>Keys</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Map&lt;String&gt;</span>
    </dt>
    <dd>Host level function keys.</dd><dt class="property-"
            title="">
        <span id="masterkey_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#masterkey_yaml" style="color: inherit; text-decoration: inherit;">master<wbr>Key</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Secret key.</dd><dt class="property-"
            title="">
        <span id="systemkeys_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#systemkeys_yaml" style="color: inherit; text-decoration: inherit;">system<wbr>Keys</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Map&lt;String&gt;</span>
    </dt>
    <dd>System keys.</dd></dl>
</pulumi-choosable>
</div>





<h2 id="package-details">Package Details</h2>
<dl class="package-details">
	<dt>Repository</dt>
	<dd><a href="https://github.com/pulumi/pulumi-azure-native">Azure Native pulumi/pulumi-azure-native</a></dd>
	<dt>License</dt>
	<dd>Apache-2.0</dd>
</dl>

