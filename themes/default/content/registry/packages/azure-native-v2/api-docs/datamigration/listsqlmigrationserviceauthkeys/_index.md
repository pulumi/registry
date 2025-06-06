
---
title: "listSqlMigrationServiceAuthKeys"
title_tag: "azure-native.datamigration.listSqlMigrationServiceAuthKeys"
meta_desc: "Documentation for the azure-native.datamigration.listSqlMigrationServiceAuthKeys function with examples, input properties, output properties, and supporting types."
layout: api
no_edit_this_page: true
---



<!-- WARNING: this file was generated by Pulumi Docs Generator. -->
<!-- Do not edit by hand unless you're certain you know what you are doing! -->

Retrieve the List of Authentication Keys for Self Hosted Integration Runtime.

Azure REST API version: 2023-07-15-preview.

Other available API versions: 2022-03-30-preview.

These versions are not included in the SDK but you can generate them as a local SDK package; see the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.




## Using listSqlMigrationServiceAuthKeys {#using}

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
><span class="k">function </span>listSqlMigrationServiceAuthKeys<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">ListSqlMigrationServiceAuthKeysArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Promise&lt;<span class="nx"><a href="#result">ListSqlMigrationServiceAuthKeysResult</a></span>></span
><span class="k">
function </span>listSqlMigrationServiceAuthKeysOutput<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">ListSqlMigrationServiceAuthKeysOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Output&lt;<span class="nx"><a href="#result">ListSqlMigrationServiceAuthKeysResult</a></span>></span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="python">
<div class="highlight"><pre class="chroma"><code class="language-python" data-lang="python"
><span class="k">def </span>list_sql_migration_service_auth_keys<span class="p">(</span><span class="nx">resource_group_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                                         <span class="nx">sql_migration_service_name</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                                         <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>ListSqlMigrationServiceAuthKeysResult</span
><span class="k">
def </span>list_sql_migration_service_auth_keys_output<span class="p">(</span><span class="nx">resource_group_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                                         <span class="nx">sql_migration_service_name</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                                         <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>Output[ListSqlMigrationServiceAuthKeysResult]</span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="go">
<div class="highlight"><pre class="chroma"><code class="language-go" data-lang="go"
><span class="k">func </span>ListSqlMigrationServiceAuthKeys<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">ListSqlMigrationServiceAuthKeysArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) (*<span class="nx"><a href="#result">ListSqlMigrationServiceAuthKeysResult</a></span>, error)</span
><span class="k">
func </span>ListSqlMigrationServiceAuthKeysOutput<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">ListSqlMigrationServiceAuthKeysOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) ListSqlMigrationServiceAuthKeysResultOutput</span
></code></pre></div>

&gt; Note: This function is named `ListSqlMigrationServiceAuthKeys` in the Go SDK.

</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="csharp">
<div class="highlight"><pre class="chroma"><code class="language-csharp" data-lang="csharp"><span class="k">public static class </span><span class="nx">ListSqlMigrationServiceAuthKeys </span><span class="p">
{</span><span class="k">
    public static </span>Task&lt;<span class="nx"><a href="#result">ListSqlMigrationServiceAuthKeysResult</a></span>> <span class="p">InvokeAsync(</span><span class="nx">ListSqlMigrationServiceAuthKeysArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="k">
    public static </span>Output&lt;<span class="nx"><a href="#result">ListSqlMigrationServiceAuthKeysResult</a></span>> <span class="p">Invoke(</span><span class="nx">ListSqlMigrationServiceAuthKeysInvokeArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="p">
}</span></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="java">
<div class="highlight"><pre class="chroma"><code class="language-java" data-lang="java"><span class="k">public static CompletableFuture&lt;<span class="nx"><a href="#result">ListSqlMigrationServiceAuthKeysResult</a></span>> </span>listSqlMigrationServiceAuthKeys<span class="p">(</span><span class="nx">ListSqlMigrationServiceAuthKeysArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
<span class="k">public static Output&lt;<span class="nx"><a href="#result">ListSqlMigrationServiceAuthKeysResult</a></span>> </span>listSqlMigrationServiceAuthKeys<span class="p">(</span><span class="nx">ListSqlMigrationServiceAuthKeysArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
</code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="yaml">
<div class="highlight"><pre class="chroma"><code class="language-yaml" data-lang="yaml"><span class="k">fn::invoke:</span>
<span class="k">&nbsp;&nbsp;function:</span> azure-native:datamigration:listSqlMigrationServiceAuthKeys
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
    <dd>Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="sqlmigrationservicename_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#sqlmigrationservicename_csharp" style="color: inherit; text-decoration: inherit;">Sql<wbr>Migration<wbr>Service<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Name of the SQL Migration Service.</dd></dl>
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
    <dd>Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="sqlmigrationservicename_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#sqlmigrationservicename_go" style="color: inherit; text-decoration: inherit;">Sql<wbr>Migration<wbr>Service<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Name of the SQL Migration Service.</dd></dl>
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
    <dd>Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="sqlmigrationservicename_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#sqlmigrationservicename_java" style="color: inherit; text-decoration: inherit;">sql<wbr>Migration<wbr>Service<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Name of the SQL Migration Service.</dd></dl>
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
    <dd>Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="sqlmigrationservicename_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#sqlmigrationservicename_nodejs" style="color: inherit; text-decoration: inherit;">sql<wbr>Migration<wbr>Service<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>Name of the SQL Migration Service.</dd></dl>
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
    <dd>Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="sql_migration_service_name_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#sql_migration_service_name_python" style="color: inherit; text-decoration: inherit;">sql_<wbr>migration_<wbr>service_<wbr>name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>Name of the SQL Migration Service.</dd></dl>
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
    <dd>Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.</dd><dt class="property-required property-replacement"
            title="Required">
        <span id="sqlmigrationservicename_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#sqlmigrationservicename_yaml" style="color: inherit; text-decoration: inherit;">sql<wbr>Migration<wbr>Service<wbr>Name</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>Name of the SQL Migration Service.</dd></dl>
</pulumi-choosable>
</div>




## listSqlMigrationServiceAuthKeys Result {#result}

The following output properties are available:



<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="authkey1_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#authkey1_csharp" style="color: inherit; text-decoration: inherit;">Auth<wbr>Key1</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The first authentication key.</dd><dt class="property-"
            title="">
        <span id="authkey2_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#authkey2_csharp" style="color: inherit; text-decoration: inherit;">Auth<wbr>Key2</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The second authentication key.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="authkey1_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#authkey1_go" style="color: inherit; text-decoration: inherit;">Auth<wbr>Key1</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The first authentication key.</dd><dt class="property-"
            title="">
        <span id="authkey2_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#authkey2_go" style="color: inherit; text-decoration: inherit;">Auth<wbr>Key2</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The second authentication key.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="authkey1_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#authkey1_java" style="color: inherit; text-decoration: inherit;">auth<wbr>Key1</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The first authentication key.</dd><dt class="property-"
            title="">
        <span id="authkey2_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#authkey2_java" style="color: inherit; text-decoration: inherit;">auth<wbr>Key2</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The second authentication key.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="authkey1_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#authkey1_nodejs" style="color: inherit; text-decoration: inherit;">auth<wbr>Key1</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The first authentication key.</dd><dt class="property-"
            title="">
        <span id="authkey2_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#authkey2_nodejs" style="color: inherit; text-decoration: inherit;">auth<wbr>Key2</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The second authentication key.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="auth_key1_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#auth_key1_python" style="color: inherit; text-decoration: inherit;">auth_<wbr>key1</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The first authentication key.</dd><dt class="property-"
            title="">
        <span id="auth_key2_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#auth_key2_python" style="color: inherit; text-decoration: inherit;">auth_<wbr>key2</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The second authentication key.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="authkey1_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#authkey1_yaml" style="color: inherit; text-decoration: inherit;">auth<wbr>Key1</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The first authentication key.</dd><dt class="property-"
            title="">
        <span id="authkey2_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#authkey2_yaml" style="color: inherit; text-decoration: inherit;">auth<wbr>Key2</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The second authentication key.</dd></dl>
</pulumi-choosable>
</div>





<h2 id="package-details">Package Details</h2>
<dl class="package-details">
	<dt>Repository</dt>
	<dd><a href="https://github.com/pulumi/pulumi-azure-native">Azure Native pulumi/pulumi-azure-native</a></dd>
	<dt>License</dt>
	<dd>Apache-2.0</dd>
</dl>

