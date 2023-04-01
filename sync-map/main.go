package main

import (
	"log"
	"sync"
)

type MyKey struct{}

func main() {
	syncMap := sync.Map{}
	normalMap := make(map[any]any)

	normalMap["MyName"] = "Maghanoy"
	normalMap[MyKey{}] = "Jim Xel"

	syncMap.Store("MyName", "Maghanoy")
	syncMap.Store(MyKey{}, "Jim Xel")

	log.Println(syncMap.Load(MyKey{}))
	log.Println(syncMap.LoadAndDelete("MyName"))
	log.Println(syncMap.Load("MyName"))
	log.Println(syncMap.LoadOrStore("MyName", "Maghanoy Jr."))
	log.Println(syncMap.Load("MyName"))
	syncMap.Delete("MyName")
	log.Println(syncMap.Load("MyName"))

	// log.Println(normalMap["MyName"])
	// log.Println(normalMap[MyKey{}])

}
