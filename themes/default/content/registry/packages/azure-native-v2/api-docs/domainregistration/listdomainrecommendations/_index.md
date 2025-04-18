
---
title: "listDomainRecommendations"
title_tag: "azure-native.domainregistration.listDomainRecommendations"
meta_desc: "Documentation for the azure-native.domainregistration.listDomainRecommendations function with examples, input properties, output properties, and supporting types."
layout: api
no_edit_this_page: true
---



<!-- WARNING: this file was generated by Pulumi Docs Generator. -->
<!-- Do not edit by hand unless you're certain you know what you are doing! -->

Description for Get domain name recommendations based on keywords.

Azure REST API version: 2024-04-01.

Other available API versions: 2022-09-01, 2023-01-01, 2023-12-01.

These versions are not included in the SDK but you can generate them as a local SDK package; see the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.




## Using listDomainRecommendations {#using}

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
><span class="k">function </span>listDomainRecommendations<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">ListDomainRecommendationsArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Promise&lt;<span class="nx"><a href="#result">ListDomainRecommendationsResult</a></span>></span
><span class="k">
function </span>listDomainRecommendationsOutput<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">ListDomainRecommendationsOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Output&lt;<span class="nx"><a href="#result">ListDomainRecommendationsResult</a></span>></span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="python">
<div class="highlight"><pre class="chroma"><code class="language-python" data-lang="python"
><span class="k">def </span>list_domain_recommendations<span class="p">(</span><span class="nx">keywords</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                                <span class="nx">max_domain_recommendations</span><span class="p">:</span> <span class="nx">Optional[int]</span> = None<span class="p">,</span>
                                <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>ListDomainRecommendationsResult</span
><span class="k">
def </span>list_domain_recommendations_output<span class="p">(</span><span class="nx">keywords</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                                <span class="nx">max_domain_recommendations</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[int]]</span> = None<span class="p">,</span>
                                <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>Output[ListDomainRecommendationsResult]</span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="go">
<div class="highlight"><pre class="chroma"><code class="language-go" data-lang="go"
><span class="k">func </span>ListDomainRecommendations<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">ListDomainRecommendationsArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) (*<span class="nx"><a href="#result">ListDomainRecommendationsResult</a></span>, error)</span
><span class="k">
func </span>ListDomainRecommendationsOutput<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">ListDomainRecommendationsOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) ListDomainRecommendationsResultOutput</span
></code></pre></div>

&gt; Note: This function is named `ListDomainRecommendations` in the Go SDK.

</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="csharp">
<div class="highlight"><pre class="chroma"><code class="language-csharp" data-lang="csharp"><span class="k">public static class </span><span class="nx">ListDomainRecommendations </span><span class="p">
{</span><span class="k">
    public static </span>Task&lt;<span class="nx"><a href="#result">ListDomainRecommendationsResult</a></span>> <span class="p">InvokeAsync(</span><span class="nx">ListDomainRecommendationsArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="k">
    public static </span>Output&lt;<span class="nx"><a href="#result">ListDomainRecommendationsResult</a></span>> <span class="p">Invoke(</span><span class="nx">ListDomainRecommendationsInvokeArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="p">
}</span></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="java">
<div class="highlight"><pre class="chroma"><code class="language-java" data-lang="java"><span class="k">public static CompletableFuture&lt;<span class="nx"><a href="#result">ListDomainRecommendationsResult</a></span>> </span>listDomainRecommendations<span class="p">(</span><span class="nx">ListDomainRecommendationsArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
<span class="k">public static Output&lt;<span class="nx"><a href="#result">ListDomainRecommendationsResult</a></span>> </span>listDomainRecommendations<span class="p">(</span><span class="nx">ListDomainRecommendationsArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
</code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="yaml">
<div class="highlight"><pre class="chroma"><code class="language-yaml" data-lang="yaml"><span class="k">fn::invoke:</span>
<span class="k">&nbsp;&nbsp;function:</span> azure-native:domainregistration:listDomainRecommendations
<span class="k">&nbsp;&nbsp;arguments:</span>
<span class="c">&nbsp;&nbsp;&nbsp;&nbsp;# arguments dictionary</span></code></pre></div>
</pulumi-choosable>
</div>



The following arguments are supported:


<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="keywords_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#keywords_csharp" style="color: inherit; text-decoration: inherit;">Keywords</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Keywords to be used for generating domain recommendations.</dd><dt class="property-optional"
            title="Optional">
        <span id="maxdomainrecommendations_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#maxdomainrecommendations_csharp" style="color: inherit; text-decoration: inherit;">Max<wbr>Domain<wbr>Recommendations</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">int</span>
    </dt>
    <dd>Maximum number of recommendations.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="keywords_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#keywords_go" style="color: inherit; text-decoration: inherit;">Keywords</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Keywords to be used for generating domain recommendations.</dd><dt class="property-optional"
            title="Optional">
        <span id="maxdomainrecommendations_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#maxdomainrecommendations_go" style="color: inherit; text-decoration: inherit;">Max<wbr>Domain<wbr>Recommendations</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">int</span>
    </dt>
    <dd>Maximum number of recommendations.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="keywords_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#keywords_java" style="color: inherit; text-decoration: inherit;">keywords</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Keywords to be used for generating domain recommendations.</dd><dt class="property-optional"
            title="Optional">
        <span id="maxdomainrecommendations_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#maxdomainrecommendations_java" style="color: inherit; text-decoration: inherit;">max<wbr>Domain<wbr>Recommendations</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Integer</span>
    </dt>
    <dd>Maximum number of recommendations.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="keywords_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#keywords_nodejs" style="color: inherit; text-decoration: inherit;">keywords</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Keywords to be used for generating domain recommendations.</dd><dt class="property-optional"
            title="Optional">
        <span id="maxdomainrecommendations_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#maxdomainrecommendations_nodejs" style="color: inherit; text-decoration: inherit;">max<wbr>Domain<wbr>Recommendations</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">number</span>
    </dt>
    <dd>Maximum number of recommendations.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="keywords_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#keywords_python" style="color: inherit; text-decoration: inherit;">keywords</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Keywords to be used for generating domain recommendations.</dd><dt class="property-optional"
            title="Optional">
        <span id="max_domain_recommendations_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#max_domain_recommendations_python" style="color: inherit; text-decoration: inherit;">max_<wbr>domain_<wbr>recommendations</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">int</span>
    </dt>
    <dd>Maximum number of recommendations.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="keywords_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#keywords_yaml" style="color: inherit; text-decoration: inherit;">keywords</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Keywords to be used for generating domain recommendations.</dd><dt class="property-optional"
            title="Optional">
        <span id="maxdomainrecommendations_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#maxdomainrecommendations_yaml" style="color: inherit; text-decoration: inherit;">max<wbr>Domain<wbr>Recommendations</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Number</span>
    </dt>
    <dd>Maximum number of recommendations.</dd></dl>
</pulumi-choosable>
</div>




## listDomainRecommendations Result {#result}

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
    <dd>Link to next page of resources.</dd><dt class="property-"
            title="">
        <span id="value_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#value_csharp" style="color: inherit; text-decoration: inherit;">Value</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#nameidentifierresponse">List&lt;Pulumi.<wbr>Azure<wbr>Native.<wbr>Domain<wbr>Registration.<wbr>Outputs.<wbr>Name<wbr>Identifier<wbr>Response&gt;</a></span>
    </dt>
    <dd>Collection of resources.</dd></dl>
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
    <dd>Link to next page of resources.</dd><dt class="property-"
            title="">
        <span id="value_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#value_go" style="color: inherit; text-decoration: inherit;">Value</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#nameidentifierresponse">[]Name<wbr>Identifier<wbr>Response</a></span>
    </dt>
    <dd>Collection of resources.</dd></dl>
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
    <dd>Link to next page of resources.</dd><dt class="property-"
            title="">
        <span id="value_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#value_java" style="color: inherit; text-decoration: inherit;">value</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#nameidentifierresponse">List&lt;Name<wbr>Identifier<wbr>Response&gt;</a></span>
    </dt>
    <dd>Collection of resources.</dd></dl>
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
    <dd>Link to next page of resources.</dd><dt class="property-"
            title="">
        <span id="value_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#value_nodejs" style="color: inherit; text-decoration: inherit;">value</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#nameidentifierresponse">Name<wbr>Identifier<wbr>Response[]</a></span>
    </dt>
    <dd>Collection of resources.</dd></dl>
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
    <dd>Link to next page of resources.</dd><dt class="property-"
            title="">
        <span id="value_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#value_python" style="color: inherit; text-decoration: inherit;">value</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#nameidentifierresponse">Sequence[Name<wbr>Identifier<wbr>Response]</a></span>
    </dt>
    <dd>Collection of resources.</dd></dl>
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
    <dd>Link to next page of resources.</dd><dt class="property-"
            title="">
        <span id="value_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#value_yaml" style="color: inherit; text-decoration: inherit;">value</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#nameidentifierresponse">List&lt;Property Map&gt;</a></span>
    </dt>
    <dd>Collection of resources.</dd></dl>
</pulumi-choosable>
</div>




## Supporting Types


<h4 id="nameidentifierresponse">Name<wbr>Identifier<wbr>Response</h4>



<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="name_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_csharp" style="color: inherit; text-decoration: inherit;">Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Name of the object.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="name_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_go" style="color: inherit; text-decoration: inherit;">Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Name of the object.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="name_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_java" style="color: inherit; text-decoration: inherit;">name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Name of the object.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="name_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_nodejs" style="color: inherit; text-decoration: inherit;">name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Name of the object.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_python" style="color: inherit; text-decoration: inherit;">name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Name of the object.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="name_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#name_yaml" style="color: inherit; text-decoration: inherit;">name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Name of the object.</dd></dl>
</pulumi-choosable>
</div>





<h2 id="package-details">Package Details</h2>
<dl class="package-details">
	<dt>Repository</dt>
	<dd><a href="https://github.com/pulumi/pulumi-azure-native">Azure Native pulumi/pulumi-azure-native</a></dd>
	<dt>License</dt>
	<dd>Apache-2.0</dd>
</dl>

