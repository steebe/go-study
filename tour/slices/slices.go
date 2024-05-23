package slices

import "fmt"

func DynamicSlice() {

	// When instantiating a slice, an array is created to reference
	// I.E: array {2, 3, 5, 7, 11, 13} is implicitly generated here
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slicing a slice to its zeroeth index makes its _length_ zero
	// However, its capacity is still the same as before this slicing,
	// as the underlying array is still unaffected so far
	s = s[:0]
	printSlice(s)

	// Extending a slice that's been "zeroed" will restore the values
	// originally present in the array at the indices restored to the slice
	s = s[:4]
	printSlice(s)

	// Slicing the beginning off of a slice actually results in loss
	// of capacity
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf(" - Slice state: len=%d cap=%d %v\n", len(s), cap(s), s)
}
