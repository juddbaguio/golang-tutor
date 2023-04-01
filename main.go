package main

import "log"

// O(1) -> Good as line by line
func O1_Function() {
	name := "Jim Xel"
	name2 := "Jim Xel"
	log.Println(name)
	log.Println(name2)
}

// O(N) -> best example: loops
func O_N_Function(n int) {
	log.Println("Starting Function")
	for i := 0; i < n; i++ {
		log.Println(i)
	}

	for i := 0; i < n; i++ {
		log.Println(i)
	} // O(N + N) = O(2N) => O(N)

	log.Println("ending the function")
}

// O(N^2) -> loop within a loop
func O_N_Sqrd_Function(i, j int) {
	for a := 0; a < i; a++ {
		for b := 0; b < j; b++ {
			log.Println(a, b)
		} // O(i*j) => O(N^2)
	}
}

func main() {
	// O(1) example
	// O1_Function()

	// O_N_Function(5)
	O_N_Sqrd_Function(1000, 1000)
}
