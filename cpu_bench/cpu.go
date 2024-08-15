package cpu_bench

import (
	"golang.org/x/crypto/bcrypt"
	"sync"
)

func seq() {
	for i := 0; i < 10; i++ {
		bcrypt.GenerateFromPassword([]byte("test"), 10)
	}
}
func parallel() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			bcrypt.GenerateFromPassword([]byte("test"), 10)
		}()
	}
	wg.Wait()
}
func limitParallel(n int) {
	ch := make(chan bool, n)
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for range ch {
				bcrypt.GenerateFromPassword([]byte("test"), 10)
			}
		}()
	}
	for i := 0; i < 10; i++ {
		ch <- true
	}
	close(ch)
	wg.Wait()
}
