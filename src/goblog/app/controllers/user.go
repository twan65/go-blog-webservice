// ユーザー生成、照会、更新、削除
package controllers

import (
	"goblog/app/models"
	"goblog/app/routes"

	"github.com/revel/revel"
)

type User struct {
	App
}

func (c User) AddUser() revel.Result {
	user := models.User{}
	return c.Render(user)
}

func (c User) CreateUser(name, username, password string) revel.Result {
	// 構造体生成
	user := models.User{Name: name, Role: "0", Username: username, Password: password}

	// DB登録
	c.Txn.Create(&user)

	// ViewにFlashメッセージを渡す。
	c.Flash.Success("会員登録完了")

	authKey := revel.Sign(user.Username)
	c.Session["authKey"] = authKey
	c.Session["username"] = user.Username
	c.Flash.Success("Welcome, " + user.Name)
	return c.Redirect(Post.Index)
}

// ユーザー更新データ取得
func (c User) EditUser(id int) revel.Result {
	var user models.User
	c.Txn.First(&user, id)

	return c.Render(user)

}

func (c User) UpdateUser(id int, name, password string) revel.Result {

	var user models.User
	// 更新データを取得
	c.Txn.First(&user, id)
	user.Name = name
	user.Password = password

	// ポスト内容を更新
	c.Txn.Save(&user)

	// ビューにFlashメッセージを渡す。
	c.Flash.Success("ユーザー情報更新完了")

	// ポスト詳細画面に遷移
	return c.Redirect(routes.Post.Index())
}
