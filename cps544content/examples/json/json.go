package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	value := []any{"car", 4, 4.5,
		map[string]int{
			"a": 5,
			"b": 100,
		},
	}

	data, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	fmt.Println("JSON: ", string(data))

	dataPretty, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println("JSON (pretty):")
	fmt.Println(string(dataPretty))

	var newValue []any
	err = json.Unmarshal(data, &newValue)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\nint is %T\nfloat is %T\nobject is %T\n", newValue, newValue[1], newValue[2], newValue[3])
}
