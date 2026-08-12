---
title: AWS Configuration Reference
meta_desc: Every configuration option accepted by the Pulumi AWS provider, with the pulumi env set and pulumi config set command for each.
layout: package
---

Every configuration option accepted by the AWS provider. Project any of them from a [Pulumi ESC](/docs/pulumi-cloud/esc/) environment under `pulumiConfig`, set them with `pulumi config set`, or pass them to an explicit [`aws.Provider`](/registry/packages/aws/api-docs/provider/).

For worked credential examples — shared credentials file, static keys, assumed roles, OIDC — see [Configuration](/registry/packages/aws/#configuration) on the Overview page.

Nested objects are shown flattened, using [structured configuration](/docs/concepts/config#structured-configuration) paths.

<div class="api-params">
<div class="api-param">
<h3 class="api-param-name" id="accesskey" data-link-title="accessKey">accessKey<a class="api-param-anchor" href="#accesskey" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
<span class="api-param-badge api-param-secret">secret</span>
</div>
<div class="api-param-description">
<p>The access key for API operations. You can retrieve this from the 'Security &amp; Credentials' section of the AWS console.</p>
<p class="api-param-meta">Config key: <code>aws:accessKey</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> --secret myorg/myproject/dev pulumiConfig.aws:accessKey AKIAIOSFODNN7EXAMPLE
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> --secret aws:accessKey AKIAIOSFODNN7EXAMPLE
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="allowedaccountids" data-link-title="allowedAccountIds">allowedAccountIds<a class="api-param-anchor" href="#allowedaccountids" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">list&lt;string&gt;</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p class="api-param-meta">Config key: <code>aws:allowedAccountIds</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:allowedAccountIds 123456789012
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> aws:allowedAccountIds 123456789012
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="assumerolewithwebidentity-duration" data-link-title="assumeRoleWithWebIdentity.duration"><span class="api-param-nested-indicator">&#8627;</span>assumeRoleWithWebIdentity.duration<a class="api-param-anchor" href="#assumerolewithwebidentity-duration" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>The duration, between 15 minutes and 12 hours, of the role session. Valid time units are ns, us (or µs), ms, s, h, or m.</p>
<p class="api-param-meta">Config key: <code>aws:assumeRoleWithWebIdentity.duration</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:assumeRoleWithWebIdentity.duration 1h
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> --path aws:assumeRoleWithWebIdentity.duration 1h
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="assumerolewithwebidentity-policy" data-link-title="assumeRoleWithWebIdentity.policy"><span class="api-param-nested-indicator">&#8627;</span>assumeRoleWithWebIdentity.policy<a class="api-param-anchor" href="#assumerolewithwebidentity-policy" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>IAM Policy JSON describing further restricting permissions for the IAM Role being assumed.</p>
<p class="api-param-meta">Config key: <code>aws:assumeRoleWithWebIdentity.policy</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:assumeRoleWithWebIdentity.policy &lt;value&gt;
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> --path aws:assumeRoleWithWebIdentity.policy &lt;value&gt;
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="assumerolewithwebidentity-policyarns" data-link-title="assumeRoleWithWebIdentity.policyArns"><span class="api-param-nested-indicator">&#8627;</span>assumeRoleWithWebIdentity.policyArns<a class="api-param-anchor" href="#assumerolewithwebidentity-policyarns" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">list&lt;string&gt;</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Amazon Resource Names (ARNs) of IAM Policies describing further restricting permissions for the IAM Role being assumed.</p>
<p class="api-param-meta">Config key: <code>aws:assumeRoleWithWebIdentity.policyArns</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:assumeRoleWithWebIdentity.policyArns &lt;value&gt;
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> --path aws:assumeRoleWithWebIdentity.policyArns &lt;value&gt;
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="assumerolewithwebidentity-rolearn" data-link-title="assumeRoleWithWebIdentity.roleArn"><span class="api-param-nested-indicator">&#8627;</span>assumeRoleWithWebIdentity.roleArn<a class="api-param-anchor" href="#assumerolewithwebidentity-rolearn" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Amazon Resource Name (ARN) of an IAM Role to assume prior to making API calls.</p>
<p class="api-param-meta">Config key: <code>aws:assumeRoleWithWebIdentity.roleArn</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:assumeRoleWithWebIdentity.roleArn arn:aws:iam::123456789012:role/GitHubActions
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> --path aws:assumeRoleWithWebIdentity.roleArn arn:aws:iam::123456789012:role/GitHubActions
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="assumerolewithwebidentity-sessionname" data-link-title="assumeRoleWithWebIdentity.sessionName"><span class="api-param-nested-indicator">&#8627;</span>assumeRoleWithWebIdentity.sessionName<a class="api-param-anchor" href="#assumerolewithwebidentity-sessionname" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>An identifier for the assumed role session.</p>
<p class="api-param-meta">Config key: <code>aws:assumeRoleWithWebIdentity.sessionName</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:assumeRoleWithWebIdentity.sessionName pulumi-session
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> --path aws:assumeRoleWithWebIdentity.sessionName pulumi-session
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="assumerolewithwebidentity-webidentitytoken" data-link-title="assumeRoleWithWebIdentity.webIdentityToken"><span class="api-param-nested-indicator">&#8627;</span>assumeRoleWithWebIdentity.webIdentityToken<a class="api-param-anchor" href="#assumerolewithwebidentity-webidentitytoken" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p class="api-param-meta">Config key: <code>aws:assumeRoleWithWebIdentity.webIdentityToken</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:assumeRoleWithWebIdentity.webIdentityToken &lt;value&gt;
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> --path aws:assumeRoleWithWebIdentity.webIdentityToken &lt;value&gt;
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="assumerolewithwebidentity-webidentitytokenfile" data-link-title="assumeRoleWithWebIdentity.webIdentityTokenFile"><span class="api-param-nested-indicator">&#8627;</span>assumeRoleWithWebIdentity.webIdentityTokenFile<a class="api-param-anchor" href="#assumerolewithwebidentity-webidentitytokenfile" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p class="api-param-meta">Config key: <code>aws:assumeRoleWithWebIdentity.webIdentityTokenFile</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:assumeRoleWithWebIdentity.webIdentityTokenFile /var/run/secrets/token
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> --path aws:assumeRoleWithWebIdentity.webIdentityTokenFile /var/run/secrets/token
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="assumeroles-duration" data-link-title="assumeRoles[].duration"><span class="api-param-nested-indicator">&#8627;</span>assumeRoles[].duration<a class="api-param-anchor" href="#assumeroles-duration" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>The duration, between 15 minutes and 12 hours, of the role session. Valid time units are ns, us (or µs), ms, s, h, or m.</p>
<p class="api-param-meta">Config key: <code>aws:assumeRoles[0].duration</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:assumeRoles[0].duration 1h
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> --path aws:assumeRoles[0].duration 1h
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="assumeroles-externalid" data-link-title="assumeRoles[].externalId"><span class="api-param-nested-indicator">&#8627;</span>assumeRoles[].externalId<a class="api-param-anchor" href="#assumeroles-externalid" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>A unique identifier that might be required when you assume a role in another account.</p>
<p class="api-param-meta">Config key: <code>aws:assumeRoles[0].externalId</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:assumeRoles[0].externalId my-external-id
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> --path aws:assumeRoles[0].externalId my-external-id
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="assumeroles-policy" data-link-title="assumeRoles[].policy"><span class="api-param-nested-indicator">&#8627;</span>assumeRoles[].policy<a class="api-param-anchor" href="#assumeroles-policy" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>IAM Policy JSON describing further restricting permissions for the IAM Role being assumed.</p>
<p class="api-param-meta">Config key: <code>aws:assumeRoles[0].policy</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:assumeRoles[0].policy &lt;value&gt;
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> --path aws:assumeRoles[0].policy &lt;value&gt;
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="assumeroles-policyarns" data-link-title="assumeRoles[].policyArns"><span class="api-param-nested-indicator">&#8627;</span>assumeRoles[].policyArns<a class="api-param-anchor" href="#assumeroles-policyarns" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">list&lt;string&gt;</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Amazon Resource Names (ARNs) of IAM Policies describing further restricting permissions for the IAM Role being assumed.</p>
<p class="api-param-meta">Config key: <code>aws:assumeRoles[0].policyArns</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:assumeRoles[0].policyArns &lt;value&gt;
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> --path aws:assumeRoles[0].policyArns &lt;value&gt;
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="assumeroles-rolearn" data-link-title="assumeRoles[].roleArn"><span class="api-param-nested-indicator">&#8627;</span>assumeRoles[].roleArn<a class="api-param-anchor" href="#assumeroles-rolearn" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Amazon Resource Name (ARN) of an IAM Role to assume prior to making API calls.</p>
<p class="api-param-meta">Config key: <code>aws:assumeRoles[0].roleArn</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:assumeRoles[0].roleArn arn:aws:iam::123456789012:role/OrganizationAccountAccessRole
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> --path aws:assumeRoles[0].roleArn arn:aws:iam::123456789012:role/OrganizationAccountAccessRole
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="assumeroles-sessionname" data-link-title="assumeRoles[].sessionName"><span class="api-param-nested-indicator">&#8627;</span>assumeRoles[].sessionName<a class="api-param-anchor" href="#assumeroles-sessionname" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>An identifier for the assumed role session.</p>
<p class="api-param-meta">Config key: <code>aws:assumeRoles[0].sessionName</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:assumeRoles[0].sessionName pulumi-session
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> --path aws:assumeRoles[0].sessionName pulumi-session
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="assumeroles-sourceidentity" data-link-title="assumeRoles[].sourceIdentity"><span class="api-param-nested-indicator">&#8627;</span>assumeRoles[].sourceIdentity<a class="api-param-anchor" href="#assumeroles-sourceidentity" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Source identity specified by the principal assuming the role.</p>
<p class="api-param-meta">Config key: <code>aws:assumeRoles[0].sourceIdentity</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:assumeRoles[0].sourceIdentity &lt;value&gt;
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> --path aws:assumeRoles[0].sourceIdentity &lt;value&gt;
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="assumeroles-tags" data-link-title="assumeRoles[].tags"><span class="api-param-nested-indicator">&#8627;</span>assumeRoles[].tags<a class="api-param-anchor" href="#assumeroles-tags" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">map&lt;string, string&gt;</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Assume role session tags.</p>
<p class="api-param-meta">Config key: <code>aws:assumeRoles[0].tags</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:assumeRoles[0].tags &lt;value&gt;
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> --path aws:assumeRoles[0].tags &lt;value&gt;
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="assumeroles-transitivetagkeys" data-link-title="assumeRoles[].transitiveTagKeys"><span class="api-param-nested-indicator">&#8627;</span>assumeRoles[].transitiveTagKeys<a class="api-param-anchor" href="#assumeroles-transitivetagkeys" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">list&lt;string&gt;</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Assume role session tag keys to pass to any subsequent sessions.</p>
<p class="api-param-meta">Config key: <code>aws:assumeRoles[0].transitiveTagKeys</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:assumeRoles[0].transitiveTagKeys &lt;value&gt;
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> --path aws:assumeRoles[0].transitiveTagKeys &lt;value&gt;
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="customcabundle" data-link-title="customCaBundle">customCaBundle<a class="api-param-anchor" href="#customcabundle" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>File containing custom root and intermediate certificates. Can also be configured using the <code>AWS_CA_BUNDLE</code> environment variable. (Setting <code>caBundle</code> in the shared config file is not supported.)</p>
<p class="api-param-meta">Config key: <code>aws:customCaBundle</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:customCaBundle /etc/ssl/certs/corp-ca.pem
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> aws:customCaBundle /etc/ssl/certs/corp-ca.pem
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="defaulttags-tags" data-link-title="defaultTags.tags"><span class="api-param-nested-indicator">&#8627;</span>defaultTags.tags<a class="api-param-anchor" href="#defaulttags-tags" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">map&lt;string, string&gt;</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Resource tags to default across all resources. Can also be configured with environment variables like <code>TF_AWS_DEFAULT_TAGS_&lt;tag_name&gt;</code>.</p>
<p class="api-param-meta">Config key: <code>aws:defaultTags.tags</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:defaultTags.tags environment=production
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> --path aws:defaultTags.tags environment=production
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="ec2metadataserviceendpoint" data-link-title="ec2MetadataServiceEndpoint">ec2MetadataServiceEndpoint<a class="api-param-anchor" href="#ec2metadataserviceendpoint" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Address of the EC2 metadata service endpoint to use. Can also be configured using the <code>AWS_EC2_METADATA_SERVICE_ENDPOINT</code> environment variable.</p>
<p class="api-param-meta">Config key: <code>aws:ec2MetadataServiceEndpoint</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:ec2MetadataServiceEndpoint http://169.254.169.254
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> aws:ec2MetadataServiceEndpoint http://169.254.169.254
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="ec2metadataserviceendpointmode" data-link-title="ec2MetadataServiceEndpointMode">ec2MetadataServiceEndpointMode<a class="api-param-anchor" href="#ec2metadataserviceendpointmode" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Protocol to use with EC2 metadata service endpoint.Valid values are <code>IPv4</code> and <code>IPv6</code>. Can also be configured using the <code>AWS_EC2_METADATA_SERVICE_ENDPOINT_MODE</code> environment variable.</p>
<p class="api-param-meta">Config key: <code>aws:ec2MetadataServiceEndpointMode</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:ec2MetadataServiceEndpointMode IPv4
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> aws:ec2MetadataServiceEndpointMode IPv4
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="endpoints" data-link-title="endpoints">endpoints<a class="api-param-anchor" href="#endpoints" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">list&lt;object&gt;</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Per-service endpoint overrides — one field per AWS service, 315 in total. Set individual services by path, as shown below.</p>
<p class="api-param-meta">Config key: <code>aws:endpoints[0].s3</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:endpoints[0].s3 http://localhost:4566
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> --path aws:endpoints[0].s3 http://localhost:4566
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="forbiddenaccountids" data-link-title="forbiddenAccountIds">forbiddenAccountIds<a class="api-param-anchor" href="#forbiddenaccountids" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">list&lt;string&gt;</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p class="api-param-meta">Config key: <code>aws:forbiddenAccountIds</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:forbiddenAccountIds 123456789012
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> aws:forbiddenAccountIds 123456789012
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="httpproxy" data-link-title="httpProxy">httpProxy<a class="api-param-anchor" href="#httpproxy" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>URL of a proxy to use for HTTP requests when accessing the AWS API. Can also be set using the <code>HTTP_PROXY</code> or <code>httpProxy</code> environment variables.</p>
<p class="api-param-meta">Config key: <code>aws:httpProxy</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:httpProxy http://proxy.example.com:3128
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> aws:httpProxy http://proxy.example.com:3128
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="httpsproxy" data-link-title="httpsProxy">httpsProxy<a class="api-param-anchor" href="#httpsproxy" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>URL of a proxy to use for HTTPS requests when accessing the AWS API. Can also be set using the <code>HTTPS_PROXY</code> or <code>httpsProxy</code> environment variables.</p>
<p class="api-param-meta">Config key: <code>aws:httpsProxy</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:httpsProxy http://proxy.example.com:3128
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> aws:httpsProxy http://proxy.example.com:3128
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="ignoretags-keyprefixes" data-link-title="ignoreTags.keyPrefixes"><span class="api-param-nested-indicator">&#8627;</span>ignoreTags.keyPrefixes<a class="api-param-anchor" href="#ignoretags-keyprefixes" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">list&lt;string&gt;</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Resource tag key prefixes to ignore across all resources. Can also be configured with the TF_AWS_IGNORE_TAGS_KEY_PREFIXES environment variable.</p>
<p class="api-param-meta">Config key: <code>aws:ignoreTags.keyPrefixes</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:ignoreTags.keyPrefixes kubernetes.io/
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> --path aws:ignoreTags.keyPrefixes kubernetes.io/
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="ignoretags-keys" data-link-title="ignoreTags.keys"><span class="api-param-nested-indicator">&#8627;</span>ignoreTags.keys<a class="api-param-anchor" href="#ignoretags-keys" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">list&lt;string&gt;</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Resource tag keys to ignore across all resources. Can also be configured with the TF_AWS_IGNORE_TAGS_KEYS environment variable.</p>
<p class="api-param-meta">Config key: <code>aws:ignoreTags.keys</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:ignoreTags.keys kubernetes.io/cluster
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> --path aws:ignoreTags.keys kubernetes.io/cluster
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="insecure" data-link-title="insecure">insecure<a class="api-param-anchor" href="#insecure" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">boolean</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Explicitly allow the provider to perform "insecure" SSL requests. If omitted, default value is <code>false</code></p>
<p class="api-param-meta">Config key: <code>aws:insecure</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:insecure true
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> aws:insecure true
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="maxretries" data-link-title="maxRetries">maxRetries<a class="api-param-anchor" href="#maxretries" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">integer</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>The maximum number of times an AWS API request is being executed. If the API request still fails, an error is thrown.</p>
<p class="api-param-meta">Config key: <code>aws:maxRetries</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:maxRetries 3
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> aws:maxRetries 3
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="noproxy" data-link-title="noProxy">noProxy<a class="api-param-anchor" href="#noproxy" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Comma-separated list of hosts that should not use HTTP or HTTPS proxies. Can also be set using the <code>NO_PROXY</code> or <code>noProxy</code> environment variables.</p>
<p class="api-param-meta">Config key: <code>aws:noProxy</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:noProxy 169.254.169.254,.internal.example.com
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> aws:noProxy 169.254.169.254,.internal.example.com
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="profile" data-link-title="profile">profile<a class="api-param-anchor" href="#profile" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>The profile for API operations. If not set, the default profile created with <code>aws configure</code> will be used.</p>
<p class="api-param-meta">Config key: <code>aws:profile</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:profile my-profile
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> aws:profile my-profile
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="region" data-link-title="region">region<a class="api-param-anchor" href="#region" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required is-required">required</span>
<span class="api-param-badge api-param-in">AWS_REGION</span>
<span class="api-param-badge api-param-in">AWS_DEFAULT_REGION</span>
</div>
<div class="api-param-description">
<p>The region where AWS operations will take place. Examples are us-east-1, us-west-2, etc.</p>
<p class="api-param-meta">Config key: <code>aws:region</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:region us-west-2
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> aws:region us-west-2
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="retrymode" data-link-title="retryMode">retryMode<a class="api-param-anchor" href="#retrymode" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Specifies how retries are attempted. Valid values are <code>standard</code> and <code>adaptive</code>. Can also be configured using the <code>AWS_RETRY_MODE</code> environment variable.</p>
<p class="api-param-meta">Config key: <code>aws:retryMode</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:retryMode adaptive
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> aws:retryMode adaptive
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="s3useast1regionalendpoint" data-link-title="s3UsEast1RegionalEndpoint">s3UsEast1RegionalEndpoint<a class="api-param-anchor" href="#s3useast1regionalendpoint" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Specifies whether S3 API calls in the <code>us-east-1</code> region use the legacy global endpoint or a regional endpoint. Valid values are <code>legacy</code> or <code>regional</code>. Can also be configured using the <code>AWS_S3_US_EAST_1_REGIONAL_ENDPOINT</code> environment variable or the <code>s3UsEast1RegionalEndpoint</code> shared config file parameter</p>
<p class="api-param-meta">Config key: <code>aws:s3UsEast1RegionalEndpoint</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:s3UsEast1RegionalEndpoint regional
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> aws:s3UsEast1RegionalEndpoint regional
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="s3usepathstyle" data-link-title="s3UsePathStyle">s3UsePathStyle<a class="api-param-anchor" href="#s3usepathstyle" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">boolean</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Set this to true to enable the request to use path-style addressing, i.e., https://s3.amazonaws.com/BUCKET/KEY. By default, the S3 client will use virtual hosted bucket addressing when possible (https://BUCKET.s3.amazonaws.com/KEY). Specific to the Amazon S3 service.</p>
<p class="api-param-meta">Config key: <code>aws:s3UsePathStyle</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:s3UsePathStyle true
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> aws:s3UsePathStyle true
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="secretkey" data-link-title="secretKey">secretKey<a class="api-param-anchor" href="#secretkey" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
<span class="api-param-badge api-param-secret">secret</span>
</div>
<div class="api-param-description">
<p>The secret key for API operations. You can retrieve this from the 'Security &amp; Credentials' section of the AWS console.</p>
<p class="api-param-meta">Config key: <code>aws:secretKey</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> --secret myorg/myproject/dev pulumiConfig.aws:secretKey wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> --secret aws:secretKey wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="sharedconfigfiles" data-link-title="sharedConfigFiles">sharedConfigFiles<a class="api-param-anchor" href="#sharedconfigfiles" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">list&lt;string&gt;</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>List of paths to shared config files. If not set, defaults to [~/.aws/config].</p>
<p class="api-param-meta">Config key: <code>aws:sharedConfigFiles</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:sharedConfigFiles ~/.aws/config
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> aws:sharedConfigFiles ~/.aws/config
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="sharedcredentialsfiles" data-link-title="sharedCredentialsFiles">sharedCredentialsFiles<a class="api-param-anchor" href="#sharedcredentialsfiles" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">list&lt;string&gt;</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>List of paths to shared credentials files. If not set, defaults to [~/.aws/credentials].</p>
<p class="api-param-meta">Config key: <code>aws:sharedCredentialsFiles</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:sharedCredentialsFiles ~/.aws/credentials
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> aws:sharedCredentialsFiles ~/.aws/credentials
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="skipcredentialsvalidation" data-link-title="skipCredentialsValidation">skipCredentialsValidation<a class="api-param-anchor" href="#skipcredentialsvalidation" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">boolean</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Skip the credentials validation via STS API. Used for AWS API implementations that do not have STS available/implemented.</p>
<p class="api-param-meta">Config key: <code>aws:skipCredentialsValidation</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:skipCredentialsValidation true
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> aws:skipCredentialsValidation true
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="skipmetadataapicheck" data-link-title="skipMetadataApiCheck">skipMetadataApiCheck<a class="api-param-anchor" href="#skipmetadataapicheck" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">boolean</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Skip the AWS Metadata API check. Used for AWS API implementations that do not have a metadata api endpoint.</p>
<p class="api-param-meta">Config key: <code>aws:skipMetadataApiCheck</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:skipMetadataApiCheck true
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> aws:skipMetadataApiCheck true
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="skipregionvalidation" data-link-title="skipRegionValidation">skipRegionValidation<a class="api-param-anchor" href="#skipregionvalidation" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">boolean</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Skip static validation of region name. Used by users of alternative AWS-like APIs or users w/ access to regions that are not public (yet).</p>
<p class="api-param-meta">Config key: <code>aws:skipRegionValidation</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:skipRegionValidation true
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> aws:skipRegionValidation true
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="skiprequestingaccountid" data-link-title="skipRequestingAccountId">skipRequestingAccountId<a class="api-param-anchor" href="#skiprequestingaccountid" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">boolean</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Skip requesting the account ID. Used for AWS API implementations that do not have IAM/STS API and/or metadata API.</p>
<p class="api-param-meta">Config key: <code>aws:skipRequestingAccountId</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:skipRequestingAccountId true
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> aws:skipRequestingAccountId true
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="stsregion" data-link-title="stsRegion">stsRegion<a class="api-param-anchor" href="#stsregion" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>The region where AWS STS operations will take place. Examples are us-east-1 and us-west-2.</p>
<p class="api-param-meta">Config key: <code>aws:stsRegion</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:stsRegion us-east-1
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> aws:stsRegion us-east-1
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="tagpolicycompliance" data-link-title="tagPolicyCompliance">tagPolicyCompliance<a class="api-param-anchor" href="#tagpolicycompliance" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>The severity with which to enforce organizational tagging policies on resources managed by this provider instance. At this time this only includes compliance with required tag keys by resource type. Valid values are "error", "warning", and "disabled". When unset or "disabled", tag policy compliance will not be enforced by the provider. Can also be configured with the TF_AWS_TAG_POLICY_COMPLIANCE environment variable.</p>
<p class="api-param-meta">Config key: <code>aws:tagPolicyCompliance</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:tagPolicyCompliance warn
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> aws:tagPolicyCompliance warn
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="token" data-link-title="token">token<a class="api-param-anchor" href="#token" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">string</span>
<span class="api-param-badge api-param-required">optional</span>
<span class="api-param-badge api-param-secret">secret</span>
</div>
<div class="api-param-description">
<p>session token. A session token is only required if you are using temporary security credentials.</p>
<p class="api-param-meta">Config key: <code>aws:token</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> --secret myorg/myproject/dev pulumiConfig.aws:token IQoJb3JpZ2luX2VjEExample
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> --secret aws:token IQoJb3JpZ2luX2VjEExample
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="tokenbucketratelimitercapacity" data-link-title="tokenBucketRateLimiterCapacity">tokenBucketRateLimiterCapacity<a class="api-param-anchor" href="#tokenbucketratelimitercapacity" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">integer</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>The capacity of the AWS SDK's token bucket rate limiter.</p>
<p class="api-param-meta">Config key: <code>aws:tokenBucketRateLimiterCapacity</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:tokenBucketRateLimiterCapacity 3
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> aws:tokenBucketRateLimiterCapacity 3
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="usedualstackendpoint" data-link-title="useDualstackEndpoint">useDualstackEndpoint<a class="api-param-anchor" href="#usedualstackendpoint" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">boolean</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Resolve an endpoint with DualStack capability</p>
<p class="api-param-meta">Config key: <code>aws:useDualstackEndpoint</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:useDualstackEndpoint true
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> aws:useDualstackEndpoint true
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="usefipsendpoint" data-link-title="useFipsEndpoint">useFipsEndpoint<a class="api-param-anchor" href="#usefipsendpoint" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">boolean</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Resolve an endpoint with FIPS capability</p>
<p class="api-param-meta">Config key: <code>aws:useFipsEndpoint</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:useFipsEndpoint true
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> aws:useFipsEndpoint true
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
<div class="api-param">
<h3 class="api-param-name" id="useragents" data-link-title="userAgents">userAgents<a class="api-param-anchor" href="#useragents" aria-label="Link to this option">#</a></h3>
<div class="api-param-header">
<span class="api-param-badge api-param-type">list&lt;string&gt;</span>
<span class="api-param-badge api-param-required">optional</span>
</div>
<div class="api-param-description">
<p>Product details to append to the User-Agent string sent in all AWS API calls.</p>
<p class="api-param-meta">Config key: <code>aws:userAgents</code></p>
<pulumi-chooser type="configsource" options="esc,stack">
<pulumi-choosable type="configsource" values="esc">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi env <span class="nb">set</span> myorg/myproject/dev pulumiConfig.aws:userAgents my-app/1.0
</span></span></code></pre></div>
</pulumi-choosable>
<pulumi-choosable type="configsource" values="stack">
<div class="highlight"><pre class="chroma"><code class="language-bash" data-lang="bash"><span class="line"><span class="cl">$ pulumi config <span class="nb">set</span> aws:userAgents my-app/1.0
</span></span></code></pre></div>
</pulumi-choosable>
</pulumi-chooser>
</div>
</div>
</div>
