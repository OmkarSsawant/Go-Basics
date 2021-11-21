package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	// short var declaration to implicitly declare type (only inside function scope)
	kotlin := "similar"
	fmt.Println(i, j, c, python, java, kotlin)
}
