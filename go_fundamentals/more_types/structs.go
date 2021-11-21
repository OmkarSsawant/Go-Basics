package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

//struct literals
var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	ver := Vertex{1, 2}
	fmt.Println(ver)
	fmt.Println(ver.X, ver.Y)

	//pointer to structs
	verp := &ver
	fmt.Println(verp, *verp)
	fmt.Println(verp.X, verp.Y, (*verp).X, (*verp).Y) //same deref

	fmt.Println(v1, p, v2, v3)
}
