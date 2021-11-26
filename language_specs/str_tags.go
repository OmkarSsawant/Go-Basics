package main

import "fmt"

type Location struct {
	x, y float64 ""                                     // an empty tag string is like an absent tag
	name string  "any string is permitted as a tag"     //take part in reflection tag can be fecthed with field name
	_    [4]byte "ceci n'est pas un champ de structure" //otherwise ignored
}

//protocal buffer struct serializable
type VideoBuff struct {
	size int    `protobuf:"1"`
	buff []byte `protobuf:"2"`
}

//json struct serializable
type ApiResponce struct {
	Count    int      `json:"count"`
	Articles []string `json:"articles"`
}

func main() {
	fmt.Println(ApiResponce{}, VideoBuff{}, Location{})
}
