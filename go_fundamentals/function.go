package main

import "fmt"

// func function-name(var1 type,var2 type,...) return-type{
// defination
// }

// a int, b int is same as a,b int
func add(a, b int) int {
	return a + b
}

func sampleMultipleReturns(a, b string) (string, string, int) {
	return b, a, len(a) + len(b)
}

func nakedReturn(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func main() {
	fmt.Println(add(2, 3))
	a, b, combinedLength := sampleMultipleReturns("a-str", "b-str")
	fmt.Println(a, b, combinedLength)
	n1, n2 := nakedReturn(20)
	fmt.Println(n1, n2)

}
