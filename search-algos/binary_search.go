package main

// Array is assumed to be sorted (preferrably, ascending order)
func binary_search(arr []int, value int) int {
	if len(arr) == 0 {
		return -1
	}

	lower_bound := 0
	higher_bound := len(arr) - 1
	var middle_index int

	for lower_bound != higher_bound {
		middle_index = (lower_bound + higher_bound) / 2

		if arr[middle_index] == value {
			return middle_index
		}

		if arr[middle_index] > value {
			higher_bound = middle_index
			continue
		}

		lower_bound = middle_index
	}

	middle_index = (lower_bound + higher_bound) / 2
	if arr[middle_index] == value {
		return middle_index
	}

	return -1
}
