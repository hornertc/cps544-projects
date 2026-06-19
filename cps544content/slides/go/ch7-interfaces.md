# Go Programming

## Interfaces

---

## Introduction

- Interface allows for generalizing behavior
  - Polymorphism
- Similar to other OOP languages except
  - **Go interfaces are satisfied implicitly**
- Interfaces are a reference type
  - act like pointers
  - almost never need a pointer to an interface

---

## Interfaces as Contracts

- Interfaces define a *weak* contract that a type must satisfy
- `interface{}` is the empty contract
  - `type any = interface{}`
- Interfaces are purely abstract/virtual
  - no implementation

```go
type Car interface {
    Straight(speed float64) float64
    Turn(radius float64) // same as Turn(float64)
}
func Drive(car Car) {
    car.Straight(10.5) // or car.Turn()
}
```

----

```go
type Mustang struct { /* fields omitted */ }
func (m *Mustang) Straight(speed float64) float64 {
    // implement it
}
func (m *Mustang) Turn(radius float64) {
    // implement it
}

type Civic struct { /* fields omitted */ }
func (c *Civic) Straight(speed float64) float64 {
    // implement it
}
func (m *Civic) Turn(r float64) {
    // implement it
}
```

----

```go
func main() {
    // Mustang does not implement Car but *Mustang does
    mustang := &Mustang{}
    Drive(mustang)
 
    civic := &Civic{}
    Drive(civic)
}
```

---

## Substitutability

- Freedom to substitute one type for another satisfying the same interface
- Key aspect of OOP
- Run-time property (not compile-time)
  - can pick type at runtime

```go
func chooseCar() Car { /* omitted */ }
func main() {
    car := chooseCar()
    // type of car is Car (an interface type)
    Drive(car) // performs "dynamic dispatch"
}
```

---

## Composability

- Can embed interfaces in an interface

```go
type Fetcher interface {
	Fetch(ctx context.Context, 
          target ocispec.Descriptor) (io.ReadCloser, error)
}

type Pusher interface {
	Push(ctx context.Context, 
         expected ocispec.Descriptor, 
         content io.Reader) error
}

type ReadOnlyStorage interface {
	Fetcher
	Exists(ctx context.Context, 
           target ocispec.Descriptor) (bool, error)
}

type Storage interface {
	ReadOnlyStorage
	Pusher
}
```

---

## Embedded Interfaces in Structs

- Can embed any type in a struct (including interfaces)

```go
type MyWriter struct {
    io.Writer
    myData int
}

// elsewhere
w := &MyWriter{
    Writer: os.Stdout
}
func (w *MyWriter) DumpResults() {
    fmt.Fprintln(w, w.myData)
}
data = /* []byte */
w.Write(data)
w.DumpResults()
```

---

## Interface Satisfaction

- Value of type `T` implements an interface `I` if it has **all** the methods of `I`
  - We say `T` is a `I`, by `var _ I = T(nil)`
- Value of type `T` does not necessarily posses all the methods of `*T`
  - even though `T` is a `I` then `*T` may not be a `I`
  - no syntactic sugar of automatic de-referencing (unlike with selectors)
- Non-pointer types may implement interfaces

----

## Conflicts

- On any given type `T`
  - method names must be unique
    - `func (t T) M()`
    - `func (t T) M(i int)` not allowed
  - implies no method overloading
  - cannot satisfy both `A` and `B`

```go
type A interface {
    M()
}
type B interface {
    M(int)
}
```

---

## Interface Values

- Interface is really two pointers
  - dynamic type
  - dynamic value
- Zero value is `nil`, `nil`
- `fmt.Printf("%T", value)` outputs the dynamic type when `value` is an interface

----

### Comparison of Interfaces

- Comparable if dynamic type is comparable
  - equal if
    - dynamics types are identical and
    - dynamic values are identical per `==`
  - panics if dynamic type is not comparable (e.g., slice)

Notes:

- Rarely need to compare interfaces
- Only compare interfaces when sure that dynamic type is comparable

----

### Gotcha

