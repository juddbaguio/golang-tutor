package main

import "log"

type Array struct {
	storage []int
}

// func newAppend(num ...int) {
// 	num
// }

// Traverse
func (a *Array) Traverse() {
	for _, value := range a.storage {
		log.Println(value)
	}
}

// Insert
func (a *Array) Insert(num int) {
	a.storage = append(a.storage, num)
}

// Delete
func (a *Array) Delete(index int) {
	if index > len(a.storage)-1 {
		log.Panicln("index is out of bounds")
	}

	_ = append(a.storage[:index], 5, 5, 5, 5, 5, 5, 5, 5)
	a.storage = append(a.storage[:index], a.storage[index+1:]...)
}

// Delete
func (a *Array) Search(index int) int {
	for idx, val := range a.storage {
		if idx == index {
			return val
		}
	}

	return -1
}

func main() {
	var arr *Array = &Array{}
	arr.Insert(5)
	arr.Insert(4)
	arr.Insert(3)
	arr.Insert(2)
	arr.Insert(1)

	// log.Println("Before Delete:")
	// arr.Traverse()

	// arr.Delete(3)
	// log.Println("After Delete:")
	// arr.Traverse()

	log.Println(arr.Search(3))

}
