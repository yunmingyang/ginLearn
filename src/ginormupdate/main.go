package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)



type User struct {
	gorm.Model
	Name string
	Age int64
	Active bool
}

func main() {
	db, err := gorm.Open("mysql", "root:1qaz9ol.@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		log.Fatalf("mysql connect failed: %v", err)
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	u1 := User{Name: "yym", Age: 18, Active: true}
	db.Create(&u1)
	u2 := User{Name: "myy", Age: 19, Active: false}
	db.Create(&u2)

	var user User
	db.First(&user)
	fmt.Printf("User: %v", user)

	// user.Name = "y"
	// user.Age = 22
	// db.Debug().Save(&user)

	// db.Debug().Model(&user).Update("name", "1")

	// db.Debug().Model(&user).Where("active = ?", false).Update("name", "default")

	// db.Debug().Model(&user).Updates(map[string]interface{}{"name": "h", "age": 0, "active": false})

	// db.Debug().Model(&user).Omit("name").Update(map[string]interface{}{"name": "1", "age": 1, "active": true})

	// db.Debug().Model(&user).Select("age").Update(map[string]interface{}{"name": "default", "age": 22, "active": true})

	n := db.Debug().Model(&user).UpdateColumn("name", "default").RowsAffected
	fmt.Printf("modify %d rows\n", n)

	db.Debug().Model(&User{}).Update("age", gorm.Expr("age + ?", 2))
}