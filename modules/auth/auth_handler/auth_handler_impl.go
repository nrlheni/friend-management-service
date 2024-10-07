package auth_handler

import (
	"friends-management-api/constants"
	"friends-management-api/modules/auth/auth_dto"
	"friends-management-api/modules/auth/auth_service"
	"friends-management-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandlerImpl struct {
	AuthService auth_service.AuthService
}

func New(authService auth_service.AuthService) AuthHandler {
	return &AuthHandlerImpl{AuthService: authService}
}

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

	utils.SuccessResponse(c, 200, constants.RegisterUser, user)
}

func (handler AuthHandlerImpl) Login(c *gin.Context) {
	var loginDTO auth_dto.LoginRequest
	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := handler.AuthService.Login(loginDTO)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	utils.SuccessResponse(c, 200, constants.LoginUser, token)
}

func (handler AuthHandlerImpl) GetAllUsers(c *gin.Context) {
	email := c.DefaultQuery("email", "")

	users, err := handler.AuthService.GetAllUsers(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	utils.SuccessResponse(c, 200, constants.GetListFriend, users)
}