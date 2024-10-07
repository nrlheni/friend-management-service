package friend_repository

import (
	"friends-management-api/modules/auth/auth_model"
	"friends-management-api/modules/friend/friend_dto"
	"friends-management-api/modules/friend/friend_model"

	"gorm.io/gorm"
)

type FriendRepositoryImpl struct {
	DB *gorm.DB
}

func New(db *gorm.DB) FriendRepository {
	return &FriendRepositoryImpl{DB: db}
}

func (r *FriendRepositoryImpl) GetFriendRequestsByEmail(email string) ([]friend_dto.FriendRequestResult, error) {
	friendRequestList := make([]friend_dto.FriendRequestResult, 0)

	subquery := r.DB.Model(&auth_model.User{}).Select("id").Where("email = ?", email)
	result := r.DB.Table("friend_requests fr").
		Select("fr.id, fr.status, u.name AS requester_name, u.email AS requester_email").
		Joins("INNER JOIN users u ON u.id = fr.requester_id").
		Where("fr.requestee_id = (?)", subquery).
		Where("status = 'pending'").
		Scan(&friendRequestList)

	if result.Error != nil {
		return nil, result.Error
	}

	return friendRequestList, nil
}

func (r *FriendRepositoryImpl) GetFriendsByEmail(email string) ([]friend_dto.FriendsResult, error) {
	friends := make([]friend_dto.FriendsResult, 0)

	subquery := r.DB.Model(&auth_model.User{}).Select("id").Where("email = ?", email)

	result := r.DB.Table("friends f").
		Select("u.id, u.name, u.email").
		Joins("INNER JOIN users u ON u.id = f.friend_id").
		Where("f.user_id = (?)", subquery).
		Scan(&friends)

	if result.Error != nil {
		return nil, result.Error
	}

	return friends, nil
}

func (r *FriendRepositoryImpl) CreateFriendRequest(friendRequest friend_model.FriendRequests) (*friend_model.FriendRequests, error) {
	result := r.DB.Create(&friendRequest)
	if result.Error != nil {
		return nil, result.Error
	}
	return &friendRequest, nil
}

func (r *FriendRepositoryImpl) CreateFriend(friend friend_model.Friends) (*friend_model.Friends, error) {
	result := r.DB.Create(&friend)
	if result.Error != nil {
		return nil, result.Error
	}
	return &friend, nil
}

func (r *FriendRepositoryImpl) BlockFriend(blockedFriend friend_model.Blocks) (*friend_model.Blocks, error) {
	result := r.DB.Create(&blockedFriend)
	if result.Error != nil {
		return nil, result.Error
	}
	return &blockedFriend, nil
}

func (r *FriendRepositoryImpl) UpdateFriendRequestStatus(friendRequest friend_model.FriendRequests) (*friend_model.FriendRequests, error) {
	result := r.DB.Save(&friendRequest)
	if result.Error != nil {
		return nil, result.Error
	}
	return &friendRequest, nil
}

func (r *FriendRepositoryImpl) GetFriendRequestByID(id int) (*friend_model.FriendRequests, error) {
	var friendRequest friend_model.FriendRequests

	result := r.DB.First(&friendRequest, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}

	return &friendRequest, nil
}

func (r *FriendRepositoryImpl) GetMutualFriends(email1, email2 string) ([]friend_dto.FriendsResult, error) {
	mutualFriends := make([]friend_dto.FriendsResult, 0)

	subquery1 := r.DB.Model(&auth_model.User{}).Select("id").Where("email = ?", email1)
	subquery2 := r.DB.Model(&auth_model.User{}).Select("id").Where("email = ?", email2)

	result := r.DB.Table("friends f1").
		Select("u.id, u.email, u.name").
		Joins("INNER JOIN friends f2 ON f1.friend_id = f2.friend_id").
		Joins("INNER JOIN users u ON u.id = f1.friend_id").
		Where("f1.user_id = (?)", subquery1).
		Where("f2.user_id = (?)", subquery2).
		Scan(&mutualFriends)

	if result.Error != nil {
		return nil, result.Error
	}

	return mutualFriends, nil
}

func (r *FriendRepositoryImpl) AreFriends(userID, friendID int) (bool, error) {
	var count int64
	result := r.DB.Model(&friend_model.Friends{}).
		Where("(user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)", userID, friendID, friendID, userID).
		Count(&count)

	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}

func (r *FriendRepositoryImpl) AlreadyBlocked(userID, friendID int) (bool, error) {
	var count int64
	result := r.DB.Model(&friend_model.Blocks{}).
		Where("(blocker_id = ? AND blocked_id = ?) OR (blocker_id = ? AND blocked_id = ?)", userID, friendID, friendID, userID).
		Count(&count)

	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}

func (r *FriendRepositoryImpl) GetBlockedFriends(blockerID int) ([]friend_dto.FriendsResult, error) {
	blockedUsers := make([]friend_dto.FriendsResult, 0)

	result := r.DB.Table("blocks b").
		Select("u.id, u.email, u.name").
		Joins("INNER JOIN users u ON b.blocked_id = u.id").
		Where("b.blocker_id = ?", blockerID).
		Scan(&blockedUsers)

	if result.Error != nil {
		return nil, result.Error
	}
	return blockedUsers, nil
}