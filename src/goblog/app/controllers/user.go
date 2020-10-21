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

// CurrentUserの権限確認
func (c User) CheckUser() revel.Result {
	// IndexとShowは権限を確認しない
	switch c.MethodName {
	case "Index", "Show", "AddUser":
		return nil
	}

	// CurrentUser情報がなければログイン画面に遷移
	if c.CurrentUser == nil {
		c.Flash.Error("ログインしてください。")
		return c.Redirect(App.Login)
	}

	// CurrentUserが管理者ではなければログイン画面に遷移
	if c.CurrentUser.Role != "1" {
		c.Response.Status = 401 // Unauthorized
		c.Flash.Error("管理者ではありません。")
		return c.Redirect(App.Login)
	}

	return nil
}

// 全てのユーザーを取得
func (c User) Index() revel.Result {
	var users []models.User
	c.Txn.Order("created_at desc").Find(&users)

	// これでレンダリングするとビューでusers変数にアクセスができる。
	return c.Render(users)
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
