package friend_repository

import (
	"friends-management-api/modules/friend/friend_dto"
	"friends-management-api/modules/friend/friend_model"
)

type FriendRepository interface {
	CreateFriendRequest(friend friend_model.FriendRequests) (*friend_model.FriendRequests, error)
	CreateFriend(friend friend_model.Friends) (*friend_model.Friends, error)
	BlockFriend(blockedFriend friend_model.Blocks) (*friend_model.Blocks, error)
	UpdateFriendRequestStatus(friendRequest friend_model.FriendRequests) (*friend_model.FriendRequests, error)
	GetFriendRequestByID(id int) (*friend_model.FriendRequests, error)
	GetFriendsByEmail(email string) ([]friend_dto.FriendsResult, error)
	GetFriendRequestsByEmail(email string) ([]friend_dto.FriendRequestResult, error)
	GetMutualFriends(email1, email2 string) ([]friend_dto.FriendsResult, error)
	AreFriends(userID, friendID int) (bool, error)
	AlreadyBlocked(userID, friendID int) (bool, error)
	GetBlockedFriends(blockerID int) ([]friend_dto.FriendsResult, error)
}