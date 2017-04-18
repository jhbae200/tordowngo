package controllers

import (
	"database/sql"

	"github.com/jinzhu/gorm"
	"github.com/revel/revel"

	"tordowngo/app/database"
)

func InitDB() {
	var err error
	err = database.InitDB()
	if err != nil {
		panic(err)
	}
}

type GormController struct {
	*revel.Controller
	Txn *gorm.DB
}

func (c *GormController) Begin() revel.Result {
	txn := database.GetDb().Begin()

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
