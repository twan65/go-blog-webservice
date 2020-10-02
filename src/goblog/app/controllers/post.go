// ポスト生成、照会、更新、削除
package controllers

import (
	"goblog/app/models"

	db "github.com/revel/modules/db/app"
	"github.com/revel/revel"
)

type Post struct {
	*revel.Controller
	db.Transactional
}

// 全てのポストを取得
func (c Post) Index() revel.Result {
	var posts []models.Post
	rows, err := c.Txn.Query("select id, title, body, created_at, updated_at from posts order by create_at desc")
	if err != nil {
		panic(err)
	}

	// 取得した全てのポストをPostに変換にセットする
	for rows.Next() {
		post := models.Post{}
		if err := rows.Scan(&post.Id, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt); err != nil {
			panic(err)
		}

		posts = append(posts, post)

	}

	// これでレンダリングするとビューでposts変数にアクセスができる。
	return c.Render(posts)
}
