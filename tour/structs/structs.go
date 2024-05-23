package structs

import "fmt"

type Node struct {
	Payload int
	Next    *Node // A struct cannot contain itself, but can contain a reference to itself
}

type Vertex struct {
	X, Y int
}

func VertexStruct() {
	// Once a struct is defined, it can be used like any other type
	point := Vertex{1, 2}
	fmt.Printf(" - Vertex {5, 7}: %v\n", point)

	// Referencing fields in a struct occurs with the `.` operator
	fmt.Printf(" - Vertex X: %d\n", point.X)

	// Go automatically dereferences struct field values if accessed thru a pointer
	p := &point
	fmt.Printf(" - Vertex Y via pointer dereference: %d\n", p.Y)

	// The `:` operator can be used to assign default values in struct instantiation
	point = Vertex{X: -1}
	fmt.Printf(" - New X val: %d\n", point.X)
}
