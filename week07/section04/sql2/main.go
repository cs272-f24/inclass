package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	db, err := sql.Open("sqlite3", "demo.db")
	if err != nil {
		log.Fatal("sql.Open returned ", err)
	}
	defer db.Close()

	sqlCreateTable := `
		CREATE TABLE IF NOT EXISTS users 
		(id INTEGER PRIMARY KEY, name TEXT, email TEXT UNIQUE)
	`
	stmt, err := db.Prepare(sqlCreateTable)
	if err != nil {
		log.Println("Prepare returned ", err)
	}
	stmt.Exec()

	sqlInsertUser := `
		INSERT INTO users(name, email)
		VALUES (?, ?)
	`
	stmt, err = db.Prepare(sqlInsertUser)
	if err != nil {
		log.Println("Prepare insert returned ", err)
	}
	stmt.Exec("Phil Peterson", "phpeterson@usfca.edu")

}
