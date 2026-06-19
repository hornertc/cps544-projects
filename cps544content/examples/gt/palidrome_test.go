package gt

import (
	"reflect"
	"testing"
)

func FuzzIsPalidrome(f *testing.F) {
	f.Add("detartrated")
	f.Add("kayak")
	f.Add("palindrome")
	// f.Add("été")
	// f.Add("A man, a plan, a canal: Panama")
	// f.Add("Et se resservir, ivresse reste.")

	f.Fuzz(func(t *testing.T, a string) {
		s := []rune(a)
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
		truth := reflect.DeepEqual([]rune(a), s)

		retval := IsPalindrome(a)
		if truth != retval {
			t.Fail()
		}
	})
}

// go test ./examples/gt -fuzz=. -fuzztime=1m
