package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func retrieve(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	resp, err := http.Get(url)
	end := time.Since(start)
	if err != nil {
		panic(err)
	}
	fmt.Println(url, resp.StatusCode, end)
}

func main() {
	var wg sync.WaitGroup
	var urls = []string{"https://godoc.org", "https://www.packtpub.com", "https://kubernetes.io/"}
	for i := range urls {
		// WaitGroup Counter++ when new goroutine is called
		wg.Add(1)
		go retrieve(urls[i], &wg)
	}
	// Wait for the collection of goroutines to finish
	wg.Wait()
}
