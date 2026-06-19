// package
package main

// imports
import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"testing"
)

// FuzzGetInput is a fuzz test for the function GetInput in windchill.go
func FuzzGetInput(f *testing.F) {
	// create/open a file for the test case log
	// print error if the file cannot be opened
	logFile, err := os.Create("fuzz_test_case_log.txt")
	if err != nil {
		fmt.Printf("Error opening log file: %v", err)
		return
	}
	defer logFile.Close() // close the file on exit

	// variable to write output to the log file
	logOutput := io.MultiWriter(logFile, os.Stdout)

	// fuzz testing method for testing random data as a slice of bytes
	f.Fuzz(func(t *testing.T, seed int64) {

		// create a random number generator using the seed
		randGen := rand.New(rand.NewSource(seed))

		// variable for number of inputs to test
		inputSize := 10

		// generate a slice of numeric (float) values between -100 and 100 and/or non-numeric values as test input
		testInput := generateSlice(randGen, inputSize, -100, 100)

		// log the random generated values
		_, err := fmt.Fprintf(logOutput, "testInput: %s\n", testInput)
		if err != nil {
			fmt.Println("Error:", err)
		}

		// create the input buffer using each test input value followed by a new line
		var inputBuffer bytes.Buffer
		for i := 0; i < inputSize; i++ {
			_, err = inputBuffer.WriteString(fmt.Sprintf("%s\n", testInput[i]))
			if err != nil {
				fmt.Println("Error:", err)
			}
		}

		// call GetInput from windchill.go using the inputBuffer test cases and add results to the log file
		temp, speed := GetInput(&inputBuffer, logOutput)

		// log the return values of GetInput
		_, err = fmt.Fprintf(logOutput, "Temp: %.2f, Speed: %.2f\n\n", temp, speed)
		if err != nil {
			fmt.Println("Error:", err)
		}
	})
}

// generateSlice generates a slice of numeric (float) and non-numeric values equal to the input size
func generateSlice(randGen *rand.Rand, size int, min, max float64) []string {
	// create an empty slice equal to the input size
	slice := make([]string, size)

	// loop through the slice and generate a random numeric (float) or non-numeric value for each index
	for i := 0; i < size; i++ {
		slice[i] = generateRandomFloat(randGen, min, max)
	}
	// return the slice
	return slice
} // end generateSlice

// generateRandomFloat generates a random numeric (float) value or calls generateRandomString
func generateRandomFloat(randGen *rand.Rand, min, max float64) string {
	// generate a random number between 0 and 1
	// if 0, generate and return a random numeric (float) value
	if randGen.Intn(2) == 0 {
		// Generate a numeric (float) value within the specified range
		return strconv.FormatFloat(randGen.Float64()*(max-min)+min, 'f', 2, 64)
	}
	// if 1, call generateRandomString
	return generateRandomString(randGen)
} // end generateRandomFloat

// generateRandomString generates a random string with non-numeric characters
func generateRandomString(randGen *rand.Rand) string {
	// variable to store the non-numeric values the random generator can choose from
	nonNum := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$%^&*()-_=+[]{}|;:'\",.<>?/"

	// generate a random string length
	length := randGen.Intn(10)

	// create a string builder to create the random string
	var result strings.Builder

	// add a random non-numeric value from nonNum to the string builder until the string length is met
	for i := 0; i < length; i++ {
		randomIndex := randGen.Intn(len(nonNum))
		err := result.WriteByte(nonNum[randomIndex])
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
	// return the created random string
	return result.String()
} // end generateRandomString
