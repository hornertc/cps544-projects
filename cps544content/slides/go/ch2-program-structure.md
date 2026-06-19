# Go Programming

## Program Structure

[Go Language Specification](https://go.dev/ref/spec)

---

### Overview

- Names
- Declarations
- Variables
- Assignments
- Type Declarations
- Packages and Files
- Scope

---

### Names

- functions, variables, constants, types, statement labels (used by GOTO, break), packages
- begins with letter or an underscore
- any number of additional letters, digits, and underscores
- case sensitive

----

### Keywords and Operators

```txt
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

```txt
+    &     +=    &=     &&    ==    !=    (    )
-    |     -=    |=     ||    <     <=    [    ]
*    ^     *=    ^=     <-    >     >=    {    }
/    <<    /=    <<=    ++    =     :=    ,    ;
%    >>    %=    >>=    --    !     ...   .    :
     &^          &^=          ~
```

----

### Predeclared

```txt
Types:
	any bool byte comparable
	complex64 complex128 error float32 float64
	int int8 int16 int32 int64 rune string
	uint uint8 uint16 uint32 uint64 uintptr

Constants:
	true false iota

Zero value:
	nil

Functions:
	append cap clear close complex copy delete imag len
	make max min new panic print println real recover
```

----

### Names and Scope

- declared within a function: local to that function
- entity declared outside of a function: visible in **all** files of the
package to which it belongs
- package names: always in lower case
- case of first letter determines visibility across
package boundaries

----

- upper-case letter: exported
  - visible and accessible outside of its own package
  - e.g.: Printf in the fmt package
- lower-case letter: un-exported (private to package)
  - un-exported fields of struct are accessible to entire package

----

### Camel Case

- prefer camel case instead of underscores
  - `QuoteRuneToASCII`, `parseRequestLine`
  - not `parse_request_line`
- uppercase acronyms
  - `HTMLEscape` not `HtmlEscape`

---

### Declarations

- declaration names program entity and specifies
some or all of its properties
- four major kinds of declarations
  - `var`
  - `const`
  - `type`
  - `func`

----

### Function Declarations

```go
func dump(i int, y float32) {
    fmt.Println(i, y)
}

func format(i int, y float32) string {
    return fmt.Sprintf("value at %d is %f", x, y)
}

func twoReturns(i, j int) (int, int) {
    return i+1, i*j
}

func namedReturns(s string) (x, y string) {
    x = s + " suffix"
    y = "prefix " + s
    return
}
```

----

### Variable Declarations

- creates a variable of a particular type (strongly typed)
- attaches a name to variable
- sets its initial value
  - **never** uninitialized
  - removes a whole class of vulnerabilities

----

- `var name type = expression`
- `type` or `= expression` maybe omitted but not both
- if no expression then *zero value* is used
  - `0` for numbers
  - `false` for booleans
  - `""` for strings
  - `nil` for reference types
    - interface, slice, pointer, map, channel, function
- zero value for all elements of aggregate types
  - struct, array
  
----

- choose zero values so they represent reasonable defaults
  - global variables
  - fields of structs
- `DisableSecurity bool` is much better than `EnableSecurity bool`
  - Want "secure by default"
  - Developer has to explicitly disable security by setting `DisableSecurity=true`

----

```go
var i, j, k int // int, int, int
var b, f, s = true, 2.3, "four" // bool, float64, string
```

- package-level variables initialized before `main()` begins

----

```go
var f, err = os.Open(name) // os.Open returns file and error
f, err := os.Open(name) // short variable declaration
```

```go
i, j := 10, 11 // multi-variable declaration (short form)
i, j = 10, 11 // tuple assignment
i, j = j, i // tuple assignment
a[i], a[j] = a[j], a[i] // swap, entire RHS evaluated first
```

```go
in, err := os.Open(infile) // declares in and err
// ...
out, err := os.Create(outfile) // declares out, err is assigned
in, out, err := f() // compilation error, no new variable LHS
in, out, err = f() // assignment only (no declarations)
```

---

### Variables

- variable: piece of storage containing a value
- variables created by declarations are identified by name, `x`
- elements of aggregate types are variables but do not have a name
  - `x[3]`, `y.f`
- if on RHS of assignment value is accessed (read)
- if on LHS of assignment value is assigned (written) to the variable

----

- pointers allow naming/renaming variables
- pointer value: address of a variable
  - location at which a value is stored
- every variable has an address
  - not every value has an address
- pointer can read/update value of a variable indirectly
  - do not need to know the name of the variable

----

### Fundamental Theorem of Software Engineering

Andrew Koenig states in [his "theorem"](https://en.wikipedia.org/wiki/Fundamental_theorem_of_software_engineering)

> We can solve any problem by introducing an extra level of indirection.

- Pointers are one level of indirection

----

```go
x := 1
p := &x // p, of type *int, points to x
fmt.Println(*p) // "1"
*p = 2
// equivalent to x = 2
fmt.Println(x) // "2"
```

----

- `p != nil` is true if `p` points to a variable

```go
var x, y int
fmt.Println(&x == &x, &x == &y, &x == nil)
```

- prints `"true false false"`

----

```go
func incr(p *int) int {
  x := *p
	*p++ // increments what p points to; does not change p
	return x
}

// elsewhere in a function
v := 1
incr(&v) 
fmt.Println(v) // prints "2"
fmt.Println(v, incr(&v), v, incr(&v)) // side effect: v is now 4
```

- prints `"4 2 4 3"` <!-- .element: class="fragment" -->

- <!-- .element: class="fragment" -->
  - `*p` is an alias for `v`
    - to find all statements that access a variable, one must find all aliases

----

### Example: Echo4 (Flags)

```go
package main

import (
  "flag"
  "fmt"
  "strings"
)

var n = flag.Bool("n", false, "omit trailing newline") //n *bool
var sep = flag.String("s", " ", "separator") // sep *string

func main() {
  flag.Parse()
  fmt.Print(strings.Join(flag.Args(), *sep))
  if !*n {
    fmt.Println()
  }
}
```

Notes:

- `go run $GOPL/ch2/echo4/main.go`

----

- safe to return the address of a local variable
  - Go uses *escape analysis* to determine how to allocate (stack vs. heap)
  - Variables exist at least as long as they are *reachable*

```go
var p = f()
func f() *int {
  v := 1
  return &v
} // v escapes f() thus heap allocated
```

- each call to `f()` returns a distinct value:

```go
fmt.Println(f() == f()) // "false"
```

Notes:

- variables may live longer before/if they are garbage collected
- hard to determine if a variable is stack or heap allocated
  - compiler dependent
  - use memory profiling to determine heap allocations

---

### `new()`

- another way to create a variable: `new(T) *T`
  - creates unnamed temporary of type `T`
  - initializes it to zero value of `T`
  - returns its address: value of type `*T`
- only a syntactic convenience

```go
x := new(float64)
// vs
var f float64
x := &f
```

----

```go
p := new(int)
q := new(int)
fmt.Println(p == q)
```

- prints `false` <!-- .element: class="fragment" -->
- <!-- .element: class="fragment" --> each call to `new()` returns a distinct variable with a unique address
  - unless it is zero-sized, then undefined
    - e.g., `struct{}`, `[0]int`

----

- `new` is a predeclared function, not a keyword
  - can redefine/shadow `new` within a function

```go
func delta(old, new int) int { return new - old }
```

- within `delta()` function `new` is shadowed and thus unavailable

---

### Assignments

```go
i = 1 // named variable
*p = true // indirect variable
person.name = "bob" // struct field
y[i] = y[i] * s // array or slice or map element
y[i] *= s // equivalent to above, compound assignment
```

----

### Postfix Increment and Decrement

```go
v := 1
v++ // same as v = v + 1
v-- // same as v = v - 1
++v // invalid! no prefix form (unlike C++)
fmt.Println(v++) // invalid! no return value (unlike C++)
```

---

### Type Declarations

- type of variable
  - size: number of bytes
  - internal representation format
  - intrinsic operations allowed
  - associated methods
  - sometimes *semantic meaning*

----

- variables and types: same representation but different semantic meaning
  - `int`: loop index, timestamp, file descriptor, month, week
  - `float64`: velocity (m/s) of an object, temperature
  - `string`: message, password, name of a color, encoded data
- `type name underlying-type` is used to define a new type
  - `type week int` declares a new type `week`
    - same representation as `int`
    - cannot be mixed with `int` or `day` or `year`

----

```go
type DayOfWeek int
type Week int
type Year int
func date(dow DayOfWeek, wk Week, yr Year) *Date { ... }

// in a function somewhere
var dow DayOfWeek = 2
var wk Week = 5
var yr Year = 2023
wk = wk + 1 // valid: intrinsic operation (add integer)
wk = wk + Week(10) // valid
yr = yr + dow // invalid: mismatched types
date(dow, wk, yr)  // valid
date(dow, wk, wk)  // invalid
fmt.Println(wk == Week(5)) // valid: true
fmt.Println(wk == yr) // invalid: mismatched types
```

Notes:

- enum design pattern defined in chapter 3

---

### Packages and Files

- packages similar to libraries in other languages
  - modularity
  - encapsulation
  - separate compilation
  - reuse
- directory with one or more `.go` files
  - package name must be the directory name except:
    - `package main`
    - `package *_test`

Notes:

- encapsulation - exported identifiers start with an upper-case letter
- separate compilation - faster compilation by way of parallelized builds
- reuse code from your project (internal dependencies) and other projects (external dependencies)

----

### GOPATH

- `$GOPATH` stores source code dependencies, binaries
  - default value `~/go` (run `go env GOPATH`)
  - can add `export GOPATH=/path/to/somewhere/goLib` to `~/.bashrc`
- Prior to go 1.11: `$GOPATH/src` contained **your** source code
  - `$GOPATH/src` location determines import path
  - `$GOPATH/src/github.com/spf13/cobra` imported as `import "github.com/spf13/cobra"`

----

### Modules

- Since go 1.11 (preferred method)
- collection of packages with a `go.mod` file
  - `go.mod` declares dependencies, module name, min Go version
- developed anywhere on your filesystem
  - not just `$GOPATH/src`
- module dependencies downloaded to `$GOPATH/pkg/mod` and versioned
  - may choose to "vendor" dependencies
- `go mod` and `go get` interact with dependencies

----

### Package Comment

- package comment is a special comment preceding `package ...`
  - only allowed in one file in package
  - `doc.go` is often that file.  Often contains no code.

```go
// Package linalg provides linear algebra construct including vector, matrix, matrix operations, and matrix decomposition functions
package linalg
```

Notes:

- show [krew](https://github.com/kubernetes-sigs/krew)
- `internal/` vs. `pkg/` vs. `cmd/`
- [package comment](https://github.com/kubernetes-sigs/krew/blob/master/internal/index/validation/validate.go)
- [import](https://github.com/kubernetes-sigs/krew/blob/master/cmd/krew/cmd/index.go)

---

### Imports

- **every** package has a unique import path
  - `github.com/spf13/cobra`
  - often includes domain name
  - even internal packages require a full import path
    - packages with `internal` in the path
    - inaccessible outside the module

----

### Versioning Imports

- Go developers expect semantic versioning of modules
  - breaking changes must change the major version
- Breaking change *should* require import change
  - `import "github.com/go-chi/chi/v5"`
  - developer makes code changes to use the new API

----

### Package's API

- Internal packages have no external API
- All exported functions, types, constants, variables
- None of the un-exported ...
- Avoid breaking changes: make API as *small* and as *stable* as possible
  - Adding func/type/const/var: *not* a breaking change
  - Removing func/type/const/var: breaking change

---

### Package Initialization

- packages form a directed acyclic graph (DAG)
  - import cycles generate errors
- packages initialized in [topological order](https://en.wikipedia.org/wiki/Topological_sorting)
- procedure to initialize a package:
  - initialize all package dependencies (guaranteed by topological sort)
  - `.go` files are initialized in sorted order
  - initialize variables, for each variable
    - initialize variables on RHS (recursively)
    - evaluate RHS
  - call `init()` if exists

----

### Globals

- `init()`
  - only for package initialization
  - cannot be called by your code
  - use sparingly, prefer variable initialization
  - useful for complex initialization
- **warning** exported package-level variable writable by **any** package that imports it
  - `var SomeVariable = 45`
  - `SomeVariable` might not be 45 when your code runs
  - `const` does not have this issue

---

### Scope

- declaration associates name with program entity
  - e.g., function or variable
- scope of declaration: part of source code where use of declared name refers to that declaration

----

- scope and lifetime are decoupled
  - scope of declaration
    - region of program text
    - known at compile time
  - lifetime of variable
    - range of time during execution when variable can be referred to
    - known at run time
    - lifetime outlives scope

----

### Syntactic Block

- sequence of statements enclosed in braces `{}`
  - surround body of function or loop
- name declared inside a syntactic block
  - only visible inside that block
- block encloses declarations and determines their scope
  - compile time property

----

### Lexical Block

- blocks generalized to include other groupings of declarations not explicitly surrounded by braces in source code
  - called lexical blocks
- lexical block for the entire source code, called
universe block:
  - for each package
  - for each file
  - for each `for`, `if`, and `switch` statement
  - for each `case` in `switch` or `select` statement
  - for each explicit syntactic block

Notes:

- lexical blocks generalize syntactic block
- a syntactic block is a lexical block
- file lexical scope for file is used for scoping imports to that file

----

### Universal Block

- declarations of built-in types, functions, and constants (e.g., `int`, `len`, and `true`):
  - in universe block
  - can be referred to throughout the entire program

----

### Package Block

- declarations outside any function (at package level) can be referred to from any file in same package

----

### File-Scoped Lexical Block

- imported packages
  - can be referred to from the same file
  - can change the file-local name of the package
  - cannot be referred to from another file in same package without another import

```go
import (
  "fmt"
  exp "example.com/exp/wowthisisalongname"
)
// can use exp.DoSomething()
```

----

### Multiple Declarations

- multiple declarations of the same name are allowed
  - each declaration must be in a different lexical block
  - called variable *shadowing*
  - somewhat discouraged, confusing to reader

----

### Name Resolution

- when compiler encounters a reference to a name:
  - looks for a declaration:
    - starting with innermost enclosing lexical block
    - working up to universe block
- if compiler finds no declaration:
  - reports an "undeclared name" error

----

```go
x := "-hello!-"
for _, x := range x {
  x := x
  if d := 'A' - 'a'; x >= 'a' && x <= 'z' {
    x := x + d
    fmt.Printf("%c", x)
  } else if x == '-' {
    fmt.Print("=")
  }
}
```

- outputs `=HELLO=`

Notes:

- 4 declarations of `x`
- lexical blocks
  - outer
  - 3 syntactic blocks
    - `for` and `if` (x2)
  - 3 lexical blocks (non-syntactic)
    - `for` and `if` (x2) first lines
    - second `if` nested within first `if` so maybe another one (but empty)

----

### Scope Implications

- limiting scope: easier to read, possibly faster

```go
if f, err := os.Open(fname); err != nil {
  return err
}
// use f???  Nope it is no longer in scope
```

```go
if f, err := os.Open(fname); err != nil {
  return err
} else {
  // use f
}

```

- however in Go we prefer to keep normal execution path (non-errors) un-nested

----

```go
f, err := os.Open(fname)
if err != nil {
  return err
}
// use f
```

- scope of `err` is larger than we would like (ignored after `return err`)
  - un-avoidable with Go

----

### Short Variable Assignment Gotcha

- declares **local** variables

```go
var x, z int
func init() {
	x, y := 1, 2
	z = x + y
}
func main() {
	fmt.Println(x, z)
}
```

- prints `0 3` <!-- .element: class="fragment" -->
- `x` and `y` are declared by `x, y := 1, 2` in the synactic scope of `init()`
- `x` in `init()` shadowed `x` at package scope
- `x` at package scope is never assigned

---

### Switch (multi-way branch)

```go
// var breed string // set elsewhere
switch breed {
  case "beagle":
    hound++
  case "bloodhound":
    hound++
  case "poodle":
    poodle++
  case "australian shepherd":
    shepherd++
  case "aussiedoodle":
    poodle++
    shepherd++
  default:
    fmt.Printf("Unknown breed %q", breed)
}
```

----

```go
// var breed string // set elsewhere
if breed == "beagle" {
  hound++
} else if breed == "bloodhound" {
    hound++
} else if breed == "poodle" {
    poodle++
} else if breed == "australian shepard" {
    shepard++
} else if breed == "aussiedoodle": {
    poodle++
    shepard++
} else {
    fmt.Printf("Unknown breed %q", breed)
}
```

----

### Tagless Switch

- switch does not need an argument
- conditions evaluated in order (top down)

```go
// var weight float64 // set elsewhere
switch {
  case weight > 40:
    large++
  case weight > 20: // implicitly weight <= 40
    medium++
  default: // implicitly weight <= 20
    small++
}
```

----

- short variable assignment is allowed in `for`, `if`, and `switch`

```go
switch hour := time.Now().Hour(); { // missing expression
case hour < 12:
	fmt.Println("Good morning!")
case hour < 17:
	fmt.Println("Good afternoon!")
default:
	fmt.Println("Good evening!")
}
```

----

### Flow Control

- `switch` statement preferred over `if-else-if` chains
- `break` will break out of the **inner most** `for`, `switch`, or `select`
- `continue` skips the remaining body and starts next iteration in **inner most** `for` loop
- labels are used to `break` other than **inner most**

---

### Select

- channel (type `chan`): (mostly) FIFO message queue
  - safe to have multiple concurrent producers and consumers
  - receive blocks if no message
- `select` lets code wait on multiple channels
  - blocks until one case can be run
    - then executes that case
    - choose one at random if multiple ready
