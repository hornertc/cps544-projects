# Go Programming

## Concurrency and Shared Variables

---

## Sequential vs Concurrent

- **Sequential program**: only one goroutine, steps happen in sequential order
- **Concurrent program**: more than one goroutine
  - Steps within a goroutine are sequential
  - Step X in goroutine 1 happen
    - before step Y in goroutine 2
    - after step Y in goroutine 2
    - simultaneously with step Y in goroutine 2
  - Cannot confidently claim step X happens before step Y.  They are **concurrent**.

----

## Concurrency-Safe

- Consider a function that works correctly in a sequential program
- Function is **concurrency-safe** if it continues to work correctly when called concurrently
  - from two or more goroutines
  - with no additional synchronization
- Can be generalized to *concurrency-safe types* if all exported methods are concurrency-safe

----

- A package can be concurrency-safe even if every type is not concurrency-safe
- Concurrent access to most variables may be avoided either
  - by confining them to a single goroutine (*confinement*) or
  - by maintaining a higher-level invariant of *mutual exclusion*
- Exported, package level functions are generally expected to be concurrency-safe
  - Access to package level variables (exported and non-exported) requires mutual exclusion

---

## Problems in Concurrent Programs

- **Deadlock** - program cannot proceed because two goroutines are waiting on resources from the other
- **Livelock** - program runs in a cycle but not making progress
- **Resource starvation** - running out of a resource such as file descriptors, memory, CPU, database locks
  - Can be caused by *unbounded parallelism*

Notes:

- Livelock example: Two resources are needed but only one can be acquired at a time, re-tries do not resolve the issue

----

- **Race condition** - program does not give the correct result for some interleaving of operations of multiple goroutines
  - may remain latent in a program, appear infrequently, only showing up
    - under heavy load
    - certain compilers
    - platforms/architectures
  - very hard to reproduce

----

- Statements are not atomic
  - `x = y - 100` can have an unexpected result unless access to `x` and `y` are synchronized in some way
- For performance, compilers store variables in CPU registers not visible by other goroutines
  - Cannot depend on `y` being the last updated value unless synchronization has occurred
  - Cannot depend on `x` being visible by other goroutines unless synchronization has occurred
  
----

## Data Race

- One type of race condition
- Data race occurs whenever two goroutines access the *same* variable concurrently and at least one of the accesses is a *write*

----

## Heisenbug

- Pun on Werner Heisenberg's use of the [observer effect](https://en.wikipedia.org/wiki/Observer_effect_(physics)) in the [Uncertainty Principle](https://en.wikipedia.org/wiki/Uncertainty_principle)
  - Quantum mechanics principle which states that the act of observing a system inevitably alters its state
