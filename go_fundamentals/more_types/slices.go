package main

import (
	"fmt"
	"strings"
)

type coordinate = struct{ x, y int }

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	//SLICE

	//slice is part of array and is refrenced to original and mutable
	//Changing the elements of a slice modifies the corresponding elements of its underlying array
	var ps []int = primes[1:4]
	fmt.Println(ps)
	ps[0] = 100
	fmt.Println(primes)

	//Literals
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	const start = 3
	const end = 5
	//if not specified start will be 0
	fmt.Println(s[:end])
	//if not specified end will be len(s)
	fmt.Println(s[start:])
	//no one specified will be whole
	fmt.Println(s[:])

	sl := []int{2, 3, 5, 7, 11, 13}
	printSlice(sl)

	// Slice the slice to give it zero length.
	sl = sl[:0]
	printSlice(sl)

	// Extend its length.
	sl = sl[:4]
	printSlice(sl)

	// Drop its first two values.
	sl = sl[2:]
	printSlice(sl)

	//with make

	a := make([]int, 5)
	printSliceMade("a", a)

	b := make([]int, 0, 5)
	printSliceMade("b", b)

	c := b[:2]
	printSliceMade("c", c)

	d := c[2:5]
	printSliceMade("d", d)

	//slices of slices
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	//append
	var s3 []int
	printSlice(s3)

	// append wor3ks on nil slices.
	s3 = append(s3, 0)
	printSlice(s3)

	// The slice grows as needed.
	s3 = append(s3, 1)
	printSlice(s3)

	// We can add more than one element at a time.
	s3 = append(s3, 2, 3, 4)
	printSlice(s3)

	//range
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
func printSlice(s []int) {
	str := fmt.Sprintf(" len=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Println(str)
}

func printSliceMade(name string, s []int) {
	str := fmt.Sprintf("%v len=%d cap=%d %v\n", name, len(s), cap(s), s)
	fmt.Println(str)
}
