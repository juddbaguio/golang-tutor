package main

import (
	somepkg "golang-tutor/structs/some_pkg"
	"log"
)

func ageIncrementer(person *somepkg.Person) {
	person.Age += 1 // by field

	// *person = somepkg.Person{} => whole struct to change
}

func main() {
	var person somepkg.Person = somepkg.Person{
		Name: "Jim",
		Age:  30,
	}

	ageIncrementer(&person)
	log.Println(person.Age)
}
