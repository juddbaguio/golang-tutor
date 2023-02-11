package main

import "log"

func main() {
	// same as traditional if else
	// if 7%2 == 0 {
	// 	log.Println("7 is even")
	// } else {
	// 	log.Println("7 is odd")
	// }

	// var sampleMap map[string]int = make(map[string]int)
	// sampleMap["jim"] = 1
	// sampleMap["xel"] = 2
	// sampleMap["maghanoy"] = 3

	// var sliceString []string = []string{"jim", "xel2", "maghanoy"}
	// for _, value := range sliceString {
	// 	if num, ok := sampleMap[value]; ok {
	// 		log.Println(num)
	// 	}
	// }

	var sliceString []string = []string{"jim", "xel", "maghanoy"}
	someVal := sliceString[0]
	switch {
	case someVal == "jim":
		log.Println("Hi I am jim")
	case someVal == "xel":
		log.Println("Hi I am xel")
	default:
		log.Println("I don't know who I am")
	}
}
