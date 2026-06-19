package main

import "fmt"

func main() {
	x := complex(1, 2) // 1+2i
	y := 3 + 4i        // 3+4i
	fmt.Println(x * y) // "(-5+10i)"
	// (1+2i)*(3+4i) = 3+4i+6i+8i^2 = 3+10i-8 = -5+10i
	fmt.Println(real(x * y))            // "-5"
	fmt.Println(imag(x * y))            // "10"
	fmt.Println(3.14+19.2i, 1i*1i, 45i) // (3.14+19.2i) (-1+0i) (0+45i)
}
