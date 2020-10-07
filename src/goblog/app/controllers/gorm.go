package controllers

import (
	"goblog/app/models"

	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

var (
	db *gorm.DB
)

// GormController定義
type GormController struct {
	*revel.Controller
	Txn *gorm.DB
}

// DB初期化
func InitDB() {
	var (
		driver, spec string
		found        bool
	)

	// Read configuration.
	if driver, found = revel.Config.String("db.driver"); !found {
		panic("No db.driver found.")
	}
	if spec, found = revel.Config.String("db.spec"); !found {
		panic("No db.spec found.")
	}

	// Open a connection.
	var err error
	db, err = gorm.Open(driver, spec)
	if err != nil {
		panic(err)
	}

	// Enable Logger
	db.LogMode(true)
	migrate()
}

// テーブル生成
func migrate() {
	db.AutoMigrate(&models.Post{}, &models.Comment{})
}

// トランザクション設定
// Begin a transaction
func (c GormController) Begin() revel.Result {
	c.Txn = db.Begin()
	return nil
}

// Rollback if it’s still going (must have panicked).
func (c GormController) Rollback() revel.Result {
	if c.Txn != nil {
		c.Txn.Rollback()
		c.Txn = nil
	}
	return nil
}

// Commit the transaction.
func (c *GormController) Commit() revel.Result {
	if c.Txn != nil {
		c.Txn.Commit()
		c.Txn = nil
	}
	return nil
}
