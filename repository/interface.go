package repository

import "github.com/wisnunu254/api-auth-golang/app/auth/model"

type UsersRepositoryInterface interface {
	ListUsersRepository() ([]*model.User, error)
	GetEmailUsersRepository(email string) (*model.User, error)
	GetIDUsersRepository(id string) (*model.User, error)
	InsertUsersRepository(user *model.UserInsert) error
	UpdateUsersRepository(user *model.User) error
	DeleteUsersRepository(id string) error
}
