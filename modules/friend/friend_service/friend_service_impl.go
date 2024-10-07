package friend_service

import (
	"errors"
	"fmt"
	"friends-management-api/constants"
	"friends-management-api/modules/auth/auth_repository"
	"friends-management-api/modules/friend/friend_dto"
	"friends-management-api/modules/friend/friend_model"
	"friends-management-api/modules/friend/friend_repository"
	"strings"

	"gorm.io/gorm"
)

type FriendServiceImpl struct {
	FriendRepository friend_repository.FriendRepository
	AuthRepository   auth_repository.AuthRepository
}

func New(
	friendRepo friend_repository.FriendRepository,
	authRepo auth_repository.AuthRepository,
) FriendService {
	return &FriendServiceImpl{
		FriendRepository: friendRepo,
		AuthRepository:   authRepo,
	}
}

func (service *FriendServiceImpl) CreateFriendRequest(dto friend_dto.FriendRequestAction) (*friend_dto.SuccessfullResponse, error) {
	requester, err := service.AuthRepository.FindByEmail(dto.Requester)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(constants.UserNotFound)
		}
		return nil, err
	}

	requestee, err := service.AuthRepository.FindByEmail(dto.To)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(constants.UserNotFound)
		}
		return nil, err
	}

	friendRequest := friend_model.FriendRequests{
		RequesterID: requester.ID,
		RequesteeID: requestee.ID,
	}

	_, err = service.FriendRepository.CreateFriendRequest(friendRequest)
	if err != nil {
		return nil, err
	}

	response := friend_dto.SuccessfullResponse{
		Success: true,
	}

	return &response, nil
}

func (service *FriendServiceImpl) GetFriendRequestList(dto friend_dto.ListRequest) (*friend_dto.FriendRequestListResponse, error) {
	friendRequests, err := service.FriendRepository.GetFriendRequestsByEmail(dto.Email)
	if err != nil {
		return nil, err
	}
	fmt.Println(friendRequests)

	list := friend_dto.FriendRequestListResponse{
		Success:  true,
		Requests: friendRequests,
		Count:    len(friendRequests),
	}

	return &list, nil
}

func (service *FriendServiceImpl) GetFriendsList(dto friend_dto.ListRequest) (*friend_dto.FriendListResponse, error) {
	friends, err := service.FriendRepository.GetFriendsByEmail(dto.Email)
	if err != nil {
		return nil, err
	}

	list := friend_dto.FriendListResponse{
		Friends: friends,
	}

	return &list, nil
}

func (service *FriendServiceImpl) UpdateFriendRequestStatus(dto friend_dto.UpdateFriendRequestStatus) (*friend_dto.SuccessfullResponse, error) {
	// Retrieve the existing friend request by ID
	existingFriendReq, err := service.FriendRepository.GetFriendRequestByID(dto.FriendRequestID)
	if err != nil {
		return nil, err
	}

	// Update the status of the friend request
	existingFriendReq.Status = dto.Status
	_, err = service.FriendRepository.UpdateFriendRequestStatus(*existingFriendReq)
	if err != nil {
		return nil, err
	}

	// If the friend request is accepted, check if they are already friends
	if dto.Status == "accepted" {
		areFriends, err := service.FriendRepository.AreFriends(existingFriendReq.RequesteeID, existingFriendReq.RequesterID)
		if err != nil {
			return nil, err
		}

		// If they are not already friends, create a bidirectional friendship
		if !areFriends {
			friend1 := friend_model.Friends{
				UserID:   existingFriendReq.RequesteeID,
				FriendID: existingFriendReq.RequesterID,
			}

			friend2 := friend_model.Friends{
				UserID:   existingFriendReq.RequesterID,
				FriendID: existingFriendReq.RequesteeID,
			}

			// Add both friends (bidirectional relationship)
			_, err := service.FriendRepository.CreateFriend(friend1)
			if err != nil {
				return nil, err
			}

			_, err = service.FriendRepository.CreateFriend(friend2)
			if err != nil {
				return nil, err
			}
		}
	}

	// Return successful response
	response := friend_dto.SuccessfullResponse{
		Success: true,
	}

	return &response, nil
}

func (service *FriendServiceImpl) GetMutualFriendsList(dto friend_dto.MutualFriendsRequest) (*friend_dto.FriendListResponse, error) {
	emails := strings.Split(dto.Emails, ",")
	friends, err := service.FriendRepository.GetMutualFriends(emails[0], emails[1])
	if err != nil {
		return nil, err
	}

	list := friend_dto.FriendListResponse{
		Friends: friends,
	}

	return &list, nil
}

func (service *FriendServiceImpl) BlockFriend(dto friend_dto.BlockFriendRequest) (*friend_dto.SuccessfullResponse, error) {
	blocker, err := service.AuthRepository.FindByEmail(dto.Requester)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(constants.UserNotFound)
		}
		return nil, err
	}

	blocked, err := service.AuthRepository.FindByEmail(dto.Block)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(constants.UserNotFound)
		}
		return nil, err
	}

	alreadyBlocked, err := service.FriendRepository.AlreadyBlocked(blocker.ID, blocked.ID)
	if err != nil {
		return nil, err
	}

	if !alreadyBlocked {
		blockedFriend := friend_model.Blocks{
			BlockerID: blocker.ID,
			BlockedID: blocked.ID,
		}

		_, err := service.FriendRepository.BlockFriend(blockedFriend)
		if err != nil {
			return nil, err
		}

		blocked.IsBlocked = true
		_, err = service.AuthRepository.UpdateUserStatus(*blocked)
		if err != nil {
			return nil, err
		}

		err = service.FriendRepository.DeleteFriendship(blocker.ID, blocked.ID)
		if err != nil {
			return nil, err
		}
	}

	response := friend_dto.SuccessfullResponse{
		Success: true,
	}

	return &response, nil
}

func (service *FriendServiceImpl) GetBlockedFriends(dto friend_dto.ListRequest) ([]friend_dto.FriendsResult, error) {
	blocker, err := service.AuthRepository.FindByEmail(dto.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(constants.UserNotFound)
		}
		return nil, err
	}

	blockedUsers, err := service.FriendRepository.GetBlockedFriends(blocker.ID)
	if err != nil {
		return nil, err
	}

	return blockedUsers, nil
}