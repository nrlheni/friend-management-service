package friend_dto

type FriendRequestAction struct {
	Requester string `json:"requester" binding:"required,email"`
	To        string `json:"to" binding:"required,email"`
}

type ListRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type UpdateFriendRequestStatus struct {
	FriendRequestID int    `json:"friendRequestID" binding:"required"`
	Status          string `json:"status" binding:"required"`
}

type BlockFriendRequest struct {
	Requester string `json:"requester" binding:"required"`
	Block     string `json:"block" binding:"required,email"`
}

type MutualFriendsRequest struct {
	Emails string `form:"emails" binding:"required"`
}