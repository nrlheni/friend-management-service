package auth

import (
	"friends-management-api/modules/auth/auth_handler"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	handler auth_handler.AuthHandler
}

func New(
	handler auth_handler.AuthHandler) *Routes {
	return &Routes{
		handler: handler,
	}
}

func (r *Routes) Setup(router *gin.Engine) {
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", r.handler.Register)
		authRoutes.POST("/login", r.handler.Login)
	}
}
