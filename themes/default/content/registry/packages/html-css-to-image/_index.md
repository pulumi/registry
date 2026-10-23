---
# WARNING: this file was fetched from https://raw.githubusercontent.com/htmlcsstoimage/pulumi-html-css-to-image/v0.2.0/docs/_index.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/htmlcsstoimage/pulumi-html-css-to-image/blob/v0.2.0/docs/_index.md
title: HTML/CSS to Image
meta_desc: Use the HTML/CSS to Image provider for Pulumi to create images and manage templates, API keys, storage destinations, proxies, and Open Graph configurations.
layout: package
---

The HTML/CSS to Image provider for Pulumi manages [HTML/CSS to Image](https://htmlcsstoimage.com) images, reusable templates, API keys, storage destinations, proxies, and Open Graph configurations alongside your application's infrastructure.

## Installation

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}

```sh
npm install @html-css-to-image/pulumi
```

{{% /choosable %}}
{{% choosable language python %}}

```sh
pip install pulumi-html-css-to-image
```

{{% /choosable %}}
{{% choosable language go %}}

```sh
go get github.com/htmlcsstoimage/pulumi-html-css-to-image/sdk/go
```

{{% /choosable %}}
{{% choosable language csharp %}}

```sh
dotnet add package HtmlCssToImage.Pulumi
```

{{% /choosable %}}
{{% choosable language java %}}

Maven:

```xml
<dependency>
    <groupId>com.htmlcsstoimage</groupId>
    <artifactId>pulumi</artifactId>
    <version>0.2.0</version>
</dependency>
```

Gradle:

```groovy
implementation 'com.htmlcsstoimage:pulumi:0.2.0'
```

{{% /choosable %}}
{{% choosable language yaml %}}

```sh
pulumi package add html-css-to-image
```

{{% /choosable %}}
{{< /chooser >}}

The SDK or package installation downloads the matching provider plugin automatically.

## Example Usage

Set your organization credentials in your Pulumi project:

```sh
pulumi config set html-css-to-image:apiId YOUR_API_ID
pulumi config set --secret html-css-to-image:apiKey YOUR_API_KEY
```

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}

```typescript
import * as hcti from "@html-css-to-image/pulumi";

const card = new hcti.ImageHtmlCss("card", {
    html: "<h1>Hello from Pulumi</h1>",
});

export const imageUrl = card.imageUrl;
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import pulumi_html_css_to_image as hcti

card = hcti.ImageHtmlCss("card", html="<h1>Hello from Pulumi</h1>")
pulumi.export("imageUrl", card.image_url)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
    hcti "github.com/htmlcsstoimage/pulumi-html-css-to-image/sdk/go"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        card, err := hcti.NewImageHtmlCss(ctx, "card", &hcti.ImageHtmlCssArgs{
            Html: pulumi.String("<h1>Hello from Pulumi</h1>"),
        })
        if err != nil {
            return err
        }
        ctx.Export("imageUrl", card.ImageUrl)
        return nil
    })
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using Pulumi;
using Hcti = Pulumi.HtmlCssToImage;

return await Deployment.RunAsync(() =>
{
    var card = new Hcti.ImageHtmlCss("card", new Hcti.ImageHtmlCssArgs
    {
        Html = "<h1>Hello from Pulumi</h1>",
    });
    return new Dictionary<string, object?> { ["imageUrl"] = card.ImageUrl };
});
```

{{% /choosable %}}
{{% choosable language java %}}

```java
import com.htmlcsstoimage.pulumi.ImageHtmlCss;
import com.htmlcsstoimage.pulumi.ImageHtmlCssArgs;
import com.pulumi.Pulumi;

public class App {
    public static void main(String[] args) {
        Pulumi.run(ctx -> {
            var card = new ImageHtmlCss("card", ImageHtmlCssArgs.builder()
                    .html("<h1>Hello from Pulumi</h1>")
                    .build());
            ctx.export("imageUrl", card.imageUrl());
        });
    }
}
```

{{% /choosable %}}
{{% choosable language yaml %}}

```yaml
name: hcti-image
runtime: yaml
resources:
  card:
    type: html-css-to-image:index:ImageHtmlCss
    properties:
      html: '<h1>Hello from Pulumi</h1>'
outputs:
  imageUrl: ${card.imageUrl}
```

{{% /choosable %}}
{{< /chooser >}}

Run `pulumi up`, then open the returned URL to render the image. Creation saves an image definition; changes to image inputs replace that definition.

## Configuration

| Name | Required? | Secret? | Description |
| --- | --- | --- | --- |
| `apiId` | Yes, through configuration or environment | No | Your organization API ID. Falls back to `HCTI_API_ID`. |
| `apiKey` | Yes, through configuration or environment | Yes | API key with permissions to manage your stack's resources. Falls back to `HCTI_API_KEY`. |
| `baseUrl` | No | No | API origin, including scheme; defaults to `https://hcti.io`. |

## Further reading

- [HTML/CSS to Image documentation](https://docs.htmlcsstoimage.com/)
- [Management API and permissions](https://docs.htmlcsstoimage.com/management-api/)
- [API reference](https://htmlcsstoimage.com/api-docs/)
- [Complete provider examples](https://github.com/htmlcsstoimage/pulumi-html-css-to-image/tree/main/examples)
