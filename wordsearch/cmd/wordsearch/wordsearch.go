// CPS 544: Assignment 5, Word Search
// Go Program that finds a combination of letters in any direction (→ ← ↓ ↑ ↘ ↗ ↙ ↖) within a matrix of letters
// Tommy Horner, thorner1

// package main
package main

// imports
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// start main
func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid command. Use 'go run ./cmd/wordsearch <word> <puzzle name>'")
		os.Exit(1)
	}

	// get arguments from command line input
	word := os.Args[1]
	puzzleName := os.Args[2]

	// variable for the puzzle file or standard input
	var puzzleFile *os.File

	// read standard input or open puzzle file from argument; print error if file cannot be opened and exit
	if puzzleName == `-` {
		puzzleFile = os.Stdin // store puzzle from standard input
	} else {
		file, err := os.Open(puzzleName)
		if err != nil {
			fmt.Printf("Error opening puzzle <%s>. Check file name.\n", puzzleName)
			os.Exit(1)
		}
		puzzleFile = file // store puzzle from file
		// defer file.Close() // close the file on program exit
	}

	scanner := bufio.NewScanner(puzzleFile) // scanner to read opened/created puzzle
	var puzzle [][]rune                     // 2D puzzle slice of runes
	var col int                             // number of columns

	// loop, read each line of the puzzle file
	for scanner.Scan() {
		row := []rune(scanner.Text()) // store the current row of the puzzle file as a slice of runes

		// assign number of columns if not already assigned
		if col == 0 {
			col = len(row)

			// print error message if the length of the current row is not equal to the assigned columns. Inconsistent line length, exit with status 3.
		} else if len(row) != col {
			_, _ = fmt.Fprintln(os.Stderr, "inconsistent line length")
			os.Exit(3)
		}

		// append the current row to 2D puzzle slice
		puzzle = append(puzzle, row)
	}

	// print error if there was a problem reading the puzzle
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading puzzle <%s>", puzzleName)
		os.Exit(2)
	}

	// close the file on program exit
	defer puzzleFile.Close()

	// store result of WordSearch
	match := WordSearch(word, puzzle)

	// sort the results in match
	lexicoSort(match)

	// loop through the results in match and print each line. Print message if no results are stored
	// result := 0
	for _, found := range match {
		fmt.Println(found)
		// result++
	}

	// if result == 0 {
	// 	fmt.Printf("'%s' was not found in <%s>\n", word, puzzleName)
	// }

} // end main

// WordSearch compares each rune of the puzzle to the first rune of the searched word.
// if a match is found calls func search to see if adjacent runes are equal to the next runes of the searched word
// requires string, [][]rune
// returns []string
func WordSearch(word string, puzzle [][]rune) []string {
	var matches []string

	// store puzzle's row and column lengths
	rows := len(puzzle)
	cols := len(puzzle[0])

	// loop through each rune of the puzzle and compare it with the first rune of the searched word
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			runes := []rune(word)
			char := puzzle[r][c]

			if char == runes[0] {

				// loop through a slice of directional structs. call func search to check if adjacent runes match the the runes in the searched word
				for _, dir := range []struct {
					dr, dc int
					arrow  string
				}{
					{0, -1, "←"}, {-1, 0, "↑"},
					{0, 1, "→"}, {1, 0, "↓"},
					{-1, -1, "↖"}, {-1, 1, "↗"},
					{1, 1, "↘"}, {1, -1, "↙"},
				} {
					// store results from search, if results are returned append the results to the matches slice
					if match := search(word, puzzle, r, c, dir.dr, dir.dc); match != "" {
						matches = append(matches, fmt.Sprintf("(%d, %d) %s", r+1, c+1, dir.arrow))
					}
				}
			}
		}
	}
	// return the result(s) in the matches slice
	return matches
} // end WordSearch

// search checks if a word exists in the puzzle in a given direction
// requires string, [][]rune, int, int, int, int
// returns string
func search(word string, puzzle [][]rune, r, c, dr, dc int) string {

	// store row and column length of the puzzle
	rows := len(puzzle)
	cols := len(puzzle[0])

	// convert word to slice of runes
	runes := []rune(word)

	// store the length of runes, the searched word
	wordLen := len(runes)

	// check if the searched word can fit in the given direction within the puzzle
	if r+dr*(wordLen-1) >= 0 && r+dr*(wordLen-1) < rows && c+dc*(wordLen-1) >= 0 && c+dc*(wordLen-1) < cols {

		// loop to check if the runes of the searched word exist in the given direction
		for i := 0; i < wordLen; i++ {
			if puzzle[r+i*dr][c+i*dc] != runes[i] {
				return "" // return an empty string if not matched
			}
		}
		// return the searched word if loop completes, match was found
		return word
	}

	// return an empty string if conditions not met
	return ""
} // end search

// lexicoSort sorts a slice of matches lexicographically based on RCD (row, column, direction)
func lexicoSort(matches []string) {
	n := len(matches)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			r1, c1, dr1 := parseRCD(matches[j])
			r2, c2, dr2 := parseRCD(matches[j+1])

			if compareRCD(r1, c1, dr1, r2, c2, dr2) > 0 {
				matches[j], matches[j+1] = matches[j+1], matches[j]
			}
		}
	}
} // end lexicoSort

// parseRCD extracts the RCD components from the matches
func parseRCD(s string) (int, int, string) {
	var r, c int
	var direction string
	_, err := fmt.Sscanf(s, "(%d, %d) %s", &r, &c, &direction)
	if err != nil {
		fmt.Printf("Error parsing match: %v", err)
	}
	return r, c, direction
} // end parseRCD

// compareRCD compares two RCD components lexicographically
func compareRCD(r1, c1 int, dr1 string, r2, c2 int, dr2 string) int {
	if r1 != r2 {
		return r1 - r2
	}
	if c1 != c2 {
		return c1 - c2
	}
	return strings.Compare(dr1, dr2)
} // end compareRCD
