package controllers

import (
	"goblog/app/models"
	"goblog/app/routes"

	"github.com/revel/revel"
)

type Comment struct {
	App
}

func (c Comment) CheckUser() revel.Result {
	if c.MethodName != "Destroy" {
		return nil
	}

	if c.CurrentUser == nil {
		c.Flash.Error("ログインしてください。")
		return c.Redirect(App.Login)
	}

	if c.CurrentUser.Role != "admin" {
		c.Response.Status = 401 // Unauthorized
		c.Flash.Error("管理者ではありません。")
		return c.Redirect(App.Login)
	}
	return nil
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
