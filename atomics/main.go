package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func adder(wg *sync.WaitGroup, ops *uint64) {
	for c := 0; c < 1000; c++ {
		// *ops += 1
		atomic.AddUint64(ops, 1)
	}
	wg.Done()
}

func main() {
	var ops uint64 // unsigned int64 >= 0 ; max positive number: 2^64 - 1
	// 64bit boundary: -2^64 up to 2^64 - 1
	// 32bit : -2^32 up to 2^64 - 1
	// 8bit: -2^8 up to 2^8 - 1

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go adder(&wg, &ops)
		// non-blocking
	}

	wg.Wait()

	fmt.Println("ops:", ops)
}
