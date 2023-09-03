package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.yahoo.com",
		"https://www.bing.com",
	}
	pages := make(chan string)

	go fetch(urls, pages)
	go parse(pages)

	fmt.Println("Press ENTER to exit")
	var input string
	fmt.Scanln(&input)
}

// Fetch the pages
func fetch(urls []string, pages chan<- string) {
	for _, url := range urls {
		fmt.Printf("Fetching %s\n", url)
		pages <- url
		fakeProcessing()
	}
}

// Pretend to parse the pages
func parse(pages <-chan string) {
	for {
		// Pretend to parse the page
		url := <-pages
		fmt.Printf("Parsing %s\n", url)
		fakeProcessing()
	}
}

// Add a rendom delay to simulate processing
func fakeProcessing() {
	sleepTime := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(sleepTime)
}
