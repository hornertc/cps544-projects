# Go Programming

## Basic Data Types

---

## Overview

- basic types
  - numbers
  - strings
  - booleans
- aggregate types: to form more complicated data types by combining values of several simpler ones
  - arrays
  - structs

----

- reference types: refer to program variables or state indirectly
  - pointers
  - slices
  - maps
  - functions
  - channels
- Interface types

Notes:

- reference types effect of an operation applied to one reference is observed by all copies of that reference

---

## Integers

- `int8`, `int16`, `int32`, `int64` are signed types
  - `int` is *native* sized integer (either 32 or 64 bit)
  - `type int int64` or `type int int32`
- `uint8`, `uint16`, `uint32`, `uint64` are unsigned types
  - `uint` is *native* sized unsigned integer (either 32 or 64 bit)

Notes:

- compiler chooses the size of `int`
  - can be different on identical hardware architecture

----

## `rune` Type

- represents a Unicode code point (i.e, unicode character/symbol)
- synonym for `int32`
  - `type rune = int32` (not `type rune int32`)
- use single quotes for `rune`
  - `'9'` is 71 (U+0039)
  - `'Ƣ'` is 418 (U+01A2)

----

## `byte` Type

- used for raw data (usually as `[]byte`)
- synonym for `uint8`
  - `type byte = uint8`

----

## Signed Numbers

- 2's-complement
  - `int8( 5)` is `00000101`
  - `int8(-5)` is `11111011`
- range $\left[ -2^{n-1}, 2^{n-1} - 1 \right]$
  - for `int8` this is $[-128, 127]$
- used for counting, indexing

Notes:

- There exist other representations
  - 1's complement
  - BCD

----

## Unsigned Numbers

- range $\left[ 0, 2^n -1 \right]$
  - for `uint8` this is $[0, 255]$
- primary used for bit masks (not counting)

----

## Unsigned vs Signed

- `int64` type for file size (not `uint64` or `int`)
- `int` is used for indexing into slices and arrays
  - cannot have more elements in array than addressable memory
  - `int64` is overkill on a 32 bit machine
- `uint` not used much, instead use sized unsigned types (e.g., `uint16`)
  - Package `math/bits` has useful functions (all operations on unsigned)
    - `bits.OnesCount8(00101111)` returns 5

Notes:

- `OnesCount8()` implemented as a table lookup (256 entries)
  - Often a CPU instructions call `popcount` implements this
  - alternative requires many CPU cycles (especially on `uint64`)
  - Trivia: claims that NSA request this instruction in CPUs
    - relates to cryptography
    - equal to hamming weight of a binary vector

----

## Bitwise Operators

- Show [bitmask code](../../examples/bitmask/main.go)

---

## Type Conversions

```go
i := int(45.67) // i is 45, truncates

googol := 1e100
i := int(googol) // implementation dependent
```

