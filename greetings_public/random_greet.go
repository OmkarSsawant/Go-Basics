package greetingspublic

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomGreet(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty Name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil

}

var formats []string = []string{
	"Hello %v , From Go",
	"Hello %v , From Go Creators",
	"Hello %v from Go API creator",
	"Hello %v, get started with Go,Go is way to GO FUTURE",
}

var formats_len = len(formats)

func randomFormat() string {
	return formats[rand.Intn(formats_len)]
}

func RandomGreets(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		greet, er := RandomGreet(name)
		if er != nil {
			log.Fatal(er)
		}
		messages[name] = greet
	}
	return messages, nil
}
