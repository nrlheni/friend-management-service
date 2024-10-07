package friend_dto

type FriendRequestAction struct {
	Requestor string `json:"requestor" binding:"required,email"`
	To        string `json:"to" binding:"required,email"`
}

type FriendListRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type UpdateFriendRequestStatus struct {
	FriendRequestID int    `json:"friendRequestID" binding:"required"`
	Status          string `json:"status" binding:"required"`
}

type BlockFriendRequest struct {
	Requestor string `json:"requestor" binding:"required"`
	Block     string `json:"block" binding:"required,email"`
}
