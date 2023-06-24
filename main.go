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
	VALUES ($1, $2, $3, $4)
	`, user.FirstName, user.LastName, user.Username, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func QueryUsers(db *sql.DB) (*[]User, error) {
	var userList []User
	rows, err := db.Query(`SELECT * FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var user User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Password); err != nil {
			return nil, err
		}

		userList = append(userList, user)
		user = User{}
	}

	return &userList, nil
}

func QueryUserByID(db *sql.DB, id int) (*User, error) {
	var user User
	row := db.QueryRow(`SELECT * FROM users WHERE id = $1`, id)

	if err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Password); err != nil {
		return nil, err
	}

	return &user, nil
}

func DeleteUser(db *sql.DB, id int) error {
	return nil
}

func main() {
	db := NewDB()
	defer db.Close()

	// userList, err := QueryUsers(db)

	user, err := QueryUserByID(db, 2)
	if err != nil {
		log.Println(err)
		return
	}

	// log.Println(*userList)
	log.Println(*user)
}
