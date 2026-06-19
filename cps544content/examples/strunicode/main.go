package main

import "fmt"

/*
// byte is an alias for uint8 and is equivalent to uint8 in all ways. It is
// used, by convention, to distinguish byte values from 8-bit unsigned
// integer values.
type byte = uint8

// rune is an alias for int32 and is equivalent to int32 in all ways. It is
// used, by convention, to distinguish character values from integer values.
type rune = int32

// any is an alias for interface{} and is equivalent to interface{} in all ways.
type any = interface{}
*/

func main() {
	s := "✔🃁\u0053"

	fmt.Println("Length", len(s))
	fmt.Printf("First item: %d %[1]x %08[1]b %[1]T\n", s[0])
	// notice the MSB=1 (so this is non-ASCII)
	// notice the type is uint8 (byte) and not a rune (int32)

	fmt.Println("\nRunes...")
	for i, c := range s {
		fmt.Printf("%d %c % 8[2]d U+%06[2]x %032[2]b %[2]T\n", i, c)
	}

	fmt.Println("\nBytes...")
	for i, b := range []byte(s) {
		fmt.Printf("%d %x %08[2]b\n", i, b)
	}
}
