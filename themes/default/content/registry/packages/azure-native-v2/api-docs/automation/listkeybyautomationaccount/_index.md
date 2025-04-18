
---
title: "listKeyByAutomationAccount"
title_tag: "azure-native.automation.listKeyByAutomationAccount"
meta_desc: "Documentation for the azure-native.automation.listKeyByAutomationAccount function with examples, input properties, output properties, and supporting types."
layout: api
no_edit_this_page: true
---



<!-- WARNING: this file was generated by Pulumi Docs Generator. -->
<!-- Do not edit by hand unless you're certain you know what you are doing! -->

Retrieve the automation keys for an account.

Azure REST API version: 2023-11-01.

Other available API versions: 2022-08-08, 2024-10-23.

These versions are not included in the SDK but you can generate them as a local SDK package; see the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.




## Using listKeyByAutomationAccount {#using}

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
><span class="k">function </span>listKeyByAutomationAccount<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">ListKeyByAutomationAccountArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Promise&lt;<span class="nx"><a href="#result">ListKeyByAutomationAccountResult</a></span>></span
><span class="k">
function </span>listKeyByAutomationAccountOutput<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">ListKeyByAutomationAccountOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Output&lt;<span class="nx"><a href="#result">ListKeyByAutomationAccountResult</a></span>></span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="python">
<div class="highlight"><pre class="chroma"><code class="language-python" data-lang="python"
><span class="k">def </span>list_key_by_automation_account<span class="p">(</span><span class="nx">automation_account_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                                   <span class="nx">resource_group_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                                   <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>ListKeyByAutomationAccountResult</span
><span class="k">
def </span>list_key_by_automation_account_output<span class="p">(</span><span class="nx">automation_account_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                                   <span class="nx">resource_group_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                                   <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>Output[ListKeyByAutomationAccountResult]</span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="go">
<div class="highlight"><pre class="chroma"><code class="language-go" data-lang="go"
><span class="k">func </span>ListKeyByAutomationAccount<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">ListKeyByAutomationAccountArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) (*<span class="nx"><a href="#result">ListKeyByAutomationAccountResult</a></span>, error)</span
><span class="k">
func </span>ListKeyByAutomationAccountOutput<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">ListKeyByAutomationAccountOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) ListKeyByAutomationAccountResultOutput</span
></code></pre></div>

&gt; Note: This function is named `ListKeyByAutomationAccount` in the Go SDK.

</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="csharp">
<div class="highlight"><pre class="chroma"><code class="language-csharp" data-lang="csharp"><span class="k">public static class </span><span class="nx">ListKeyByAutomationAccount </span><span class="p">
{</span><span class="k">
    public static </span>Task&lt;<span class="nx"><a href="#result">ListKeyByAutomationAccountResult</a></span>> <span class="p">InvokeAsync(</span><span class="nx">ListKeyByAutomationAccountArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="k">
    public static </span>Output&lt;<span class="nx"><a href="#result">ListKeyByAutomationAccountResult</a></span>> <span class="p">Invoke(</span><span class="nx">ListKeyByAutomationAccountInvokeArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="p">
}</span></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="java">
<div class="highlight"><pre class="chroma"><code class="language-java" data-lang="java"><span class="k">public static CompletableFuture&lt;<span class="nx"><a href="#result">ListKeyByAutomationAccountResult</a></span>> </span>listKeyByAutomationAccount<span class="p">(</span><span class="nx">ListKeyByAutomationAccountArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
<span class="k">public static Output&lt;<span class="nx"><a href="#result">ListKeyByAutomationAccountResult</a></span>> </span>listKeyByAutomationAccount<span class="p">(</span><span class="nx">ListKeyByAutomationAccountArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
</code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="yaml">
<div class="highlight"><pre class="chroma"><code class="language-yaml" data-lang="yaml"><span class="k">fn::invoke:</span>
<span class="k">&nbsp;&nbsp;function:</span> azure-native:automation:listKeyByAutomationAccount
<span class="k">&nbsp;&nbsp;arguments:</span>
<span class="c">&nbsp;&nbsp;&nbsp;&nbsp;# arguments dictionary</span></code></pre></div>
</pulumi-choosable>
</div>



The following arguments are supported:


<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="automationaccountname_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#automationaccountname_csharp" style="color: inherit; text-decoration: inherit;">Automation<wbr>Account<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the automation account.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_csharp" style="color: inherit; text-decoration: inherit;">Resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Name of an Azure Resource group.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="automationaccountname_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#automationaccountname_go" style="color: inherit; text-decoration: inherit;">Automation<wbr>Account<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the automation account.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_go" style="color: inherit; text-decoration: inherit;">Resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Name of an Azure Resource group.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="automationaccountname_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#automationaccountname_java" style="color: inherit; text-decoration: inherit;">automation<wbr>Account<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The name of the automation account.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_java" style="color: inherit; text-decoration: inherit;">resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Name of an Azure Resource group.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="automationaccountname_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#automationaccountname_nodejs" style="color: inherit; text-decoration: inherit;">automation<wbr>Account<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the automation account.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_nodejs" style="color: inherit; text-decoration: inherit;">resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Name of an Azure Resource group.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="automation_account_name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#automation_account_name_python" style="color: inherit; text-decoration: inherit;">automation_<wbr>account_<wbr>name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The name of the automation account.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resource_group_name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resource_group_name_python" style="color: inherit; text-decoration: inherit;">resource_<wbr>group_<wbr>name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Name of an Azure Resource group.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="automationaccountname_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#automationaccountname_yaml" style="color: inherit; text-decoration: inherit;">automation<wbr>Account<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The name of the automation account.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="resourcegroupname_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resourcegroupname_yaml" style="color: inherit; text-decoration: inherit;">resource<wbr>Group<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Name of an Azure Resource group.</dd></dl>
</pulumi-choosable>
</div>




## listKeyByAutomationAccount Result {#result}

The following output properties are available:



<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="keys_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#keys_csharp" style="color: inherit; text-decoration: inherit;">Keys</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#keyresponse">List&lt;Pulumi.<wbr>Azure<wbr>Native.<wbr>Automation.<wbr>Outputs.<wbr>Key<wbr>Response&gt;</a></span>
    </dt>
    <dd>Lists the automation keys.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="keys_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#keys_go" style="color: inherit; text-decoration: inherit;">Keys</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#keyresponse">[]Key<wbr>Response</a></span>
    </dt>
    <dd>Lists the automation keys.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="keys_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#keys_java" style="color: inherit; text-decoration: inherit;">keys</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#keyresponse">List&lt;Key<wbr>Response&gt;</a></span>
    </dt>
    <dd>Lists the automation keys.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="keys_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#keys_nodejs" style="color: inherit; text-decoration: inherit;">keys</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#keyresponse">Key<wbr>Response[]</a></span>
    </dt>
    <dd>Lists the automation keys.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="keys_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#keys_python" style="color: inherit; text-decoration: inherit;">keys</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#keyresponse">Sequence[Key<wbr>Response]</a></span>
    </dt>
    <dd>Lists the automation keys.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="keys_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#keys_yaml" style="color: inherit; text-decoration: inherit;">keys</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#keyresponse">List&lt;Property Map&gt;</a></span>
    </dt>
    <dd>Lists the automation keys.</dd></dl>
</pulumi-choosable>
</div>




## Supporting Types


<h4 id="keyresponse">Key<wbr>Response</h4>



<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="keyname_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#keyname_csharp" style="color: inherit; text-decoration: inherit;">Key<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Automation key name.</dd><dt class="property-required"
            title="Required">
        <span id="permissions_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#permissions_csharp" style="color: inherit; text-decoration: inherit;">Permissions</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Automation key permissions.</dd><dt class="property-required"
            title="Required">
        <span id="value_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#value_csharp" style="color: inherit; text-decoration: inherit;">Value</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Value of the Automation Key used for registration.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="keyname_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#keyname_go" style="color: inherit; text-decoration: inherit;">Key<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Automation key name.</dd><dt class="property-required"
            title="Required">
        <span id="permissions_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#permissions_go" style="color: inherit; text-decoration: inherit;">Permissions</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Automation key permissions.</dd><dt class="property-required"
            title="Required">
        <span id="value_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#value_go" style="color: inherit; text-decoration: inherit;">Value</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Value of the Automation Key used for registration.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="keyname_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#keyname_java" style="color: inherit; text-decoration: inherit;">key<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Automation key name.</dd><dt class="property-required"
            title="Required">
        <span id="permissions_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#permissions_java" style="color: inherit; text-decoration: inherit;">permissions</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Automation key permissions.</dd><dt class="property-required"
            title="Required">
        <span id="value_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#value_java" style="color: inherit; text-decoration: inherit;">value</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Value of the Automation Key used for registration.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="keyname_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#keyname_nodejs" style="color: inherit; text-decoration: inherit;">key<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Automation key name.</dd><dt class="property-required"
            title="Required">
        <span id="permissions_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#permissions_nodejs" style="color: inherit; text-decoration: inherit;">permissions</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Automation key permissions.</dd><dt class="property-required"
            title="Required">
        <span id="value_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#value_nodejs" style="color: inherit; text-decoration: inherit;">value</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Value of the Automation Key used for registration.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="key_name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#key_name_python" style="color: inherit; text-decoration: inherit;">key_<wbr>name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Automation key name.</dd><dt class="property-required"
            title="Required">
        <span id="permissions_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#permissions_python" style="color: inherit; text-decoration: inherit;">permissions</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Automation key permissions.</dd><dt class="property-required"
            title="Required">
        <span id="value_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#value_python" style="color: inherit; text-decoration: inherit;">value</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Value of the Automation Key used for registration.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="keyname_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#keyname_yaml" style="color: inherit; text-decoration: inherit;">key<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Automation key name.</dd><dt class="property-required"
            title="Required">
        <span id="permissions_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#permissions_yaml" style="color: inherit; text-decoration: inherit;">permissions</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Automation key permissions.</dd><dt class="property-required"
            title="Required">
        <span id="value_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#value_yaml" style="color: inherit; text-decoration: inherit;">value</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Value of the Automation Key used for registration.</dd></dl>
</pulumi-choosable>
</div>





<h2 id="package-details">Package Details</h2>
<dl class="package-details">
	<dt>Repository</dt>
	<dd><a href="https://github.com/pulumi/pulumi-azure-native">Azure Native pulumi/pulumi-azure-native</a></dd>
	<dt>License</dt>
	<dd>Apache-2.0</dd>
</dl>

