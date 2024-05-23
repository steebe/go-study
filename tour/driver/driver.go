package main

import (
	"fmt"

	"steebe.dev/tour/functions"
	"steebe.dev/tour/pointers"
	"steebe.dev/tour/slices"
	"steebe.dev/tour/structs"
)

func main() {
	fmt.Println("- Functions...")
	functions.DeferPrint()

	fmt.Println("- Pointers...")
	pointers.IntPointers()

	fmt.Println(" ----- ")

	fmt.Println("- Structs...")
	structs.VertexStruct()

	fmt.Println(" ----- ")

	fmt.Println("- Slices...")
	slices.DynamicSlice()
}

/*
	Basic Types

	bool
	string
	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr
	byte // alias for uint8
	rune // alias for int32 & represents a Unicode code point
	float32 float64
	complex64 complex128

	The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems.
	When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.
*/

/*
	Zero Values

	Variables declared without an explicit initial value are given their zero value.

	The zero value is:

    0 for numeric types,
    false for the boolean type, and
    "" (the empty string) for strings.

*/
