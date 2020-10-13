// ポスト生成、照会、更新、削除
package controllers

import (
	"goblog/app/models"
	"goblog/app/routes"

	"github.com/revel/revel"
)

type Post struct {
	App
}

// CurrentUserの権限確認
func (c Post) CheckUser() revel.Result {
	// IndexとShowは権限を確認しない
	switch c.MethodName {
	case "Index", "Show":
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

// 全てのポストを取得
func (c Post) Index() revel.Result {
	var posts []models.Post
	c.Txn.Order("created_at desc").Find(&posts)

	// これでレンダリングするとビューでposts変数にアクセスができる。
	return c.Render(posts)
}

// Empty Postを生成して画面にレンダリング
func (c Post) New() revel.Result {
	post := models.Post{}
	return c.Render(post)
}

// ポスト取得
func (c Post) Show(id int) revel.Result {
	var post models.Post
	c.Txn.First(&post, id)
	c.Txn.Where(&models.Comment{PostId: id}).Find(&post.Comments)

	return c.Render(post)
}

// ポストをDBに保存
func (c Post) Create(title, body string) revel.Result {
	// add Database
	post := models.Post{Title: title, Body: body}
	c.Txn.Create(&post)

	// ViewにFlashメッセージを渡す。
	c.Flash.Success("ポスト作成を完了")

	return c.Redirect(routes.Post.Index())
}

// ポスト更新データ取得
func (c Post) Edit(id int) revel.Result {
	var post models.Post
	c.Txn.First(&post, id)

	return c.Render(post)

}

func (c Post) Update(id int, title, body string) revel.Result {

	var post models.Post
	// 更新データを取得
	c.Txn.First(&post, id)
	post.Title = title
	post.Body = body

	// ポスト内容を更新
	c.Txn.Save(&post)

	// ビューにFlashメッセージを渡す。
	c.Flash.Success("更新完了")

	// ポスト詳細画面に遷移
	return c.Redirect(routes.Post.Show(id))
}

// ポスト削除
func (c Post) Delete(id int) revel.Result {

	c.Txn.Where("post_id = ?", id).Delete(&models.Comment{})
	c.Txn.Where("id = ?", id).Delete(&models.Post{})

	c.Flash.Success("削除完了")

	return c.Redirect(routes.Post.Index())
}
