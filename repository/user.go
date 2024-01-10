package repository

import (
	"log"
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

	// Assuming queries.SelectEmail is a constant string with a valid SQL SELECT statement
	query := queries.SelectEmail

	// If you are using placeholders, make sure to use them properly
	// For example, if your query is "SELECT * FROM users WHERE email = $1", use:
	// query := "SELECT * FROM users WHERE email = $1"
	log.Printf("Executing query: %s\n", queries.SelectEmail)
	err := repo.db.DB.Get(&user, query, email)
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}

	return &user, nil
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
