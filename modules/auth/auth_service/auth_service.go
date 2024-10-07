package auth_service

import (
	"friends-management-api/modules/auth/auth_dto"
)

type AuthService interface {
	Register(registerRequest auth_dto.RegisterRequest) (*auth_dto.SuccessfullResponse, error)
	Login(loginRequest auth_dto.LoginRequest) (*auth_dto.LoginResponse, error)
	GetAllUsers(email string) ([]auth_dto.UsersResponse, error)
}
