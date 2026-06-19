# Go Programming

## Goroutines and Channels

---

## Introduction

- concurrent programming: expression of program as composition of several autonomous activities
- examples:
  - web servers handle requests from thousands of clients at once
  - phone apps simultaneously
    - render animations in user interface
    - performing computation
    - network requests in the background

----

## Embedded Computing Example

- Assume the hardware only has one core
- Need to write a program that does the following
  - Every second set a pin 1 high for 50 us
  - Every 0.4 seconds set pin 2 high for 200 ms
  - When pin 3 goes high then set the pulse width on pin 2 to 300 ms
- One could write a sequential program for this but it would be hard to maintain
- Ideally want to describe each task separately from one another -> *concurrent programming*

Notes:

- **All** concurrent programs can be re-written as sequential programs
  - After all, concurrent programs must run on sequential hardware

----

![Concurrency vs Parallelism](../img/concurrency.png)

----

![CSP](https://user-images.githubusercontent.com/785824/200107637-8ba8cb54-2ff0-473a-89b9-50ec8f7ec6fb.png)

---

## Concurrent Programming in Go

- Goroutines and channels: support communicating sequential processes (CSP)
  - model of concurrency in which values are passed between independent activities (goroutines)
  - variables are for the most part confined to a single activity
- More traditional model of shared memory multithreading (chapter 9)

Notes:

- CSP from Tony Hoare

---

## Goroutines

- goroutine: each concurrently executing activity
- example: a program that has two functions:
  - one that does some computation
  - one that writes some output
  - assume that neither function calls the other
- in concurrent program with two or more goroutines: calls to both functions can be active at the same time

----

- when program starts, goroutine created for function `main`
  - called the main goroutine
- new goroutines created by the `go` statement
- `go` statement function or method call prefixed by `go`
  - function called in newly created goroutine
  - `go` statement itself **completes immediately**
  - return values are ignored

```go
f() // call f() and wait for result
go f() // call f() in new goroutine, don't wait
```

----

## Goroutines In Action

### Spinner

- spinner example (gopl.io/ch8/spinner), main goroutine computes the 45th Fibonacci number
- program is expressed as two autonomous activities
- only a single core is needed to run this program
  - `GOMAXPROCS=1 ./spinner`

Note:

- One could update the spinner in the `fib()` by polling the clock
  - A terrible programming style
  - Does not allow separation of concerns

----

### Concurrency In Servers

- networking is a natural domain in which to use concurrency
- servers typically handle many connections from their clients at once
- each client is often independent of the others
- `net` package provides the components for building networked client and server programs that communicate over TCP, UDP, or Unix domain sockets
  - `net/http` package builds on `net` package

----

- Clock 1
  - See `gopl.io/ch8/clock1`
  - Writes the current time to the client once per second
  - **sequential** clock server

- Clock 2
  - See `gopl.io/ch8/clock2`
  - Same as Clock 1, except
  - **concurrent** clock server

Notes:

- `go run ./ch8/clock1`
- In new terminal, `nc -v localhost 8000`
- Run a few clients

---

## Channels

- goroutines are activities of concurrent Go program
- channels are connections between them
- channel is communication mechanism that lets one goroutine send values to another goroutine
- each channel is conduit for values of a particular type
  - channel’s element type
  - e.g., type of a channel whose elements have type `int` is written `chan int`

----

## Creating Channels

- `ch := make(chan float64)` then `ch` has type `chan float64`
- must be constructed before values are sent or received over channel
- is a reference type
- zero value is `nil`

----

## Comparing Channels

- Channels are comparable (may be used as map keys)
- `==` can be used to compare channels of the same type
  - comparison is true if both are references to same channel data structure
  - e.g., pointers to the same address
- channel may also be compared to `nil`

----

## Channel Operations

- *send* puts a value on the channel
- *receive* gets a value from a channel
- both use `<-` operator
  - difference is which side the channel is on

```go
ch <- x   // send
x = <- ch // receive
<- ch     // receive and discard value
```

----

## Closing Channels

- `close(ch)` closes channel `ch`
- Sets flag indicating that no more values will ever be sent on channel
  - subsequent attempts to send (or close) will `panic`
- Receive operations on closed channel yield the values that have been sent
- After channel is empty, any receive operations
  - complete immediately
  - yield the zero value of the element type

----

## Producers and Consumers

- Producer is goroutine that sends
- Consumer is goroutine that receives
- Channels are multi-producer / multi-consumer channels
- Unicast - value send on a channel is delivered to **one** consumer
- Channel values/messages are received generally in FIFO order
  - FIFO is ill-defined with multiple producers and consumers without extra synchronization

Notes:

- Other langues tend to have specialized channels each case
  - single producer / single consumer
  - single producer / multiple consumer
  - multi-producer / single consumer

----

## Channel Types

- `make(chan T)` create an *unbuffered* channel
  - a.k.a., *synchronous* channel
- `make(chan T, n)` creates a *buffered* channel of size `n` if $n \neq 0$
  - a.k.a., *asynchronous* channel
- `len(ch)` is the number of elements in the channel
- `cap(ch)` is the buffer size of the channel

---

### Unbuffered Channels

- Send operation blocks sending goroutine
  - Until another goroutine executes a receive
- If receive operation is attempted first then receiving is blocked
  - Until another goroutine performs a send
- Then, value is transmitted
  - Receiver is awaken
  - Sender is awaken

----

- Sending data over an unbuffered channel causes sending and receiving goroutines to synchronize
  - *synchronous* channels
- Receipt of the value happens *before* reawakening sending goroutine

----

### X Happens Before Y

- X happens before Y means
  - guaranteed that X occurs earlier in time than Y and
  - all its prior effects (e.g., updates to variables) are complete and
  - they can be relied on by other goroutines
- When X neither happens before Y nor after Y, then X is
concurrent with Y
  - means that nothing can be assumed about their ordering

----

### Example: echo client

- `ch8/netcat1` only reads from TCP and writes to STDOUT
  - missing reading from STDIN and writing to TCP
- `ch8/netcat2` terminates as soon as STDIN is closed (Ctrl-D)
- `ch8/netcat3` terminates after both sides are closed
  - channel is used for synchronization
- `ch8/netcat4` symmetric example of `netcat3`

----

## Pipelines

- A form of concurrency where operations (called *stages*) are chained together
  - Results from one stage is used by the next stage
- A -> B -> C -> D
  - Once the pipeline is "full" tasks A, B, C, and D all run concurrently (often in parallel)

----

## Example Pipelines

- Counter -> Squarer -> Printer
- `ch8/pipeline1` runs forever
- `ch8/pipeline2` use `close()` to terminate the pipeline
  - sender calls `close()`
- `ch8/pipeline3` broken out into functions with unidirectional channels
- Unbounded parallelism is rarely a good idea

Notes:

- Show `$GOPL/ch8/cake`

----

- No way to test if a channel has been closed
- To test if a channel has been closed and drained use two result version of receive

```go
for x := range ch {
    // do something with x
}

// equivalent to
for {
    x, ok := <- ch
    if !ok {
        break
    }
    // do something with x
}
```

----

## Closing Channels: Details

- Not necessary to close every channel when finished with it
  - only necessary when need to tell receiving goroutines that all data have been sent
- Garbage collector will reclaim channels once unreachable (closed or not)
  - Unlike `*os.File` that needs to be closed or it leaks OS resources
- Closing a closed or `nil` channel will panic
- `close()` is a broadcast operation
  - all consumers "see" the "close"

Notes:

- Contrast `close()` with send, which is a unicast operation

----

## Unidirectional Channel Types

- Type `chan T` is "bidirectional" in the sense that you can send and receive
- Type `chan<- T` (send-only channel of `T`) allows send and close but not receives
- Type `<-chan T` (receive-only channel of `T`) allows receives but not sends
- Violations are detected at compile time
- Bidirectional channels are implicitly converted to undirectional

---

## Buffered Channels

- They have a FIFO queue of elements/values

```go
ch := make(chan string, 5)
cap(ch) // 5
len(ch) // 0
```

----

- send operation inserts an element at the back of the queue
- receive operation removes an element from the front
- if full, send operation blocks until space is made available by *another* goroutine’s receive
- if empty, receive operation blocks until value is sent by another goroutine

----

```go
ch := make(chan string, 3)
cap(ch) // 3
ch <- "A"
ch <- "B"
len(ch) // 2
ch <- "C"
// ch <- "D" // would block

fmt.Println(<-ch) // A
fmt.Println(<-ch) // B
fmt.Println(<-ch) // C
// <-ch would block because len(ch) == 0
```

----

```go
func mirroredQuery(mirrors []string) string {
    // leaks len(mirrors) - 1 goroutines
    // responses := make(chan string)

    responses := make(chan string, len(mirrors))
    for _, mirror := range mirrors {
        go func(m string) {
            // does not do what you expect (go < 1.22)
            // responses <- request(mirror)
            responses <- request(m)
        } (mirror)
    }
    return <-responses
}
```

----

## Goroutine Leaks

- Goroutines not terminating (deadlocked) when their work is complete is a **leak** of resources
- It is a software **bug**
- Garbage collector does not help
  - Only garbage collects variables not referenced anywhere
- So leaking goroutines is actually two problems
  - Consumes extra Go scheduler resources (waste of memory and CPU)
  - Does not free up its data (stack and variables) to be garbage collected (waste of memory)

----

## Buffered vs Unbuffered

- Unbuffered channels give stronger synchronization guarantees
  - every send is synchronized with its corresponding receive
- Buffered channels decouple send and receive

---

## Looping in Parallel

- Order of processing does not matter
- Problem can be broken down into independent subproblems
- *embarassingly* parallel
  - easiest form of concurrency
  - scales linearly with amount of parallelism
- See `ch8/thumbnail/thumbnail_test.go`

----

![makeThumbnail6](../img/goroutines-thumbnail6.png)

Notes:

- From page 239 for the diagram of `makeThumbnail5`

---

## Multiplexing with `select`

- What if a goroutine needs to wait on more than one channel?
  - Once a send/receive is blocked there is no way to unblock
- Need multiplexing
- `select` provides multiplexing for channels
  - wait on send
  - wait on receive
  - don't wait

----

```go
ticker := time.NewTicker(1 * time.Second)
abort := make(chan struct{})
go func(){
    os.Stdin.Read(make([]byte, 1)) // read a single byte
    abort <- struct{}{}
}
loop:
for {
  x := f()
  select {
  case t := <-ticker.C:
    // print status update at time t
  case <-abort:
    break loop
  case ch2 <- x:
    // process x
  // default:
  //     // if nothing is ready bail
  } 
}
ticker.Stop() // terminate tickers goroutine
```

Notes:

- Can only select on channels
  - Not on `io.Writer` or `sync.Mutex`

----

- `select` waits for the first case statement that is ready
  - if multiple are ready then chosen at random
- `default` case is *always* ready and is used if no other case statement is ready
  - allows for non-blocking send/receive
- `select{}` waits forever

```go
// non-blocking receive
select {
    case x := <-ch:
       // do something with x
    default:
}
// If no value is on ch then return immediately
```

----

## Non-Blocking Loop

- Drain a channel but not block until it is closed

```go
loop:
  for {
    select {
    case x, ok := <- ch:
      if !ok {
        break loop
      }
      // do something with x
    default:
      break loop
    }
  }
```

---

## Example: Chat Server

- `ch8/chat/chat.go`

---

## Contexts

- `context` package provides a universal `Context` type
- generally passed as the first argument to all IO related functions
  - not stored in structs
- See [API docs](https://pkg.go.dev/context)

----

## Context Deadlines

```go
d := time.Now().Add(shortDuration)
ctx, cancel := context.WithDeadline(context.Background(), d)

// cancel just to be safe
defer cancel()

select {
case <-neverReady:
	fmt.Println("ready")
case <-ctx.Done():
	fmt.Println(ctx.Err())
}
```

- Outputs "context deadline exceeded"

----

## Context Cancelling

```go
package main

import (
	"context"
	"fmt"
)

func main() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
```

----

## Context Values

```go
type favContextKey string

f := func(ctx context.Context, k favContextKey) {
    if v := ctx.Value(k); v != nil {
        fmt.Println("found value:", v)
        return
    }
    fmt.Println("key not found:", k)
}

k := favContextKey("language")
ctx := context.WithValue(context.Background(), k, "food")

f(ctx, k) // prints "food"
f(ctx, favContextKey("color")) // prints "key not found: color"
```

- Allows dependency injection

----

## Example: Contexts

```go
type Fetcher interface {
	Fetch(ctx context.Context, 
          target ocispec.Descriptor) (io.ReadCloser, error)
}
```

----

## HTTP Requests with Contexts

```go
func NewRequestWithContext(ctx context.Context, 
                           method, url string, 
                           body io.Reader)(*Request, error)
```

- When used with a `http.Client.Do(req)` call you can use the context to cancel the IO operation (deadline or explicitly)
- Actually stores the `context.Context` in the `Request` struct which is bad practice
  - Done for backwards compatibility reasons

---

## Testing with Concurrency

- `testing.T` methods `Log*()` and `Error*()` (and `Fail()`) may be called from created goroutines
- `Fatal*()`, `FailNow()`, `Skip*()` must only be called within a test function
  - Not a goroutine that the test function spawns
- Therefore, must collect errors from spawned goroutines in test functions
  - Channels are one way of collecting errors

Notes:

- [docs](https://pkg.go.dev/testing#T)

---

## Reminders and Lose Ends

- Only senders may close the channel
- At most, one sender may close a channel
- `nil` channel
  - sending and receiving to block forever
  - `close()` panics
- A send on a channel is received by at most one goroutine (i.e. **unicast**)
- Closing a channel **broadcasts** that information to all goroutines receiving on the channel

----

- It is not necessary to `close()` every channel
  - Only needed when the sender needs to indicate that no more data is to be sent
- It is not necessary to drain every channel
  - GC will reclaim the memory when the channel is no longer referenced (just like a slice/array)
- `select {}` blocks forever
