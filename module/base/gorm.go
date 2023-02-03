package base

import (
	"fiberun/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var client *gorm.DB

func ConnectDB(dsn string) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic(err.Error())
	}

	client = db
}

func GetDB() *gorm.DB {
	return client
}
