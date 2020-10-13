package models

type User struct {
	Id       int64  `db:"id"`
	Name     string `db:"name"`
	Role     string `db:"role"`
	Username string `db:"username"`
	Password string `db:"password"`
}
