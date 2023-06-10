package main

import (
	"database/sql"
	"log"
	"os"
)

func CreateTable(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
		id bigserial primary key,
		first_name varchar,
		last_name varchar,
		username varchar,
		password varchar
		)
	`)

	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}

type User struct {
	ID        int
	FirstName string
	LastName  string
	Username  string
	Password  string
}

func InsertRowToUsers(db *sql.DB, user User) error {
	_, err := db.Exec(`
	INSERT INTO users (first_name, last_name, username, password)
	VALUES ($1, $2, $4, $3)
	`, user.FirstName, user.LastName, user.Username, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	db := NewDB()
	defer db.Close()

	CreateTable(db)
	err := InsertRowToUsers(db, User{
		FirstName: "Jim Xel",
		LastName:  "Maghanoy",
		Username:  "jim123",
		Password:  "123456",
	})

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
