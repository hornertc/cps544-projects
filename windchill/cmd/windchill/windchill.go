// CPS 544: Assignment 4, Wind Chill
// Go Program that computes wind chill temperature based on specific user input
// Tommy Horner, thorner1

// package main
package main

// imports
import (
	"errors"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

var tempInput string
var speedInput string
var table bool

// start main
func main() {

	flag.BoolVar(&table, "table", false, "Output as a table")

	flag.Parse()

	if table {
		displayTable()
		os.Exit(0)
	}

	// Create/open a log file for error logging
	logFile, err := os.Create("error_log.txt")
	if err != nil {
		fmt.Printf("Error opening log file: %v", err)
		os.Exit(1)
	}
	defer logFile.Close()

	// Set the log output to the file
	logOutput := logFile

	temp, speed := GetInput(os.Stdin, logOutput)

	// Loop1:
	// 	for {

	// 		fmt.Print("Enter the temperature in °F: ")
	// 		_, err := fmt.Scanln(&tempInput)
	// 		if err != nil {
	// 			fmt.Println("Error:", err)
	// 		}
	// 		temp, err := strconv.ParseFloat(tempInput, 64)
	// 		if err != nil {
	// 			fmt.Println("Error:", err)
	// 			goto Loop1
	// 		}
	// 		if temp < -58 || temp > 48 {
	// 			fmt.Print(temp, " is out of range. Valid temperatures are between -58°F and 41°F.\n")
	// 		} else {
	// 			break
	// 		}
	// 	}

	// Loop2:
	// 	for {
	// 		fmt.Print("Enter the wind speed in miles per hour: ")
	// 		_, err := fmt.Scanln(&speedInput)
	// 		if err != nil {
	// 			fmt.Println("Error:", err)
	// 		}
	// 		speed, err := strconv.ParseFloat(speedInput, 64)
	// 		if err != nil {
	// 			fmt.Println("Error:", err)
	// 			goto Loop2
	// 		}
	// 		if speed < 2 {
	// 			fmt.Print(speed, " is out of range. Valid wind speeds are above 2 mph.\n")
	// 		} else {
	// 			break
	// 		}
	// 	}

	// 	temp, err := strconv.ParseFloat(tempInput, 64)
	// 	if err != nil {
	// 		fmt.Println("Error:", err)
	// 		return
	// 	}

	// 	speed, err := strconv.ParseFloat(speedInput, 64)
	// 	if err != nil {
	// 		fmt.Println("Error:", err)
	// 		return
	// 	}

	fmt.Printf("The wind chill index is %.2f°F\n", windchill(temp, speed))

} // end main

// GetInput is a function that retrieves user input
func GetInput(reader io.Reader, logOutput io.Writer) (float64, float64) {
	for {

		fmt.Print("Enter the temperature in °F: ")
		_, err := fmt.Fscanln(reader, &tempInput)
		if errors.Is(err, io.EOF) {
			return 0, 0
		} else if err != nil {
			fmt.Println("Error:", err)
			// _, err := fmt.Fprintln(logOutput, "Error:", err)
			// if err != nil {
			// 	fmt.Println("Error:", err)
			// }
			continue
		}
		temp, err := strconv.ParseFloat(tempInput, 64)
		if err != nil {
			fmt.Println("Error:", err)
			// _, err := fmt.Fprintln(logOutput, "Error:", err)
			// if err != nil {
			// 	fmt.Println("Error:", err)
			// }
			continue
		}
		if temp < -58 || temp > 48 {
			fmt.Print(temp, " is out of range. Valid temperatures are between -58°F and 41°F.\n")
			_, err := fmt.Fprint(logOutput, temp, " is out of range. Valid temperatures are between -58°F and 41°F.\n")
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			break
		}
	}

	for {
		fmt.Print("Enter the wind speed in miles per hour: ")
		_, err := fmt.Fscanln(reader, &speedInput)
		if errors.Is(err, io.EOF) {
			return 0, 0
		} else if err != nil {
			fmt.Println("Error:", err)
			// _, err := fmt.Fprintln(logOutput, "Error:", err)
			// if err != nil {
			// 	fmt.Println("Error:", err)
			// }
			continue
		}
		speed, err := strconv.ParseFloat(speedInput, 64)
		if err != nil {
			fmt.Println("Error:", err)
			// _, err := fmt.Fprintln(logOutput, "Error:", err)
			// if err != nil {
			// 	fmt.Println("Error:", err)
			// }
			continue
		}
		if speed < 2 {
			fmt.Print(speed, " is out of range. Valid wind speeds are above 2 mph.\n")
			// _, err := fmt.Fprint(logOutput, speed, " is out of range. Valid wind speeds are above 2 mph.\n")
			// if err != nil {
			// 	fmt.Println("Error:", err)
			// }
		} else {
			break
		}
	}

	temp, err := strconv.ParseFloat(tempInput, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return 0, 0
	}

	speed, err := strconv.ParseFloat(speedInput, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return 0, 0
	}

	return temp, speed
}

func windchill(temp float64, speed float64) float64 {

	chill := (35.74 + (0.6215 * temp) - (35.75 * math.Pow(speed, 0.16)) + (0.4275 * temp * math.Pow(speed, 0.16)))

	return chill
}

func displayTable() {

	data := make([]string, 216)
	i := 0
	for x := 40.00; x >= -45.00; x -= 5.00 {
		for y := 5.00; y <= 60.00; y += 5.00 {
			data[i] = fmt.Sprintf("%.1f", windchill(x, y))
			i++
		}
	}

	table := make([][]string, 12)

	for row := 0; row < 12; row++ {
		dataIndex := row
		table[row] = make([]string, 18)
		for col := 0; col < 18; col++ {
			table[row][col] = data[dataIndex]
			dataIndex += 12
		}
	}

	for _, value := range table {
		fmt.Println(strings.Join(value, "\t"))
	}
}
