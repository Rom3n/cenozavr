package database

import (
	"cenozavr/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbase *gorm.DB

func Init() *gorm.DB {
	dsn := "host=localhost user=admin password=admin dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Goods{})
	return db
}

func GetDB() *gorm.DB {
	if dbase == nil {
		dbase = Init()
		for dbase == nil {
			fmt.Println("DataBase is unavailable. Trying to connect")
			dbase = Init()
		}
	}
	return dbase
}
