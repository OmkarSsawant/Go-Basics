package main

// way 1 to import
import (
	"fmt"
	"math"
)

// way 2 to import
// import "fmt"
// import "math"

func main() {
	// error as var/func starts with nonCAP name ie, unexported
	// fmt.Println("The pie value is  %f",math.pi)
	// no-error as var/func starts with CAP name ie, exported
	fmt.Println(math.Pi)

}
