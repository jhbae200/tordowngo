package controllers

import (
	"database/sql"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/revel/revel"

	"tordowngo/app/model"
)

var (
	Db *gorm.DB
)

func InitDB() {
	driver := revel.Config.StringDefault("db.driver", "sqlite3")
	connect_string := revel.Config.StringDefault("db.connect", "tordown.db")
	var err error
	Db, err = gorm.Open(driver, connect_string)
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&model.Member{})
	Db.AutoMigrate(&model.File{})
}

type GormController struct {
	*revel.Controller
	Txn *gorm.DB
}

func (c *GormController) Begin() revel.Result {
	txn := Db.Begin()

	if txn.Error != nil {
		panic(txn.Error)
	}

	c.Txn = txn
	return nil
}

func (c *GormController) Commit() revel.Result {
	if c.Txn == nil {
		return nil
	}
	c.Txn.Commit()
	if err := c.Txn.Error; err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GormController) Rollback() revel.Result {
	if c.Txn == nil {
		return nil
	}
	c.Txn.Rollback()
	if err := c.Txn.Error; err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}
