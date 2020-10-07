package controllers

import (
	"goblog/app/models"
	"goblog/app/routes"

	"github.com/revel/revel"
)

type Comment struct {
	GormController
}

func (c Comment) Create(postId int, body, commenter string) revel.Result {
	// コメント登録
	comment := models.Comment{PostId: postId, Body: body, Commenter: commenter}
	c.Txn.Create(&comment)

	c.Flash.Success("コメント作成完了")

	return c.Redirect(routes.Post.Show(postId))
}

func (c Comment) Delete(postId, id int) revel.Result {
	c.Txn.Where("id = ?", id).Delete(&models.Comment{})

	c.Flash.Success("コメント削除完了")

	return c.Redirect(routes.Post.Show(postId))
}
