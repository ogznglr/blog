package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dbPort = "yourport"
const dbPassword = "yourpassword"
const hostName = "yourhostname"

var DB *gorm.DB

func Connection() {
	str := fmt.Sprintf("root:%s@tcp(%s:%s)/blog?charset=utf8mb4&parseTime=True&loc=Local", dbPassword, hostName, dbPort)
	db, err := gorm.Open(mysql.Open(str), &gorm.Config{})
	if err != nil {
		panic("Couldn't connect to the database")
	}
	DB = db
}

func Migrate(d interface{}) error {

	err := DB.AutoMigrate(d)
	if err != nil {
		panic("Database couldn't be updated")
	}
	return nil
}
