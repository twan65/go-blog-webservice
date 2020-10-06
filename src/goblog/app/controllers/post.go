// ポスト生成、照会、更新、削除
package controllers

import (
	"database/sql"
	"fmt"
	"goblog/app/models"
	"goblog/app/routes"
	"time"

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
	rows, err := c.Txn.Query("select id, title, body, created_at, updated_at from posts order by created_at desc")
	if err != nil {
		panic(err)
	}

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

// Empty Postを生成して画面にレンダリング
func (c Post) New() revel.Result {
	post := models.Post{}
	return c.Render(post)
}

// ポストをDBに保存
func (c Post) Create(title, body string) revel.Result {
	// add Database
	_, err := c.Txn.Exec("insert into posts(title, body, created_at, updated_at) values(?,?,?,?)", title, body, time.Now(), time.Now())

	if err != nil {
		panic(err)
	}

	// ViewにFlashメッセージを渡す。
	c.Flash.Success("ポスト作成を完了")

	return c.Redirect(routes.Post.Index())
}

// ポスト取得
func (c Post) Show(id int) revel.Result {
	post, err := getPost(c.Txn, id)
	if err != nil {
		panic(err)
	}

	return c.Render(post)
}

// ポスト更新データ取得
func (c Post) Edit(id int) revel.Result {
	post, err := getPost(c.Txn, id)
	if err != nil {
		panic(err)
	}

	return c.Render(post)
}

func (c Post) Update(id int, title, body string) revel.Result {
	// ポスト内容を更新
	if _, err := c.Txn.Exec("update posts set title=?, body=?, updated_at=? where id=?", title, body, time.Now(), id); err != nil {
		panic(err)
	}

	// ビューにFlashメッセージを渡す。
	c.Flash.Success("更新完了")

	// ポスト詳細画面に遷移
	return c.Redirect(routes.Post.Show(id))
}

func getPost(txn *sql.Tx, id int) (models.Post, error) {
	post := models.Post{}
	err := txn.QueryRow("select id, title, body, created_at, updated_at from posts where id=?", id).
		Scan(&post.Id, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt)

	switch {
	case err == sql.ErrNoRows:
		return post, fmt.Errorf("No post with that ID - %d.", id)
	case err != nil:
		return post, err
	}
	return post, nil
}
