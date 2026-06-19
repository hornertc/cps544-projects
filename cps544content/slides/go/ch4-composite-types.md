# Go Programming

## Composite Types

---

## Overview

- arrays
- slices
- maps
- structs

----

## Types

- arrays: fixed size, homogeneous
  - all elements have same type
- structs: fixed size, heterogenous
  - elements of possibly different types
- slices:
  - dynamically sized arrays
  - homogeneous
- maps
  - dynamic fields (like a dynamic struct)
  - heterogeneous

---

## Arrays

- fixed length sequence of zero or more elements of a particular type
- rarely used directly due to fixed length
- slices much more versatile
  - can grow and shrink

----

- Element access `myarray[5]`, `myarray[i]`
- `len(myarray)` returns length
- `cap(myarray)`, same as `len(myarray)`
- ellipses `...` can be used to avoid writing the size in initializers
- size of array is part of type `[5]int`

```go
var x3 [3]float32 = [3]float32{3.14, 2.7, 1.41}
y3 := [...]float32{3.14, 2.7, 1.41} // without explicit size
x3 = [4]float32{1, 2, 43, 4.5} // compilation error
```

----

- literal syntax is similar for arrays, slices, maps, and structs
  - this is intentional
  - but can be confusing
- allowed to specify index constants like maps

```go
type Dimension int

const (
    X Dimension = iota
    Y
    Z
)

vector1 := [...]float64{X: 45.2, Y: 0, Z: 92} // map-like syntax
vector2 := [...]float64{45.2, 0, 92}          // equivalent
vector3 := [...]float64{Z: 92, X: 45.2}       // equivalent
```

----

- indices can appear in any order
- may omit indices
  - default to zero value

```go
foo := [...]int{9: 2}
len(foo) // 10
cap(foo) // 10
foo[0]   // 0
foo[9]   // 2

bar := [20]int{9: 2}
len(bar) // 20
cap(bar) // 20
bar[0]   // 0
bar[9]   // 2
```

----

## Comparable

- can compare `==` and `!=` arrays of the same type (and size)
  - said to be "Comparable" thus usable as keys in maps
- not orderable
  - `<` not valid for arrays

----

## Copying

- passing array into a function (or returning an array)
  - copies the entire array
  - not a reference type
  - can pass a pointer to an array

```go
func modifyArray(ptr *[8]int) {
    // notice that (*ptr)[0] is not needed
    ptr[0] = ptr[1] + ptr[2]
}
```

Notes:

- Differs from other languages where arrays are just pointers
  - copy is cheap

---

## Slices

- variable length sequence of elements of same type
- `[]T` where `T` is element type
  - like array but no size
- slice is lightweight data structure
  - gives access to some or all of the elements of an array

----

- components
  - pointer `*T`
    - points to element of an array
    - value is the first element of slice
  - length `int`
    - number of elements
    - cannot exceed capacity
  - capacity `int`
    - number of elements between the start of slice and end of array
    - but may be shorter than end of array

----

```go
// T can be any type
type Slice[T any] struct {
    data *T
    len, cap int
}
```

- `len(myslice)` - length
- `cap(myslice)` - capacity

```go
months := [...]string{1: "January", /* omitted */, 12: "December"}
```

----

![slices of months](../img/slices-months.png)

----

## Extending Slices

- slicing beyond cap(s) causes panic
- slicing beyond len(s) extends the slice
- result may have larger length than original

```go
months := [...]string{1: "January", /* omitted */, 12: "December"}
summer := months[6:9] // len = 3, cap = 7
endlessSummer := summer[:5] // len = 5, cap = 7
```

----

- `x[m:n]` yields a `string` if `x` is a `string`
- `x[m:n]` yields `[]byte` if `x` is a `[]byte`
- `x[m:n]` yields `[]byte` if `x` is a `[N]byte` (`N` constant)
  - recall `[N]byte` is the type of an array
  - slicing an array yields a slice (not a new array)
  - `x[:]` yields slice of the entire array `x`
- copying slice creates alias for underlying array
  - reference type

----

## Shrinking Slices

- `x[m:n:k]`
  - `m` - starting index (default 0)
  - `n` - ending index, exclusive, default `len(x)`
  - `k` - max index, exclusive, default `cap(x)`
- data pointer is `&x[m]`
- length is `n-m`
- capacity is `k-m`
  - allows one to limit the capacity so that when passed to functions like `append` you can protect elements in the array (but not in the slice)

Notes:

- Run `go run ./examples/slices/main.go`

----

## Comparison

- unlike arrays, slices are not comparable
  - cannot use `==` with two slices
  - cannot use slices as keys in maps
- `bytes.Equal()` can be used to compare `[]byte`
- Comparing slice to nil is valid
  - `x == nil` and `x != nil`

----

### Why can't slices be compared using `==` operator?

1. unlike array elements, elements of slice are indirect
it's possible for a slice to contain itself
2. because slice elements are indirect, fixed slice value
may contain different elements at different times
since contents of underlying array are modified

Notes:

- these are weak reasons but reasons nonetheless

