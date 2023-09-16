## Channels

Channels are the Go way for goroutines to **communicate** with one another, and to synchronize their work.

As Rob Pike said: "don't let computations communicate by sharing memory, let them share memory by communicating".

The communication can be uni- or bi-directional. Assuming `T` is the data type of the data passed through the channel:

- `chan T` is a bidirectional channel
- `chan<- T` is a unidirectional channel, send-only
- `<-chan T` is a unidirectional channel, receive-only

### Fake web crawler

Let's implement a fake version of the web crawler we mentionned earlier. We first make a channel called `pages`. It will be used by our two workers. We shall separate fetching from parsing a page. The program terminates when the user presses Enter. `fetch()` and `parse()` are invoked as goroutines (with `go`).

We pass the channel as a parameter. Notice the slight difference in declaration:

- `chan string`, bidirectional channel in `main`
- `chan<- string`, unidirectional channel, send-only in `fetch`
- `<-chan string`, unidirectional channel, receive-only in `parse`

The bidirectional channel declared in `main` can be "specialized" into a send-only or receive-only channel when passed as an argument. This cannot be done the other way around; you cannot turn you unidirectional channel into a bidirectional channel.

```go
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
```
