package friend_service

import (
	"errors"
	"friends-management-api/constants"
	"friends-management-api/modules/auth/auth_repository"
	"friends-management-api/modules/friend/friend_dto"
	"friends-management-api/modules/friend/friend_model"
	"friends-management-api/modules/friend/friend_repository"
	"time"
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
	requester, err := service.AuthRepository.FindByEmail(dto.Requestor)
	if err != nil {
		return nil, errors.New(constants.UserNotFound)
	}

	requestee, err := service.AuthRepository.FindByEmail(dto.To)
	if err != nil {
		return nil, errors.New(constants.UserNotFound)
	}

	friendRequest := friend_model.FriendRequests{
		RequesterID: requester.ID,
		RequesteeID: requestee.ID,
		CreatedAt: time.Now(),
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

func (service *FriendServiceImpl) GetFriendRequestList(dto friend_dto.FriendListRequest) (*friend_dto.FriendRequestListResponse, error) {
	friendRequests, err := service.FriendRepository.GetFriendRequestsByEmail(dto.Email)
	if err != nil {
		return nil, err
	}

	list := friend_dto.FriendRequestListResponse{
		Success:  true,
		Requests: friendRequests,
		Count:    len(friendRequests),
	}

	return &list, nil
}

func (service *FriendServiceImpl) GetFriendsList(dto friend_dto.FriendListRequest) (*friend_dto.FriendListResponse, error) {
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
	existingFriendReq, err := service.FriendRepository.GetFriendRequestByID(dto.FriendRequestID)
	if err != nil {
		return nil, err
	}

	existingFriendReq.Status = dto.Status
	_, err = service.FriendRepository.UpdateFriendRequestStatus(*existingFriendReq)
	if err != nil {
		return nil, err
	}

	if dto.Status == "accepted" {
		friend := friend_model.Friends{
			UserID:   existingFriendReq.RequesteeID,
			FriendID: existingFriendReq.RequesterID,
			CreatedAt: time.Now(),
		}

		_, err := service.FriendRepository.CreateFriend(friend)
		if err != nil {
			return nil, err
		}
	}

	response := friend_dto.SuccessfullResponse{
		Success: true,
	}
	return &response, nil
}
