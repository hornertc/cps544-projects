# Go Programming

## JavaScript Object Notation

---

## Introduction

- [JSON](https://www.json.org/json-en.html) is the lingua franca for machine to machine communication
  - much to the chagrin of XML and binary format advocates
  - simplicity, (human) readability, and universal support
- Only support primitive types
  - boolean, number, string, object, arrays

----

```json
{
    "field1": 5,
    "other field": {"x": 4.5},
    "list": ["a", -1, 4.5, "Hello\u29A7"]
}
```

- No trailing commas

----

![JSON Object](https://www.json.org/img/object.png)

----

![JSON Array](https://www.json.org/img/array.png)

----

![JSON Value](https://www.json.org/img/value.png)

---

## JSON in Go

- [encoding/json](http://golang.org/pkg/encoding/json) implements JSON marshalling and unmarshalling
- Marshalling:
  - Go data types to a `[]byte` encoded as JSON
  - encoding to JSON
- Unmarshalling:
  - `[]byte` encoded as JSON to Go data types
  - decoding from JSON

----

## Default Mapping to Go Types

- number to `float64` (even if integer)
- boolean to `bool`
- string to `string` (with escaping, e.g., `\n\t`)
- array to `slice`
- object to `map[string]any`

---

## Struct Tags

- String data attached to fields
- Stored in binary (unlike comments)

```go
type Record struct {
    A string `json:"day" yaml:"day"`
    B int `json:"target,omitempty"`
    c int `docs:"This does something"`
}
```

- By convention, string is broken into the format
  - `key1:"value1" key2:"value2"`
  - `key1` can easily be extracted independent of other data in the struct tag

---

## Marshalling

- Can marshal most any type
  - excludes channels, functions, etc.
- Structs are converted to JSON objects
  - Only exported fields are encoded/decoded
  - Struct tags may be used to denote field names
- Reflection provides type information
- Map keys are sorted

```go
value := []int{1, 2, 3}
data, err := json.Marshal(value) // [1, 2, 3]
```

---

## Unmarshalling

- Pass a pointer of the data type you want the JSON coerced into
- Reflection is used to determine the types
- Values are zero-initialized *then* populated by available JSON data

```go
var value map[string][]int
err := json.Unmarshal(data, &value)
// value is now non-nil
```

Notes:

- Show `$GOPL/ch4/movie/main.go`
- Show `curl -v https://xkcd.com/571/info.0.json | jq`
  - `jq .alt`
  - Can write Go code to parse it

---

## Best Practices

- Object should be the top level type
  - even if encoding a single array it is better to have an object at the top

```json
{
    "stockPrices": [15.2, 2, 3, 4.5]
}
```

----

## Binary Data

- `[]byte` is encoded into a string in JSON encoded as base64

- base64 encodes 6 bits ($2^6=64$) at a time
  - using characters `A...Za...z0-9+/` (26+26+10+2=64)
  - `=` is used as padding
  - compact and safe encoding of binary data

----

## Detecting Un-set Values

- When unmarshalling types get their zero-value

```go
// {"day": "Friday", "target": 45}, Target is 45
// {"day": "Friday", "target": 0}, Target is 0
// {"day": "Friday"}, Target is 0
type LaunchMissileAction struct {
    Day string `json:"day"`
    Target int `json:"target"`
}

// {"day": "Friday", "target": 0}, Target is pointer to 0, new(0)
// {"day": "Friday"}, Target is nil
type LaunchMissileAction struct {
    Day string `json:"day"`
    Target *int `json:"target"`
}
```

---

## Comparison to Other Languages

- Assume the top level value is an object...
- Javascript/ECMAScript's `JSON.parse()` returns a **new** object
- Python's `json.load()` returns a `dict`
- C++ (with libraries) parses to an abstract `value`
  - Must use generic type-safe getters and setters

Notes:

- None of the mentioned languages unmarshal to native types

----

- Go unmarshalls to native or custom Go types
  - Accessing fields is same as always
- Rust uses macros to generate code to de/serialize types
  - Can use native or custom types just like Go

---

## Alternatives to JSON

- Parsing JSON can be slow
- Go natively [supports many other encodings](https://pkg.go.dev/encoding#section-directories)
- Raw binary formats
- Protobuf
- Msgpack
