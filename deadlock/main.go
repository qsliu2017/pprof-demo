package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

var (
	A  = &sync.Mutex{}
	B  = &sync.Mutex{}
	wg = &sync.WaitGroup{}
)

func worker0() {
	defer wg.Done()
	println("goroutine 0 requests A")
	A.Lock()
	defer A.Unlock()
	time.Sleep(1000 * time.Millisecond)
	println("goroutine 0 requests B")
	B.Lock()
	defer B.Unlock()
}

func worker1() {
	defer wg.Done()
	println("goroutine 1 requests B")
	B.Lock()
	defer B.Unlock()
	time.Sleep(1000 * time.Millisecond)
	println("goroutine 1 requests A")
	A.Lock()
	defer A.Unlock()
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	wg.Add(2)
	go worker0()
	go worker1()

	wg.Wait()
}