- `nil` interface value, `var i I`
  - zero initialized (`nil` dynamic type and value)
  - `i != nil` is false
- not same as `var i I = (*G)(nil)`
  - non-`nil` dynamic type `*G`
  - `nil` dynamic value
  - `i == nil` is **false**
    - dynamic types are not identical
  - thus `i != nil` is **true**!

Notes:

- `i == nil` can be interpreted as "is `i` equal to the interface's zero value (`nil`, `nil`)"
- Show `go run ./examples/interfaces`

---

## Use Cases

- [sort.Interface](https://pkg.go.dev/sort#example-package)
- [http.Handler](https://pkg.go.dev/net/http#Handler)
  - [http.HandlerFunc](https://pkg.go.dev/net/http#Handler)

Notes:

- Show how `http.ListenAndServe()` works

---

## Type Assertions

- Syntax `y := x.(T)`, assert that `x` is of type `T`, then assign value to `y`
  - type of `y` is `T` at compile time
- *Concrete* type `T`: if `x` is of type `T` then assign to `y` otherwise panic
- *Interface* type `T`: if `x` has dynamic type `T` then assign dynamic value to `y` otherwise panic
  - dynamic type of `y` is the same as dynamic type of `x` (unchanged)
- Use `y, ok := x.(T)` to distinguish a `nil` value from a `nil` interface and avoid panics

Notes:

- `go run examples/type_assertions`

----

## Type Switches

```go
func output(x any) {
	switch x := x.(type) {
	case int:
		fmt.Println(x + 3) // x is int
	case uint, uint8:
		fmt.Printf("%x", x) // x is any
	case string:
		fmt.Println(x) // x is string
	default:
		panic(fmt.Sprintf("unknown type %T: %v", x, x)) // x is any
	}
}
```

- discriminated union

---

## Recommendations

- Functions should accept interfaces and return concrete types
  - Make the function often more general and testable
  - Can mock out inputs
  - Returning concrete types does not artifically restrict the caller's use of the value

----

- Interfaces should be as small as possible
  - See `fs.FS`
- Parameter types for functions should be the minimal acceptable interface
  - Implementation can type assert to promote types to enable optional behavior
  - See `io.Copy()`

---

## Weak Contracts

- Implicitly satisfied
  - adding a method to a type might make the type implement another interface
  - type assertions can be used to change behavior based on available methods

----

- Only name and signature matters
  - parameter names need not match

```go
type I interface {
	F(dt float32)
}

type G struct{}
func (g *G) F(distance float32) {
	// do something with distance
}
var _ I = (*G)(nil)
```

----

- Documentation only
  - contractual requirements for the implementors is only described in comments
  - not enforced at runtime or testing time

- **Solution**: more mature libraries will include exported test helper to ensure that the contract is satisified
  - See [testing/fstest](https://pkg.go.dev/testing/fstest#TestFS)
  - See [testing/slogtest](https://pkg.go.dev/testing/slogtest#TestHandler)

---

## Interfaces Reimagined (Go 1.18 and Beyond)

- "< 1.18", interface is a set of methods
- ">= 1.18", interface is a set of types
  - completely backwards compatible with "set of methods"

---

## Type Systems

- *type safety* - should not be able to add a string to an integer
- *static typing* - type of variable/value is known at compile time
- *dynamic typing* - type of variable/value is only know at runtime
- *runtime type information* - information stored at runtime about a type (of a variable)
- *gradual typing* - types may be added later
- *duck typing* - if the type has the correct methods then it must be the correct type

Notes:

- duck test - "If it walks like a duck and it quacks like a duck, then it must be a duck"

----

## Go's Type System

- *Statically typed* with the ability to use interfaces for *dynamic typing*
- Each variable has a single type known at compile time
  - Interfaces store the type in the value not the variable
- *Runtime type information* available through the reflection API for all types
- Supports *duck typing* within static typing through interfaces

----

## POSIX Shell's Type System

- Every variable has the same type (string)
  - Integers are stored as strings `$((x + 1))`
- No *type safety* because one can add a string to an integer without error
