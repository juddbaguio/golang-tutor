package main

import "log"

// the channel argument is always the sender (chan <- [data_type])
func pingFunc(senderChan chan<- string, name string) {
	senderChan <- name // sender
}

// the channel argument is always the receiver (<- chan [data_type])
func pongFunc(senderChan <-chan string, receiverChan chan<- string) {
	name := <-senderChan

	receiverChan <- name
}

func main() {
	ping := make(chan string, 1)
	pong := make(chan string, 1)
	pingFunc(ping, "judd")
	pongFunc(ping, pong)

	log.Println(<-pong)
}
