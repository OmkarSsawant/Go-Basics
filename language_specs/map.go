package main

import "fmt"

func main() {
	future := make(map[string]string, 2)

	future[`Lang`] = "Go"
	future[`AI`] = "Upcoming"
	future[`UI`] = "Flutter"

	fmt.Println(future, len(future))
}
