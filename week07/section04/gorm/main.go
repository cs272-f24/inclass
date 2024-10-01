package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Customer struct {
	gorm.Model
	Name   string `gorm:"unique"`
	Height int
}

func main() {
	// Open the database
	db, err := gorm.Open(sqlite.Open("demo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("gorm.Open returned ", err)
	}

	// Create or migrate tables using Go type reflection
	db.AutoMigrate(&Customer{})

	cust := Customer{Name: "John Smith", Height: 65}

	// Insert row into DB
	db.Create(&cust)
}
