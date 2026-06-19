package main

import "fmt"

func dump(prefix string, x []int) {
	fmt.Println(prefix, "len:", len(x), "cap:", cap(x), "slice:", x, "array:", x[:cap(x)], "location:", &x[0], &x[1])
}

func main() {
	{
		x := []int{5, 6} // len: 2, cap: 2
		dump("A1", x)

		x1 := append(x, 7, 8) // allocation
		dump("A2", x)
		dump("A3", x1)
	}

	{
		x := make([]int, 2, 4) // len: 2, cap: 4 to avoid memory allocations below
		x[0] = 5
		x[1] = 6
		dump("B1", x)

		x1 := append(x, 7, 8) // no allocations
		dump("B2", x)
		dump("B3", x1)
	}

	{
		insert := []int{10, 11}
		x := []int{4, 5, 6, 7} // len: 4, cap: 4
		dump("C1", x)

		// compile error
		// x = append(x[:2], 10, 11, x[2:]...) // 4, 5, 10, 11, 6, 7

		// insert 10 and 11 after the 5 (before the 6)
		y := append(insert, x[2:]...) // 10, 11, 6, 7
		dump("C2", y)
		x = append(x[:2], y...) // 4, 5, 10, 11, 6, 7
		dump("C3", x)
	}

	{
		insert := []int{10, 11}
		x := []int{4, 5, 6, 7}

		y := make([]int, 6)

		// no allocations
		copy(y[:2], x[:2])
		dump("D1", y)

		copy(y[2:4], insert)
		dump("D2", y)

		copy(y[4:], x[2:])
		dump("D3", y)
	}

	// copy overlapping
	{
		x := []int{10, 11, 12, 13, 14, 15}
		copy(x[2:5], x[1:4])
		dump("Copied", x)
	}

	{
		x := []int{10, 11, 12, 13, 14, 15}
		naiveCopy(x[2:5], x[1:4])
		dump("Copied (naive)", x)
	}
}

// does not handle overlap properly
func naiveCopy(dst, src []int) {
	// ignore mismatch lengths
	for i := range src {
		dst[i] = src[i]
	}
}
