// package
package main

// imports
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// start main
func main() {
	// scanner to read input
	scanner := bufio.NewScanner(os.Stdin)

	// Loop, scan each line for input
	for scanner.Scan() {
		input := scanner.Text()
		// fmt.Println("input:", input)

		// create a slice of strings from input
		inputSlice := strings.Fields(input)
		// fmt.Println("inputSlice:", inputSlice)

		// check to make sure input was provided
		if len(inputSlice) < 3 {
			inputErr := "Error: No valid expression provided\n"
			_, writeErr := os.Stderr.WriteString(inputErr)
			if writeErr != nil {
				fmt.Println("Error writing to stderr:", writeErr)
			}
			os.Exit(1)
		}

		// create an empty slice of ints with length equal to inputSlice
		numbers := make([]int, len(inputSlice))
		// fmt.Println("numbers:", numbers)

		// loop through slice of strings, convert each element to int, and add to slice of ints (numbers)
		// if element cannot be converted to int, print error and exit
		for i, stringSlice := range inputSlice {
			num, err := strconv.Atoi(stringSlice)
			if err != nil {
				convertErr := "Error: Invalid input, whole numbers only. " + err.Error() + "\n"
				_, writeErr := os.Stderr.WriteString(convertErr)
				if writeErr != nil {
					fmt.Println("Error writing to stderr:", writeErr)
				}
				os.Exit(1)
			}
			numbers[i] = num
		}

		// fmt.Println("numbers:", numbers)

		// create an empty slice of strings to store found solutions
		solutions := []string{}
		// fmt.Println("solutions:", solutions)

		// store the answer to the expression, the last element of the slice of ints (numbers)
		answer := numbers[len(numbers)-1]
		// fmt.Println("answer:", answer)

		// store the starting number of the expression as a string, the first element of the slice of ints (numbers)
		startExpression := strconv.Itoa(numbers[0])
		// fmt.Println("startExpression:", startExpression)
		// fmt.Println("searchSolutions:", numbers[0:len(numbers)-1], answer, startExpression, 1, numbers[0], &solutions)
		// fmt.Println()

		// call the searchSolutions function to find all possible solutions of the input
		// passes:
		// all numbers on the LHS of the equation / all but the last number of user input
		// the number on the RHS of the equation / the answer / the last number of user input
		// the starting number of the expression / the first number of user input as a string
		// starting index of 1, the second number of the expression
		// the current result of the expression / the first number of user input
		// pointer to the solutions variable
		SearchSolutions(numbers[0:len(numbers)-1], answer, startExpression, 1, numbers[0], &solutions)

		// fmt.Println(solutions)
		// fmt.Println()

		// if solutions were found, sort them, and print them out joined by a ',' if there are multiple
		// print empty line if no solutions found
		if len(solutions) > 0 {
			SortAsc(solutions)
			fmt.Println(strings.Join(solutions, ", "))
		} else {
			fmt.Println()
		}
	} // end scanner loop

	// check for errors with the scanner
	if err := scanner.Err(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
} // end main

// SearchSolutions is a recursive function to find all possible solutions for the input provided by the user
// parameters:
// numbers []int = slice of ints that includes all numbers on the LHS of the equation
// answer int = the number the expression should equal to / last number of the user input
// currExpression string = the current expression build on the LHS of the equation
// currIndex int = the index of the current element being added to the expression
// currResult int = the answer to the current expression
// solutions *[]string= pointer to a slice of strings to store all found solutions
func SearchSolutions(numbers []int, answer int, currExpression string, currIndex int, currResult int, solutions *[]string) {
	// fmt.Println("STARTanswer:", answer)
	// fmt.Println("STARTcurrExpression:", currExpression)
	// fmt.Println("STARTcurrIndex:", currIndex)
	// fmt.Println("STARTcurrResult:", currResult)
	// fmt.Println("STARTsolutions:", solutions)
	// fmt.Println("STARTlen(numbers):", len(numbers))
	// fmt.Println()

	// if the current index is equal to the length of numbers and the current result is equal to the answer, a solution was found. Append solution is the solutions variable
	// if the current result is not equal to the answer, no solution was found, return.
	if currIndex == len(numbers) {
		if currResult == answer {
			*solutions = append(*solutions, currExpression+" = "+strconv.Itoa(answer))
			// fmt.Println("solutions1", solutions)
			// fmt.Println()
		}
		return
	}

	// store the number at the current index to use as the next number of the expression
	num := numbers[currIndex]
	// fmt.Println("num:", num)
	// fmt.Println()

	// call searchSolutions with addition to check for a final solution or to include the next number of the user input
	// numbers, answer, and solutions remain constant
	// update currExpression to include the next number, currExpression + num
	// increase the index by 1
	// update currResult to equal the answer to the current expression, currResult + num
	// fmt.Println("ADDfunc: numbers{", numbers, "} answer{", answer, "} currExpression+num{", currExpression+fmt.Sprintf(" + %d", num), "} currIndex+1{", currIndex+1, "} currResult+num{", currResult+num, "} solutions{", solutions, "}")
	SearchSolutions(numbers, answer, currExpression+fmt.Sprintf(" + %d", num), currIndex+1, currResult+num, solutions)

	// call searchSolutions with subtraction to check for a final solution or to include the next number of the user input
	// numbers, answer, and solutions remain constant
	// update currExpression to include the next number, currExpression - num
	// increase the index by 1
	// update currResult to equal the answer to the current expression, currResult - num
	// fmt.Println("SUBfunc: numbers{", numbers, "} answer{", answer, "} currExpression-num{", currExpression+fmt.Sprintf(" - %d", num), "} currIndex+1{", currIndex+1, "} currResult-num{", currResult-num, "} solutions{", solutions, "}")
	SearchSolutions(numbers, answer, currExpression+fmt.Sprintf(" - %d", num), currIndex+1, currResult-num, solutions)

	// call searchSolutions with multiplication to check for a final solution or to include the next number of the user input
	// numbers, answer, and solutions remain constant
	// update currExpression to include the next number, currExpression * num
	// increase the index by 1
	// update currResult to equal the answer to the current expression, currResult * num
	// fmt.Println("MULTfunc: numbers{", numbers, "} answer{", answer, "} currExpression*num{", currExpression+fmt.Sprintf(" * %d", num), "} currIndex+1{", currIndex+1, "} currResult*num{", currResult*num, "} solutions{", solutions, "}")
	SearchSolutions(numbers, answer, currExpression+fmt.Sprintf(" * %d", num), currIndex+1, currResult*num, solutions)

	// if num is not 0 (cannot divide by 0) and the current result can be divided by the number at the current index evenly (whole numbers only):
	// call searchSolutions with division to check for a final solution or to include the next number of the user input
	// numbers, answer, and solutions remain constant
	// update currExpression to include the next number, currExpression / num
	// increase the index by 1
	// update currResult to equal the answer to the current expression, currResult / num
	// fmt.Println("currResult%num:", currResult, "%", num, "=", currResult%num)
	if num != 0 && currResult%num == 0 {
		// fmt.Println("DIVfunc: numbers{", numbers, "} answer{", answer, "} currExpression/num{", currExpression+fmt.Sprintf(" / %d", num), "} currIndex+1{", currIndex+1, "} currResult/num{", currResult/num, "} solutions{", solutions, "}")
		SearchSolutions(numbers, answer, currExpression+fmt.Sprintf(" / %d", num), currIndex+1, currResult/num, solutions)
	}
} // end searchSolutions

// SortAsc is a function to sort the results of the found solutions in ascending order
// parameters:
// solutions []string = slice of strings / the found solutions
func SortAsc(solutions []string) {

	// store the length of the slice of strings / found solutions
	solutionLength := len(solutions)
	updated := true // declare boolean as true
	for updated {   // loop until no swaps are made, until updated remains false
		updated = false                         // set updated to false
		for i := 0; i < solutionLength-1; i++ { // loop to compare solutions in the slice of strings

			// if solution at index[i] is greater than the solution at index[i+1], swap the solutions and set updated to true
			if solutions[i] > solutions[i+1] {
				solutions[i], solutions[i+1] = solutions[i+1], solutions[i]
				updated = true
			}
		} // end compare loop
	} // end updated loop
} // end sortSolutions
