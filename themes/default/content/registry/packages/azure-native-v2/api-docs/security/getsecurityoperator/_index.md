
---
title: "getSecurityOperator"
title_tag: "azure-native.security.getSecurityOperator"
meta_desc: "Documentation for the azure-native.security.getSecurityOperator function with examples, input properties, output properties, and supporting types."
layout: api
no_edit_this_page: true
---



<!-- WARNING: this file was generated by Pulumi Docs Generator. -->
<!-- Do not edit by hand unless you're certain you know what you are doing! -->

Get a specific security operator for the requested scope.

Azure REST API version: 2023-01-01-preview.




## Using getSecurityOperator {#using}

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
><span class="k">function </span>getSecurityOperator<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">GetSecurityOperatorArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Promise&lt;<span class="nx"><a href="#result">GetSecurityOperatorResult</a></span>></span
><span class="k">
function </span>getSecurityOperatorOutput<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">GetSecurityOperatorOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Output&lt;<span class="nx"><a href="#result">GetSecurityOperatorResult</a></span>></span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="python">
<div class="highlight"><pre class="chroma"><code class="language-python" data-lang="python"
><span class="k">def </span>get_security_operator<span class="p">(</span><span class="nx">pricing_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                          <span class="nx">security_operator_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                          <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>GetSecurityOperatorResult</span
><span class="k">
def </span>get_security_operator_output<span class="p">(</span><span class="nx">pricing_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                          <span class="nx">security_operator_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                          <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>Output[GetSecurityOperatorResult]</span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="go">
<div class="highlight"><pre class="chroma"><code class="language-go" data-lang="go"
><span class="k">func </span>LookupSecurityOperator<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">LookupSecurityOperatorArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) (*<span class="nx"><a href="#result">LookupSecurityOperatorResult</a></span>, error)</span
><span class="k">
func </span>LookupSecurityOperatorOutput<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">LookupSecurityOperatorOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) LookupSecurityOperatorResultOutput</span
></code></pre></div>

&gt; Note: This function is named `LookupSecurityOperator` in the Go SDK.

</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="csharp">
<div class="highlight"><pre class="chroma"><code class="language-csharp" data-lang="csharp"><span class="k">public static class </span><span class="nx">GetSecurityOperator </span><span class="p">
{</span><span class="k">
    public static </span>Task&lt;<span class="nx"><a href="#result">GetSecurityOperatorResult</a></span>> <span class="p">InvokeAsync(</span><span class="nx">GetSecurityOperatorArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="k">
    public static </span>Output&lt;<span class="nx"><a href="#result">GetSecurityOperatorResult</a></span>> <span class="p">Invoke(</span><span class="nx">GetSecurityOperatorInvokeArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="p">
}</span></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="java">
<div class="highlight"><pre class="chroma"><code class="language-java" data-lang="java"><span class="k">public static CompletableFuture&lt;<span class="nx"><a href="#result">GetSecurityOperatorResult</a></span>> </span>getSecurityOperator<span class="p">(</span><span class="nx">GetSecurityOperatorArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
<span class="k">public static Output&lt;<span class="nx"><a href="#result">GetSecurityOperatorResult</a></span>> </span>getSecurityOperator<span class="p">(</span><span class="nx">GetSecurityOperatorArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
</code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="yaml">
<div class="highlight"><pre class="chroma"><code class="language-yaml" data-lang="yaml"><span class="k">fn::invoke:</span>
<span class="k">&nbsp;&nbsp;function:</span> azure-native:security:getSecurityOperator
<span class="k">&nbsp;&nbsp;arguments:</span>
<span class="c">&nbsp;&nbsp;&nbsp;&nbsp;# arguments dictionary</span></code></pre></div>
</pulumi-choosable>
</div>



The following arguments are supported:


<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="pricingname_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#pricingname_csharp" style="color: inherit; text-decoration: inherit;">Pricing<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>name of the pricing configuration</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="securityoperatorname_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#securityoperatorname_csharp" style="color: inherit; text-decoration: inherit;">Security<wbr>Operator<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>name of the securityOperator</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="pricingname_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#pricingname_go" style="color: inherit; text-decoration: inherit;">Pricing<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>name of the pricing configuration</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="securityoperatorname_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#securityoperatorname_go" style="color: inherit; text-decoration: inherit;">Security<wbr>Operator<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>name of the securityOperator</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="pricingname_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#pricingname_java" style="color: inherit; text-decoration: inherit;">pricing<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>name of the pricing configuration</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="securityoperatorname_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#securityoperatorname_java" style="color: inherit; text-decoration: inherit;">security<wbr>Operator<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>name of the securityOperator</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="pricingname_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#pricingname_nodejs" style="color: inherit; text-decoration: inherit;">pricing<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>name of the pricing configuration</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="securityoperatorname_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#securityoperatorname_nodejs" style="color: inherit; text-decoration: inherit;">security<wbr>Operator<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>name of the securityOperator</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="pricing_name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#pricing_name_python" style="color: inherit; text-decoration: inherit;">pricing_<wbr>name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>name of the pricing configuration</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="security_operator_name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#security_operator_name_python" style="color: inherit; text-decoration: inherit;">security_<wbr>operator_<wbr>name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>name of the securityOperator</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-required property-replacement"
            title="Required">
        <span id="pricingname_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#pricingname_yaml" style="color: inherit; text-decoration: inherit;">pricing<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>name of the pricing configuration</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="securityoperatorname_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#securityoperatorname_yaml" style="color: inherit; text-decoration: inherit;">security<wbr>Operator<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>name of the securityOperator</dd></dl>
</pulumi-choosable>
</div>




## getSecurityOperator Result {#result}

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
    <dd>Resource Id</dd><dt class="property-"
            title="">
        <span id="name_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_csharp" style="color: inherit; text-decoration: inherit;">Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Resource name</dd><dt class="property-"
            title="">
        <span id="type_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_csharp" style="color: inherit; text-decoration: inherit;">Type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Resource type</dd><dt class="property-"
            title="">
        <span id="identity_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#identity_csharp" style="color: inherit; text-decoration: inherit;">Identity</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#identityresponse">Pulumi.<wbr>Azure<wbr>Native.<wbr>Security.<wbr>Outputs.<wbr>Identity<wbr>Response</a></span>
    </dt>
    <dd>Identity for the resource.</dd></dl>
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
    <dd>Resource Id</dd><dt class="property-"
            title="">
        <span id="name_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_go" style="color: inherit; text-decoration: inherit;">Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Resource name</dd><dt class="property-"
            title="">
        <span id="type_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_go" style="color: inherit; text-decoration: inherit;">Type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Resource type</dd><dt class="property-"
            title="">
        <span id="identity_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#identity_go" style="color: inherit; text-decoration: inherit;">Identity</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#identityresponse">Identity<wbr>Response</a></span>
    </dt>
    <dd>Identity for the resource.</dd></dl>
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
    <dd>Resource Id</dd><dt class="property-"
            title="">
        <span id="name_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_java" style="color: inherit; text-decoration: inherit;">name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Resource name</dd><dt class="property-"
            title="">
        <span id="type_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_java" style="color: inherit; text-decoration: inherit;">type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Resource type</dd><dt class="property-"
            title="">
        <span id="identity_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#identity_java" style="color: inherit; text-decoration: inherit;">identity</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#identityresponse">Identity<wbr>Response</a></span>
    </dt>
    <dd>Identity for the resource.</dd></dl>
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
    <dd>Resource Id</dd><dt class="property-"
            title="">
        <span id="name_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_nodejs" style="color: inherit; text-decoration: inherit;">name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Resource name</dd><dt class="property-"
            title="">
        <span id="type_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_nodejs" style="color: inherit; text-decoration: inherit;">type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Resource type</dd><dt class="property-"
            title="">
        <span id="identity_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#identity_nodejs" style="color: inherit; text-decoration: inherit;">identity</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#identityresponse">Identity<wbr>Response</a></span>
    </dt>
    <dd>Identity for the resource.</dd></dl>
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
    <dd>Resource Id</dd><dt class="property-"
            title="">
        <span id="name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_python" style="color: inherit; text-decoration: inherit;">name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Resource name</dd><dt class="property-"
            title="">
        <span id="type_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_python" style="color: inherit; text-decoration: inherit;">type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Resource type</dd><dt class="property-"
            title="">
        <span id="identity_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#identity_python" style="color: inherit; text-decoration: inherit;">identity</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#identityresponse">Identity<wbr>Response</a></span>
    </dt>
    <dd>Identity for the resource.</dd></dl>
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
    <dd>Resource Id</dd><dt class="property-"
            title="">
        <span id="name_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_yaml" style="color: inherit; text-decoration: inherit;">name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Resource name</dd><dt class="property-"
            title="">
        <span id="type_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_yaml" style="color: inherit; text-decoration: inherit;">type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Resource type</dd><dt class="property-"
            title="">
        <span id="identity_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#identity_yaml" style="color: inherit; text-decoration: inherit;">identity</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#identityresponse">Property Map</a></span>
    </dt>
    <dd>Identity for the resource.</dd></dl>
</pulumi-choosable>
</div>




## Supporting Types


<h4 id="identityresponse">Identity<wbr>Response</h4>



<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="principalid_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#principalid_csharp" style="color: inherit; text-decoration: inherit;">Principal<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The principal ID of resource identity.</dd><dt class="property-required"
            title="Required">
        <span id="tenantid_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#tenantid_csharp" style="color: inherit; text-decoration: inherit;">Tenant<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The tenant ID of resource.</dd><dt class="property-optional"
            title="Optional">
        <span id="type_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_csharp" style="color: inherit; text-decoration: inherit;">Type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The identity type.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="principalid_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#principalid_go" style="color: inherit; text-decoration: inherit;">Principal<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The principal ID of resource identity.</dd><dt class="property-required"
            title="Required">
        <span id="tenantid_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#tenantid_go" style="color: inherit; text-decoration: inherit;">Tenant<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The tenant ID of resource.</dd><dt class="property-optional"
            title="Optional">
        <span id="type_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_go" style="color: inherit; text-decoration: inherit;">Type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The identity type.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="principalid_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#principalid_java" style="color: inherit; text-decoration: inherit;">principal<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The principal ID of resource identity.</dd><dt class="property-required"
            title="Required">
        <span id="tenantid_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#tenantid_java" style="color: inherit; text-decoration: inherit;">tenant<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The tenant ID of resource.</dd><dt class="property-optional"
            title="Optional">
        <span id="type_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_java" style="color: inherit; text-decoration: inherit;">type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The identity type.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="principalid_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#principalid_nodejs" style="color: inherit; text-decoration: inherit;">principal<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The principal ID of resource identity.</dd><dt class="property-required"
            title="Required">
        <span id="tenantid_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#tenantid_nodejs" style="color: inherit; text-decoration: inherit;">tenant<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The tenant ID of resource.</dd><dt class="property-optional"
            title="Optional">
        <span id="type_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_nodejs" style="color: inherit; text-decoration: inherit;">type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The identity type.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="principal_id_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#principal_id_python" style="color: inherit; text-decoration: inherit;">principal_<wbr>id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The principal ID of resource identity.</dd><dt class="property-required"
            title="Required">
        <span id="tenant_id_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#tenant_id_python" style="color: inherit; text-decoration: inherit;">tenant_<wbr>id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The tenant ID of resource.</dd><dt class="property-optional"
            title="Optional">
        <span id="type_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_python" style="color: inherit; text-decoration: inherit;">type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The identity type.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="principalid_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#principalid_yaml" style="color: inherit; text-decoration: inherit;">principal<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The principal ID of resource identity.</dd><dt class="property-required"
            title="Required">
        <span id="tenantid_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#tenantid_yaml" style="color: inherit; text-decoration: inherit;">tenant<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The tenant ID of resource.</dd><dt class="property-optional"
            title="Optional">
        <span id="type_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#type_yaml" style="color: inherit; text-decoration: inherit;">type</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The identity type.</dd></dl>
</pulumi-choosable>
</div>





<h2 id="package-details">Package Details</h2>
<dl class="package-details">
	<dt>Repository</dt>
	<dd><a href="https://github.com/pulumi/pulumi-azure-native">Azure Native pulumi/pulumi-azure-native</a></dd>
	<dt>License</dt>
	<dd>Apache-2.0</dd>
</dl>

