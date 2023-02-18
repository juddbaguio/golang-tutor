package main

import "log"

// num = 0x00cadadada (memory-address)
func incrementer(num *int) {
	// log.Println(num) // memory-address
	*num += 1
}

func main() {
	var someInt int = 5
	// memAddressOfSumInt := &someInt
	incrementer(&someInt)
	log.Println(someInt)

	// invalid dangling pointers
	// var danglingInt *int // zero-value: nil
	// incrementer(danglingInt)

	// valid use case
	var notDanglingInt *int = &someInt
	incrementer(notDanglingInt)
	log.Println(*notDanglingInt)
	log.Println(someInt)

	//
	var doubleReferenceInt **int = &notDanglingInt
	// *doubleReferenceInt => *int
	**doubleReferenceInt += 2 // dereference
	log.Println(someInt)
}
