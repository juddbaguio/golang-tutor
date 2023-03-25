package main

import (
	"fmt"
	"sync"
)

func main() {
	m := map[int]int{}

	mux := &sync.RWMutex{} // use if you have a lot of readers

	go writeLoop(m, mux) // Lock()
	// first group of reads
	// writeLoop(m,mux) Unlock()
	go readLoop(m, mux) // RLock()
	go readLoop(m, mux) // RLock()
	// writeLoop(m,mux) Lock()

	// second second group
	// writeLoop(m,mux) Unlock()
	go readLoop(m, mux) // Rlock()
	go readLoop(m, mux) // Rlock()
	// writeLoop(m,mux) Lock()

	// stop program from exiting, must be killed
	block := make(chan struct{})
	<-block
}

func writeLoop(m map[int]int, mux *sync.RWMutex) {
	for {
		for i := 0; i < 100; i++ {
			mux.Lock()
			m[i] = i
			mux.Unlock()
		}
	}
}

func readLoop(m map[int]int, mux *sync.RWMutex) {
	for {
		mux.RLock()
		for k, v := range m {
			fmt.Println(k, "-", v)
		}
		mux.RUnlock()
	}
}
