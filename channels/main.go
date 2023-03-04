package main

import (
	"log"
	"time"
)

func printMyNameNTimes(name string, n int, isDone chan bool) {
	for i := 0; i < n; i++ {
		log.Println("My name is ", name)
		time.Sleep(1 * time.Second)
	}

	isDone <- true // to send
}

func logEvenNumber(numChan chan int, isDone chan bool) {
	for num := range numChan { // < - 0, 1 , 2, 3, 5
		if num%2 != 0 {
			log.Println("I am not an even number")
			time.Sleep(1 * time.Second)
			continue
		}
		log.Println("I am an even number", num)
		time.Sleep(1 * time.Second)
	}

	isDone <- true
}

func main() {
	// sampleBoolChan := make(chan bool)
	// go printMyNameNTimes("judd", 10, sampleBoolChan)

	// <-sampleBoolChan // to receive, will block until there is value

	numChan := make(chan int)
	boolChan := make(chan bool)

	go logEvenNumber(numChan, boolChan)

	for i := 0; i < 20; i++ {
		numChan <- i
	}

	close(numChan)

	<-boolChan

}
