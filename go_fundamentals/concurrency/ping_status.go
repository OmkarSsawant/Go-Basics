package main

import (
	"log"
	"net/http"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {
	defer waitGroup.Wait()
	log.SetFlags(0)
	urls := []string{
		"https://www.google.com",
		"https://www.fb.com",
		"https://www.instagram.com",
	}
	for _, url := range urls {
		go status(url)
		waitGroup.Add(1)
	}
	// time.Sleep(5 * time.Second)
}

func status(web string) {

	defer waitGroup.Done()
	res, err := http.Get(web)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(web, res.Status)

}
