package main

import (
	"fmt"
	"math"
)

func main() {
	// goLang := "future"
	// if goLang == "future" {
	// 	fmt.Println("Hello Dev")
	// }

	// additionally varial required can be declared in
	// if scope

	if goLang := "future"; goLang == "future" {
		fmt.Println("Hello Dev")
	}
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}
