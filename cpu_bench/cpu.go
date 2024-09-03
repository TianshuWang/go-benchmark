package cpu_bench

import (
	"fmt"
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

func sendData(ch chan<- bool) {
	for i := 0; i < 10; i++ {
		ch <- true
	}
	close(ch)
}

func receiveData(wg *sync.WaitGroup, threadId int, ch <-chan bool) {
	//defer wg.Done()
	for data := range ch {
		bcrypt.GenerateFromPassword([]byte("test"), 10)
		fmt.Printf("Thread: %d\n", threadId)
		fmt.Println("Data Received:", data)
	}
	wg.Done()
}

func limitParallel(n int) {
	ch := make(chan bool)
	go sendData(ch)
	var wg sync.WaitGroup
	//wg.Add(n)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go receiveData(&wg, i, ch)
	}

	wg.Wait()
}
