package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)



type Test struct {
	ID int
	Name string
	Comment string
}

func main() {
	db, err := gorm.Open("mysql", "root:1qaz9ol.@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("connect failed: %v\n", err)
	}
	defer db.Close()

	// db.AutoMigrate(&Test{})

	t1 := new(Test)
	t1.Name = "name"
	t1.Comment = "t"
	db.Create(t1)

	var t2 Test
	db.First(&t2, "id = ?", "7")
	fmt.Printf("%v\n", t2)

	db.Model(&t2).Update("comment", "t2")

	db.Delete(&t2, "comment = ?", "t2")
}