package db

import (
	"bannayuu-web-admin/constants"
	"github.com/kaewdungdee2538/ouanfunction/numeric"
	"github.com/kaewdungdee2538/ouanfunction/pg_db"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func SetupDB(dbHost string, dbUserName string, dbPassword string, dbName string, dbPort string, maxIdleStr string, maxConnectionsStr string) {
	maxIdle := numeric.ConvertStringToInt(maxIdleStr)
	maxConnections := numeric.ConvertStringToInt(maxConnectionsStr)
	if (maxConnections == 0){
		maxConnections = 20
	}
	// connect postgresql databse by ouanfunction/pg_db
	database, err := pg_db.SetupDB(dbHost, dbUserName, dbPassword, dbName, dbPort, maxIdle, maxConnections)
	if err != nil {
		panic(err)
	}
	// set database reference to global variable
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
