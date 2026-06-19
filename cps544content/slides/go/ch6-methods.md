# Go Programming

## Methods

---

## Introduction

- Methods defined by adding an extra parameter before the function name
  - `func (obj Object) Move(x, y float32)`
  - `obj` is called the *receiver*
- Must be declared at package level scope
- `Object` may be *any* type declared in *current* package
- All rules of functions apply
  - receiver passed by value
  - export rules

Notes:

- Cannot add methods to types in another package
- Can "inherit" methods from other types and "extend" them
  - make a new type in current package
  - `type Type other.Type`

---

```go
type Meal []string

func (f Meal) Healthy() bool {
	for _, food := range f {
		switch food {
		case "carrot", "broccoli":
			return true
		}
	}
	return false
}

func main() {
	lunch := Meal([]string{"candy", "carrot"})
	fmt.Println(lunch.Healthy()) // true
}
```

---

## Selectors

- `lunch.Healthy` is a *selector*
- selects the appropriate method of receiver `lunch`
- another type in same package might define a method called `Healthy()` but that will not conflict with `Meal`'s `Healthy` method
- Method names tend to be shorter
  - `Distance` instead of `DistanceBetweenPoints`

---

## Methods With Pointer Receivers

- `func (obj *Object) Mutate(x int)`
- `Mutate()` can update variables/fields of `obj`
- Convention: all or no methods should be pointer receivers
  - even if some methods do not change the receiver

```go
o := Object{}
op := &o
op.Mutate(4)
o.Mutate(4) // short hand for (&o).Mutate(4)
op.NonMutate() // short hand for (*op).NonMutate()
```

----

## `nil` is Valid Receiver Value

- `obj` can be `nil`
- only the type (at compile time) is used to resolve method selector to function
- can be useful to represent object state
- rarely utilized

---

## Composing Types

- Methods from embedded fields are "promoted" to enclosing type
  - Can be overwritten
  - Identical to how fields of embedded structs are handled
- Powerful way to compose *behavior*
  - "has a" not "is a"

----

## Selectors Revisited

- *selector* is resolved to a function by looking for methods of:
  1. the receiver
  2. any embedded field
  3. any embedded field's, embedded field
  4. and so on...

- field access `obj.X` also a selector (selects the appropriate field from embedded structs)
- Selectors for fields are resolved in the same way as methods

Notes:

- Selector resolution is done at compile-time (not-runtime)

---

## Method Values

- `lunch.Healthy` is a *method value*
  - signature is `func () bool`
- can be used detached from the receiver as a regular function
  - first class functions

```go
meal := Meal([]string{/*...*/})
mv := meal.Healthy // method value
mv()
```

---

## Method Expressions

- `Meal.Healthy` is a *method expression*
  - signature is `func(Meal) bool`
- can be used as a function but need the receiver as the first argument
  - first class functions

```go
meal := Meal([]string{/*...*/})
me := Meal.Healthy // method expression
me(meal)
```

Notes:

- Show `code $GOPL/ch6/coloredpoint/main.go`

---

## Encapsulation

- Information hiding
- Only one mechanism in Go
  - exported and un-exported fields, methods, types
  - no private, protected, public, friends
- Unit of encapsulation is package
  - not type
- *Getter* and *setter* methods
  - can expose un-exported fields

---

## Method Naming

- Prefer shortest name
- Setters are named `SetYYY`
  - `transform.SetScale(13.5)`
- Getters are named `YYY`
  - `transform.Scale()` returns the scale
  - `duration.Seconds()` returns seconds