- [Googol](https://en.wikipedia.org/wiki/Googol)

---

## Formatting

- `fmt.*printf()` family, format integers, floats, etc.
  - package `strconv` is the underlying package

- review the [docs](https://pkg.go.dev/fmt#hdr-Printing)

---

## Floating-Point Numbers

- `float32`, `float64`
  - `3.14159`, `NaN`, `Inf`, `-Inf`
- package `math` has many useful [functions](https://pkg.go.dev/math)
- rule of thumb: do not use `==` to test equality of two floating-point numbers
  - instead use `math.Abs(x - y) < 0.001`

---

## Complex Numbers

- `complex64` and `complex128`
- `complex(re, im)` constructs a `complex128`
- package `math/cmplx` has useful functions paralleling package `math`

```go
x := complex(1, 2)                  // 1+2i
y := 3 + 4i                         // 3+4i
fmt.Println(x * y)                  // "(-5+10i)" 
// (1+2i)*(3+4i) = 3+4i+6i+8i^2 = 3+10i-8 = -5+10i
fmt.Println(real(x * y))            // "-5"
fmt.Println(imag(x * y))            // "10"
fmt.Println(3.14+19.2i, 1i*1i, 45i) // (3.14+19.2i) (-1+0i) (0+45i)
```

---

## Booleans

- value of type `bool`
- can be `true` or `false`
- comparison operators (`==`, `<`, `>=`) produce `bool`
- only `bool` allowed in conditionals (`if`, `for`)
- operators
  - `!x` boolean NOT
  - `x && y` boolean AND of `x` and `y`
  - `x || y` boolean OR of `x` and `y`

----

- `&&` and `||` have short circuit behavior
  - if result is already determined by value of first argument do not evaluate second argument
- `len(s) != 0 && s[0] == 'x'`
- `doNormal(i) || doFallback(i, j)`
- `&&` has higher precedence than `||`
  - no parentheses required for conditions of the form:

```go
if 'a' <= c && c <= 'z' ||
'A' <= c && c <= 'Z' ||
'0' <= c && c <= '9' {
}
```

Notes:

- function calls do not have short-circuit behavior
  - all arguments are evaluated first, then the function
  - `&&` and `||` could **not** be implemented as a function

---

## Strings

- **immutable** sequence of **bytes**
- contain arbitrary data including bytes with value 0 (NULL, `\0`)
- usually contains human-readable text
- conventionally interpreted as UTF-8 encoded
  - sequence of Unicode code points (runes)
- `len(s)`: number of *bytes* (as an `int`) of string `s`
  - `cap(s)` not defined for strings

----

## String Indexing

- `s[5]`: byte (not character/rune) at index 5
- `s[len(s)]` panics due to being out of range
- strings can be sliced, `s[i:j]`, is $j-i$ bytes (not runes)
  - `s[:4]` is elements 0, 1, 2, 3
  - `s[2:]` is elements 2, 3, 4, ..., `len(s)-1`
  - very efficient
- `s[5] = 'L'` compilation error

----

- concatenation
  - `s + w` new string by concatenating `s` and `w`
  - recall strings are immutable (no appending to a string)
  - `s += w` new string (`s + w`) and assigns it to variable `s`
  - always involves a single memory allocation
  - consider `strings.Builder` for building strings efficiently
- comparison
  - `==` is byte for byte comparison
  - `<` is natural lexicographic ordering (by byte)

----

From [here](https://en.wikipedia.org/wiki/UTF-8)

> The chosen values of the leading bytes means that a list of UTF-8 strings can be sorted in code point order by sorting the corresponding byte sequences.

----

## Representation

- string is really
  - data `*byte`
  - length `int`
- many strings can point to the same array (or parts of it)

----

## Literals

- `.go` files are UTF-8
  - no need to escape unicode in strings
- escape characters similar to other languages `\n\tx\"y\'`
- arbitrary bytes
  - `\xhh` hexadecimal byte value
  - `\oooo` octal byte value
  - `\uhhhh` unicode code point in hexadecimal code point

----

- raw (non-escaped) string literal uses backquotes

```go
x := `"hello
world" and \n` // on two lines with double quotes and slashes followed by n
```

----

## Character vs Byte

- $i^{th}$ byte not necessarily $i^{th}$ character
- UTF-8 encoding of non-ASCII characters uses two or more bytes

---

## Unicode

- collects all of the characters in all of the world’s writing systems, plus:
  - accents and other diacritical marks
  - control codes like tab and carriage return
  - plenty of esoterica
  - emojis
  - playing cards
  - chess symbols
- assigns each character standard number called Unicode code point

----

- Unicode version 8 defines code points for over 120,000 characters in well over 100 languages and scripts
- Go uses `int32`: natural data type to hold a single `rune`
- UTF-16 - two bytes, 65k not enough code points to represent all characters
- UTF-32 - 4 bytes, waste of memory most of the time
- UTF-8 - compact but can store all code points

----

## UTF-8 Encoding (To the Rescue)

- invented by Ken Thompson and Rob Pike, two of the inventors of Go!
- **variable-length encoding** of Unicode code points as bytes
- uses between 1 and 4 bytes to represent each rune
  - uses 1 byte for ASCII characters
  - uses 2 or 3 bytes for most code points in common use
  - use of 4 bytes is rare

----

- high-order bits of first byte of rune encoding indicate how many bytes follow
  - high-order 0 indicates 7-bit ASCII (each code point takes 1 byte)
    - identical to conventional ASCII
  - high-order 110 indicates that code point takes 2 bytes
    - second byte begins with 10
- larger runes have analogous encodings

----

| First   | Last     | Byte 1   | Byte 2   | Byte 3   | Byte 4   |
|---------|----------|----------|----------|----------|----------|
| U+0000  | U+007F   | 0xxxxxxx |          |          |          |
| U+0080  | U+07FF   | 110xxxxx | 10xxxxxx |          |          |
| U+0800  | U+FFFF   | 1110xxxx | 10xxxxxx | 10xxxxxx |          |
| U+10000 | U+10FFFF | 11110xxx | 10xxxxxx | 10xxxxxx | 10xxxxxx |

- Lexicographic sorting

----

- Unicode characters [tables](https://en.wikipedia.org/wiki/List_of_Unicode_characters)
- See [example](../../examples/strunicode/main.go)

----

- advantages:
  - backward compatible with ASCII
    - ASCII encoded file is a UTF-8 encoded file
  - compact: $E[B / C]$ is smaller than UTF-32
  - self-synchronizing
    - can find the start of a code point by looking back no more than 3 bytes
  - prefix code
    - can be processed in order without lookahead (efficient computationally)
- disadvantage:
  - finding $i^{th}$ character is $O(i)$ complexity
  - invalid UTF-8 encoding is possible

---

## `[]byte`

- string contains an array of bytes
  - once created, it is immutable
- elements of a byte slice can be freely modified
- strings can be converted to byte slices and back again

----

```go
s := "abc" // s points to "static" memory in the program
b := []byte(s) // copy
s2 := string(b) // copy
```

- conceptually, `[]byte(s)` conversion
  - allocates new byte array holding copy of bytes of s
  - returns slice that references entirety of that array
- conversion from byte slice back to string with `string(b)` also makes a copy
  - ensures immutability of `s2` (by `b`)

----

- package `strings` has many functions
- package `bytes` parallels `strings` package
  - use correct one to avoid excess copying
- `bytes.Buffer` useful for efficiently manipulating bytes (e.g., appending)

---

## Constants

- Constants are like variables but they cannot change
- Benefits:
  - immutable: easier to understand code
  - optionally untyped constants
  - compiler/linker/loader may put constants in a page of memory and mark it as non-writable

----

## Iota Generator

- `const` declaration may use constant generator `iota`
- used to create sequence of related values without spelling out each one explicitly
- `iota` begins at zero, increments by one for each item in the sequence
- useful for defining enumerations

----

```go
type Month int
const (
    January Month = iota
    February
    March
    April
    May
    June
    July
    August
    September
    October
    November
    December // 11
)
```

----

```go
type Verb uint8
const (
    VerbGet Verb = 1 << iota // 0b000001
    VerbList                  // 0b000010
    VerbWatch                 // 0b000100
    VerbUpdate                // 0b001000
    VerbPatch                 // 0b010000
    VerbDelete                // 0b100000
)
```

----

## Untyped Constants

- allow extremely high precision arithmetic at compile time
  - at least 256 bits
- usable in more expressions than typed constants/variables

```go
const daysOfWeek = 7                  // untyped
const pi = 3.141592654                // untyped
const e float64 = 2.71828182845904523 // typed
```
