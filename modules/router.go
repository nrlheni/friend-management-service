package router

import (
	"friends-management-api/modules/auth"
	"friends-management-api/modules/friend"

	"github.com/gin-gonic/gin"
)

type Route interface {
	Setup(router *gin.Engine)
}

type Routes []Route

func NewRoutes(
	auth *auth.Routes,
	friend *friend.Routes,
) Routes {
	return Routes{
		auth,
		friend,
	}
}

func (r Routes) Setup(router *gin.Engine) {
	for _, route := range r {
		route.Setup(router)
	}
}
