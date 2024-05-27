package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite3 driver
)

type User struct {
	ID   int
	Name string
}

func main() {
	db, err := gorm.Open("sqlite3", "user.db")
	if err != nil {
		panic("failed to connect")
	}
	db.SingularTable(true)
	defer db.Close()

	db.Exec(`PRAGMA foreign_keys=ON`)

	// Creates tables, columns, indexes. Does not delete or modify existing.
	db.AutoMigrate(&User{})

	db.Create(&User{Name: "tom"})

	var u User
	db.Where("name = ?", "tom").First(&u)

	fmt.Println(u)
}
