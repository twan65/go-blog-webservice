package controllers

import (
	"goblog/app/models"

	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
)

type App struct {
	GormController
}

func (c App) Login() revel.Result {
	return c.Render()
}

func (c App) CreateSession(username, password string) revel.Result {
	var user models.User

	// usernameでユーザー取得
	c.Txn.Where(&models.User{Username: username}).First(&user)

	// CompareHashAndPassword関数でパスワード比較
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	// パスワードが合致している場合、セッションを生成し、ポストホーム画面に遷移
	if err == nil {
		authKey := revel.Sign(user.Username)
		c.Session["authKey"] = authKey
		c.Session["username"] = user.Username
		c.Flash.Success("Welcome, " + user.Name)
		return c.Redirect(Post.Index)
	}

	// 全てのセッション情報を削除し、ホームに遷移
	for k := range c.Session {
		delete(c.Session, k)
	}
	c.Flash.Out["username"] = username
	c.Flash.Error("ログインに失敗しました。")
	return c.Redirect(Home.Index)
}

// ログアウト処理
func (c App) DestroySession() revel.Result {
	// clear session
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(Home.Index)
}
