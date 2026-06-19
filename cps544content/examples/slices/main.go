package main

import "fmt"

func inspect(name string, s []int) {
	fmt.Printf("%-10s is %-17s len=%d cap=%d\n", name, fmt.Sprintf("%v", s), len(s), cap(s))
}

func main() {
	x := []int{10, 11, 12, 13, 14}
	inspect("x", x)

	inspect("x[1:]", x[1:])
	inspect("x[1:3]", x[1:3])
	inspect("x[1:3:5]", x[1:3:5]) // same as above
	inspect("x[1:3:4]", x[1:3:4])
	inspect("x[1:3:3]", x[1:3:3])
	inspect("x[1:3][:3]", x[1:3][:3])

	inspect("x[0:3]", x[0:3])
	// y := x[1:3:4]
	// fmt.Println(y[0:3])

	// Slicing an array
	z := [5]int{10, 11, 12, 13, 14} // array
	inspect("z[1:5]", z[1:5])       // slice
}

// [10]T{}[:5]
