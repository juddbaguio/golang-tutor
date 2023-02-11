package main

import "log"

func main() {
	// var intSlice []int = []int{1, 2, 3, 4}
	// traditional loop

	// for i := 0; i < len(intSlice); i++ {
	// 	log.Println(intSlice[i])
	// }

	// var i int
	// for i < len(intSlice) {
	// 	log.Println(intSlice[i])
	// 	i++
	// }

	// with range loop
	// for index, value := range intSlice {
	// 	log.Println("The value is - ", value)
	// 	log.Println("The index is - ", index)
	// }

	// looping with maps
	var sampleMap map[string]int = make(map[string]int)
	sampleMap["jim"] = 1
	sampleMap["xel"] = 2
	sampleMap["maghanoy"] = 3

	for key, value := range sampleMap {
		// continue, break also exists with th same purpose
		if key == "xel" {
			continue
		}
		log.Printf("The key is %s and the value is %v\n", key, value)
		if key == "jim" {
			break
		}
	}
}