- [Heisenbug](https://en.wikipedia.org/wiki/Heisenbug) is a bug that seems to disappear or alter its behavior when one attempts to study it.

----

## Other Fun Bug Names

- **bohrbug** is a "good, solid bug". Like the deterministic Bohr atom model, they do not change their behavior and are relatively easily detected

- **mandelbug** (named after Benoît Mandelbrot's fractal) is a bug whose causes are so complex it defies repair, or makes its behavior appear chaotic or even non-deterministic

----

- **schrödinbug** (named after Erwin Schrödinger) is a bug that manifests itself in running software after a programmer notices that the code should never have worked in the first place

- **hindenbug** (named after the Hindenburg disaster) is a bug with catastrophic behavior

- **higgs-bugson** (named after the Higgs boson particle) is a bug that is predicted to exist based upon other observed conditions (e.g., vaguely related log entries and anecdotal user reports) but is difficult, if not impossible, to artificially reproduce in a development environment.

---

## Avoiding Data Races: Option 1

- Do not write to the variable
  - Reads are perfectly fine to occur concurrently
  - Often possible by defining a `func init()` that initializes variables
    - Thereafter they are only read
  - In general, populate common variables **before** creating goroutines that need read-only access to those variables

Notes:

- These options are in priority order.  Option 1 is preferred over 2 and over 3.

----

## Avoiding Data Races: Option 2

- For variables that do need to be read and written
- Can *confine* access to the variable to a **single** goroutine
  - **confinement** done by communicating actions/behavior to the *broker* goroutine (via channels) in the form of messages

> Do not communicate by sharing memory; instead, share memory by communicating

----

- Confinement can also be accomplished when the variable *moved* between goroutines
- Called **serial confinement**
  - Pipelines can easily do this by passing the variable over a channel
    - sender/producer releases ownership of the variable
    - receiver/consumer acquires ownership of the variable
  - No two goroutines *own* the same variable at the same time

Notes:

- Go has not formal mechanisms to transfer ownership
  - Onerous is on the developer to not access the variable after sending it
- Rust does have such mechanisms
- See `$GOPL/ch9/bank1/bank.go`

----

## Avoiding Data Races: Option 3

- Allow many goroutines to access a variable but only *one at a time*
- Buffered channel of size $n$ can be used as a **counting semaphore**
  - Ensures no more than $n$ goroutines can access a variable
  - Access is gained by receiving (alternatively sending) a *token* on the channel
  - Access is released by sending (alternatively receiving) a *token* to the channel

Notes:

- Conceptually easier to think about "receiving" as acquiring a token
- Implementation is easier if "sending" is used to acquire a token
  - No need to pre-fill the channel with token(s)

----

- $n=1$ implies only a single goroutine can access a variable, **binary semaphore**
  - Provides *mutual exclusion* - one goroutine has access to the variable at the same time
- **mutual exclusion** is so common the `sync.Mutex` does exactly this
  - More efficient than using a channel

Notes:

- See `$GOPL/ch9/bank2/bank.go`

----

## `sync.Mutex`

- Call `mutex.Lock()` to acquire the token
- Read or modify the variable(s)
  - called the **critical section**
- Call `mutex.Unlock()` to release the token
  - Deadlock occurs if not unlocked
  - Often `defer mutex.Unlock()` is used to ensure unlock is called for *all* logic paths

Notes:

- See `$GOPL/ch9/bank3/bank.go`

----

- Low level primitive
  - Channel contains a mutex
- Mutex must not be copied after first use
- `mutex.TryLock()` can be used to acquire the lock if available
  - Returns a `bool` indicating if it got the lock
- Mutex is not **re-entrant**
  - If lock already acquired by the goroutine, another call to `mutex.Lock()` is deadlock
  - Re-entrant locks allow a single goroutine to acquire the lock many times
  - Care must be taken to not acquire the lock and then call a method that also acquires the lock

----

## `sync.Mutex` Best Practices

- Use `defer mutex.Unlock()` but keep the critical section as small as possible
- Define the variables guarded by a mutex immediately following the mutex
- Acquire the lock only in exported methods
  - Never acquire the lock in non-exported methods
  - Never call an exported function from exported methods
- Ensure that exported methods are atomic (complete and consistent)

----

## `sync.RWMutex`

- Often variables have many more reads than writes
- `sync.RWMutex` is a multi-reader, single writer lock
  - Readers are non-exclusive
    - Can have any number of them
  - Writer is exclusive
    - When acquired there must be no readers or writers

---

## Writing to Stdout Concurrently

- `fmt.Println(...)` is not concurrency-safe
  - Same as `fmt.Fprintln(os.Stdout, ...)`
  - Data race on `os.Stdout`
  - Calls to `.Write()` will be interleaved and garble output
- `log` is concurrency-safe
  - formatted logs, `printf`
- `slog` is concurrency-safe
  - structured logs, JSON lines

---

## Memory Synchronization

- There are two reasons `Balance()` needs synchronization
  - *mutual exclusion* - ensure that `Withdraw()` is not done concurrently with `Balance()`
  - *memory synchronization* - ensure that the change to `balance` is seen by `Balance`

Notes:

- Show [bank.go](../../examples/memsync/bank.go)` (add a main function) in [assembly](https://godbolt.org/)
- Show details of Power64 instruction [LWSYNC](https://www.ibm.com/docs/en/xl-c-aix/13.1.2?topic=functions-lwsync-iospace-lwsync)

----

- Modern computer there may have dozens of processors
- Each processor has its own local cache of the main memory
- For efficiency, writes to memory are
  - buffered within each processor and
  - flushed out to main memory only when necessary
  - may be committed to main memory in a different order
- Synchronization primitives like channel and mutex operations cause the processor to flush and commit all cached writes

---

## Concurrency-safe Initialization

- Can use a `sync.Mutex` but inefficent because only ever one write
- Mostly reads so a `sync.RWMutex` is a better choice
  - For efficiency, grab `RLock()` first and see if it is initialized.  If not acquire exclusive (expensive) writer `Lock()` and initialize.

----

- `sync.Once` is a simple way to implement this pattern that is more efficient

```go
var once sync.Once
var x ExpensiveTypeToInitialize
func initX() {
    x = /* ... */
}
func X() ExpensiveTypeToInitialize {
    once.Do(initX)
    return x
}
```

- `initX` is guaranteed
  - to be called at most once
  - called before `x` is returned from `X()`

---

## Atomic Types

- `sync/atomic` provides atomic types
  - Concurrency safe integer-like types
  - Does not use a mutex
  - More efficient than a Mutex
  - Does not block
- `Bool`, `Int32/64`, `Uint32/64`, `Pointer`, `Value` (`any`)
- Only for primitive types
  - However the `Pointer` and `Value` can point to complex types

Notes:

- Show [assembly](https://godbolt.org/) for `examples/atomic/atomic.go`

---

## Race Detector

- Even with careful coding, data races do occur
- Go provides a sophisticated race detector to dynamically analyze a program
- Simply add `-race` to `go build` or `go run` or `go test`
- Builds a modified version of the executable that records
  - all access to shared variables that occurred during execution
  - identify the goroutines that read or wrote the variable

----

- Also detects synchronization events
  - Channel recv/send
  - Mutex `Lock()`
  - WaitGroup `Wait()`
  - `go` statements
- Analyzes the stream of records/events and detects cases where one goroutine reads or writes a shared variable that was most recently written by a different goroutine without an intervening synchronization operation

----

- Tool reports
  - declaration location and type of variable
  - reading and writing goroutines
    - stack of variable access
    - stack when goroutine was created

----

- race detector reports **all** data races that were actually executed
- however, it can only detect race conditions that occur during a run
- it cannot prove that none will ever occur
- for best results, tests should exercise packages using concurrency

----

- race detector runs slow and uses more memory
  - due to extra bookkeeping
  - overhead is often 2-20x slower and 5-10x memory overhead
  - often worth running all the time for tests
  - sometimes used in production

Notes:

- `go run ./examples/race/` and `go run -race ./examples/race/`
  - Also show with `time.Sleep` and `log` (which adds synchronization)

---

## Channel Implementation Details

- Internal implementation [details](https://go101.org/article/channel.html)

---

## Distributed Systems

- *Distributed system* is a system that is composed of multiple programs often running on many different networked computers
  - Inherently concurrent
- *Distributed deadlock* - similar to deadlock but in a distributed system
- Lots of interesting algorithms in this space
  - Distributed consensus - Paxos, Raft
  - Distributed transactions - Three-Phase Commit (3PC)
