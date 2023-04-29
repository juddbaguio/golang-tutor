package main

import "log"

func main() {
	var ll *LinkedList = &LinkedList{}

	// ll.DeleteAt(5)

	ll.Insert(5)
	ll.Insert(3)
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(9)

	ll.Traverse()

	ll.DeleteAt(2)

	log.Println("after deleting index 2 node")
	ll.Traverse()

}
