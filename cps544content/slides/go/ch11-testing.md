# Software Testing

---

## Comprehensive Automated Testing Benefits

- Rapid development, test, and deploy cycles
- Increased confidence in refactoring
- Less defects
- Improved software design (due to more refactoring)

---

## Classes of Testing

- White box testing (a.k.a., clear box)
  - Tester knows everything inside the "box" being tested
  - Source code is available
- Black box testing
  - Tester has incomplete knowledge about what is inside the "box"

---

### Categories of Testing

- Unit
- Randomized (Fuzzing)
- Functional
- Integration
- Regression

---

### Testing Philosophies

- Test driven development (TDD)
  - Red-Green-Refactor
- Behavior driven development (BDD)

----

## TDD

1. Issue created to identify bug

1. ❌ Write test and break the build (red build)
1. ✅ Implement minimal fix (do not modify the test) (green build)
1. ✅ Improve the implementation (refactor)

1. Submit pull request for approval

Notes:

- Ideally the tests are not modified after creation but that is rarely true in reality
- Understanding of the issue changes as the implementation is done

----

## BDD

- Slightly more focused on black-box testing
- How would a user of this component expect it to behave
- Behaviors are higher level than tests
- Conformance or compliance testing

---

### Philosophy

- Agile Manifest (Principle): "**Working** software is the primary measure of progress."
- Agile does not go far enough
- We really want "**Maintainable** software that is working correctly..."
- Thus refactoring is important, enabled by automated testing

---

### Testing in Go

- Go has a builtin test runner
  - invoked with `go test ./...`
- Minimalistic but feature complete
- Does not include an assertion library
  - one can be added easily
- For illustration below consider a Go package called `pkg`

---

### White Box Testing

- Write tests in files matching `pkg/*_test.go`
- Declare tests in `package pkg`
- Provides access to un-exported items
- Definitions in `*_test.go` cannot be imported by other packages

---

### Black Box Testing

- Write tests in files matching `pkg/*_test.go`
- Declare tests in `package pkg_test`
- Tests are not in the same package as the package under test
- Provides access to only exported items in the same way other packages will use this code
- Definitions in `*_test.go` cannot be imported by other packages

---

### Types of Tests

- Example
- Test
- Benchmark
- Fuzz

---

### Example

- Demonstrates how to properly use the library
- `go test` Compiles and runs the code
- If `// Output:` provided then also compares output
- Has the form `func ExampleXxxxx()`
- Best done as black box tests

```go
package strings_test
func ExampleJoin() {
    eqn:= strings.Join([]string{"a", "b", "c"}, "=")
    fmt.Println(eqn)
    // Output: a=b=c
}
```

---

### Test

- Primarily for unit testing of functions
- `go test` compile and runs the test
- `func TestXxxxx(t *testing.T)`
- Methods of `testing.T` are used to log and indicate failure cases

```go
func TestAdd(t *testing.T){
    got := Add(4, 6) // Add() is DUT/SUT
    want := 10
    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}
```

----

### Test (Table-Driven)

```go
func TestAdd(t *testing.T) {
	type testCase struct {
		name             string
		arg1, arg2, want int
	}
	var testCases = []testCase{
		{"basic", 2, 3, 5},
		{"large", 400, 800, 1200},
		{"negative", -3, 10, 7},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if output := Add(tt.arg1, tt.arg2); output != tt.want {
				t.Errorf("Output %q not equal to %q", output, tt.want)
			}
		})
	}
}
```

---

### Benchmark

- Used to micro-benchmark performance critical code segments
- `func BenchmarkXxxxx(b *testing.B)`
- `go test -bench=.` compile and runs the benchmark
- Benchmark runner adaptively determines the best value of `b.N`

```go
func BenchmarkFib(b *testing.B) {
    // run the Fib function b.N times
    for n := 0; n < b.N; n++ {
        Fib(10)
    }
}
```

----

### Benchmark (Sweep)

- Sweeps are often used to justify a particular parameter value (e.g. buffer size)
- `b.ReportAllocs()` can be used to report memory allocations
- `b.SetBytes()` can be used to report B/s performance information

----

### Table-Driven Sweep

```go
func BenchmarkFib(b *testing.B) {
	type testCase struct {
		name string
		arg1 int
	}
	var testCases = []testCase{
		{"basic", 2},
		{"large", 40},
		{"very large", 1000},
	}
	for _, tt := range testCases {
		b.Run(tt.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				Fib(tt.arg1)
			}
		})
	}
}
```

---

### Fuzz Test

- Randomized testing is used to automatically discover _interesting_ test cases.
- `func FuzzXxxxxx(f *testing.F)`
- Black box testing: fuzzer knows nothing of the implementation

```go
func FuzzParse(f *testing.F) {
    f.Add(`null`)
    f.Add(`21.5.6`)
    f.Add("{}")
    f.Add(`"not closed`)
    f.Fuzz(func(t *testing.T, a string) {
        // often do not need to (or cannot) check if ParseJSON() behaves properly
        // just waiting for ParseJSON() to panic()
        ParseJSON(a)
    })
}
```

----

### Fuzz Test Seeds

- Seed corpus is provided by `f.Add()`
  - limited to the types `[]byte`, `string`, `bool`, integers
  - multi-dimensional `f.Add(3.14, "dog")`
- New test cases, found by fuzzing, are added to the corpus file
  - Saved to directory `testdata/fuzz/FuzzParse`
- Seed + file corpus used to run regular `go test`

----

### Running the Fuzzer

- `go test -fuzz=. -fuzztime=1h`
- _hammers_ your software to find _new_ failure test cases
  - Computationally expensive
  - Runs a separate process per core
- Somewhat intelligently tries to sweep the entire parameter space
  - may be multi-dimensional
  - curse of dimensionality

---

### Test Execution

- Each package
  - compiled into a separate executable
  - run in parallel
- In addition, tests within the same package can opt-in to be run in parallel
  - calling `t.Parallel()`
- Data for tests is placed in `testdata` directory within the package
- Code coverage
- Profiling (CPU, memory/heap, blocking)

---

### Mocking

- Mocks are test objects used as surrogates during testing
- Usable when function under test accepts interfaces or function pointers (not concrete types)
- Mock need not implement the functionality
  - only needs to ensure the correct returns are provides for the expected parameters (table lookup is fine)
