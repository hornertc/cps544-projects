package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Print("Enter a name:")
	fmt.Scan(&name)

	var count int
	fmt.Print("Enter an integer: ")
	fmt.Scan(&count)

	var float_value float32
	var bool_value bool
	fmt.Println("Enter a float and a boolean")
	fmt.Scan(&float_value, &bool_value)

	fmt.Printf("We got %s %d %g %t", name,
		count, float_value, bool_value)
}
