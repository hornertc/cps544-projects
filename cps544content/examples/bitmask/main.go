package main

import (
	"fmt"
	"math/bits"
)

func main() {
	p := func(desc string, value uint8) {
		// %b is binary
		// 08 is adverb that zero pads to exactly 8 digits
		fmt.Printf("%08b -> %s\n", value, desc)
	}

	var x uint8 = 1<<1 | 1<<6 | 1<<3
	var y uint8 = 1<<1 | 1<<2

	var x1 uint8 = 0b0100_1010
	var x2 uint8 = 0x4A // or 0X4A
	var x3 uint8 = 0112
	if x != x1 || x != x2 || x != x3 {
		panic("What!")
	}

	fmt.Println("76543210")
	p("x", x)
	p("y", y)
	// x&&y does not compile
	p("AND, intersection", x&y)
	p("OR, union", x|y)
	p("XOR, symmetric difference", x^y)
	p("NOT, negation", ^x) // note ~x is not negation like other languages
	p("AND NOT, bit clear, difference", x&^6)
	p("AND NOT again", x&(^uint8(6)))

	for i := 0; i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println("Bit set", i)
		}
	}

	p("shift left", x<<2)  // same as x*4 if no overflow
	p("shift right", x>>2) // same as x/4 if no overflow
	p("circular shift left", bits.RotateLeft8(x, 2))
	p("circular shift right", bits.RotateLeft8(x, -2))

	fmt.Println("Count ones", bits.OnesCount8(x))
}
