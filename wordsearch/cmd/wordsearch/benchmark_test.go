// package
package main

// imports
import (
	"bufio"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// BenchmarkWordSearch is a benchmark test for the function WordSearch in wordsearch.go
func BenchmarkWordSearch(b *testing.B) {
	// puzzles to test, comment/uncomment to select a puzzle
	// puzzleFile := "puzzle1.txt" // puzzle1.txt
	puzzleFile := "puzzle2.txt" // puzzle2.txt

	// variable to store the path to the puzzle file
	puzzlePath := filepath.Join("C:/Users/thorner1/Desktop/MCSClasses/CPS544/GitHub/wordsearch-thorner1/", puzzleFile)

	// open the puzzle file
	// print error if the file cannot be opened
	file, err := os.Open(puzzlePath)
	if err != nil {
		b.Fatalf("Error opening puzzle file: %v", err)
	}
	defer file.Close() // close the file on exit

	// read and store the puzzle file contents
	var puzzle [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := []rune(scanner.Text())
		puzzle = append(puzzle, row)
	}

	// print error message if the file cannot be read
	if err := scanner.Err(); err != nil {
		b.Fatalf("Error reading puzzle file: %v", err)
	}

	// specific words to search for each puzzle file
	// searchWords := []string{"œke", "dŮp", "tsrq", "abcdefghijklmop"} // puzzle1.txt
	searchWords := []string{"x", "fx", "cxx1", "abcdefghijklmop"} // puzzle2.txt

	// shuffle the order of 'searchWords' to get an unbiased benchmark test
	shuffleSearchWords(searchWords)

	b.ReportAllocs() // report memory allocation
	b.ResetTimer()   // reset benchmark time, removes setup time from results

	// loop to run the WordSearch function from wordsearch.go using 'searchWords' and 'puzzle'
	// creates a sub-benchmark for each word in 'searchWords'
	for _, wordToSearch := range searchWords {
		b.Run(wordToSearch, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				WordSearch(wordToSearch, puzzle)
			}
		})
	}
} // end BenchmarkWordSearch

// function to shuffle the order of 'searchWords' for an unbiased benchmark test
func shuffleSearchWords(words []string) {

	// create a new rand.Source to generate a random order each call
	source := rand.NewSource(time.Now().UnixNano())

	// create a new rand.Rand instance using 'source'
	r := rand.New(source)

	// shuffle the order of 'words', the search words passed into the function
	r.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
} // end shuffleSearchWords
