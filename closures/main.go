package main

import "log"

// CTRL + Space
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// func modifiedSeq() func(int) int {
// 	return func(i int) int {}
// }

var num int = 0

func incrementer() int {
	num++
	return num
}

func main() {
	someFn := intSeq()
	log.Println(someFn()) // 1
	log.Println(someFn()) // expectation? 2
	log.Println(someFn())
	log.Println("~~~~~~~~~~~~~~~~~~~~~~~~~")
	log.Println(incrementer())
	log.Println(incrementer())
	log.Println(incrementer())

}
