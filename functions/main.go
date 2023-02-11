package main

// func add(a int, b int) int {
// 	return a + b
// }

// func add(a int, b int) (sum int) {
// 	sum = a + b
// 	return sum
// }

// equivalent of second example
// func add(a int, b int) int {
// var sum int
// 	sum = a + b
// 	return sum
// }

// func withMultipleValues(key string) (string, int) {
// 	var sampleMap map[string]int = make(map[string]int)
// 	sampleMap["jim"] = 1
// 	sampleMap["xel"] = 2
// 	sampleMap["maghanoy"] = 3

// 	value, ok := sampleMap[key]
// 	if !ok {
// 		return "key doesn't exists", 0
// 	}

// 	return key, value
// }

// variadic functions
// func sampleVariadicAdder(a ...int) int {
// 	var sum int
// 	for _, value := range a {
// 		sum += value
// 	}
// 	return sum
// }

// invalid variadic function
// func invalidVariadicAdder(a ...int, b string) {}

// valid variadic function
// func validVariadicAdder(b string, a ...int) {}

func main() {
	// log.Println("the sum is ", add(1, 2))
	// result_one, result_two := withMultipleValues("jim")
	// log.Println(result_one, result_two)
	// log.Println(sampleVariadicAdder(1, 2, 3, 4))
}
