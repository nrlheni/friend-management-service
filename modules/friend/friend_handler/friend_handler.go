package friend_handler

import (
	"github.com/gin-gonic/gin"
)

type FriendHandler interface {
	CreateFriendRequest(c *gin.Context)
	UpdateFriendRequestStatus(c *gin.Context)
	GetFriendRequestList(c *gin.Context)
	GetFriendsList(c *gin.Context)
	GetMutualFriendsList(c *gin.Context)
	GetBlockedFriendsList(c *gin.Context)
	BlockFriend(c *gin.Context)
}
