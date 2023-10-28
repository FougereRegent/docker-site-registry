package utils

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbConnector *gorm.DB = nil
var connectionString string = "user-service.db"

func GetConnector() *gorm.DB {

	if dbConnector != nil {
		return dbConnector
	}

	dbConnector, err := gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
	if err != nil {
		err_sentence := fmt.Sprintf("An Error was encoutered : %s", err)
		fmt.Println(err_sentence)
	}

	return dbConnector
}
