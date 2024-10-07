package friend

import (
	"friends-management-api/modules/friend/friend_handler"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	handler friend_handler.FriendHandler
}

func New(
	handler friend_handler.FriendHandler) *Routes {
	return &Routes{
		handler: handler,
	}
}

func (r *Routes) Setup(router *gin.Engine) {
	authRoutes := router.Group("/friend")
	{
		authRoutes.POST("/request", r.handler.CreateFriendRequest)
		authRoutes.POST("/request/update", r.handler.UpdateFriendRequestStatus)
		authRoutes.GET("/request", r.handler.GetFriendRequestList)
		authRoutes.GET("/", r.handler.GetFriendsList)
	}
}
