package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)



type User struct {
	gorm.Model
	Name sql.NullString
	Age sql.NullInt64
}

func main() {
	db, err := gorm.Open("mysql", "root:1qaz9ol.@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		log.Fatalf("Connect failed: %v", err)
	}
	defer db.Close()

	if !db.HasTable("users") {
		db.AutoMigrate(&User{})
	}

	// // Create two test data
	// u1 := User{Name: sql.NullString{String: "", Valid: true}, Age: sql.NullInt64{Int64: 10, Valid: true}}
	// db.Debug().Create(&u1)

	// u2 := User{Name: sql.NullString{String: "rrrrr", Valid: true}, Age: sql.NullInt64{Int64: 0, Valid: true}}
	// db.Debug().Create(&u2)

	// var user User
	u1 := new(User)
	// The first record, primary key should be a number
	db.First(&u1)
	fmt.Printf("First: %v\n", u1)
	// The last record, primary key should be a number
	// TODO: a bug, if use the same variable of the previous steps, the variable will not be changed
	// even though it is a pointer, to avoid this, we should re-apply a variable
	var u2 User
	db.Last(&u2)
	fmt.Printf("Last: %v\n", u2)
	// A random record
	var u3 User
	db.Debug().Take(&u3)
	fmt.Printf("Take: %v\n", u3)
	// All record
	users := make([]User, 0)
	db.Find(&users)
	fmt.Printf("All: %v\n", users)
	// Find a record with primary key is a integer
	var u4 User
	db.Debug().First(&u4, 4)
	fmt.Printf("First with ID: %v\n", u4)

	var u5 User
	db.Debug().FirstOrInit(&u5, User{Name: sql.NullString{String: "u5", Valid: true}})
	fmt.Printf("FirstorInit: %v\n", u5)

	var u6 User
	db.Debug().FirstOrCreate(&u6, User{Name: sql.NullString{String: "u6", Valid: true}})
	fmt.Printf("FirstOrCreate: %v\n", u6)
}