package main

import "log"

func main() {
	var intSlice []int = []int{6, 3, 2, 4, 8, 9, 0, 1}

	log.Println(QuickSort(intSlice))
}
