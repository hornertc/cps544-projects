package main

type DayOfWeek int
type Week int
type Year int

func date(dow DayOfWeek, wk Week, yr Year) string {
	return ""
}

func main() {
	var dow DayOfWeek = 2
	var wk Week = 5
	var yr Year = 2023
	wk = wk + 1        // valid
	wk = wk + Week(10) // valid
	yr = yr + dow      // invalid: mismatched types
	date(dow, wk, yr)  // valid
	date(dow, wk, wk)  // invalid!
}