----

## `nil` Slice

- zero value of a slice type is `nil`
- `nil` slice has no underlying array
- `nil` slice has length and capacity zero
- there are also non-nil slices of length and capacity
zero
  - such as `[]int{}` or `make([]int, 3)[3:]`
- to test whether slice is *empty*:
  - use `len(s) == 0`, not `s == nil`
- generally Go treats `nil` slices identically to zero-length slices

----

## Making Slices

- built-in function `make(type, len, cap)` creates slice of a specified element type, length, and capacity

```go
x := make([]T, 5)      // len=5, cap=5
x := make([]T, 5, 10)  // len=5, cap=10
x := make([]T, 10)[:5] // len=5, cap=10, same as above
// With compile time capacity we do not need make()
a := [10]int{}; x := a[:5] // len=5, cap=5, same as above
```

----

## Copy

- `copy(dst, src)` copies slices `src` into `dst`
- handles overlap properly
- very efficient
- works for all slice types (generic function)

----

## Append

- built-in `append()` function appends items to slices

```go
str := "string with +/- and crazy symbols +✔"
var runes []rune
// runes := make([]rune, 0, len(str)) // preallocation
for _, r := range str {
    if r != '+' {
        runes = append(runes, r)
    }
}
```

- if just converting, not filtering, string to `[]rune` can do `[]rune(str)`

----

- `append()` may use a more sophisticated growth strategy
  - might reallocate array before all capacity is used
  - might grow my more/less a factor of 2
- nearly always of the form `x = append(x, ...)`
- updating slice variable (as above) required when
  - change length or capacity of slice
  - refer to different underlying array
- slices are *reference types* however pointer, length, and capacity are passed by **value**

Notes:

- before slide
  - `appendInt()` and `appendSlice()` from $GOPL/ch4/append/main.go
- after slide
  - show `prepend()` [source code](../../examples/prepend/main.go)

----

## In-Place Techniques

- avoid memory allocation

```go
// no need to return slice
// since the slice itself it not modified (only the data)
func incAll(ints []int) {
    for i := range ints {
        ints[i] += 1
    }
}
```

```go
func incAllBroken(ints []int) {
    // ints is actually un-modified
    for i, x := range ints {
        x += 1
    }
}
```

----

## In-Place (fancy) Techniques

```go
// potentially modifies strings 
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
```

----

```go
func nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
```

---

## Maps

- maps implemented as a hash table
  - one of the most ingenious and versatile of all data structures
- unordered collection of key/value pairs
- all the keys are distinct
- value associated with given key can be retrieved, updated,
or removed
- (amortized) constant-time, $\mathcal{O}(1)$ for key lookup
  - no matter how large the hash table
- compare to binary trees: $\mathcal{O}(\log_2 N)$ lookup

----

- `map[K]V`
  - `K` is the Comparable (`k1 == k2` defined) key type
  - `V` is any value type
- hash function needed for keys
  - Go defines a reasonable hash function for you
  - cannot be overwritten

----

## Empty Maps

```go
x := make(map[string]int)     // empty map
x := map[string]int{}         // equivalent
x := make(map[string]int, 50) // with capacity hint
len(x) // 0
cap(x) // not defined for maps
```

- zero-value of map is `nil`
- length of nil map is zero

----

## Initialized Maps

```go
price := map[string]float32{
    "hamburger":     3.50,
    "cheese burger": 4.00,
    "vegetables":     0.50,
}
prices2 := make(map[string]float32)
prices2["hamburger"] =    3.50
prices2["cheese burger"] = 4.00
prices2["vegetables"] =   0.50
```

----

- map access is `prices["hamburger"]`
  - zero-value returned if key not in map
  - `prices["tofu"] += 2.50`
- deletion of a key/value pair
  - `delete(prices, "vegetables")`
  - no-op if key not in map

----

## Map Element ≠ Variable

- cannot take address of map value
  - `&prices["vegetables"]`
- why? insertion into map might cause rehashing of existing elements into new
storage locations
  - potentially invalidating address
- if wanting address of map value then
  - apply "Fundamental Theorem of Software Engineering"
  - `map[string]*float32`

----

## Iterating

```go
for name, price := range prices {
    fmt.Printf("%s=%g", name, price)
}
```

- "unordered map" - random order each time
  - forces developers to robust across Go implementations
- if deterministic order is required
  - extract keys of map into a slice, sort, then iterate on slice

----

## `nil` maps

- most operations work as expected
  - lookup (return zero-value)
  - delete (no-op)
  - len (return 0)
  - range loops (nothing in map)
- **however** assigning to a nil map will panic
  - need to make the map first

```go
var x map[string]int
x["something"] = 45    // panics

y := map[string]int{} // note the curly braces
y["something"] = 45    // OK
```

----

## Checking Keys

- subscripting always yields value
  - if key present, get corresponding value
  - if not, get zero value for element type
- sometimes can test value against zero-value
  - does not work if the zero-value is a valid value in the map
    - often the case for numeric values

