package main

import "fmt"

func f(x int, y float64) (i, j int) {
	i += x + 1
	j = x + int(y)
	return
}

func g(z int, _ float64) (int, int) {
	return z + 10, 14
}

func main() {
	fmt.Printf("%T\n", f) // func(int, float64) (int, error)
	fmt.Printf("%T\n", g) // func(int, float64) (int, error)
}
