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

func fetch(urls []string, pages chan<- string) {
	for _, url := range urls {
		// Pretend to fetch the page
		fmt.Printf("Fetching %s\n", url)
		pages <- url
	}
}

func parse(pages <-chan string) {
	for {
		// Pretend to parse the page
		url := <-pages
		fmt.Printf("Parsing %s\n", url)
		sleepTime := time.Duration(rand.Intn(2000)) * time.Millisecond
		time.Sleep(sleepTime)
	}
}
