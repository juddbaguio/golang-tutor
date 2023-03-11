package main

import (
	"fmt"
	"log"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	// close(done) // closing the connection of our channel

	_, isChannelConnectionOpen := <-done // blocking part

	log.Println(isChannelConnectionOpen)
}
