# Go Programming

## Functions

---

## Introduction

- function wraps a sequence of statements as a callable unit
- allows breaking code into parts:
  - smaller
  - testable
- hides implementation details and complexity

Notes:

- Walk through `code $GOPL/ch5` web crawler code

---

## Declarations

```go
func name(/*parameter-list*/) (/*result-list*/) {
    // body
}

func f(x, y float32, i *int) (w string, x1, x2 float32) {
    // body
}

// same type as above
func f2(x float32, y float32, i *int) (string, float32, float32) {
    // body
}
```

- returns can be named (just like parameters)
- if no or single unnamed return then parenthesis can/should be omitted

----

## Function Signature

- type of a function is called its signature
- names do not impact signature

```go
func f(x int, y float64) (i, j int) {
    i += x+1
    j = x+int(y)
    return
}

func g(z int, _ float64) (int, int) {
    return z+10, 14
}

fmt.Printf("%T\n", f) // func(int, float64) (int, int)
fmt.Printf("%T\n", g) // func(int, float64) (int, int)
```

---

## Parameters

- no default parameters
- no way to specify parameters by name
- parameters (and named results) are local variables within the body of the function
  - same lexical scope as function's outer most local variables
- passed by value

---

## Stacks

- each function call adds a *stack frame* to the stack
- local variables added to the stack or heap
  - depends on escape analysis
- each Goroutine has a variable stack size
  - starts at 2 KiB
  - grows to ~1 GiB
  - 100k+ Goroutines can be used thus initially small and dynamic stack sizes are important
- recursion is allowed
  - no *tail* recursion optimization

Notes:

- stacks in C default to about 1 MiB
- heap is shared by all Goroutines

---

## Error Handling

- conventionally errors are reported to caller by returning them as last argument

```go
func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %w", url, err)
	}
	return visit(nil, doc), nil
}
```

Notes:

- response body must be closed (as per the docs for `http.Get()`)
  - GC only frees memory
  - GC will **not** close file descriptors for you
- always handle error first
  - then continue with successful path un-indented
- URL added to error parsing error message
- `http.Get()` is nearly deprecated
  - prefer `http.NewRequestWithContext()`, make own `http.Client`, call `client.Do(req)`

----

## Defining Errors

- `error` (built-in) interface has one method `Error() string`

```go
var ErrBadness = errors.New("badness happened")

func f (x string) error {
    if strings.HasPrefix(x, "xyz") {
        // return ErrBadness // good
        return fmt.Errorf("%q missing prefix xyz: %w", x, ErrBadness) // best
    }
    _, err := os.Stat(x)
    // return err // good
    if err != nil {
      return fmt.Errorf("checking file: %w", err) // best
    }
    return nil // return nil when no error
}
```

Notes:

- IO errors from Go include the file path

----

## Wrapped Errors

- `fmt.Errorf("context %w", err)` wraps `err` and provides more context
  - `picking up food: example.com/food returned 404`
- wrapped errors have a method `Unwrap() error`
  - returns the cause of the error
- linked list pointing to the "cause" of the error

```go
// if err == ErrBadness {} // won't work
if errors.Is(err, ErrBadness) {
    // handle someError
}
```

----

## Error Data

```go
type ErrParse struct {
    Line, Column int
}

// Error implements error
func (e *ErrParse) Error() string {
    return fmt.Sprintf("%d:%d", e.Line, e.Column)
}

func main() {
    err := f()
    var errParse *ErrParse
    if errors.As(err, &errParse) {
        fmt.Println(errParse.Line)
    }
}
```

----

## Error Messages

- extracted with `err.Error()` as a string
- message strings should
  - not be capitalized
  - not contain newlines
- large chained/wrapped errors might be long
  - but easily `grep`-able
- unhandled errors are returns all the way up to main to be handled
  - often written to STDERR and exiting the program with a non-zero exit code

----

## Error vs Bool

- return `bool` such as `val, ok := obj.Get("food")` to indicate existence
- return `(*Value, error)` if an error can occur
  - non-`nil` error means something bad happened
  - `nil` error and non-`nil` value means "found" and here it is
  - `nil` error and `nil` value means "not found"

