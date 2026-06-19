package gt

import "testing"

func TestAdd(t *testing.T) {
	type testCase struct {
		name             string
		arg1, arg2, want int
	}
	var testCases = []testCase{
		{"basic", 2, 3, 5},
		{"large", 400, 800, 1200},
		{"negative", -3, 10, 7},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if output := Add(tt.arg1, tt.arg2); output != tt.want {
				t.Errorf("Output %q not equal to %q", output, tt.want)
			}
		})
	}
}

func BenchmarkFib(b *testing.B) {
	type testCase struct {
		name string
		arg1 int
	}
	var testCases = []testCase{
		{"basic", 2},
		{"large", 10},
		{"very large", 40},
	}
	for _, tt := range testCases {
		b.Run(tt.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				Fib(tt.arg1)
			}
		})
	}
}
