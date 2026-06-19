package main

import (
	"fmt"
	"time"
)

func main() {
	var x int
	go func() {
		for {
			// time.Sleep(100 * time.Millisecond)
			// log.Print("+")
			x += 2
		}
	}()

	go func() {
		for {
			// time.Sleep(150 * time.Millisecond)
			// log.Print("-")
			x -= 1
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("x is", x)
	// log.Println("x is", x)
}
