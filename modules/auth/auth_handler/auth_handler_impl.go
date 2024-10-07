package auth_handler

import (
	"friends-management-api/modules/auth/auth_dto"
	"friends-management-api/modules/auth/auth_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandlerImpl struct {
	AuthService auth_service.AuthService
}

func New(authService auth_service.AuthService) AuthHandler {
	return &AuthHandlerImpl{AuthService: authService}
}

// Register handles user registration
func (handler AuthHandlerImpl) Register(c *gin.Context) {
	var registerDTO auth_dto.RegisterRequest
	if err := c.ShouldBindJSON(&registerDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := handler.AuthService.Register(registerDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// Login handles user login
func (handler AuthHandlerImpl) Login(c *gin.Context) {
	var loginDTO auth_dto.LoginRequest
	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := handler.AuthService.Login(loginDTO)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success!",
		"data": user,
	})
}

// Other setup (e.g. Logout, Forgot Password)
func Logout(c *gin.Context) {
	// Implement logout logic here
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
