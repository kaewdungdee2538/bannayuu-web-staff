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

func SaveTransactionDB(query string,value map[string]interface{}) (bool, string) {
	tx := GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return true, err.Error()
	}

	if err := tx.Exec(query,value).Error; err != nil {
		return true, err.Error()
	}

	cmt := tx.Commit().Error
	if cmt != nil {
		return true, cmt.Error()
	}else{
		return false,constants.MessageSuccess
	}

}
