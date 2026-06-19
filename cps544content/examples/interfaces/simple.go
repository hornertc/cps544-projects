package main

import "fmt"

type I interface {
	F(dt float32)
}

type G struct{}

func (g *G) F(distance float32) {
	// do something with distance
	fmt.Println(g, "arg is", distance)
}

var _ I = (*G)(nil)

func main() {
	f := "(%T : %v) : i == nil is %v therefore i != nil is %v\n"
	{
		var i I
		fmt.Printf(f, i, i, i == nil, i != nil) //false
		// i.F(40) // panics
	}
	{
		var i I = (*G)(nil)
		fmt.Printf(f, i, i, i == nil, i != nil) //true!
		i.F(45)                                 // works
	}
}