----

## Error Handling Strategies

- (a) return the error to the caller
  - first add context with `fmt.Errorf()`
- (b) handle the error
  - retry a network operation
  - ignore if an optimization or cleanup (e.g., cache, rm temp dir)
- (c) `panic(err)` will dump stack trace
  - this is recoverable
- (d) print the error to STDERR and call `os.Exit(1)`
  - not recoverable

Notes:

- show `$GOPL/ch5/wait/wait.go`

----

## Handling End of File (EOF)

```go
func readHeader(r io.Reader) (*Header, error) {
  buf := make([]byte, 48)
  _, err := r.Read(buf)
  if errors.Is(err, io.EOF) {
    return nil, nil
  }
  if err != nil {
    return nil, fmt.Errorf("reading header: %w", err)
  }
  return &Header{
    from: buf[:24],
    to: buf[24:]
  }, nil
}
```

----

## Multiple Errors

- concurrency can cause multiple errors
- options
  - return just first/one/any error
  - return all errors

```go
//func errors.Join(errs ...error) error

return errors.Join(listOfErrors...)
```

---

## Function Values

- functions are first-class values; like other values:
  - function values have types
  - may be assigned to variables
  - may be passed to/returned from functions
- function value may be called like any other function
- zero value is `nil`
- not comparable
- reference type

----

```go
func transform(x []float64, f func (float64) float64) {
  for i := range x {
    x[i] = f(x[i])
  }
}

func square(y float64) float64 { return x*x }
func double(y float64) float64 { return 2*x }

func main() {
  // omitting error check
  var f func(float64) float64 // signature
  switch os.Args[1] {
    case "d":
      f = double
    case "sq":
      f = square
  }
  if f == nil {
    f = double
  }
  // get x
  transform(x, f)
  // do something with x
}
```

----

## Benefits of Function Values

- allow functions to be parameterized over data and behavior
- `strings.Map` applies a function over each rune of a string

----

## Object Factory Pattern

```go
type Action func (block *Block) error
var actions = map[string]Action {
  "forward": forwardAction,
  "right": func (b *Block) error {
    b.Orientation -= math.Pi/2
    return nil
  }
}

func createAction(name string) Action { return actions[name] }
func addAction(name string, action Action) { actions[name] = action }
```

----

## Anonymous Functions

- function without a name (anonymous)
- function literal
- defined inline as an expression

```go
x := func (x int) bool {
  // definition
}

fmt.Println(x(42))
```

----

## Returning Functions

- function generators

```go
func counter(x int) func () int {
  return func () int { // closure
    x++
    return x - 1
  }
}

count := counter(10)
fmt.Println(count()) // 10
fmt.Println(count()) // 11
```

----

## Example: Topological Sorting

See `code $GOPL/ch5/toposort/main.go`

---

## Variadic Functions

- function's *arity* is number of arguments
- *variadic* means variable number of arguments
- well-known example `fmt.Printf()`

```go
func sum(vals ...int) (total int) {
  for _, val := range vals { // vals is a []int
    total += val
  }
  return
}

// sum(3, 4, 5) is 12
// sum(1) is 1
// sum() is 0
```

----

- compiler implicitly
  - allocates array in caller
  - copies the arguments into it
  - passes a slice of the entire array to function

- only one variadic argument type allowed
  - `...` can only occur once
- pass a slice directly to a variadic argument

```go
values := []int{3, 4, 5}
sum(values...) // 12, does not make a copy of values
```

Notes:

