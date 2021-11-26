package main

import "fmt"

type T struct {
	x, y int
}

type X struct {
	a, b int
}

//define method for XI
func (xi X) Sample() string {
	return fmt.Sprintf("Sample Method of %T", xi)
}

type TI interface {
	Sample() string
}

func main() {

	var x X = X{}
	var t T = T{}
	// t = x invalid as one of it should not be defined
	fmt.Println(x, t)

	//point  1 //same underlying type , so valid

	//x has method Sample() string , hence implements TI
	// T is an interface type and x implements T.
	var xi X = X{}
	var ti TI = xi

	fmt.Println(ti.Sample())

	//channel
	/*
		xc is a bidirectional channel value, T is a channel type,
		 x's type V and T have identical element types,
		and at least one of V or T is not a defined type.
	*/
	xc := make(chan T)
	var tc chan X
	// tc = xc invalid -> both are defined types
	fmt.Println(tc)
	close(xc)

	//nil is assignable to all

	var ut T
	/*
		x's type V and T have identical underlying types and
		at least one of V or T is not a defined type.
	*/
	//x(12) is an untyped constant representable by a value of type T
	ut = T{}
	fmt.Println(ut)
}
