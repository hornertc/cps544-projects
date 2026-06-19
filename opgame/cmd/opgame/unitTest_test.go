// package
package main

// imports
import (
	"reflect"
	"testing"
)

// TestSortAsc is a unit test for the function SortAsc in opgame.go
func TestSortAsc(t *testing.T) {
	// create a slice of strings with three elements of different equations
	solutions := []string{"3 + 2 = 5", "1 + 2 = 3", "4 + 2 = 6"}

	// create a slice of strings that is the correct sort ascending order of solutions
	expected := []string{"1 + 2 = 3", "3 + 2 = 5", "4 + 2 = 6"}

	// call SortAsc function from opgame.go using 'solutions'
	SortAsc(solutions)

	// verify the sorted 'solutions' slice is equal to the 'expected' slice
	// print error message is they are not equal, test fails
	// test passes if they are equal
	if !reflect.DeepEqual(solutions, expected) {
		t.Errorf("Expected %v, but got %v", expected, solutions)
	}
} // end TestSortAsc
