package model

type User struct {
	ID        int64  `db:"id"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	Type      string `db:"type"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
	DeletedAt string `db:"deleted_at"`
}

func UserModel() *User {
	return &User{}
}

type AuthModelLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
