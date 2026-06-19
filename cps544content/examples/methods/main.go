package main

import "fmt"

type Meal []string

func (f Meal) Healthy() bool {
	for _, food := range f {
		switch food {
		case "carrot", "broccoli":
			return true
		}
	}
	return false
}

func main() {
	lunch := Meal([]string{"candy", "carrot"})
	fmt.Println(lunch.Healthy()) // true

	f := Meal.Healthy
}
