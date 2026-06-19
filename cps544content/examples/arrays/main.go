package main

import (
	"fmt"
	"reflect"
)

type Dimension int

const (
	X Dimension = iota
	Y
	Z
)

func main() {
	vector1 := [...]float64{X: 45.2, Y: 0, Z: 92} // map-like syntax
	vector2 := [...]float64{45.2, 0, 92}          // equivalent
	vector3 := [...]float64{Z: 92, X: 45.2}       // equivalent

	if !reflect.DeepEqual(vector1, vector2) {
		panic("Not equal")
	}
	if !reflect.DeepEqual(vector1, vector3) {
		panic("Not equal")
	}

	x := [5]byte{}
	y := x[1:3]
	fmt.Printf("%T %T\n", x, y)

	w := "hello✔"
	wr := []rune(w)
	fmt.Printf("%q\n", wr)
}

func modifyArray(ptr *[8]int) {
	// notice that (*ptr)[0] is not needed
	ptr[0] = ptr[1] + ptr[2]
}
