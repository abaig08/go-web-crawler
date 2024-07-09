package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

var wg sync.WaitGroup

func fetch(url string) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("URL: %s\n%s\n", url, body)
}

func main() {
	var urls []string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter URLs (one per line), and press Enter twice to finish:")

	// Read URLs from user input
	for scanner.Scan() {
		url := scanner.Text()
		if url == "" {
			break
		}
		urls = append(urls, url)
	}

	if len(urls) == 0 {
		fmt.Println("No URLs provided.")
		return
	}

	// Fetch each URL concurrently
	for _, url := range urls {
		wg.Add(1)
		go fetch(url)
	}
	wg.Wait()
}
