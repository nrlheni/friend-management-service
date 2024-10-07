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
	friendRoutes := router.Group("/friend")
	{
		friendRoutes.POST("/request", r.handler.CreateFriendRequest)
		friendRoutes.POST("/request/update", r.handler.UpdateFriendRequestStatus)
		friendRoutes.GET("/request", r.handler.GetFriendRequestList)
		friendRoutes.GET("/", r.handler.GetFriendsList)
		friendRoutes.GET("/mutual", r.handler.GetMutualFriendsList)
		friendRoutes.GET("/block", r.handler.GetBlockedFriendsList)
	}
}
