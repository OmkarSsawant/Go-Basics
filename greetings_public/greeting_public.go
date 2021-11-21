package greetingspublic

import (
	"errors"
	"fmt"
)

func Greet(name string) (string, error) {

	if name == "" {
		return "", errors.New("no name specified")
	}
	//to fail test
	// message := fmt.Sprint("Hello %v, get started with Go,Go is way to GO FUTURE")

	message := fmt.Sprintf("Hello %v, get started with Go,Go is way to GO FUTURE", name)
	return message, nil
}
