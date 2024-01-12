package repository

import (
	"database/sql"
	"fmt"
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

	if err != nil {
		if err == sql.ErrNoRows {
			// Handle the case where no rows were found for the given email
			return nil, fmt.Errorf("user not found with email %s", email)
		}
		// Handle other errors
		return nil, err
	}
	return &user, nil
}

func (repo *UsersRepository) GetIDUsersRepository(id string) (*model.User, error) {
	user := model.User{}
	err := repo.db.DB.Get(&user, queries.SelectEmail, id)
	return &user, err
}

func (repo *UsersRepository) InsertUsersRepository(user *model.UserInsert) error {

	_, err := repo.db.DB.Exec(queries.InsertUsers, user.Email, user.Password)
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
