# Generics

Added in Go 1.18 (not in the textbook)

## Overview

- Motivation
- Examples

---

## Motivation

- Static typing is wonderful
  - Fast machine code
  - Compile time errors
- Want to be able to define generic algorithms that accept many different types
  - sort, reverse, find, delete

----

- Go supports dynamic typing with interfaces, so what is the problem?
  - Runtime cost
    - Optimizations cannot be made through the virtual dispatch
  - Limited static analysis
- Sometimes, the type is known at compile time, so why not take advantage of that for extra compile time checking and performance

----

- Generics (called templates in other languages) provide compile time polymorphism
  - Interfaces provide runtime polymorphism
- Generics functions/types are compiled by a process called **monomorphization**
  - **Monomorphization** is the process of turning generic code into specific code by filling in the concrete types that are used when compiled
- Compiler writes an optimized procedure for each unique instantiation of the generic function/type
  - Allows compiler optimizations based on the type
  - Bloats the binary slightly

---

```go
// F is a function that takes an argument of type T
// where T has an underlying type of int
func F[T ~int] f(p T) { … }

// T has an underlying type of int or string or is comparable
func F[T ~int | ~string | comparable] g(p T) { … }

type Cache[T any] struct {  
    m sync.Map
}

func (c *Cache[T]) Set(key string, value T) {
    c.m.Store(key, item[T]{ 
        value: value,   
    })
}
```

---
---

## Examples

- Standard library
  - [slices](https://pkg.go.dev/slices#Delete)
  - [maps](https://pkg.go.dev/maps@go1.21.4#Equal)

Notes:

- Compare `slices.Sort` to `sort.Sort`
- TODO build out a memoization example

---

## Resources

- A nice guide is [here](https://itnext.io/a-comprehensive-guide-to-generics-in-go-5a9dcda5669c)
