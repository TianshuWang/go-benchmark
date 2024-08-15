package io_bench

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func BenchmarkSyncRead(b *testing.B) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Millisecond)
		w.WriteHeader(200)
	}))
	defer srv.Close()
	urls := generateLinks(srv.URL)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		syncCrawl(urls)
	}
}
func BenchmarkConcurrentRead(b *testing.B) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Millisecond)
		w.WriteHeader(200)
	}))
	defer srv.Close()
	urls := generateLinks(srv.URL)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		concurrentCrawl(urls)
	}
}
