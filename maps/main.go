package main

import "log"

func main() {
	// 1.) using make()
	// map format - map[data_type_1]data_type_2
	var sampleMap map[string]int = make(map[string]int)
	log.Println("without values", sampleMap)
	sampleMap["jim"] = 1
	sampleMap["xel"] = 2
	sampleMap["maghanoy"] = 3
	log.Println("with values:", sampleMap)
	log.Println(sampleMap["jim"])
	log.Println(sampleMap["xel"])
	log.Println(sampleMap["maghanoy"])

	delete(sampleMap, "maghanoy")
	log.Println("with deleted-key (maghanoy) values:", sampleMap)
	log.Println("accessing deleted key - ", sampleMap["maghanoy"])

	_, ok := sampleMap["maghanoy"]
	// optional
	// var num int
	// var validator bool
	// num, validator = sampleMap["maghanoy"]
	log.Println("does the value exist? ", ok)

	val, ok := sampleMap["jim"]
	log.Println("the value is ", val)
	log.Println("does the value exist? ", ok)

	var zeroValueMap map[string]int
	log.Println("zero-value of map", zeroValueMap)
	zeroValueMap["key"] = 1
	log.Println(zeroValueMap["key"])
}
