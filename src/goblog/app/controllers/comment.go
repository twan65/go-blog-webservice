package controllers

import (
	"goblog/app/routes"
	"time"

	db "github.com/revel/modules/db/app"
	"github.com/revel/revel"
)

type Comment struct {
	*revel.Controller
	db.Transactional
}

func (c Comment) Create(postId int, body, commenter string) revel.Result {
	// コメント登録
	_, err := c.Txn.Exec("insert into comments(body, commenter, post_id, created_at, updated_at) values(?,?,?,?,?)", body, commenter, postId, time.Now(), time.Now())
	if err != nil {
		panic(err)
	}

	c.Flash.Success("コメント作成完了")

	return c.Redirect(routes.Post.Show(postId))
}
