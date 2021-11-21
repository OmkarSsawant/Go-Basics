package main

import (
	"fmt"
)

//defer stack is like event loop
//just diff is LIFO principle
func stackDiffered() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

func main() {
	fmt.Println("before")

	// defer log.Println("middle-differed")
	stackDiffered()

	fmt.Println("after")
}
