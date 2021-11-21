package main

import "fmt"

func main() {
	for i := 0; i < 120; i++ {
		fmt.Println(i)
	}
	// init and post statements are optional
	// which acts similar to C while
	requestThreshold := 500
	for requestThreshold > 0 {
		fmt.Println("remaining Requests", requestThreshold)
		requestThreshold--
	}
	//infinite loop be  like
	for {
		// break
	}
}
