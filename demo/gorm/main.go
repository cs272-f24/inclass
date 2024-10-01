package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"uniqueIndex"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("demo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("gorm.Open returned ", err)
	}

	db.AutoMigrate(&User{})

	user := User{
		Name:  "Phil Peterson",
		Email: "phpeterson@usfca.edu",
	}

	result := db.Create(&user)
	if result.Error != nil {
		log.Println("db.Create returned ", err)
	}

	var user2 User
	db.First(&user2, 1)
	log.Println("User found: ", user2)
}
