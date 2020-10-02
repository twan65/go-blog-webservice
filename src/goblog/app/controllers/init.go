package controllers

import (
	db "github.com/revel/modules/db/app"
	"github.com/revel/revel"
)

// アプリケーション起動時、DBを初期化
func InitDB() {
	db.Init()
	schema := `
		CREATE TABLE IF NOT EXISTS posts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT	NOT NULL,
			body TEXT	NOT NULL,
			created_at DATETIME	NOT NULL,
			updated_at DATETIME	NOT NULL
		);
	`

	db.Db.Exec(schema)
}

func init() {
	// アプリケーション起動時、DBを初期化
	revel.OnAppStart(InitDB)
}
