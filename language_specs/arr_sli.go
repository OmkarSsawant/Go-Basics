package main

import "fmt"

type any interface{}

func main() {

	arr := []int{7, 6, 5, 4, 3, 2, 1}
	sli := arr[1:5]
	sli[0] = -1 //same as arr[1] = -1
	printSlice(&sli)
	fmt.Println("arr : ", arr, ` I was Changed because 
			slice(sli) takes my storage reference`)

	var unarr [3]any
	fmt.Println(unarr) //nil , except for primitives

	sli = append(sli, 3, 8, 9)
	printSlice(&sli) //doubles the capacity when exceded

	//manually create slice
	mSlice := make([]int, 4, 6)
	//same as
	mSlice = new([6]int)[:4]
	printSlice(&mSlice)

	//can be multidimentional
	mulSlice := make([][]int, 6)
	fmt.Println(mulSlice)

	//to create (6 x 6 matrix)
	for i := 0; i < 6; i++ {
		//*inside values must be intialized individually
		mulSlice[i] = make([]int, 6)
		for j := 0; j < 6; j++ {
			mulSlice[i][j] = 6
		}
	}
	fmt.Println(mulSlice)

}

func printSlice(s *[]int) {
	fmt.Println("s : ", s, " cap: ", cap(*s), " len: ", len(*s))
}
