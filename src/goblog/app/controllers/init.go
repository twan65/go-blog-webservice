package controllers

import (
	"github.com/revel/revel"
)

func init() {
	// アプリケーション起動時、DBを初期化
	revel.OnAppStart(InitDB)
	revel.InterceptMethod((*GormController).Begin, revel.BEFORE)
	revel.InterceptMethod((*GormController).Commit, revel.AFTER)
	revel.InterceptMethod((*GormController).Rollback, revel.FINALLY)

	// 全てのアクションに対してsetCurrentUser関数が行われるように設定
	revel.InterceptMethod((*App).setCurrentUser, revel.BEFORE)

	// checkUserをインタセプターとして登録
	revel.InterceptMethod(Post.CheckUser, revel.BEFORE)
	revel.InterceptMethod(Comment.CheckUser, revel.BEFORE)

	revel.InterceptMethod(User.CheckUser, revel.BEFORE)

	// TODO:ユーザー情報更新は該当ユーザーのみとする。
}
