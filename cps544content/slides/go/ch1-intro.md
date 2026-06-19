# The Go Programming Language

## Chapter 1

---

### Origins of Go

- “C-like language”
- “C for the 21st century”
- Created by Robert Griesemer, Rob Pike, and Ken Thompson
- Inherited from C:
  - Expression syntax
  - Control-flow statements
  - Basic data types
  - Call-by-value parameter passing - pointers
  - Programs that
    - Compile to efficient machine code
    - Cooperate naturally with abstractions of current OSs

---

<!-- ![Go ancestors](../img/go-ancestors.png) -->
<img src="../img/go-ancestors.png" alt="Go ancestors" width="500"/>

Notes:

- Alef was developed for use in Plan9 (designed by Rob Pike, Ken Thompson, creators of Go)

---

### Go Guiding Design Principles

- Easy to understand language
  - This course will cover the **entire** language
  - Easy to read (intuitive syntax)
- Strong concurrency
- Fast compilation
  - Aggressive caching
  - Multi-threaded
  - Complete encapsulation

> Go was created while waiting for large server builds - Rob Pike

---

### Hello World

Save the following to `hello.go`

```Go
package main
import "fmt"
func main() {
  fmt.Println("Hello, World")
}
```

- Go is a compiled language
- Go toolchain is available through the `go` program
- Compile, link, and run the example with `go run hello.go`

Notes:

`go run` compiles so fast it feels like an interpretive language but without the REPL.

----

```Go [1|2|3|4]
package main
import "fmt"
func main() {
  fmt.Println("Hello, World")
}
```

- Go code is organized into packages
  - Similar to libraries or modules in other languages
- Package consists of one or more `.go` source files in a single directory that define what the package does
- Each source file begins with a package declaration
  - States to which package the file belongs

----

### `main` Package

- Package `main` is special
- Defines a stand-alone executable program, not a library
- Within package main the function `main` is also special
  - It’s where execution of the program begins
  - `main` will normally call upon functions in other packages to do much of the work, e.g., `fmt.Println()`
  - Program completes when `main()` completes

----

### Importing Packages

- `import` declaration (always follows the package declaration)
tells compiler what packages are needed by source file
- import declaration must import exactly the packages needed; program will not compile
  - if there are missing imports or
  - if there are unnecessary imports
    - Prevents references to unused packages from accumulating as programs evolve

----

### After importing

- After the import statement, program consists of declarations of:
  - functions (keyword `func`)
  - variables (keyword `var`)
  - constants (keyword `const`)
  - types (keyword `type`)
- order of declarations does not matter

----

### Code Formatting

- Go takes a strong or opinionated stance on code formatting
- `gofmt` tool rewrites code into the standard format
- `goimports` is an improved tool that formats code and sorts imports
- Neither is part of the standard distribution, but both can be obtained using the command `go install golang.org/x/tools/cmd/goimports@latest`

----

### Command-Line Arguments

- One mechanism of program input
- Variable `os.Args` (in `os` package)
  - slice of strings, type is `[]string`
  - `os.Args[0]` - program name
  - Rest are arguments as separated by the calling program (often a shell)

Notes:

- Other ways to get input into a program:
  - environment variables
  - well-known files
  - network sockets
  - signals
  - shared memory

----

- Example `best a wonderful life --movie`
  - `best` program name
  - `os.Args` would be `[]string{"best", "a", "wonderful", "life", "--movie"}`
  - options/flags must be interpreted by the program (not the shell)

- Example `best "a wonderful life" --movie`
  - `os.Args` would be `[]string{"best", "a wonderful life", "--movie"}`
  - quotes force movie names into one argument

---

### Slices Overview

- Slice - dynamically sized sequence of array elements
  - individual elements can be accessed as `s[i]`
  - contiguous subsequence as `s[m:n]`
- Number of elements: `len(s)`
- All indexing in Go uses half-open intervals
  - include the first index but exclude the last
  - slice `s[m:n]`, where 0 ≤ m ≤ n ≤ len(s), contains n-m elements
  - `[4,9)` in mathematical notation

