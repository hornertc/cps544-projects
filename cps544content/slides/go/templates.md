# Go Programming

## Templates

---

## Introduction

- Formatting structured data (in increasing expressiveness)
  - JSON/YAML/TOML
  - `fmt.Printf`
  - Templates

---

## Go Templates

- `text/template` and `html/template`
- Provides mechanism for substituting values of variables into text or HTML
- *template* - string or file containing text (or HTML) and double braces `{{ ... }}` called actions
- Everything not an action is printed literally
- Actions trigger various behaviors

----

- Each action contains an expression in the template
language; simple but powerful notation for:
  - printing values
  - selecting struct fields
  - calling functions and methods
  - expressing control flow such as if-else statements and range loops
  - instantiating other templates

----

- Template is parsed then executed with values of type `any`
- Values can be anything but often a struct
- Template may access exported fields and methods of values

---

## Examples

- GitHub Issues
  - text
  - HTML

Notes:

- Show `$GOPL/ch4/github`
- `go run ./ch4/issuesreport repo:golang/go is:open json decoder`
- `go run ./ch4/issueshtml repo:golang/go commenter:gopherbot json encoder`

---

## Built-in Functions

- `len` - returns length
- `slice` - slices a string, slice, or array
  - `slice x 1 5` is like Go's `x[1:5]`
- `index` - index into a map, slice, array
  - `index x "food"` is like Go's `x["food"]`
- And a dozen more

----

## Custom Functions

- Can add your own functions and call methods
- Can take any number of arguments
- Must return either
  - one value
  - one value followed by an `error`
    - if non-nil error then templating stops

---

## Another Example

`go run ./examples/tmpl examples/tmpl`

---

## HTML Templates

- `html/template` package:
  - uses same API and expression language as `text/template`
  - adds features for automatic and context-appropriate escaping of strings appearing within HTML, JavaScript, CSS, or URLs

----

- These features help avoid an *injection attack*
  - perennial security problem with HTML generation tools
  - an adversary crafts a string value like the title of an issue to include malicious code that, when improperly escaped by a template, gives them control over the page

----

```shell
go run ./ch4/issueshtml repo:golang/go 3133 10535 json encoder
```

- Notice the escaping
  - `<` converted to `&lt`
  - `>` converted to `&gt`

----

## Disable Escaping

- Sometimes the value comes from a trusted source and needs to be rendered *as is*
- Can suppress auto-escaping in HTML templates
  - Use `template.HTML` instead of `string`

----

## Go Documentation Generation

- [Documentation conventions](https://go.dev/doc/comment)
- Text and HTML templates are used in `go doc`
- `go doc PKGNAME` dumps the exported API to STDOUT
  - `go doc -all PKGNAME` also outputs the documentation comments
- `godoc` is another tool
  - install with `go install golang.org/x/tools/cmd/godoc@latest`
  - serves HTTP website similar to `https://pkg.go.dev`

---

## More Examples

- Customizing output
  - `kubectl get pods -o go-template='{{range .items}}Pod {{.metadata.name}} is on node {{.spec.nodeName}}, {{end}}'`
- Generating test data (see [ASCE Data Telemetry](https://gitlab.com/act3-ai/asce/data/telemetry/-/tree/main/testdata))
  - Even generating binary data is possible (see [ASCE Data Tool](https://gitlab.com/act3-ai/asce/data/tool/-/tree/main/internal/mirror/testing))

----

### User defined functions or mappings

- definition is a file or string (template)
- argument is a complex type (template values)
- returns a relatively simple type (rendered)
  - string
  - binary data (as a string, thus non-UTF-8 encoded bytes)
  - more complex data encoded as a string (e.g., JSON)

----

### User Defined Mapping

`oci.Descriptor` to list of OCI image references

```json
{
  "ref": "docker.io/library/busybox",
  {
    "digest": "sha256:5837e622ab93579f37257b1c999f9e7e2bd438385ca0d9593a77a8ea7024081d",
    "mediaType": "application/vnd.oci.image.manifest.v1+json",
    "size": 610
  }
}
```

```txt
myreg.private.local:5000/busybox
myotherreg.other/verybusybox
```
