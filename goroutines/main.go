package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func greeter(s string) {
	for i := 0; i < 4; i++ {
		fmt.Println(s)
		time.Sleep(3 * time.Millisecond)
	}
}

func getStatusCode(endPoint string) {
	res, err := http.Get(endPoint)

	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Printf("%d status code for %s\n", res.StatusCode, endPoint)

	defer wg.Done()
}

func main() {
	// A goroutine is a lightweight thread of execution.
	fmt.Println("Goroutines in Go")

	go greeter("Hello")
	greeter("World")

	website := []string{
		"https://github.com/A158-debug",
		"https://google.com",
		"https://facebook.com",
		"https://wikipedia.org",
		"https://amazon.com",
	}
    
	// The sync package provides a way to group goroutines and wait for them to finish.
	for _, web := range website {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()

}
