package main

import "log"

// factorial: 8! = 8*7! => n! = n*(n-1)!
// factorial: 0! = 1
// base case - stopping point sa recursive function
func factorial(n int) int {
	if n == 0 {
		return 1
	} // base case

	return n * factorial(n-1)
}

func main() {
	log.Println(factorial(8))
}
