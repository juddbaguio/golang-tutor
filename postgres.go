package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Tx interface {
	Commit() error
	Rollback() error
}

func NewDB() *sql.DB {
	connString := "postgres://biyaheroes:123456@localhost:5432/biyaheroes?sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func IsViableTx(tx Tx) bool {
	_, ok := tx.(*sql.Tx)

	return ok
}
