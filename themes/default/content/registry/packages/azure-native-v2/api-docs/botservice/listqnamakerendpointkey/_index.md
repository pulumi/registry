
---
title: "listQnAMakerEndpointKey"
title_tag: "azure-native.botservice.listQnAMakerEndpointKey"
meta_desc: "Documentation for the azure-native.botservice.listQnAMakerEndpointKey function with examples, input properties, output properties, and supporting types."
layout: api
no_edit_this_page: true
---



<!-- WARNING: this file was generated by Pulumi Docs Generator. -->
<!-- Do not edit by hand unless you're certain you know what you are doing! -->

Lists the QnA Maker endpoint keys

Azure REST API version: 2023-09-15-preview.

Other available API versions: 2022-09-15.

These versions are not included in the SDK but you can generate them as a local SDK package; see the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.




## Using listQnAMakerEndpointKey {#using}

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
><span class="k">function </span>listQnAMakerEndpointKey<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">ListQnAMakerEndpointKeyArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Promise&lt;<span class="nx"><a href="#result">ListQnAMakerEndpointKeyResult</a></span>></span
><span class="k">
function </span>listQnAMakerEndpointKeyOutput<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">ListQnAMakerEndpointKeyOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Output&lt;<span class="nx"><a href="#result">ListQnAMakerEndpointKeyResult</a></span>></span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="python">
<div class="highlight"><pre class="chroma"><code class="language-python" data-lang="python"
><span class="k">def </span>list_qn_a_maker_endpoint_key<span class="p">(</span><span class="nx">authkey</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                                 <span class="nx">hostname</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                                 <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>ListQnAMakerEndpointKeyResult</span
><span class="k">
def </span>list_qn_a_maker_endpoint_key_output<span class="p">(</span><span class="nx">authkey</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                                 <span class="nx">hostname</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                                 <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>Output[ListQnAMakerEndpointKeyResult]</span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="go">
<div class="highlight"><pre class="chroma"><code class="language-go" data-lang="go"
><span class="k">func </span>ListQnAMakerEndpointKey<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">ListQnAMakerEndpointKeyArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) (*<span class="nx"><a href="#result">ListQnAMakerEndpointKeyResult</a></span>, error)</span
><span class="k">
func </span>ListQnAMakerEndpointKeyOutput<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">ListQnAMakerEndpointKeyOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) ListQnAMakerEndpointKeyResultOutput</span
></code></pre></div>

&gt; Note: This function is named `ListQnAMakerEndpointKey` in the Go SDK.

</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="csharp">
<div class="highlight"><pre class="chroma"><code class="language-csharp" data-lang="csharp"><span class="k">public static class </span><span class="nx">ListQnAMakerEndpointKey </span><span class="p">
{</span><span class="k">
    public static </span>Task&lt;<span class="nx"><a href="#result">ListQnAMakerEndpointKeyResult</a></span>> <span class="p">InvokeAsync(</span><span class="nx">ListQnAMakerEndpointKeyArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="k">
    public static </span>Output&lt;<span class="nx"><a href="#result">ListQnAMakerEndpointKeyResult</a></span>> <span class="p">Invoke(</span><span class="nx">ListQnAMakerEndpointKeyInvokeArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="p">
}</span></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="java">
<div class="highlight"><pre class="chroma"><code class="language-java" data-lang="java"><span class="k">public static CompletableFuture&lt;<span class="nx"><a href="#result">ListQnAMakerEndpointKeyResult</a></span>> </span>listQnAMakerEndpointKey<span class="p">(</span><span class="nx">ListQnAMakerEndpointKeyArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
<span class="k">public static Output&lt;<span class="nx"><a href="#result">ListQnAMakerEndpointKeyResult</a></span>> </span>listQnAMakerEndpointKey<span class="p">(</span><span class="nx">ListQnAMakerEndpointKeyArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
</code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="yaml">
<div class="highlight"><pre class="chroma"><code class="language-yaml" data-lang="yaml"><span class="k">fn::invoke:</span>
<span class="k">&nbsp;&nbsp;function:</span> azure-native:botservice:listQnAMakerEndpointKey
<span class="k">&nbsp;&nbsp;arguments:</span>
<span class="c">&nbsp;&nbsp;&nbsp;&nbsp;# arguments dictionary</span></code></pre></div>
</pulumi-choosable>
</div>



The following arguments are supported:


<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="authkey_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#authkey_csharp" style="color: inherit; text-decoration: inherit;">Authkey</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Subscription key which provides access to this API.</dd><dt class="property-optional"
            title="Optional">
        <span id="hostname_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#hostname_csharp" style="color: inherit; text-decoration: inherit;">Hostname</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>the host name of the QnA Maker endpoint</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="authkey_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#authkey_go" style="color: inherit; text-decoration: inherit;">Authkey</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Subscription key which provides access to this API.</dd><dt class="property-optional"
            title="Optional">
        <span id="hostname_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#hostname_go" style="color: inherit; text-decoration: inherit;">Hostname</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>the host name of the QnA Maker endpoint</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="authkey_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#authkey_java" style="color: inherit; text-decoration: inherit;">authkey</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Subscription key which provides access to this API.</dd><dt class="property-optional"
            title="Optional">
        <span id="hostname_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#hostname_java" style="color: inherit; text-decoration: inherit;">hostname</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>the host name of the QnA Maker endpoint</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="authkey_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#authkey_nodejs" style="color: inherit; text-decoration: inherit;">authkey</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Subscription key which provides access to this API.</dd><dt class="property-optional"
            title="Optional">
        <span id="hostname_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#hostname_nodejs" style="color: inherit; text-decoration: inherit;">hostname</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>the host name of the QnA Maker endpoint</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="authkey_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#authkey_python" style="color: inherit; text-decoration: inherit;">authkey</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Subscription key which provides access to this API.</dd><dt class="property-optional"
            title="Optional">
        <span id="hostname_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#hostname_python" style="color: inherit; text-decoration: inherit;">hostname</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>the host name of the QnA Maker endpoint</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="authkey_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#authkey_yaml" style="color: inherit; text-decoration: inherit;">authkey</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Subscription key which provides access to this API.</dd><dt class="property-optional"
            title="Optional">
        <span id="hostname_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#hostname_yaml" style="color: inherit; text-decoration: inherit;">hostname</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>the host name of the QnA Maker endpoint</dd></dl>
</pulumi-choosable>
</div>




## listQnAMakerEndpointKey Result {#result}

The following output properties are available:



<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="installedversion_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#installedversion_csharp" style="color: inherit; text-decoration: inherit;">Installed<wbr>Version</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Current version of runtime.</dd><dt class="property-"
            title="">
        <span id="laststableversion_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#laststableversion_csharp" style="color: inherit; text-decoration: inherit;">Last<wbr>Stable<wbr>Version</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Latest version of runtime.</dd><dt class="property-"
            title="">
        <span id="primaryendpointkey_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#primaryendpointkey_csharp" style="color: inherit; text-decoration: inherit;">Primary<wbr>Endpoint<wbr>Key</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Primary Access Key.</dd><dt class="property-"
            title="">
        <span id="secondaryendpointkey_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#secondaryendpointkey_csharp" style="color: inherit; text-decoration: inherit;">Secondary<wbr>Endpoint<wbr>Key</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Secondary Access Key.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="installedversion_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#installedversion_go" style="color: inherit; text-decoration: inherit;">Installed<wbr>Version</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Current version of runtime.</dd><dt class="property-"
            title="">
        <span id="laststableversion_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#laststableversion_go" style="color: inherit; text-decoration: inherit;">Last<wbr>Stable<wbr>Version</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Latest version of runtime.</dd><dt class="property-"
            title="">
        <span id="primaryendpointkey_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#primaryendpointkey_go" style="color: inherit; text-decoration: inherit;">Primary<wbr>Endpoint<wbr>Key</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Primary Access Key.</dd><dt class="property-"
            title="">
        <span id="secondaryendpointkey_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#secondaryendpointkey_go" style="color: inherit; text-decoration: inherit;">Secondary<wbr>Endpoint<wbr>Key</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Secondary Access Key.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="installedversion_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#installedversion_java" style="color: inherit; text-decoration: inherit;">installed<wbr>Version</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Current version of runtime.</dd><dt class="property-"
            title="">
        <span id="laststableversion_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#laststableversion_java" style="color: inherit; text-decoration: inherit;">last<wbr>Stable<wbr>Version</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Latest version of runtime.</dd><dt class="property-"
            title="">
        <span id="primaryendpointkey_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#primaryendpointkey_java" style="color: inherit; text-decoration: inherit;">primary<wbr>Endpoint<wbr>Key</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Primary Access Key.</dd><dt class="property-"
            title="">
        <span id="secondaryendpointkey_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#secondaryendpointkey_java" style="color: inherit; text-decoration: inherit;">secondary<wbr>Endpoint<wbr>Key</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Secondary Access Key.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="installedversion_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#installedversion_nodejs" style="color: inherit; text-decoration: inherit;">installed<wbr>Version</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Current version of runtime.</dd><dt class="property-"
            title="">
        <span id="laststableversion_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#laststableversion_nodejs" style="color: inherit; text-decoration: inherit;">last<wbr>Stable<wbr>Version</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Latest version of runtime.</dd><dt class="property-"
            title="">
        <span id="primaryendpointkey_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#primaryendpointkey_nodejs" style="color: inherit; text-decoration: inherit;">primary<wbr>Endpoint<wbr>Key</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Primary Access Key.</dd><dt class="property-"
            title="">
        <span id="secondaryendpointkey_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#secondaryendpointkey_nodejs" style="color: inherit; text-decoration: inherit;">secondary<wbr>Endpoint<wbr>Key</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Secondary Access Key.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="installed_version_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#installed_version_python" style="color: inherit; text-decoration: inherit;">installed_<wbr>version</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Current version of runtime.</dd><dt class="property-"
            title="">
        <span id="last_stable_version_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#last_stable_version_python" style="color: inherit; text-decoration: inherit;">last_<wbr>stable_<wbr>version</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Latest version of runtime.</dd><dt class="property-"
            title="">
        <span id="primary_endpoint_key_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#primary_endpoint_key_python" style="color: inherit; text-decoration: inherit;">primary_<wbr>endpoint_<wbr>key</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Primary Access Key.</dd><dt class="property-"
            title="">
        <span id="secondary_endpoint_key_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#secondary_endpoint_key_python" style="color: inherit; text-decoration: inherit;">secondary_<wbr>endpoint_<wbr>key</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Secondary Access Key.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="installedversion_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#installedversion_yaml" style="color: inherit; text-decoration: inherit;">installed<wbr>Version</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Current version of runtime.</dd><dt class="property-"
            title="">
        <span id="laststableversion_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#laststableversion_yaml" style="color: inherit; text-decoration: inherit;">last<wbr>Stable<wbr>Version</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Latest version of runtime.</dd><dt class="property-"
            title="">
        <span id="primaryendpointkey_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#primaryendpointkey_yaml" style="color: inherit; text-decoration: inherit;">primary<wbr>Endpoint<wbr>Key</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Primary Access Key.</dd><dt class="property-"
            title="">
        <span id="secondaryendpointkey_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#secondaryendpointkey_yaml" style="color: inherit; text-decoration: inherit;">secondary<wbr>Endpoint<wbr>Key</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Secondary Access Key.</dd></dl>
</pulumi-choosable>
</div>





<h2 id="package-details">Package Details</h2>
<dl class="package-details">
	<dt>Repository</dt>
	<dd><a href="https://github.com/pulumi/pulumi-azure-native">Azure Native pulumi/pulumi-azure-native</a></dd>
	<dt>License</dt>
	<dd>Apache-2.0</dd>
</dl>

