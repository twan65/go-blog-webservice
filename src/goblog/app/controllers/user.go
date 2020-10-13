// ユーザー生成、照会、更新、削除
package controllers

import (
	"goblog/app/models"

	"github.com/revel/revel"
)

type User struct {
	App
}

func (c User) AddUser() revel.Result {
	user := models.User{}
	return c.Render(user)
}
