package main

import "fmt"

func main() {
	// Shows how to write a traditional for loop with goto statements

	fmt.Println("Traditional loop")
	n := 5
	// traditional for loop
	for i := 0; i < n; i++ {
		fmt.Println(i * i)
	}

	fmt.Println("Goto loop")
	{
		i := 0
	start:
		if !(i < n) {
			goto done
		}

		// body of loop
		{
			fmt.Println(i * i)
		}

		i++
		goto start
	done:
	}

	fmt.Println("For loop with case statement")
	for i := 0; i < n; i++ {
		fmt.Println("i =", i)
		switch {
		case i == 0:
			fmt.Println("continue")
			continue
		case i%3 == 0:
			fmt.Println("break")
			break
		}
	}

	fmt.Println("Nested loops")
	// outer:
	for i := 0; i < n; i++ {
		fmt.Println("outer i =", i)
		for j := 0; j < n; j++ {
			fmt.Println("inner j =", j)
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
			if i*j > 10 {
				fmt.Println("breaking out")
				break
				// break outer
			}
			// more code ...
		}
	}
	// NOTE can do the same with labelled continue statements
	// continue mylabel

	// NOTE labelled break statements can be used to break out of case, select, and for, when you want to break out of more than just the nearest case, select, or for

	fmt.Println("Exit handling")
	fmt.Println(f(n))
}

func f(n int) string {
	var ij int
	for i := 0; i < n; i++ {
		if ij%9 == 0 {
			goto success
		}
		for j := 0; j < n; j++ {
			ij = i * j
			if ij == 12 {
				goto success
			}
			if ij == 13 {
				goto failure
			}
		}
	}

success:
	return fmt.Sprintf("found it %d", ij)

failure:
	return fmt.Sprintf("found the enemy first %d", ij)
}
