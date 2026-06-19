package main

import "fmt"

func main() {
	x := "-hello!-"
	for _, x := range x {
		x := x
		if d := 'A' - 'a'; x >= 'a' && x <= 'z' {
			x := x + d
			fmt.Printf("%c", x)
		} else if x == '-' {
			fmt.Print("=")
		}
	}
}
