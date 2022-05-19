package main

import (
	// "fmt"
	"database/sql"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Define model
type User struct {
	ID int64
	Name sql.NullString `gorm:"default:'ddd'"` // *string
	Age int64
}

func main() {
	db, err := gorm.Open("mysql", "root:1qaz9ol.@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("Connect database failed: %v\n", err)
	}
	defer db.Close()
	db.SingularTable(true)

	// If modify the format of the table, need to drop the table first
	if !db.HasTable("user"){
		db.AutoMigrate(User{})
	}

	// u := User{Age: 58}
	// if db.NewRecord(&u) {
	// 	db.Create(&u)
	// }

	// u = User{Name:"", Age: 38}
	// if db.NewRecord(&u) {
	// 	db.Debug().Create(&u)
	// }

	// // pointer and sql.Nullxx(struct) could resolve the problem, that
	// // "" , 0 or nil is the default zero value for the type so that these
	// // values are considered as non-assignment
	// u := User{Name: new(string), Age: 99}
	// if db.NewRecord(&u) {
	// 	db.Debug().Create(&u)
	// }

	u := User{Age: 88}
	if db.NewRecord(&u) {
		db.Debug().Create(&u)
	}

	u = User{Name: sql.NullString{String: "", Valid: true}, Age: 77}
	if db.NewRecord(&u) {
		db.Debug().Create(&u)
	}

	u = User{Name: sql.NullString{String: "", Valid: false}, Age: 77}
	if db.NewRecord(&u) {
		db.Debug().Create(&u)
	}
}