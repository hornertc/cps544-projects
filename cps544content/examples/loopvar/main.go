package main

import "fmt"

func main() {
	items := []string{"a", "b", "c"}
	n := len(items)

	iRefs := make([]*int, n)
	itemRefs := make([]*string, n)
	for i, item := range items {
		// i := i
		// item := item
		iRefs[i] = &i
		itemRefs[i] = &item
		fmt.Println(i, item)
	}
	fmt.Println("After loop")
	for i := 0; i < n; i++ {
		fmt.Println(*iRefs[i], *itemRefs[i])
	}

	fmt.Println("Because...")
	fmt.Println(iRefs)
	fmt.Println(itemRefs)

	explanation(items)
}

func explanation(items []string) {
	n := len(items)

	iRefs := make([]*int, n)
	itemRefs := make([]*string, n)

	{
		var i int
		var item string
		for j := 0; j < n; j++ {
			i = j
			item = items[j]
			// start explicit loop body
			iRefs[i] = &i
			itemRefs[i] = &item
			fmt.Println(i, item)
			// end explicit loop body
		}
	}
}
