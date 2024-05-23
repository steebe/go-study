package functions

import "fmt"

// Returns a closure that yields the next number in a Fibonacci sequence
func Fibonacci() func() int {
	lastValue := 0
	value := 1

	return func() int {
		temp := value
		value = lastValue + value
		lastValue = temp
		return value
	}
}

// Multiple retvals
// Naked return
func SumAndProduct(a, b int) (sum, product int) {
	// These are assignments, not short assignments, since the return declares them
	sum = a + b
	product = a * b
	return
}

func DeferPrint() {
	defer fmt.Println("^Deferred execution Stack^")

	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Me first")
}