```go
price, ok := prices["hamburger"]
if ok {
  fmt.Println("hamburgers are in stock!", price)
} else {
  fmt.Println(`hamburgers are out of stock
    time to butcher more cows.`)
}
```

----

## Comparison of Maps

- only valid comparison is with `nil`
  - not another map
- but we can write a comparison

```go
func equal(x, y map[string]int) bool {
  if len(x) != len(y) { // short circuit
    return false
  }
  for k, xv := range x {
    // if y[k] != xv { // Does not work
    if yv, ok := y[k]; !ok || yv != xv {
      return false
    }
  }
  return true
}
```

----

## Sets in Go

- no set type in Go
- use a `map[K]struct{}`
  - `struct{}` is an empty structure (zero-length)
  - constructed with `struct{}{}`
  - alternatively could use `map[K]bool`

Notes:

- Show dedup code `$GOPL/ch4/dedup/main.go`

----

## Comparability Revisited

- comparable types
  - numeric types (int, float32, ...), booleans
  - strings
  - arrays
  - pointers, channels
  - array (if element is comparable)
  - structs (if all fields are comparable)
- (always) non-comparable types
  - slices, maps
  - functions

----

## Orderable

- orderable if `<`, `<=`, `>`, `>=` are defined
- orderable types
  - float, int
  - strings
- not orderable (incomplete list)
  - al non-comparable types
  - bool
  - complex
  - array, slice
  - channel
  - structs

Notes:

- `sort.Interface` can be used to define an ordering for non-natively orderable types (e.g., structs)
- then can sort with then wit `sort.Sort()`

----

## Non-comparable Key Types

- trick to make them comparable
  - define `func k(key nonComparable) comparable`
  - define map `map[comparable]valueType`
  - insert like so `m[k(key)] = value`
- function must ensure distinctiveness of keys
  - often just remove non-comparable fields

Notes:

- show `$GOPL/ch4/charcount/main.go`
- show `$GOPL/ch4/graph/main.go`

---

## Structures

- named fields with arbitrary types
  - copied as a unit
  - passed to functions
  - returned from functions
  - stored in arrays

```go
type PizzaOrder struct {
  Size int // inches
  FirstName, LastName string
  Toppings []string
  Price float32
  Delivery *Address
}
```

----

- structs initialization
  - `MyStructType{Field: value, ...}`
  - `MyStructType{value1, value2, ...}`
    - must include *all* fields

```go
order := PizzaOrder{
    FirstName: "John",
    LastName: "Doe",
    Toppings: []string{"pepperoni", "sausage"},
    Price: 13.75,
    Size: 14,
    // no delivery
}
order2 := PizzaOrder{14, "John", "Doe", 
  []string{"pepperoni", "sausage"}, 13.75, nil} // equivalent
```

----

- fields are variables
  - can take address

```go
fmt.Printf("%s %s ordered a %d inch pizza", 
  order.FirstName, order.LastName, order.Size)
if order.Delivery != nil {
    fmt.Println("Street", order.Delivery.Street)
}
superSizeMe(&order.Size)
```

----

- exported fields begin with capital letter

```go
package pizzashop

type ShadyPizzaOrder struct {
    Size int // inches
    FirstName, LastName string
    Toppings []string
    Price float32
    Delivery *Address
    upSellTarget bool // visible only inside package pizzashop
    prettyShoes  bool // visible only inside package pizzashop
}
```

----

## Self-Referential Structures

- struct `S` cannot declare field of same type `S`
  - aggregate cannot contain itself
  - "Fundamental Theorem of Software Engineering"
    - Use `*S` as field type (fixed size)
- compiler must be able to determine structure size

```go
type S struct {
    Value float64
    Child *S
}
```

Notes:

- `$GOPL/ch4/treesort/main.go`

----

## Zero-Value

- zero-value of each field type
- desirable that zero-value be sensible/natural/secure default
  - same as raw variables

----

## Struct Embedding

- embed one struct inside another
- called an *anonymous field*
  - has no name, only type
  - accessed by type name

```go
type Automobile struct {
    VIN string
    GVW float32
}
type Motorcycle struct {
    Automobile
    LengthOfHandleBars float32
}
cycle := Motorcycle{}
cycle.VIN = "123456-ABC"
cycle.Automobile.VIN = "123456-ABC" // equivalent
```

Notes:

- Embedded structs are similar to OOP "is a" relationship

----

## Initializing Embedded Fields

```go
cycle := Motorcycle {
    Automobile: Automobile {
        VIN: "123",
        GVW: 6230.0,
    },
    LengthOfHandleBars: 13.4,
}
```

Notes:

- run `$GOPL/ch4/embed/main.go`
  - show `%#v` is Go-syntax

----

## Passing Structures

```go
type F struct {a, b int}
func g1(f F) string { // f is copied
    return fmt.Sprintln(f.a)
}
func g2(fp *F) string { // pointer to f copied
    return fmt.Sprintln(fp.a)
    // return fmt.Sprintln((*fp).a) // (*fp) not needed
}
x := &F{10, 11} // x is a *F
g2(x)
```
