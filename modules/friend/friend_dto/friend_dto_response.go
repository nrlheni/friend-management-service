package friend_dto

type SuccessfullResponse struct {
	Success bool `json:"success"`
}

type FriendRequestListResponse struct {
	Success  bool                  `json:"success"`
	Requests []FriendRequestResult `json:"requests"`
	Count    int                   `json:"count"`
}

type FriendListResponse struct {
	Friends []FriendsResult `json:"friends"`
}

type FriendRequestResult struct {
	ID             int    `json:"id"`
	Status         string `json:"status"`
	RequesterName  string `json:"requesterName"`
	RequesterEmail string `json:"requesterEmail"`
}

type FriendsResult struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}