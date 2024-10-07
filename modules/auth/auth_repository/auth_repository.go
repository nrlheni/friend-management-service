package auth_repository

import (
	"friends-management-api/modules/auth/auth_model"
)

type AuthRepository interface {
	FindByEmail(email string) (*auth_model.User, error)
	CreateUser(user auth_model.User) (*auth_model.User, error)
	GetAllUsers(email string) ([]auth_model.User, error)
}