---

### Echo (v1)

```Go
package main

import (
  "fmt"
  "os"
)

func main() {
  var s, sep string
  for i := 1; i < len(os.Args); i++ {
    s += sep + os.Args[i]
    sep = " "
  }
  fmt.Println(s)
}
```

----

### Echo (v2)

```Go
package main

import (
  "fmt"
  "os"
)

func main() {
  s, sep := "", ""
  for _, arg := range os.Args[1:] {
    s += sep + arg
    sep = " "
  }
  fmt.Println(s)
}
```

Notes:

- variable declarations
- "for each" loop
- string concatenation may be costly

----

### Echo (v3)

```Go
package main

import (
  "fmt"
  "os"
  "strings"
)

//!+
func main() {
  fmt.Println(strings.Join(os.Args[1:], " "))
}
```

Notes:

- Avoids expensive string concatenation
- Show implementation of `strings.Join()`

---

### `for` loop

- `for` loop is the only loop statement
- it has a number of forms, one of which is illustrated here:

```Go
for initialization; condition; post {
  // zero or more statements
}
```

----

- optional initialization statement is executed before the loop starts
  - If present, must be simple statement:
    - short variable declaration
    - Incrementer
    - assignment statement
- condition is a boolean expression evaluated at the beginning of each iteration of the loop
- if it evaluates to true, the statements controlled by the loop are executed
- post statement is executed after the body of the loop, then the condition is
evaluated again
- loop ends when condition becomes false

Notes:

- parentheses are never used around the three components of a for loop
- braces are mandatory, and opening brace must be on the same line as the post statement

----

- Any of these parts may be omitted
- If there is no initialization and no post, the semicolons may also be omitted:

```Go
// a traditional "while" loop 
for condition {
// ... }
```

- If the condition is omitted entirely: `for { // ... }` the loop is infinite
  - loop may be terminated with a `break` or `return` statement

---

### `for each` loop

```Go
for index, value := range mySlice {
  // use index and value
}
```

- iterates over a range of values from a few data types like a string or slice
- each iteration, `range` produces pair of values:
  - index
  - value of the element at that index
- If index is not used replace with `_`
- If value is not used omit

Notes:

- Just an index is useful for looping from `0` to `len(mySlice) - 1`
- The 'for each' loop produces a pair of values for each iteration: the index and the value of the element at that index.
- If the index is not used, it can be replaced with `_` to avoid creating unnecessary variables.

---

### Increment/Decrement

- increment statement `i++` adds 1 to `i`
- corresponding decrement statement `i--` that subtracts 1 from `i`
- these are statements (do not represent a value), not expressions (represent a value)
  - so, `j = i++` is illegal
- also, they are postfix only
  - so `--i` is illegal

Notes:

- In C++ `i++` is an expression (returns a value).  Also `++i` is valid and returns a different value.
- Go is limited to supporting postfix increment and decrement ('i++' and 'i--').

---

### Variable Declarations

- Four syntaxes
  - `s := ""` (short variable declaration)
    - most compact
    - can be used only within a function
  - `var s string`
    - default initializations to zero value. ("") for strings
  - `var s = ""`
    - rarely used except when declaring multiple variables

----

- Four syntaxes (cont.)
  - `var s string = ""`
    - explicit variable type
    - redundant when type is same as that of initial value
    - necessary when the type is not the same
- In practice, one of the first two forms should generally be used, with:
  - explicit initialization to say that the initial value is important
  - implicit/default initialization to say that the initial value doesn’t matter

---

### Input Example

```go
var name string
fmt.Print("Enter a name: ")
fmt.Scan(&name) // ignoring error

var count int
fmt.Print("Enter an integer: ")
fmt.Scan(&count) // ignoring error
```

Notes:

- `go run ./examples/input`

---

### Finding Duplicates (Example)

```go
import "bufio"

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 { // Duplicate
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
```

---

### Other Examples

- `dup1`
- `server1`
- `server2`

Notes:

- `go run $GOPL/ch1/dup1/main.go`
