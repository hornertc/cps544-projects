package main

import "fmt"

func incr(p *int) int {
	x := *p
	*p++ // increments what p points to; does not change p
	return x
}

func main() {
	v := 1
	incr(&v)
	fmt.Println(v)
	fmt.Println(v, incr(&v), v, incr(&v)) // side effect: v is now 4
}
