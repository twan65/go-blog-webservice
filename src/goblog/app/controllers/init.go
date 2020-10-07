package controllers

import (
	"github.com/revel/revel"
)

// アプリケーション起動時、DBを初期化
// func InitDB() {
// 	db.Init()
// 	schema := `
// 		CREATE TABLE IF NOT EXISTS posts (
// 			id INTEGER PRIMARY KEY AUTOINCREMENT,
// 			title TEXT	NOT NULL,
// 			body TEXT	NOT NULL,
// 			created_at DATETIME	NOT NULL,
// 			updated_at DATETIME	NOT NULL
// 		);

// 		CREATE TABLE IF NOT EXISTS comments (
// 			id          INTEGER PRIMARY KEY AUTOINCREMENT,
// 			body        TEXT                NOT NULL,
// 			commenter   TEXT                NOT NULL,
// 			post_id     INT                 NOT NULL,
// 			created_at  DATETIME            NOT NULL,
// 			updated_at  DATETIME            NOT NULL
// 		);
// 	`

// 	db.Db.Exec(schema)
// }

func init() {
	// アプリケーション起動時、DBを初期化
	revel.OnAppStart(InitDB)
	revel.InterceptMethod((*GormController).Begin, revel.BEFORE)
	revel.InterceptMethod((*GormController).Commit, revel.AFTER)
	revel.InterceptMethod((*GormController).Rollback, revel.FINALLY)
}
