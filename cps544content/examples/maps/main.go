package main

import "fmt"

func main() {
	x := make(map[string]int, 50) // with capacity hint
	fmt.Println(len(x), cap(x))
}
