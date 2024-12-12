# Documentation generation (docsgen)

This package supports generating documentation for Pulumi packages from their
schema.

## Testing

To run the tests in this package, use `go test`:

```
go test ./...
```

Test data for `GeneratePackage` (see `TestGeneratePackage/*`) is stored under
`testdata/TestGeneratePackage`. Each subdirectory represents a test case, and
contains the following:

* `schema.json` or `schema.yaml` -- a valid Pulumi schema that will serve as
  input to GeneratePackage
* `docs` -- a directory containing the expected output of GeneratePackage. To
  generate this on a first run, or whenever you have made changes that need to
  be reflected in expected outputs, set `PULUMI_ACCEPT` to any value before
  running the tests.

### Adding a new test case for `GeneratePackage`

To add a test case for `GeneratePackage`, create a new subdirectory under
`testdata/TestGeneratePackage` (e.g.
`testdata/TestGeneratePackage/my-test-case`) and update `TestGeneratePackage`'s
inputs as follows:

* Add a new entry to the `testCases` array. The `directory` should match exactly
  the name of the directory you created (e.g. `my-test-case`), and the
  `description` should be a suitable human-readable description of the case. For
  each of navigation and search, keep the array sorted alphabetically by
  `directory`.
* Place a `schema.json` or `schema.yaml` in the new directory.
* If your schema references providers that are not already present in the
  `testCaseProviders` array, you will need to add both an entry to this array
  and a schema for the referenced provider(s). Schema files for providers should
  be placed in `testdata/TestGeneratePackage/_schema` and named
  `<name>-<version>.json`, where `<name>` and `<version>` match the `name` and
  `version` added to the `testCaseProviders` array.

To run the tests and save the generated output for validating future runs, use
`go test` with the `PULUMI_ACCEPT` environment variable set:

```
PULUMI_ACCEPT=1 go test ./... -run TestGeneratePackage/my-test-case
```

The `-run` flag is optional but should make things faster.

### Updating test cases with golden outputs

If you have made changes that affect test cases with golden outputs, use
`PULUMI_ACCEPT` and run all the tests/the tests which are failing, e.g.:

```
PULUMI_ACCEPT=1 go test ./...
```

## Implementation

### Templates

This package makes use of Go's built-in `html/template` package to produce HTML
from templates and input data. Although we are using the `html/template`
package, it has the same exact interface as the
[`text/template`](https://golang.org/pkg/text/template) package, except for some
HTML specific things. Therefore, all of the functions available in the
`text/template` package are also available with the `html/template` package.

Cheatsheet:

* Data can be injected using `{{.PropertyName}}`.
* Nested properties can be accessed using the dot notation, i.e.
  `{{.Property1.Property2}}`.
* Templates can inject other templates using the `{{template "template_name"}}`
  directive. For this to work, you will need to first define the named template
  using `{{define "template_name"}}`.
* You can pass data to nested templates by simply passing an argument after the
  template's name.
* To remove whitespace from injected values, use the `-` in the template tags.
  For example, `{{if .SomeBool}} some text {{- else}} some other text {{-
  end}}`. Note the use of `-` to eliminate whitespace from the enclosing text.
  Read more [here](https://golang.org/pkg/text/template/#hdr-Text_and_spaces).
* To render un-encoded content use the custom global function `htmlSafe`.
  **Note**: This should only be used if you know for sure you are not injecting
  any user-generated content, as it bypasses HTML encoding.
* To render strings to Markdown, use the custom global function `markdownify`.
* To print regular strings, that share the same syntax as the Go templating
  engine, use the built-in global function `print`
  [function](https://golang.org/pkg/text/template/#hdr-Functions).

Learn more from here:
https://curtisvermeeren.github.io/2017/09/14/Golang-Templates-Cheatsheet
