package database

import (
	"tordowngo/app/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/revel/revel"
)

var (
	Db *gorm.DB
)

func InitDB() error {
	driver := revel.Config.StringDefault("db.driver", "sqlite3")
	connect_string := revel.Config.StringDefault("db.connect", "tordown.db")
	var err error
	Db, err = gorm.Open(driver, connect_string)
	if err != nil {
		return err
	}
	Db.AutoMigrate(&model.Member{})
	Db.AutoMigrate(&model.File{})
	return nil
}

func GetDb() *gorm.DB {
	return Db
}
