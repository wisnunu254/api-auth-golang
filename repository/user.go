package repository

import (
	"time"

	"github.com/wisnunu254/api-auth-golang/app/auth/model"
	"github.com/wisnunu254/api-auth-golang/pkg/db"
	"github.com/wisnunu254/api-auth-golang/repository/queries"
)

type UsersRepository struct {
	db *db.DB
}

func UsersRepositorys(db *db.DB) *UsersRepository {
	return &UsersRepository{db}
}

func (repo *UsersRepository) ListUsersRepository() ([]*model.User, error) {
	var users []*model.User
	err := repo.db.DB.Select(&users, queries.ListUsers)
	return users, err
}

func (repo *UsersRepository) GetEmailUsersRepository(email string) (*model.User, error) {
	user := model.User{}
	err := repo.db.DB.Get(&user, queries.SelectEmail, email)
	return &user, err
}

func (repo *UsersRepository) GetIDUsersRepository(id string) (*model.User, error) {
	user := model.User{}
	err := repo.db.DB.Get(&user, queries.SelectEmail, id)
	return &user, err
}

func (repo *UsersRepository) InsertUsersRepository(user *model.User) error {
	_, err := repo.db.DB.NamedExec(queries.InsertUsers, user)
	return err
}

func (repo *UsersRepository) UpdateUsersRepository(user *model.User) error {
	_, err := repo.db.DB.NamedExec(queries.UpdateUsers, user)
	return err
}

func (repo *UsersRepository) DeleteUsersRepository(id string) error {
	_, err := repo.db.DB.Exec(queries.DeleteUsers, id, time.Now())
	return err
}
