package main

import (
	"fmt"
	greetingspublic "greetings_public"
	"log"
)

func main() {
	names := []string{"Gladys", "Samantha", "Darrin"}

	log.SetPrefix("greet:")
	log.SetFlags(0)
	// message, er := greetingspublic.RandomGreet("Omkar")
	randomGreets, er := greetingspublic.RandomGreets(names)
	if er != nil {
		log.Fatal(er)
	}

	fmt.Println(randomGreets)
	// log.Println(greetingspublic.RandomGreets(names))
}
