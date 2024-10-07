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
	var friendRequestList []friend_dto.FriendRequestResult

	subquery := r.DB.Model(&auth_model.User{}).Select("id").Where("email = ?", email)
	result := r.DB.Table("friend_requests fr").
		Select("fr.id, fr.status, u.name AS requester_name, u.email AS requester_email").
		Joins("INNER JOIN users u ON u.id = fr.requester_id").
		Where("fr.requestee_id = (?)", subquery).
		Scan(&friendRequestList)

	if result.Error != nil {
		return nil, result.Error
	}

	return friendRequestList, nil
}

func (r *FriendRepositoryImpl) GetFriendsByEmail(email string) ([]string, error) {
	var friendEmails []string

	subquery := r.DB.Model(&auth_model.User{}).Select("id").Where("email = ?", email)

	result := r.DB.Table("friends f").
		Select("u.email").
		Joins("INNER JOIN users u ON u.id = f.friend_id").
		Where("f.user_id = (?)", subquery).
		Pluck("u.email", &friendEmails)

	if result.Error != nil {
		return nil, result.Error
	}

	return friendEmails, nil
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
