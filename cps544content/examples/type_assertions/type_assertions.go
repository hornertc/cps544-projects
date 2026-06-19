package main

import (
	"fmt"
	"io"
)

type Incrementor interface {
	Inc()
}

type Decrementor interface {
	Dec()
}

type Both interface {
	Incrementor
	Decrementor
}

type Counter int

func (c *Counter) Inc() {
	*c += 1
}

func (c *Counter) Dec() {
	*c -= 1
}

func main() {
	x := Counter(10)

	var i Incrementor = &x
	var d Decrementor = &x
	var b Both = &x

	fmt.Printf("x (%[1]T : %[1]v)\n", &x)
	fmt.Printf("i (%[1]T : %[1]v)\n", i)
	fmt.Printf("d (%[1]T : %[1]v)\n", d)
	fmt.Printf("b (%[1]T : %[1]v)\n", b)

	// Concrete type assertion
	y1, ok := i.(*Counter)
	fmt.Printf("y1 (%[1]T : %[1]v) %v\n", y1, ok)

	// Concrete type assertion (not that type, but implements the methods)
	y2, ok := i.(*Counter2)
	fmt.Printf("y2 (%[1]T : %[1]v) %v\n", y2, ok)

	// panics
	// y2b := i.(*Counter2)
	// fmt.Printf("y2b (%[1]T : %[1]v)\n", y2b)

	// Interface type assertion (implements)
	y3, ok := i.(Both) // same for Decrementor
	fmt.Printf("y3 (%[1]T : %[1]v) %v\n", y3, ok)

	// Interface type assertion (does not implement)
	y4, ok := i.(io.Reader)
	fmt.Printf("y4 (%[1]T : %[1]v) %v\n", y4, ok)

	// panics
	// y4b := i.(io.Reader)
	// fmt.Printf("y4b (%[1]T : %[1]v)\n", y4b)

	// weak contracts
	type up interface {
		Inc()
	}

	y5, ok := i.(up)
	fmt.Printf("y5 (%[1]T : %[1]v) %v\n", y5, ok)
}

type Counter2 int

func (c *Counter2) Inc() {
	*c += 1
}

func (c *Counter2) Dec() {
	*c -= 1
}
