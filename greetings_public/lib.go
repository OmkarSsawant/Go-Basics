package greetingspublic

import (
	"errors"
	"fmt"
)

func Greet(name string) (string, error) {
	if name == "" {
		return "", errors.New("No Name Specified")
	}
	message := fmt.Sprintf("Hello %v, get started with Go,Go is way to GO FUTURE", name)
	return message, nil
}
