package main

// Time complexity O(n)
func linear_search(arr []int, value int) (idx int) {
	for idx, el := range arr {
		if value == el {
			return idx
		}
	}

	return -1
}
