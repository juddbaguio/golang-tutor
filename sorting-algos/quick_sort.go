package main

func QuickSort(intSlice []int) []int {
	if len(intSlice) <= 1 {
		return intSlice
	}

	pivot := intSlice[len(intSlice)-1]
	var leftArray []int  // for less than or equal of pivot
	var rightArray []int // for greater than or equal of pivot

	for i := 0; i < len(intSlice)-1; i++ {
		if intSlice[i] <= pivot {
			leftArray = append(leftArray, intSlice[i])
		} else {
			rightArray = append(rightArray, intSlice[i])
		}
	}

	leftArray = QuickSort(leftArray)
	rightArray = QuickSort(rightArray)

	intSlice = leftArray
	intSlice = append(intSlice, pivot)
	intSlice = append(intSlice, rightArray...)

	return intSlice
}
