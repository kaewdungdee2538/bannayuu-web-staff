package db

import (
	constants "bannayuu-web-admin/constants"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func SetupDB() {
	dsn := fmt.Sprintf("host=%s user=cit password=db13apr dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", constants.DbHost, constants.DbName, constants.DbPort)
	fmt.Println(dsn)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("connect to database failed")
	}
	db = database
}
