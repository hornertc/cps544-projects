// package
package main

// imports
import "fmt"

// ExampleSearchSolutions is an example test for the function SearchSolutions in opgame.go
func ExampleSearchSolutions() {
	// set variables to be used as the parameters for the SearchSolutions function
	numbers := []int{2, 3, 5} // numbers on the LHS of the equation
	answer := 10              // number on the RHS of the equation
	startExpression := "2"    // starting number
	currIndex := 1            // starting index
	currResult := 2           // current answer at the start of the function/the first number in 'numbers'
	var solutions []string    // empty slice of strings to store found solutions

	// call SearchSolutions from opgame.go with the test function variables
	// print the solutions
	// verifies 'solutions'is equal to the output documented below
	// test passes if output matches, test fails if they do not
	SearchSolutions(numbers, answer, startExpression, currIndex, currResult, &solutions)
	fmt.Println(solutions)
	// Output: [2 + 3 + 5 = 10]
} // end ExampleSearchSolutions
