package main

import "fmt"

var x, z int

func init() {
	x, y := 1, 2
	z = x + y
}

func main() {
	fmt.Println(x, z)
}
