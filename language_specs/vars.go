package main

import "fmt"

type AllTypes struct {
	i    int8
	i1   int16
	i2   int32
	i3   int64
	f    float32
	f1   float64
	b    bool
	str  string
	char rune
}

func main() {
	var obj AllTypes
	fmt.Println(obj)

}
