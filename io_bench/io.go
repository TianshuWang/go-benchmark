package io_bench

import (
	"net/http"
	"sync"
)

func generateLinks(url string) []string {
	urls := make([]string, 10)
	for i := 0; i < 10; i++ {
		urls[i] = url
	}
	return urls
}
func syncCrawl(urls []string) {
	for _, url := range urls {
		http.Get(url)
	}
}
func concurrentCrawl(urls []string) {
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		go func(url string) {
			defer wg.Done()
			http.Get(url)
		}(url)
	}
	wg.Wait()
}
