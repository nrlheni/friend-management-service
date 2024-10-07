package friend_handler

import (
	"friends-management-api/constants"
	"friends-management-api/modules/friend/friend_dto"
	"friends-management-api/modules/friend/friend_service"
	"friends-management-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FriendHandlerImpl struct {
	FriendService friend_service.FriendService
}

func New(friendService friend_service.FriendService) FriendHandler {
	return &FriendHandlerImpl{FriendService: friendService}
}

func (handler FriendHandlerImpl) CreateFriendRequest(c *gin.Context) {
	var dto friend_dto.FriendRequestAction
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	friendRequest, err := handler.FriendService.CreateFriendRequest(dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.SuccessResponse(c, 200, constants.CreateFriendRequest, friendRequest)
}

func (handler FriendHandlerImpl) UpdateFriendRequestStatus(c *gin.Context) {
	var dto friend_dto.UpdateFriendRequestStatus
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedFriendRequest, err := handler.FriendService.UpdateFriendRequestStatus(dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.SuccessResponse(c, 200, constants.UpdateFriendRequestStatus, updatedFriendRequest)
}

func (handler FriendHandlerImpl) GetFriendRequestList(c *gin.Context) {
	email := c.Query("email")

	dto := friend_dto.ListRequest{
		Email: email,
	}

	friendRequest, err := handler.FriendService.GetFriendRequestList(dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.SuccessResponse(c, 200, constants.GetListFriendRequest, friendRequest)
}

func (handler FriendHandlerImpl) GetFriendsList(c *gin.Context) {
	email := c.Query("email")

	dto := friend_dto.ListRequest{
		Email: email,
	}

	friends, err := handler.FriendService.GetFriendsList(dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.SuccessResponse(c, 200, constants.GetListFriend, friends)
}

func (handler FriendHandlerImpl) GetMutualFriendsList(c *gin.Context) {
	var dto friend_dto.MutualFriendsRequest
	if err := c.ShouldBindQuery(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	friends, err := handler.FriendService.GetMutualFriendsList(dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.SuccessResponse(c, 200, constants.GetListFriend, friends)
}

func (handler FriendHandlerImpl) GetBlockedFriendsList(c *gin.Context) {
	email := c.Query("email")

	dto := friend_dto.ListRequest{
		Email: email,
	}

	blockedFriends, err := handler.FriendService.GetBlockedFriends(dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.SuccessResponse(c, 200, constants.GetBlockedListFriend, blockedFriends)
}

func (handler FriendHandlerImpl) BlockFriend(c *gin.Context) {
	var dto friend_dto.BlockFriendRequest
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blockedFriend, err := handler.FriendService.BlockFriend(dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.SuccessResponse(c, 200, constants.BlockFriend, blockedFriend)
}