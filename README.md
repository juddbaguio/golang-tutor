// Create an HR application
// that can do the following:

// - add Employee which may be one of the following roles: Doctor, Engineer, Lawyer

// - let the employees introduce their selves, one by one,
// which prints out in the following format using a method called WhoAmI()
// "My name is {name} and I am {age} year(s) old. My role is {role}."

# Possible Solution to Session-4 Assignment:
```go
package main

import "log"

type Employee interface {
	WhoAmI()
}

type HR struct {
	Employees []Employee
}

func (hr *HR) AddEmployee(employee Employee) {
	hr.Employees = append(hr.Employees, employee)
}

func (hr *HR) IntroduceEmployees() {
	for _, employee := range hr.Employees {
		employee.WhoAmI()
	}
}

type Person struct {
	Name string
	Age  int
}

type Doctor struct {
	Person
}

func (d Doctor) WhoAmI() {
	log.Printf("My name is %s and I am %d year(s) old. My role is doctor.\n", d.Name, d.Age)
}

type Engineer struct {
	Person
}

func (e Engineer) WhoAmI() {
	log.Printf("My name is %s and I am %d year(s) old. My role is engineer.\n", e.Name, e.Age)
}

type Lawyer struct {
	Person
}

func (l Lawyer) WhoAmI() {
	log.Printf("My name is %s and I am %d year(s) old. My role is lawyer.\n", l.Name, l.Age)
}

func main() {
	var hr HR = HR{}
	hr.AddEmployee(Doctor{
		Person: Person{
			Name: "Judd",
			Age:  26,
		},
	})

	hr.AddEmployee(Lawyer{
		Person: Person{
			Name: "Jim",
			Age:  26,
		},
	})

	hr.AddEmployee(Engineer{
		Person: Person{
			Name: "Xel",
			Age:  26,
		},
	})

	hr.IntroduceEmployees()
}

```