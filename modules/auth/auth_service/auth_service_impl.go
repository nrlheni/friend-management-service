package auth_service

import (
	"errors"
	"friends-management-api/constants"
	"friends-management-api/modules/auth/auth_dto"
	"friends-management-api/modules/auth/auth_model"
	"friends-management-api/modules/auth/auth_repository"
	"friends-management-api/utils"
	"time"
)

type AuthServiceImpl struct {
	AuthRepository auth_repository.AuthRepository
}

func New(authRepo auth_repository.AuthRepository) AuthService {
	return &AuthServiceImpl{AuthRepository: authRepo}
}

func (service *AuthServiceImpl) Register(dto auth_dto.RegisterRequest) (*auth_dto.SuccessfullResponse, error) {
	// Check if user already exists
	existingUser, _ := service.AuthRepository.FindByEmail(dto.Email)
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(dto.Password)
	if err != nil {
		return nil, err
	}

	// Create new user
	user := auth_model.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: hashedPassword,
		CreatedAt: time.Now(),
	}

	_, err = service.AuthRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	userResponse := auth_dto.SuccessfullResponse{
		Success: true,
	}

	return &userResponse, nil
}

func (service *AuthServiceImpl) Login(dto auth_dto.LoginRequest) (*auth_dto.LoginResponse, error) {
	user, err := service.AuthRepository.FindByEmail(dto.Email)
	if err != nil || user == nil {
		return nil, errors.New(constants.InvalidCredentials)
	}

	if !utils.CheckPasswordHash(dto.Password, user.Password) {
		return nil, errors.New(constants.UserNotFound)
	}

	response := auth_dto.LoginResponse{
		ID: uint(user.ID),
		Name: user.Name,
		Email: user.Email,
	}

	return &response, nil
}
