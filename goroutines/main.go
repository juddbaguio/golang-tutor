package main

import (
	"log"
	"time"
)

func printMyName(name string) {
	log.Println("My name is", name)
}

func printMyNameNTimes(name string, times int) {
	for i := 0; i < 10; i++ {
		log.Println("My name is", name)
		time.Sleep(1 * time.Second)
	} // approximately {{times}} seconds to end the function
}

func main() {
	// go printMyName("judd") // in the background
	log.Println("I am here in the main thread")
	go printMyNameNTimes("judd", 10)
	// sleep
	time.Sleep(11 * time.Second)
}
