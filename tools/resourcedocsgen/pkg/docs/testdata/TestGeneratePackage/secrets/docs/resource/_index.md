
---
title: "Resource"
title_tag: "mypkg.Resource"
meta_desc: "Documentation for the mypkg.Resource resource with examples, input properties, output properties, lookup functions, and supporting types."
layout: api
no_edit_this_page: true
---



<!-- WARNING: this file was generated by test. -->
<!-- Do not edit by hand unless you're certain you know what you are doing! -->




## Create Resource Resource {#create}

Resources are created with functions called constructors. To learn more about declaring and configuring resources, see [Resources](/docs/concepts/resources/).

### Constructor syntax
<div>
<pulumi-chooser type="language" options="csharp,go,typescript,python,yaml,java"></pulumi-chooser>
</div>


<div>
<pulumi-choosable type="language" values="javascript,typescript">
<div class="no-copy"><div class="highlight"><pre class="chroma"><code class="language-typescript" data-lang="typescript"><span class="k">new </span><span class="nx">Resource</span><span class="p">(</span><span class="nx">name</span><span class="p">:</span> <span class="nx">string</span><span class="p">,</span> <span class="nx">args</span><span class="p">:</span> <span class="nx"><a href="#inputs">ResourceArgs</a></span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#CustomResourceOptions">CustomResourceOptions</a></span><span class="p">);</span></code></pre></div>
</div></pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<div class="no-copy"><div class="highlight"><pre class="chroma"><code class="language-python" data-lang="python"><span class=nd>@overload</span>
<span class="k">def </span><span class="nx">Resource</span><span class="p">(</span><span class="nx">resource_name</span><span class="p">:</span> <span class="nx">str</span><span class="p">,</span>
             <span class="nx">args</span><span class="p">:</span> <span class="nx"><a href="#inputs">ResourceArgs</a></span><span class="p">,</span>
             <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.ResourceOptions">Optional[ResourceOptions]</a></span> = None<span class="p">)</span>
<span></span>
<span class=nd>@overload</span>
<span class="k">def </span><span class="nx">Resource</span><span class="p">(</span><span class="nx">resource_name</span><span class="p">:</span> <span class="nx">str</span><span class="p">,</span>
             <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.ResourceOptions">Optional[ResourceOptions]</a></span> = None<span class="p">,</span>
             <span class="nx">config</span><span class="p">:</span> <span class="nx">Optional[ConfigArgs]</span> = None<span class="p">,</span>
             <span class="nx">config_array</span><span class="p">:</span> <span class="nx">Optional[Sequence[ConfigArgs]]</span> = None<span class="p">,</span>
             <span class="nx">config_map</span><span class="p">:</span> <span class="nx">Optional[Mapping[str, ConfigArgs]]</span> = None<span class="p">,</span>
             <span class="nx">foo</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
             <span class="nx">foo_array</span><span class="p">:</span> <span class="nx">Optional[Sequence[str]]</span> = None<span class="p">,</span>
             <span class="nx">foo_map</span><span class="p">:</span> <span class="nx">Optional[Mapping[str, str]]</span> = None<span class="p">)</span></code></pre></div>
</div></pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<div class="no-copy"><div class="highlight"><pre class="chroma"><code class="language-go" data-lang="go"><span class="k">func </span><span class="nx">NewResource</span><span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">name</span><span class="p"> </span><span class="nx">string</span><span class="p">,</span> <span class="nx">args</span><span class="p"> </span><span class="nx"><a href="#inputs">ResourceArgs</a></span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#ResourceOption">ResourceOption</a></span><span class="p">) (*<span class="nx">Resource</span>, error)</span></code></pre></div>
</div></pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="csharp">
<div class="no-copy"><div class="highlight"><pre class="chroma"><code class="language-csharp" data-lang="csharp"><span class="k">public </span><span class="nx">Resource</span><span class="p">(</span><span class="nx">string</span><span class="p"> </span><span class="nx">name<span class="p">,</span> <span class="nx"><a href="#inputs">ResourceArgs</a></span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.CustomResourceOptions.html">CustomResourceOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span></code></pre></div>
</div></pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<div class="no-copy"><div class="highlight"><pre class="chroma">
<code class="language-java" data-lang="java"><span class="k">public </span><span class="nx">Resource</span><span class="p">(</span><span class="nx">String</span><span class="p"> </span><span class="nx">name<span class="p">,</span> <span class="nx"><a href="#inputs">ResourceArgs</a></span><span class="p"> </span><span class="nx">args<span class="p">)</span>
<span class="k">public </span><span class="nx">Resource</span><span class="p">(</span><span class="nx">String</span><span class="p"> </span><span class="nx">name<span class="p">,</span> <span class="nx"><a href="#inputs">ResourceArgs</a></span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">CustomResourceOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
</code></pre></div></div>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<div class="no-copy"><div class="highlight"><pre class="chroma"><code class="language-yaml" data-lang="yaml">type: <span class="nx">mypkg:Resource</span><span class="p"></span>
<span class="p">properties</span><span class="p">: </span><span class="c">#&nbsp;The arguments to resource properties.</span>
<span class="p"></span><span class="p">options</span><span class="p">: </span><span class="c">#&nbsp;Bag of options to control resource&#39;s behavior.</span>
<span class="p"></span>
</code></pre></div></div>
</pulumi-choosable>
</div>

#### Parameters

<div>
<pulumi-choosable type="language" values="javascript,typescript">

<dl class="resources-properties"><dt
        class="property-required" title="Required">
        <span>name</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The unique name of the resource.</dd><dt
        class="property-required" title="Required">
        <span>args</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#inputs">ResourceArgs</a></span>
    </dt>
    <dd>The arguments to resource properties.</dd><dt
        class="property-optional" title="Optional">
        <span>opts</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#CustomResourceOptions">CustomResourceOptions</a></span>
    </dt>
    <dd>Bag of options to control resource&#39;s behavior.</dd></dl>

</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">

<dl class="resources-properties"><dt
        class="property-required" title="Required">
        <span>resource_name</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The unique name of the resource.</dd><dt
        class="property-required" title="Required">
        <span>args</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#inputs">ResourceArgs</a></span>
    </dt>
    <dd>The arguments to resource properties.</dd><dt
        class="property-optional" title="Optional">
        <span>opts</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="/docs/reference/pkg/python/pulumi/#pulumi.ResourceOptions">ResourceOptions</a></span>
    </dt>
    <dd>Bag of options to control resource&#39;s behavior.</dd></dl>

</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">

<dl class="resources-properties"><dt
        class="property-optional" title="Optional">
        <span>ctx</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span>
    </dt>
    <dd>Context object for the current deployment.</dd><dt
        class="property-required" title="Required">
        <span>name</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The unique name of the resource.</dd><dt
        class="property-required" title="Required">
        <span>args</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#inputs">ResourceArgs</a></span>
    </dt>
    <dd>The arguments to resource properties.</dd><dt
        class="property-optional" title="Optional">
        <span>opts</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#ResourceOption">ResourceOption</a></span>
    </dt>
    <dd>Bag of options to control resource&#39;s behavior.</dd></dl>

</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="csharp">

<dl class="resources-properties"><dt
        class="property-required" title="Required">
        <span>name</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The unique name of the resource.</dd><dt
        class="property-required" title="Required">
        <span>args</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#inputs">ResourceArgs</a></span>
    </dt>
    <dd>The arguments to resource properties.</dd><dt
        class="property-optional" title="Optional">
        <span>opts</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.CustomResourceOptions.html">CustomResourceOptions</a></span>
    </dt>
    <dd>Bag of options to control resource&#39;s behavior.</dd></dl>

</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">

<dl class="resources-properties"><dt
        class="property-required" title="Required">
        <span>name</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The unique name of the resource.</dd><dt
        class="property-required" title="Required">
        <span>args</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#inputs">ResourceArgs</a></span>
    </dt>
    <dd>The arguments to resource properties.</dd><dt
        class="property-optional" title="Optional">
        <span>options</span>
        <span class="property-indicator"></span>
        <span class="property-type">CustomResourceOptions</span>
    </dt>
    <dd>Bag of options to control resource&#39;s behavior.</dd></dl>

</pulumi-choosable>
</div>



### Constructor example

The following reference example uses placeholder values for all [input properties](#inputs).
<div>
<pulumi-chooser type="language" options="csharp,go,typescript,python,yaml,java"></pulumi-chooser>
</div>


<div>
<pulumi-choosable type="language" values="csharp">

```csharp
var resourceResource = new Mypkg.Resource("resourceResource", new()
{
    Config = new Mypkg.Inputs.ConfigArgs
    {
        Foo = "string",
    },
    ConfigArray = new[]
    {
        new Mypkg.Inputs.ConfigArgs
        {
            Foo = "string",
        },
    },
    ConfigMap = 
    {
        { "string", new Mypkg.Inputs.ConfigArgs
        {
            Foo = "string",
        } },
    },
    Foo = "string",
    FooArray = new[]
    {
        "string",
    },
    FooMap = 
    {
        { "string", "string" },
    },
});
```

</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="go">

```go
example, err := mypkg.NewResource(ctx, "resourceResource", &mypkg.ResourceArgs{
	Config: &mypkg.ConfigArgs{
		Foo: pulumi.String("string"),
	},
	ConfigArray: mypkg.ConfigArray{
		&mypkg.ConfigArgs{
			Foo: pulumi.String("string"),
		},
	},
	ConfigMap: mypkg.ConfigMap{
		"string": &mypkg.ConfigArgs{
			Foo: pulumi.String("string"),
		},
	},
	Foo: pulumi.String("string"),
	FooArray: pulumi.StringArray{
		pulumi.String("string"),
	},
	FooMap: pulumi.StringMap{
		"string": pulumi.String("string"),
	},
})
```

</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="java">

```java
var resourceResource = new Resource("resourceResource", ResourceArgs.builder()
    .config(ConfigArgs.builder()
        .foo("string")
        .build())
    .configArray(ConfigArgs.builder()
        .foo("string")
        .build())
    .configMap(Map.of("string", ConfigArgs.builder()
        .foo("string")
        .build()))
    .foo("string")
    .fooArray("string")
    .fooMap(Map.of("string", "string"))
    .build());
```

</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="python">

```python
resource_resource = mypkg.Resource("resourceResource",
    config={
        "foo": "string",
    },
    config_array=[{
        "foo": "string",
    }],
    config_map={
        "string": {
            "foo": "string",
        },
    },
    foo="string",
    foo_array=["string"],
    foo_map={
        "string": "string",
    })
```

</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="typescript">

```typescript
const resourceResource = new mypkg.Resource("resourceResource", {
    config: {
        foo: "string",
    },
    configArray: [{
        foo: "string",
    }],
    configMap: {
        string: {
            foo: "string",
        },
    },
    foo: "string",
    fooArray: ["string"],
    fooMap: {
        string: "string",
    },
});
```

</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="yaml">

```yaml
type: mypkg:Resource
properties:
    config:
        foo: string
    configArray:
        - foo: string
    configMap:
        string:
            foo: string
    foo: string
    fooArray:
        - string
    fooMap:
        string: string
```

</pulumi-choosable>
</div>



## Resource Resource Properties {#properties}

To learn more about resource properties and how to use them, see [Inputs and Outputs](/docs/intro/concepts/inputs-outputs) in the Architecture and Concepts docs.

### Inputs

<pulumi-choosable type="language" values="python">
<p>
In Python, inputs that are objects can be passed either as <a href="/docs/languages-sdks/python/#inputs-and-outputs">argument classes or as dictionary literals</a>.
</p>
</pulumi-choosable>

The Resource resource accepts the following [input](/docs/intro/concepts/inputs-outputs) properties:



<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="config_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#config_csharp" style="color: inherit; text-decoration: inherit;">Config</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#config">Config</a></span>
    </dt>
    <dd></dd><dt class="property-required"
            title="Required">
        <span id="configarray_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#configarray_csharp" style="color: inherit; text-decoration: inherit;">Config<wbr>Array</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#config">List&lt;Config&gt;</a></span>
    </dt>
    <dd></dd><dt class="property-required"
            title="Required">
        <span id="configmap_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#configmap_csharp" style="color: inherit; text-decoration: inherit;">Config<wbr>Map</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Dictionary&lt;string, Config<wbr>Args&gt;</span>
    </dt>
    <dd></dd><dt class="property-required"
            title="Required">
        <span id="foo_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#foo_csharp" style="color: inherit; text-decoration: inherit;">Foo</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd></dd><dt class="property-required"
            title="Required">
        <span id="fooarray_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#fooarray_csharp" style="color: inherit; text-decoration: inherit;">Foo<wbr>Array</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">List&lt;string&gt;</span>
    </dt>
    <dd></dd><dt class="property-required"
            title="Required">
        <span id="foomap_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#foomap_csharp" style="color: inherit; text-decoration: inherit;">Foo<wbr>Map</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Dictionary&lt;string, string&gt;</span>
    </dt>
    <dd></dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="config_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#config_go" style="color: inherit; text-decoration: inherit;">Config</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#config">Config<wbr>Args</a></span>
    </dt>
    <dd></dd><dt class="property-required"
            title="Required">
        <span id="configarray_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#configarray_go" style="color: inherit; text-decoration: inherit;">Config<wbr>Array</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#config">[]Config<wbr>Args</a></span>
    </dt>
    <dd></dd><dt class="property-required"
            title="Required">
        <span id="configmap_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#configmap_go" style="color: inherit; text-decoration: inherit;">Config<wbr>Map</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">map[string]Config<wbr>Args</span>
    </dt>
    <dd></dd><dt class="property-required"
            title="Required">
        <span id="foo_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#foo_go" style="color: inherit; text-decoration: inherit;">Foo</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd></dd><dt class="property-required"
            title="Required">
        <span id="fooarray_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#fooarray_go" style="color: inherit; text-decoration: inherit;">Foo<wbr>Array</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">[]string</span>
    </dt>
    <dd></dd><dt class="property-required"
            title="Required">
        <span id="foomap_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#foomap_go" style="color: inherit; text-decoration: inherit;">Foo<wbr>Map</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">map[string]string</span>
    </dt>
    <dd></dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="config_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#config_java" style="color: inherit; text-decoration: inherit;">config</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#config">Config</a></span>
    </dt>
    <dd></dd><dt class="property-required"
            title="Required">
        <span id="configarray_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#configarray_java" style="color: inherit; text-decoration: inherit;">config<wbr>Array</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#config">List&lt;Config&gt;</a></span>
    </dt>
    <dd></dd><dt class="property-required"
            title="Required">
        <span id="configmap_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#configmap_java" style="color: inherit; text-decoration: inherit;">config<wbr>Map</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Map&lt;String,Config<wbr>Args&gt;</span>
    </dt>
    <dd></dd><dt class="property-required"
            title="Required">
        <span id="foo_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#foo_java" style="color: inherit; text-decoration: inherit;">foo</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd></dd><dt class="property-required"
            title="Required">
        <span id="fooarray_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#fooarray_java" style="color: inherit; text-decoration: inherit;">foo<wbr>Array</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">List&lt;String&gt;</span>
    </dt>
    <dd></dd><dt class="property-required"
            title="Required">
        <span id="foomap_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#foomap_java" style="color: inherit; text-decoration: inherit;">foo<wbr>Map</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Map&lt;String,String&gt;</span>
    </dt>
    <dd></dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="config_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#config_nodejs" style="color: inherit; text-decoration: inherit;">config</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#config">Config</a></span>
    </dt>
    <dd></dd><dt class="property-required"
            title="Required">
        <span id="configarray_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#configarray_nodejs" style="color: inherit; text-decoration: inherit;">config<wbr>Array</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#config">Config[]</a></span>
    </dt>
    <dd></dd><dt class="property-required"
            title="Required">
        <span id="configmap_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#configmap_nodejs" style="color: inherit; text-decoration: inherit;">config<wbr>Map</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">{[key: string]: Config<wbr>Args}</span>
    </dt>
    <dd></dd><dt class="property-required"
            title="Required">
        <span id="foo_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#foo_nodejs" style="color: inherit; text-decoration: inherit;">foo</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd></dd><dt class="property-required"
            title="Required">
        <span id="fooarray_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#fooarray_nodejs" style="color: inherit; text-decoration: inherit;">foo<wbr>Array</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string[]</span>
    </dt>
    <dd></dd><dt class="property-required"
            title="Required">
        <span id="foomap_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#foomap_nodejs" style="color: inherit; text-decoration: inherit;">foo<wbr>Map</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">{[key: string]: string}</span>
    </dt>
    <dd></dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="config_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#config_python" style="color: inherit; text-decoration: inherit;">config</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#config">Config<wbr>Args</a></span>
    </dt>
    <dd></dd><dt class="property-required"
            title="Required">
        <span id="config_array_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#config_array_python" style="color: inherit; text-decoration: inherit;">config_<wbr>array</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#config">Sequence[Config<wbr>Args]</a></span>
    </dt>
    <dd></dd><dt class="property-required"
            title="Required">
        <span id="config_map_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#config_map_python" style="color: inherit; text-decoration: inherit;">config_<wbr>map</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Mapping[str, Config<wbr>Args]</span>
    </dt>
    <dd></dd><dt class="property-required"
            title="Required">
        <span id="foo_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#foo_python" style="color: inherit; text-decoration: inherit;">foo</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd></dd><dt class="property-required"
            title="Required">
        <span id="foo_array_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#foo_array_python" style="color: inherit; text-decoration: inherit;">foo_<wbr>array</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Sequence[str]</span>
    </dt>
    <dd></dd><dt class="property-required"
            title="Required">
        <span id="foo_map_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#foo_map_python" style="color: inherit; text-decoration: inherit;">foo_<wbr>map</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Mapping[str, str]</span>
    </dt>
    <dd></dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="config_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#config_yaml" style="color: inherit; text-decoration: inherit;">config</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#config">Property Map</a></span>
    </dt>
    <dd></dd><dt class="property-required"
            title="Required">
        <span id="configarray_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#configarray_yaml" style="color: inherit; text-decoration: inherit;">config<wbr>Array</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#config">List&lt;Property Map&gt;</a></span>
    </dt>
    <dd></dd><dt class="property-required"
            title="Required">
        <span id="configmap_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#configmap_yaml" style="color: inherit; text-decoration: inherit;">config<wbr>Map</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Map&lt;Property Map&gt;</span>
    </dt>
    <dd></dd><dt class="property-required"
            title="Required">
        <span id="foo_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#foo_yaml" style="color: inherit; text-decoration: inherit;">foo</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd></dd><dt class="property-required"
            title="Required">
        <span id="fooarray_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#fooarray_yaml" style="color: inherit; text-decoration: inherit;">foo<wbr>Array</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">List&lt;String&gt;</span>
    </dt>
    <dd></dd><dt class="property-required"
            title="Required">
        <span id="foomap_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#foomap_yaml" style="color: inherit; text-decoration: inherit;">foo<wbr>Map</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Map&lt;String&gt;</span>
    </dt>
    <dd></dd></dl>
</pulumi-choosable>
</div>


### Outputs

All [input](#inputs) properties are implicitly available as output properties. Additionally, the Resource resource produces the following output properties:



<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="id_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_csharp" style="color: inherit; text-decoration: inherit;">Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The provider-assigned unique ID for this managed resource.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="id_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_go" style="color: inherit; text-decoration: inherit;">Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The provider-assigned unique ID for this managed resource.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="id_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_java" style="color: inherit; text-decoration: inherit;">id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The provider-assigned unique ID for this managed resource.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="id_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_nodejs" style="color: inherit; text-decoration: inherit;">id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The provider-assigned unique ID for this managed resource.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="id_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_python" style="color: inherit; text-decoration: inherit;">id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The provider-assigned unique ID for this managed resource.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="id_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_yaml" style="color: inherit; text-decoration: inherit;">id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The provider-assigned unique ID for this managed resource.</dd></dl>
</pulumi-choosable>
</div>







## Supporting Types



<h4 id="config">
Config<pulumi-choosable type="language" values="python,go" class="inline">, Config<wbr>Args</pulumi-choosable>
</h4>

<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="foo_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#foo_csharp" style="color: inherit; text-decoration: inherit;">Foo</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd></dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="foo_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#foo_go" style="color: inherit; text-decoration: inherit;">Foo</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd></dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="foo_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#foo_java" style="color: inherit; text-decoration: inherit;">foo</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd></dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="foo_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#foo_nodejs" style="color: inherit; text-decoration: inherit;">foo</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd></dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="foo_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#foo_python" style="color: inherit; text-decoration: inherit;">foo</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd></dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="foo_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#foo_yaml" style="color: inherit; text-decoration: inherit;">foo</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd></dd></dl>
</pulumi-choosable>
</div>


<h2 id="package-details">Package Details</h2>
<dl class="package-details">
	<dt>Repository</dt>
	<dd><a href="">mypkg </a></dd>
	<dt>License</dt>
	<dd></dd>
</dl>

