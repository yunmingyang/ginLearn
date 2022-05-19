package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)



type User struct {
	gorm.Model
	Name string
	Age sql.NullInt64
	Birthday *time.Time
	Email string `gorm:"type:varchar(100);unique_index"`
	Role string `gorm:"size:255"`
	MemberNumber *string `gorm:"unique;not null"`
	Num int `gorm:"AUTO_INCREMENT"`
	Addresss string `gorm:"index:addr"`
	Ignore int `gorm:"-"`
}

type Animal struct {
	AnimalID int64 `gorm:"primary_key"`
	Name string
	Age int64 
}

// Default table name will use this name even though db.SingularTable
// func (Animals) TableName() string {
// 	return "aaaa"
// }

func main() {
	// Modify the default table name
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "SMS_" + defaultTableName
	}

	db, err := gorm.Open("mysql", "root:1qaz9ol.@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("Connect database failed: %v\n", err)
	}
	defer db.Close()
	db.SingularTable(true)

	// db.AutoMigrate(&User{})
	// // both a pointer and a instance is allowed
	db.AutoMigrate(Animal{})

	// db.Table("test123").CreateTable(&User{})

	// mn := "asdasd"
	// u1 := &User{Name: "test", Addresss: "sdasdasd", MemberNumber: &mn}
	// // If table name is not the default one, the table should be set manuanlly before selecting or some other operations
	// db.Table("test123").Create(u1)
	// Soft delete, delete_at field will be set, and gorm could not find this line, but in mysql,
	// the line is still there
	// db.Where("name = ?", "test").Delete(u1)
	var u2 User
	// db.Where("name = ?", "test").First(&u2)
	db.Select("name").Where("name = ?", "test").First(&u2)
	log.Printf("%v\n",u2)
}