# Go Programming

## Goroutines vs Threads

---

## Hardware

- Modern computing hardware
  - Processor (qty. 1, 2, or 4) - integrated circuit
    - CPU cores (qty. 1, 4, 12, ..., 64) - dedicated ALU, fetch, decode, program counter (PC)
      - Threads - 2 per core with Hyper-threading - is an extra PC creates another virtual core
- Sharing memory and cache along the way (L3, sometimes L2, not L1)

Notes:

- `cat /proc/cpuinfo` shows 12 cores, E-2176M
  - But it is actually 6 cores with hyperthreading (12 threads)

## Operating System Threads (pthreads)

- Has a fixed and some what large stack size (e.g., 2MB)
- Managed by the kernel
- Kernel schedules _OS_ _threads_ to _CPU_ _cores_

----

## Context Switch

- Context switch of an OS thread occurs when
  - Thread calls a blocking syscall, or
  - Hardware timer interrupts CPU core when allotted time slice of thread expires
    - Timer configured by the kernel thread scheduler

----

- Causes a kernel scheduler to be invoked:
  - suspends currently executing thread and saves its registers in memory
  - looks over list of threads (of all programs) and decides which one should run next
  - restores that thread’s registers from memory
  - resumes execution of that thread

----

- OS threads are scheduled by kernel, so passing control from one thread to another requires a full context switch
  - saving the state of one user thread to memory
  - restoring the state of another
  - updating the scheduler’s data structures
- Full context switch is slow
  - poor memory locality
  - large number of memory load/store

---

## Go Scheduler

- Go runtime contains its own scheduler that uses a technique known as m:n scheduling
  - multiplexes (or schedules) $m$ goroutines on $n$ OS threads.
- Go scheduler is analogous to that of the kernel scheduler
  - however, it is concerned only with the goroutines of a **single** Go program

----

- Go scheduler is invoked by timers and implicitly by certain
Go language constructs
  - e.g., when goroutine calls time.Sleep or blocks in a channel or mutex operation
    - Scheduler puts it to sleep and runs another goroutine
- Go scheduler doesn’t need a switch to kernel context
  - Therefore switching goroutines is much cheaper than rescheduling a thread

----

- Go scheduler determines which goroutines will run on the thread pool
- `GOMAXPROCS` environment variable sets the size of the thread pool
- Can run thousands of goroutines on a single OS thread
- goroutines that are sleeping or blocked in communication do not need a thread at all
- goroutines that are blocked in system calls or are calling non-Go functions do need an OS thread
  - `GOMAXPROCS` need not account for them

Notes:

- `GOMAXPROCS=1 go run ./examples/gomaxprocs`
- `GOMAXPROCS=2 go run ./examples/gomaxprocs`

---

## Go Runtime

- Go runtime includes two primary components
  - Garbage collector
  - Goroutine scheduler
- Also manages the OS level non-blocking IO event loop
  - `select`, `kqueue`, `epoll`

---

## Goroutines

- Each Goroutine has a variable stack size
  - starts at 2 KiB
  - grows to ~1 GiB
  - 100k+ Goroutines can be used thus initially small and dynamic stack sizes are important
- Variable stack size is very useful
  - hard to determine how much stack a goroutine will needed a prior

----

## Goroutines and Thread Identity

- OS threads:  current thread has a unique identity accessible to the program
  - Similar to `Thread.GetCurrentID()` and returns an `int`
- Makes it easy to implement thread-local storage
  - essentially a global map keyed by thread identity
  - each thread can store and retrieve values independent of other threads

----

- Goroutines have no identity (accessible by a Go program) by design
  - Go scheduler does assign integer IDs to the goroutines
- Thread local storage tends to be abused
  - Leads to function behavior that is different based on what thread is being used
  - Function's behavior is not dependent on its arguments alone

----

- Go encourages simpler style of programming
  - parameters that affect the behavior of a function are explicit
- Make programs easier to read
  - Programs are read much more than they are written

---

## Polling vs Eventing

- Polling
  - Loop forever testing a condition then sleeping
  - Busy wait
  - Power inefficient
- Eventing (event-driven loops)
  - Loop forever but block on a call to `select`/`kqueue`/`epoll` (C functions)
  - There is no sleep
  - When one of the events passed to `select` occurs, `select` returns

Notes:

- Polling is similar to "Are we there yet"