- [Example](https://go.dev/play/p/20JJg3Mq8j) showing that a copy is not made of `values`

----

- argument `...int` not same as `[]int`
- types `func(...int)` is distinct from `func([]int)`
  - eventhough within the body they are identical (both slices)

- `fmt.Printf(format string, ...any)`
  - `any` is an alias for `interface{}`
  - represents any value (number, any struct, function)

---

## Deferred Function Calls

- syntactically: defer statement is ordinary function or
method call prefixed by the keyword `defer`
- function and argument expressions are evaluated when
statement is executed
- however, actual call is deferred until function that contains
the defer statement has finished either
  - normally (executing a return statement or falling off the end) or
  - abnormally (by panicking)

----

- any number of calls may be deferred
  - they are executed in the reverse of the order in which they were
deferred
- often used to ensure resources are cleaned up
  - file/socket is closed
- typical flow
  1. acquire resource
  1. check error
  1. defer release of resource
- Go does not have destructors

----

```go
func bigSlowOperation() {
  defer trace("bigSlowOperation")() // don't forget the extra parentheses
  // f := trace("bigSlowOperation")
  // defer f() // equivalent to above

  // ...lots of work...
}

func trace(msg string) func() {
  start := time.Now()
  log.Printf("enter %s", msg)
  return func() {
    log.Printf("exit %s (%s)", msg, time.Since(start))
  }
}
```

----

- deferred functions run after return statements have
updated function’s result variables
- anonymous function can access its enclosing
function’s variables including named results
- deferred anonymous function can observe function’s
results (and change it)

----

## Never Defer In a Loop

```go
func processFiles(filenames []string) error {
  for _, filename := range filenames {
    f, err := os.Open(filename)
    if err != nil {
      return err
    }
    defer f.Close()
    // process f
  }
  return nil
} // f1000.Close(), f999.Close(), ..., f1.Close()
```

- could run out of file descriptors
- bad use of scarce kernel resources

----

- instead move the loop into a function

```go
func processFile(filename string) error {
  f, err := os.Open(filename)
  if err != nil {
    return err
  }
  defer f.Close()
  // process f
  return nil
} // f.Close()

func processFiles(filenames []string) error {
  for _, filename := range filenames {
    if err := processFile(filename); err != nil {
      return err
    }
  }
  return nil
} 
```

---

## Panic

- Go tries to check errors at compile time
- Those not checked at compile time are checked at run time
  - index out of bounds
  - nil pointer dereference
- Go runtime will issue a panic for run time errors

----

- during typical panic:
  - normal execution stops
  - all deferred function calls in that goroutine are executed
  - panic works its way up the stack and if not recovered
  - program crashes with a log message
  - log message includes panic value
- logs error message and a stack trace (for each
goroutine) showing stack of function calls active at the time of the
panic

----

## Errors vs Panic

- panics resemble exceptions in other languages
- `panic()` for violated assumptions
  - precondition: argument `x int` must be between 5 and 10
- **panic() should never be called in a healthy/bug-free program**
- return `error` for conditions that may occur in a healthy/bug-free program

----

- `panic()` is automatically called when:
  - dereferencing a nil pointer
  - index out of bounds
- can call `panic(msg)` manually and add context
  - allows the user of the program to see the violation
  - avoid calling it for automatically detected run time errors

----

## Recovering from a Panic

- often program termination is the correct response to a panic
- `recover()` can be called in a deferred function to recover from a panic
  - similar to catching an exception

```go
func Parse(input string) (s *Syntax, err error) {
  defer func() {
    if p := recover(); p != nil {
      err = fmt.Errorf("internal error: %w", p
    }
  }()
  // do parsing but code might do index out of bounds access
  s, err = parse(input)
  return
}
```

----

- recovering indiscriminately from panics is a dubious practice
- state of a package’s variables after a panic is rarely well
defined or documented
- for example:
  - critical update to data structure was incomplete
  - file or network connection was opened but not closed
  - lock was acquired but not released

----

- program crashes indicate bugs very clearly
- indiscriminate log messages tend to hide bugs
- while a crashing program indicates a bug, often better than a resilient program hiding the bug

Notes:

- `code $GOPL/ch5/title3` show panic recovery with a un-exported type

---

## Options Pattern

- Go has no default or optional function arguments
- Use options design pattern instead
- Pass a variadic number of `Option` to a function
- Within function apply each option in turn
- `Option` is often a simple function

Notes:

- See [example](../../examples/options/options.go)
