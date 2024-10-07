package friend_service

import (
	"friends-management-api/modules/friend/friend_dto"
)

type FriendService interface {
	CreateFriendRequest(dto friend_dto.FriendRequestAction) (*friend_dto.SuccessfullResponse, error)
	UpdateFriendRequestStatus(dto friend_dto.UpdateFriendRequestStatus) (*friend_dto.SuccessfullResponse, error)
	GetFriendRequestList(dto friend_dto.ListRequest) (*friend_dto.FriendRequestListResponse, error)
	GetFriendsList(dto friend_dto.ListRequest) (*friend_dto.FriendListResponse, error)
	GetMutualFriendsList(dto friend_dto.MutualFriendsRequest) (*friend_dto.FriendListResponse, error)
	GetBlockedFriends(dto friend_dto.ListRequest) ([]friend_dto.FriendsResult, error)
	BlockFriend(dto friend_dto.BlockFriendRequest) (*friend_dto.SuccessfullResponse, error)
}
