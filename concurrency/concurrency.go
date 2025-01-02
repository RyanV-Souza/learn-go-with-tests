package concurrency

import (
	"fmt"
)

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func() {
			fmt.Println("Checking:", url)
			resultChannel <- result{url, wc(url)}
		}()
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		fmt.Printf("Result for %s: %t\n", r.string, r.bool)
		results[r.string] = r.bool
	}

	return results
}