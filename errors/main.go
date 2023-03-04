package main

import (
	"errors"
	"fmt"
	"log"
)

func AddNumber(num int) error {
	if num > 10 {
		return errors.New("the number is too large")
	}

	return nil
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Error() string {
	return fmt.Sprintf("I have an error because my age is %d", p.Age)
}

func logError(err error) {
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func main() {
	var person Person
	person.Age = 10

	logError(person)
}
