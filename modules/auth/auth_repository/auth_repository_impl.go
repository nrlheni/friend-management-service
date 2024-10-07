package auth_repository

import (
	"friends-management-api/modules/auth/auth_model"

	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	DB *gorm.DB
}

func New(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{DB: db}
}

func (r *AuthRepositoryImpl) FindByEmail(email string) (*auth_model.User, error) {
	var user auth_model.User
	result := r.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *AuthRepositoryImpl) CreateUser(user auth_model.User) (*auth_model.User, error) {
	result := r.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
