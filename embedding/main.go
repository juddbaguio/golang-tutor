package main

import (
	"fmt"
	"log"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type PersonCapabilities interface {
	CelebrateBirthday()
	GetFullName() string
	GetAge() int
}

func (p *Person) CelebrateBirthday() {
	p.Age += 1
}

func (p Person) GetFullName() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

func (p Person) GetAge() int {
	return p.Age
}

type Student struct {
	Year             int
	Subjects         []string
	SliceOfInterface []interface{}
	Person           // Person Person
}

func (s *Student) AddStudent(firstName, lastName string, age, year int) {
	s.Year = year
	s.Person = Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

type StudentWithEmbedInterface struct {
	Year int
	PersonCapabilities
}

func main() {
	var student Student
	student.AddStudent("Jim", "Maghanoy", 30, 2)
	log.Println("Before", student)

	student.FirstName = "Judd"
	student.Subjects = []string{"Math", "Science"}
	student.SliceOfInterface = []interface{}{"String", 1, false}
	log.Println("After", student)

	var studentWithInterface StudentWithEmbedInterface = StudentWithEmbedInterface{
		Year: 2,
		PersonCapabilities: &Person{
			FirstName: "Jim Xel",
			LastName:  "Maghanoy",
			Age:       30,
		},
	}

	log.Println(studentWithInterface.GetFullName())
	log.Println(studentWithInterface.GetAge())
}
