package model

type User struct {
	ID       int64  `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

type UserInsert struct {
	Email    string `db:"email"`
	Password string `db:"password"`
}

func UserModel() *User {
	return &User{}
}

type AuthModelLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthModelRegister struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
