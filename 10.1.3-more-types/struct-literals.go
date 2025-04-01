package main

import "fmt"

type Vertex struct {
	X, Y int
}


var (
	v1 = Vertex{1, 2} 	// has type Vertex
	v2 = Vertex{X: 1}	// Y:0 is implied as it is the default value for an int
	v3 = Vertex{}		// X & Y == 0 because defaults yo'
	p = &Vertex{1, 2}	// has type "pointer to vertex"  *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}