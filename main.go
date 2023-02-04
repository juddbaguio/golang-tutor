package main

import "fmt"

const (
	name string = "Jim Xel"
	age  int    = 33
)

// some_name := "Jim Xel"

// var (
// 	name string = "Jim Xel New"
// 	age  int    = 25
// )

func main() {
	// var name string = "Jim Xel"
	// var age int = 33
	var some_slice []int = make([]int, 0)
	for i := 0; i <= 10; i++ {
		some_slice = append(some_slice, i)
	}
	fmt.Println(some_slice)
	fmt.Println(some_slice[0])
	fmt.Println(some_slice[0:5]) // slice[0:x] -> values from index 0 to x - 1
	fmt.Println(some_slice[0:])
	fmt.Println(some_slice[0:0])

	var some_array [3]int = [3]int{}
	fmt.Println(some_array)

	// 	fmt.Printf("Hello, %s, my age is %d\n", name, age)
	// }

	// var sampleFloat float32 = 33
	// var floatAge float32 = float32(age)

	// fmt.Println("Hello world - FMT")
	// fmt.Printf("Hello, %s, my age is %d\n", name, age)
	// calculatorPrinter()
}
