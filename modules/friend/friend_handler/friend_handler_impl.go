package friend_handler

import (
	"friends-management-api/modules/friend/friend_dto"
	"friends-management-api/modules/friend/friend_service"
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

	c.JSON(http.StatusCreated, friendRequest)
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

	c.JSON(http.StatusOK, updatedFriendRequest)
}

func (handler FriendHandlerImpl) GetFriendRequestList(c *gin.Context) {
	email := c.Query("email")

	dto := friend_dto.FriendListRequest{
		Email: email,
	}

	friendRequest, err := handler.FriendService.GetFriendRequestList(dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, friendRequest)
}

func (handler FriendHandlerImpl) GetFriendsList(c *gin.Context) {
	email := c.Query("email")

	dto := friend_dto.FriendListRequest{
		Email: email,
	}

	friends, err := handler.FriendService.GetFriendsList(dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, friends)
}
