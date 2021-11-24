package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 8; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world") //what it esentially does is cpu switches between light-weight threads(go-routines)
	say("hello")
	// time.Sleep(200 * time.Millisecond) gets anticipated result

}
