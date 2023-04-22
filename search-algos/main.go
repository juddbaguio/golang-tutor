package main

import "log"

func main() {
	// unsorted
	// var intSlice []int = []int{5, 2, 3, 1, 0, 23, 18}
	// log.Println(linear_search(intSlice, 3))

	// sorted
	var sortedIntSlice []int = []int{30, 32, 33, 36, 37, 40, 41, 42, 45, 50}
	log.Println(binary_search(sortedIntSlice, 9))
	log.Println(binary_search(sortedIntSlice, 33))
}
