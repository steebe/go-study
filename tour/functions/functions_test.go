package functions

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	seq := make(map[int]int)
	seq[0] = 1
	seq[1] = 2
	seq[2] = 3
	seq[3] = 5
	seq[4] = 8
	seq[5] = 13
	seq[6] = 21
	seq[7] = 34
	seq[8] = 55
	seq[9] = 89

	f := Fibonacci()

	for i := 0; i < len(seq); i++ {
		result := f()
		if result != seq[i] {
			t.Fatalf("Expected %v, got %v", seq[i], result)
		}
	}
}
