{{- $content := . -}}
{{- /* Phase 1: Convert Chroma syntax-highlighted HTML to fenced code blocks */ -}}
{{- $content = replaceRE `<div[^>]*>\s*<pre[^>]*><code class="language-([^"]*)"[^>]*>` "```$1\n" $content -}}
{{- $content = replaceRE `</code></pre>\s*</div>` "\n```\n\n" $content -}}
{{- $content = replaceRE `<pre[^>]*><code>` "```\n" $content -}}
{{- $content = replaceRE `</code></pre>` "\n```\n\n" $content -}}
{{- /* Phase 1b: Convert pulumi-chooser/choosable web components to comment markers */ -}}
{{- $content = replaceRE `<pulumi-chooser[^>]*type="([^"]*)"[^>]*>` "\n<!-- chooser: $1 -->\n" $content -}}
{{- $content = replaceRE `</pulumi-chooser>` "\n<!-- /chooser -->\n" $content -}}
{{- $content = replaceRE `<pulumi-choosable[^>]*values="([^",]*)[^"]*"[^>]*>` "\n<!-- option: $1 -->\n" $content -}}
{{- $content = replaceRE `</pulumi-choosable>` "\n<!-- /option -->\n" $content -}}
{{- /* Phase 1c: Convert definition lists to markdown before stripping */ -}}
{{- $content = replaceRE `<dt>([^<]*)</dt>` "\n**$1:** " $content -}}
{{- $content = replaceRE `</?d[dl][^>]*>` "" $content -}}
{{- /* Phase 1d: Convert list items to markdown bullets before stripping */ -}}
{{- $content = replaceRE `<li[^>]*>` "\n- " $content -}}
{{- $content = replaceRE `</li>` "" $content -}}
{{- $content = replaceRE `</?[ou]l[^>]*>` "" $content -}}
{{- /* Phase 2: Strip all block-level and decorative tags in consolidated passes. */ -}}
{{- $content = replaceRE `</?(?:span|label|div|p|blockquote|pre|section|table|thead|tbody|tr|td|th|details|summary)[^>]*>` "" $content -}}
{{- $content = replaceRE `(?:<i[^>]*></i>|<input[^>]*>)` "" $content -}}
{{- /* Strip non-functional HTML comments (keep chooser/option markers) */ -}}
{{- $content = replaceRE `<!--\s*markdownlint[^>]*-->` "" $content -}}
{{- $content = replaceRE `<!-- WARNING:.*?-->` "" $content -}}
{{- $content = replaceRE `<!-- Do not edit by hand.*?-->` "" $content -}}
{{- /* Phase 3: Normalize whitespace so inline conversions can match cleanly */ -}}
{{- $content = replaceRE `\n{3,}` "\n\n" $content -}}
{{- /* Phase 4: Convert inline HTML to markdown.
     Note: The link regex matches plain-text content only ([^<]*). Links containing nested tags
     (e.g. <code> inside <a>) won't match here and fall through to Phase 5's catch-all strip,
     losing the URL. This is an inherent trade-off of regex-based HTML conversion. */ -}}
{{- $content = replaceRE `<a[^>]*href="([^"]*)"[^>]*>([^<]*)</a>` "[$2]($1)" $content -}}
{{- /* Collapse whitespace inside markdown link brackets (from multi-line <a> tags) */ -}}
{{- $content = replaceRE `\[\s+` "[" $content -}}
{{- $content = replaceRE `\s+\]\(` "](" $content -}}
{{- $content = replaceRE `<code>([^<]*)</code>` "`$1`" $content -}}
{{- $content = replaceRE `<strong>([^<]*)</strong>` "**$1**" $content -}}
{{- $content = replaceRE `<em>([^<]*)</em>` "*$1*" $content -}}
{{- /* Phase 5: Convert heading tags to markdown headings */ -}}
{{- $content = replaceRE `<h1[^>]*>([^<]*)</h1>` "\n# $1\n" $content -}}
{{- $content = replaceRE `<h2[^>]*>([^<]*)</h2>` "\n## $1\n" $content -}}
{{- $content = replaceRE `<h3[^>]*>([^<]*)</h3>` "\n### $1\n" $content -}}
{{- $content = replaceRE `<h4[^>]*>([^<]*)</h4>` "\n#### $1\n" $content -}}
{{- $content = replaceRE `<h5[^>]*>([^<]*)</h5>` "\n##### $1\n" $content -}}
{{- $content = replaceRE `<h6[^>]*>([^<]*)</h6>` "\n###### $1\n" $content -}}
{{- /* Strip any remaining HTML tags not yet handled */ -}}
{{- $content = replaceRE `</?a[^>]*>` "" $content -}}
{{- /* Strip lone +/- lines left from accordion toggle spans */ -}}
{{- $content = replaceRE `(?m)^[+\-]\n` "" $content -}}
{{- /* Phase 6: Decode HTML entities and final cleanup */ -}}
{{- $content = $content | htmlUnescape -}}
{{- $content = replaceRE `\n{3,}` "\n\n" $content -}}
{{- return $content -}}
