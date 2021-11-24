package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//can only define methods whose type is defined
// in same package
type String string

func (str String) println() {
	fmt.Println(str)
}

func main() {
	var s String = "Hello"
	s.println()

	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
