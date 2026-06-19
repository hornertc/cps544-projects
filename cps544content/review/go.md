# Go Review Sheet

- Short tutorial/review of Go is in [Effective Go](https://go.dev/doc/effective_go)

- Do not break your public API which consists of all exported declarations (var, const, func, and type).
- Prefer not to export declarations until you are sure they will not change and needed by other packages
- Wrap errors from external packages
- Strings are immutable
- Slices are references to a contiguous "slice" of an array
- Loop variables define a new implicit scope.  At each iteration the loop body creates a new explicit lexical scope with the braces.
- Should not call `defer` in a loop

- Types
  - Basic types
    - `int`, `uint` (and their 8, 16, 32, 64 bit variants)
    - `byte`
    - `rune`
    - `uintptr` (used for like a `void *` in other languages, rarely used)
  - Aggregate types (fixed size)
    - arrays - homogeneous
    - structs - heterogeneous
  - Reference types (the only types that seem like they are passed by reference, but really they just have pointers to the real data, and thus are still passed by value)
    - slices - dynamically sized array of a single type.  A homogeneous mapping from [0, N) -> T
    - maps - dynamically sized homogeneous mapping C -> T where C is comparable
    - functions - first class functions
    - channels - send/receive, send-only, receive-only
  - Interface types - provide limited polymorphism by way of duck-typing

- Use `int` when referring to the size of memory (only 32 bits is accessible on a 32 bit machine).
- Use `int64` when referring to size of file on disk, or counting a potentially large number of items.

- `[]byte(str)` makes a copy since the `[]byte` must be mutable and string is immutable.
- `string(bytes)` makes a copy to ensure that the string is immutable (even if the `[]byte` is never modified)
- As such it is important to use either the `bytes` package or the `strings` to avoid conversion.
- slices of strings index bytes (not runes) and do not make a copy (so the operation is fast)

- Testing
  - White box testing - Know everything (including source code) for the package under test.  Handled by writing tests in the package itself so it has access to un-exported items.
  - Black box testing - Know nothing about the package under test.  This is handled with testing packages with the name "xxxxxx_test".
- Types of tests in Go that are supported by their test runner are
  - Example code.  Has the form `func ExampleXxxxx()`
  - Unit test.  Has the form `func TestXxxxx(t *testing.T)`
  - Benchmark test.  Has the form `func BenchmarkXxxxx(b *testing.B)`
  - Fuzzing test.  Has the form `func FuzzXxxxxx(f *testing.F)`
- Go testing is minimalistic and does not include an assertion library but one can be added easily

- Getters should be named without the "Get" prefix, e.g., `obj.Item()`.
- Setters should be named with the "Set" prefix, e.g., `obj.SetItem()`.

- Fundamental theorem of software engineering... "Every problem can be solved with exactly one more level of indirection".

- Functions should accept interfaces (where applicable) and return structs/values (not interfaces)
- Interfaces should be as small as possible and composed
- Pass `context.Context` as the first argument in functions, do not store contexts in structs
- Do not communicate by sharing memory; instead, share memory by communicating.

- Channels
  - Only senders may close the channel.  At most one sender may close a channel.
  - `nil` channel causes sending and receiving to block forever, and `close()` to panic
  - Closing a channel **broadcasts** that information to all goroutines receiving on the channel
  - A send on a channel is received by at most one goroutine.  It is **unicast**.
  - It is not necessary to `close()` every channel.  Only needed when the sender needs to indicate that no more data is to be received.
  - It is not necessary to drain every channel.  The GC will reclaim the memory when the channel is no longer referenced (just like a slice/array).
  - `select {}` blocks forever
  - All the test functions (e.g., `t.Fatal()`, `t.Error()`) must be called within a test function and not a goroutine that the test function spawns.  So you must collect errors from spawned goroutines in test functions.
  - More details on channel internal implementation can be found [here](https://go101.org/article/channel.html)
