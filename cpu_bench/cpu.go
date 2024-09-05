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

func sendData(ch chan<- bool) {
	for i := 0; i < 10; i++ {
		ch <- true
	}
	close(ch) // when close the channel, the channel will not be able to receive any more data, but still permit the consumption of the data
}

func receiveData(wg *sync.WaitGroup, threadId int, ch <-chan bool) {
	defer wg.Done() // make sure that this thread will be finished after the execution of this function
	for range ch {
		bcrypt.GenerateFromPassword([]byte("test"), 10)
	}
	//wg.Done()
}

func limitParallel(n int) {
	ch := make(chan bool)
	go sendData(ch)
	var wg sync.WaitGroup
	//wg.Add(n) when already knows how many threads will be running
	for i := 0; i < n; i++ {
		wg.Add(1) //when don't know how many threads will be run
		go receiveData(&wg, i, ch)
	}
	wg.Wait()
}
