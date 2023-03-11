package main

import "log"

func main() {
	amountChan := make(chan int, 3) // default 1
	amountChan <- 1                 // [1]
	amountChan <- 2                 // [1] <- 2
	amountChan <- 3                 // [1] <- 2 <- 3
	log.Println(<-amountChan)
	log.Println(<-amountChan)
	log.Println(<-amountChan)
}
