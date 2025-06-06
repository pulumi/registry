
---
title: "getConsole"
title_tag: "azure-native.portal.getConsole"
meta_desc: "Documentation for the azure-native.portal.getConsole function with examples, input properties, output properties, and supporting types."
layout: api
no_edit_this_page: true
---



<!-- WARNING: this file was generated by Pulumi Docs Generator. -->
<!-- Do not edit by hand unless you're certain you know what you are doing! -->

Gets the console for the user.

Azure REST API version: 2018-10-01.




## Using getConsole {#using}

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
><span class="k">function </span>getConsole<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">GetConsoleArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Promise&lt;<span class="nx"><a href="#result">GetConsoleResult</a></span>></span
><span class="k">
function </span>getConsoleOutput<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">GetConsoleOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Output&lt;<span class="nx"><a href="#result">GetConsoleResult</a></span>></span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="python">
<div class="highlight"><pre class="chroma"><code class="language-python" data-lang="python"
><span class="k">def </span>get_console<span class="p">(</span><span class="nx">console_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>GetConsoleResult</span
><span class="k">
def </span>get_console_output<span class="p">(</span><span class="nx">console_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>Output[GetConsoleResult]</span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="go">
<div class="highlight"><pre class="chroma"><code class="language-go" data-lang="go"
><span class="k">func </span>LookupConsole<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">LookupConsoleArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) (*<span class="nx"><a href="#result">LookupConsoleResult</a></span>, error)</span
><span class="k">
func </span>LookupConsoleOutput<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">LookupConsoleOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) LookupConsoleResultOutput</span
></code></pre></div>

&gt; Note: This function is named `LookupConsole` in the Go SDK.

</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="csharp">
<div class="highlight"><pre class="chroma"><code class="language-csharp" data-lang="csharp"><span class="k">public static class </span><span class="nx">GetConsole </span><span class="p">
{</span><span class="k">
    public static </span>Task&lt;<span class="nx"><a href="#result">GetConsoleResult</a></span>> <span class="p">InvokeAsync(</span><span class="nx">GetConsoleArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="k">
    public static </span>Output&lt;<span class="nx"><a href="#result">GetConsoleResult</a></span>> <span class="p">Invoke(</span><span class="nx">GetConsoleInvokeArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="p">
}</span></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="java">
<div class="highlight"><pre class="chroma"><code class="language-java" data-lang="java"><span class="k">public static CompletableFuture&lt;<span class="nx"><a href="#result">GetConsoleResult</a></span>> </span>getConsole<span class="p">(</span><span class="nx">GetConsoleArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
<span class="k">public static Output&lt;<span class="nx"><a href="#result">GetConsoleResult</a></span>> </span>getConsole<span class="p">(</span><span class="nx">GetConsoleArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
</code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="yaml">
<div class="highlight"><pre class="chroma"><code class="language-yaml" data-lang="yaml"><span class="k">fn::invoke:</span>
<span class="k">&nbsp;&nbsp;function:</span> azure-native:portal:getConsole
<span class="k">&nbsp;&nbsp;arguments:</span>
<span class="c">&nbsp;&nbsp;&nbsp;&nbsp;# arguments dictionary</span></code></pre></div>
</pulumi-choosable>
</div>



The following arguments are supported:


<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="consolename_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#consolename_csharp" style="color: inherit; text-decoration: inherit;">Console<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the console</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="consolename_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#consolename_go" style="color: inherit; text-decoration: inherit;">Console<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the console</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="consolename_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#consolename_java" style="color: inherit; text-decoration: inherit;">console<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The name of the console</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="consolename_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#consolename_nodejs" style="color: inherit; text-decoration: inherit;">console<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The name of the console</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="console_name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#console_name_python" style="color: inherit; text-decoration: inherit;">console_<wbr>name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The name of the console</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="consolename_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#consolename_yaml" style="color: inherit; text-decoration: inherit;">console<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The name of the console</dd></dl>
</pulumi-choosable>
</div>




## getConsole Result {#result}

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
        <span class="property-type"><a href="#consolepropertiesresponse">Pulumi.<wbr>Azure<wbr>Native.<wbr>Portal.<wbr>Outputs.<wbr>Console<wbr>Properties<wbr>Response</a></span>
    </dt>
    <dd>Cloud shell console properties.</dd></dl>
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
        <span class="property-type"><a href="#consolepropertiesresponse">Console<wbr>Properties<wbr>Response</a></span>
    </dt>
    <dd>Cloud shell console properties.</dd></dl>
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
        <span class="property-type"><a href="#consolepropertiesresponse">Console<wbr>Properties<wbr>Response</a></span>
    </dt>
    <dd>Cloud shell console properties.</dd></dl>
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
        <span class="property-type"><a href="#consolepropertiesresponse">Console<wbr>Properties<wbr>Response</a></span>
    </dt>
    <dd>Cloud shell console properties.</dd></dl>
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
        <span class="property-type"><a href="#consolepropertiesresponse">Console<wbr>Properties<wbr>Response</a></span>
    </dt>
    <dd>Cloud shell console properties.</dd></dl>
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
        <span class="property-type"><a href="#consolepropertiesresponse">Property Map</a></span>
    </dt>
    <dd>Cloud shell console properties.</dd></dl>
</pulumi-choosable>
</div>




## Supporting Types


<h4 id="consolepropertiesresponse">Console<wbr>Properties<wbr>Response</h4>



<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="ostype_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#ostype_csharp" style="color: inherit; text-decoration: inherit;">Os<wbr>Type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The operating system type of the cloud shell.</dd><dt class="property-required"
            title="Required">
        <span id="provisioningstate_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#provisioningstate_csharp" style="color: inherit; text-decoration: inherit;">Provisioning<wbr>State</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Provisioning state of the console.</dd><dt class="property-required"
            title="Required">
        <span id="uri_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#uri_csharp" style="color: inherit; text-decoration: inherit;">Uri</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Uri of the console.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="ostype_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#ostype_go" style="color: inherit; text-decoration: inherit;">Os<wbr>Type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The operating system type of the cloud shell.</dd><dt class="property-required"
            title="Required">
        <span id="provisioningstate_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#provisioningstate_go" style="color: inherit; text-decoration: inherit;">Provisioning<wbr>State</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Provisioning state of the console.</dd><dt class="property-required"
            title="Required">
        <span id="uri_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#uri_go" style="color: inherit; text-decoration: inherit;">Uri</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Uri of the console.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="ostype_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#ostype_java" style="color: inherit; text-decoration: inherit;">os<wbr>Type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The operating system type of the cloud shell.</dd><dt class="property-required"
            title="Required">
        <span id="provisioningstate_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#provisioningstate_java" style="color: inherit; text-decoration: inherit;">provisioning<wbr>State</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Provisioning state of the console.</dd><dt class="property-required"
            title="Required">
        <span id="uri_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#uri_java" style="color: inherit; text-decoration: inherit;">uri</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Uri of the console.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="ostype_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#ostype_nodejs" style="color: inherit; text-decoration: inherit;">os<wbr>Type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The operating system type of the cloud shell.</dd><dt class="property-required"
            title="Required">
        <span id="provisioningstate_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#provisioningstate_nodejs" style="color: inherit; text-decoration: inherit;">provisioning<wbr>State</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Provisioning state of the console.</dd><dt class="property-required"
            title="Required">
        <span id="uri_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#uri_nodejs" style="color: inherit; text-decoration: inherit;">uri</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Uri of the console.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="os_type_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#os_type_python" style="color: inherit; text-decoration: inherit;">os_<wbr>type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The operating system type of the cloud shell.</dd><dt class="property-required"
            title="Required">
        <span id="provisioning_state_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#provisioning_state_python" style="color: inherit; text-decoration: inherit;">provisioning_<wbr>state</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Provisioning state of the console.</dd><dt class="property-required"
            title="Required">
        <span id="uri_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#uri_python" style="color: inherit; text-decoration: inherit;">uri</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Uri of the console.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="ostype_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#ostype_yaml" style="color: inherit; text-decoration: inherit;">os<wbr>Type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The operating system type of the cloud shell.</dd><dt class="property-required"
            title="Required">
        <span id="provisioningstate_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#provisioningstate_yaml" style="color: inherit; text-decoration: inherit;">provisioning<wbr>State</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Provisioning state of the console.</dd><dt class="property-required"
            title="Required">
        <span id="uri_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#uri_yaml" style="color: inherit; text-decoration: inherit;">uri</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Uri of the console.</dd></dl>
</pulumi-choosable>
</div>





<h2 id="package-details">Package Details</h2>
<dl class="package-details">
	<dt>Repository</dt>
	<dd><a href="https://github.com/pulumi/pulumi-azure-native">Azure Native pulumi/pulumi-azure-native</a></dd>
	<dt>License</dt>
	<dd>Apache-2.0</dd>
</dl>

