package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	str := "root:password@tcp(db:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(str), &gorm.Config{})

	DB = db
}

func Migrate(d interface{}) error {

	DB.AutoMigrate(d)

	return nil
}
