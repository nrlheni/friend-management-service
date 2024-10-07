package friend_repository

import (
	"friends-management-api/modules/friend/friend_dto"
	"friends-management-api/modules/friend/friend_model"
)

type FriendRepository interface {
	CreateFriendRequest(friend friend_model.FriendRequests) (*friend_model.FriendRequests, error)
	CreateFriend(friend friend_model.Friends) (*friend_model.Friends, error)
	UpdateFriendRequestStatus(friendRequest friend_model.FriendRequests) (*friend_model.FriendRequests, error)
	GetFriendRequestByID(id int) (*friend_model.FriendRequests, error)
	GetFriendsByEmail(email string) ([]string, error)
	GetFriendRequestsByEmail(email string) ([]friend_dto.FriendRequestResult, error)
}
