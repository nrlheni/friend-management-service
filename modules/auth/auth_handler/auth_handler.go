package auth_handler

import (
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	GetAllUsers(c *gin.Context)
}
