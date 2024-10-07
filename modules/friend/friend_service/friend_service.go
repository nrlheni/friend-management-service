package friend_service

import (
	"friends-management-api/modules/friend/friend_dto"
)

type FriendService interface {
	CreateFriendRequest(dto friend_dto.FriendRequestAction) (*friend_dto.SuccessfullResponse, error)
	UpdateFriendRequestStatus(dto friend_dto.UpdateFriendRequestStatus) (*friend_dto.SuccessfullResponse, error)
	GetFriendRequestList(dto friend_dto.FriendListRequest) (*friend_dto.FriendRequestListResponse, error)
	GetFriendsList(dto friend_dto.FriendListRequest) (*friend_dto.FriendListResponse, error)
}
