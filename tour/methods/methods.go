package methods

type Node struct {
	payload int
	next    *Node
}

// This is a method, because it contains a receiver argument before its stub
// Calling this method can only occur from an instance of Node
// If Node were not part of the package containing the method, the method would not compile
func (n Node) Insert(index int, payload int) {
	// TODO:
}

// Methods do not only apply to structs
// They can be created to support any custom type
type MyInt int

func (i MyInt) Abs() int {
	if i < 0 {
		return int(i) * -1
	}
	return int(i)
}

type Vertex struct {
	X, Y int
}

// Pointer receivers are receivers that are capable of adjusting the caller's state.
// These are heaviliy important, as methods operating on a pointer receiver perform
// mutations on the actual instance of an object in memory, rather than require the calling
// code to perform additional behavior after receiving the method's retval to mutate the instance.
func (v *Vertex) Scale(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}
