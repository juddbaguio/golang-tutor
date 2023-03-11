package main

import (
	"log"
	"time"
)

func firstFunc(timesUpChan chan string) {
	time.Sleep(2 * time.Second)
	timesUpChan <- "I'm from first func"
}

func secondFunc(timesUpChan chan string) {
	time.Sleep(3 * time.Second)
	timesUpChan <- "I'm from second func"
}

func main() {
	firstChan := make(chan string)
	secondChan := make(chan string)

	go firstFunc(firstChan)
	go secondFunc(secondChan)

	// similar to switch-case keyword but cases are channels
	select {
	case val := <-firstChan:
		log.Println(val)
	case val := <-secondChan:
		log.Println(val)
	}
}
