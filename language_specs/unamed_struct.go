package main

import (
	"fmt"
)

type class struct {
	name       string
	annotaions []string
	methods    []string
	variables  []string
}

func (c *class) printName() {
	fmt.Println(c.name)
}

type Unamed struct {
	x, y int
	class
}

func main() {
	var us Unamed
	us.name = "Omkar"
	us.printName()
}
