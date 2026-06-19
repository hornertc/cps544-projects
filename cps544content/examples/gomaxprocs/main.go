package main

import "fmt"

func main() {
	for {
		// NOTE living dangerously by writing to os.Stdout concurrently
		go fmt.Print(0)
		fmt.Print(1)
	}
}
