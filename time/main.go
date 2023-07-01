package main

import (
	"log"
	"time"
)

func main() {
	// currentTime := time.Now()
	// log.Println(currentTime.Format(time.RFC3339))                // local time format within system
	// log.Println(currentTime.UTC().Format(time.RFC3339))          // UTC
	// log.Println(currentTime.UTC().Format("2006-01-02 15:04:05")) // UTC YYYY-MM-DD HH:mm:ss

	timeString := "2023-07-01 06:44:07"

	// parsedTime, err := time.Parse("2006-01-02 15:04:05", timeString)
	parsedTime, err := time.Parse(time.RFC850, timeString)
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(parsedTime)
	}
}
