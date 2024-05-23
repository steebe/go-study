package pointers

import "fmt"

func IntPointers() {
	i, j := 42, 2701

	// `&` operator performs the point to a memory address
	// This makes its assignee a pointer, and this is called "referencing"
	p := &i
	fmt.Printf(" - Memory Address of a pointer (p) with simple reference by pointer var name: %v\n", p)
	fmt.Printf(" - Value that a pointer (p) is pointing to via `*`: %v\n", *p) // read i through the pointer

	// `*` operator performs a reference to the memory address that the pointer points to
	// It can be used to set the underlying value stored at the memory address
	// This is called "dereferencing" or "indirecting"
	*p = 21
	fmt.Printf(" - Updated value of initial variable (i) that the pointer manipulated: %d\n", i)

	// Pointers can be reassigned to other addresses
	p = &j
	*p = *p / 37
	fmt.Printf(" - Updated address and value of pointer (p) after reassignment: %d\n", j)
}
