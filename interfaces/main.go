package main

import "log"

// blueprint
type Shape interface {
	Area() float32
}

type Square struct {
	Side int
}

// implementation detail
func (s Square) Area() float32 {
	return float32(s.Side) * float32(s.Side)
}

type Triangle struct {
	Base   int
	Height int
}

// area of triangle = (1/2)*base*height
func (t Triangle) Area() float32 {
	return 0.5 * float32(t.Base) * float32(t.Height)
}

type UserRepository interface {
	GetAgeByName(name string) int
}

type UserMysqlRepo struct {
	databaseName string
}

type MyApp struct {
	ListOfDatabases []interface {
		WhatIsMyDatabase() string
	}
}

func (repo UserMysqlRepo) GetAgeByName(name string) int {
	log.Println("Running with the following database engine: ", repo.databaseName)

	return 40
}

func (repo UserMysqlRepo) WhatIsMyDatabase() string {
	return repo.databaseName
}

type UserPostgresRepo struct {
	databaseName string
}

func (repo UserPostgresRepo) GetAgeByName(name string) int {
	log.Println("Running with the following database engine: ", repo.databaseName)

	return 50
}

func (repo UserPostgresRepo) WhatIsMyDatabase() string {
	return repo.databaseName
}

func main() {
	// var shape Shape = Square{
	// 	Side: 5,
	// }

	// log.Println("my shape's area is:", shape.Area())

	// var secondShape Shape = Triangle{
	// 	Base:   10,
	// 	Height: 25,
	// }

	// log.Println("my second shape's area is:", secondShape.Area())

	// var userRepo UserRepository = UserMysqlRepo{
	// 	databaseName: "MySQL",
	// }

	// log.Println(userRepo.GetAgeByName("Judd"))

	// userRepo = UserPostgresRepo{
	// 	databaseName: "PostgreSQL",
	// }

	// log.Println(userRepo.GetAgeByName("Judd"))

	var app MyApp = MyApp{
		[]interface{ WhatIsMyDatabase() string }{UserMysqlRepo{
			databaseName: "MySQL",
		}, UserPostgresRepo{
			databaseName: "PostgreSQL",
		}},
	}

	for _, database := range app.ListOfDatabases {
		log.Println("My database -", database.WhatIsMyDatabase())
	}

}
