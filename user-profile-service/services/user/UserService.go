package user

import (
	userModel "github.com/Aranyak-Ghosh/spotigo/user_profile/models/database/user"
	httpModels "github.com/Aranyak-Ghosh/spotigo/user_profile/models/http"
	errorhandling "github.com/Aranyak-Ghosh/spotigo/user_profile/utils/errorHandling"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type UserService interface {
	GetUser(id string, transactionId string) (userModel.User, *httpModels.ErrorResponse)
	ListUsers(searchParam string, limit int, offset int, transactionId string) (httpModels.PagedResult, *httpModels.ErrorResponse)
	CreateUser(user *userModel.User, transactionId string) *httpModels.ErrorResponse
	UpdateUser(user *userModel.User, transactionId string) *httpModels.ErrorResponse
	DeleteUser(id string, transactionId string) *httpModels.ErrorResponse
	FollowUser(follower_id string, following_id string, transactionId string) *httpModels.ErrorResponse
	UnFollowUser(follower_id string, following_id string, transactionId string) *httpModels.ErrorResponse
	ListFollowers(user_id string, searchParam string, limit int, offset int, transactionId string) (httpModels.PagedResult, *httpModels.ErrorResponse)
	ListFollowing(user_id string, searchParam string, limit int, offset int, transactionId string) (httpModels.PagedResult, *httpModels.ErrorResponse)
}

type userService struct {
	userRepository *userModel.UserRepository
	logger         *zap.SugaredLogger
}

func (service userService) GetUser(id string, transactionId string) (userModel.User, *httpModels.ErrorResponse) {
	service.logger.Infow("GetUser", "TransactionId", transactionId, "id", id)
	data, err := service.userRepository.GetById(id)

	var errorResponse *httpModels.ErrorResponse = nil
	if err != nil {
		errorResponse = errorhandling.HandleDatabaseError(err, transactionId)
	}

	return data, errorResponse
}

func (service userService) ListUsers(searchParam string, limit int, offset int, transactionId string) (httpModels.PagedResult, *httpModels.ErrorResponse) {
	service.logger.Infow("ListUsers", "TransactionId", transactionId, "searchParam", searchParam, "limit", limit, "offset", offset)
	data, count, err := service.userRepository.List(searchParam, offset, limit)

	var errorResponse *httpModels.ErrorResponse = nil
	if err != nil {
		errorResponse = errorhandling.HandleDatabaseError(err, transactionId)
		return httpModels.PagedResult{}, errorResponse
	}

	return httpModels.PagedResult{Count: count, Data: data}, errorResponse
}

func (service userService) CreateUser(user *userModel.User, transactionId string) *httpModels.ErrorResponse {
	service.logger.Infow("CreateUser", "TransactionId", transactionId, "user", user)
	err := service.userRepository.Create(user)

	var errorResponse *httpModels.ErrorResponse = nil
	if err != nil {
		errorResponse = errorhandling.HandleDatabaseError(err, transactionId)
	}

	return errorResponse
}

func (service userService) UpdateUser(user *userModel.User, transactionId string) *httpModels.ErrorResponse {
	service.logger.Infow("UpdateUser", "TransactionId", transactionId, "user", user)
	userToUpdate, err := service.userRepository.GetById(user.ID)

	var errorResponse *httpModels.ErrorResponse = nil

	if err != nil {
		errorResponse = errorhandling.HandleDatabaseError(err, transactionId)
		return errorResponse
	}

	user.CreatedAt = userToUpdate.CreatedAt

	err = service.userRepository.Update(user)

	if err != nil {
		errorResponse = errorhandling.HandleDatabaseError(err, transactionId)
	}

	return errorResponse
}

func (service userService) DeleteUser(id string, transactionId string) *httpModels.ErrorResponse {
	service.logger.Infow("DeleteUser", "TransactionId", transactionId, "id", id)
	err := service.userRepository.Delete(&userModel.User{ID: id})

	var errorResponse *httpModels.ErrorResponse = nil
	if err != nil {
		errorResponse = errorhandling.HandleDatabaseError(err, transactionId)
	}

	return errorResponse
}

func (service userService) FollowUser(follower_id string, following_id string, transactionId string) *httpModels.ErrorResponse {
	service.logger.Infow("FollowUser", "TransactionId", transactionId, "follower_id", follower_id, "following_id", following_id)
	err := service.userRepository.Follow(&userModel.User{ID: follower_id}, &userModel.User{ID: following_id})

	var errorResponse *httpModels.ErrorResponse = nil
	if err != nil {
		errorResponse = errorhandling.HandleDatabaseError(err, transactionId)
	}

	return errorResponse
}

func (service userService) UnFollowUser(follower_id string, following_id string, transactionId string) *httpModels.ErrorResponse {
	service.logger.Infow("UnFollowUser", "TransactionId", transactionId, "follower_id", follower_id, "following_id", following_id)
	err := service.userRepository.UnFollow(&userModel.User{ID: follower_id}, &userModel.User{ID: following_id})

	var errorResponse *httpModels.ErrorResponse = nil
	if err != nil {
		errorResponse = errorhandling.HandleDatabaseError(err, transactionId)
	}

	return errorResponse
}

func (service userService) ListFollowers(user_id string, searchParam string, limit int, offset int, transactionId string) (httpModels.PagedResult, *httpModels.ErrorResponse) {
	service.logger.Infow("ListFollowers", "TransactionId", transactionId, "user_id", user_id, "searchParam", searchParam, "limit", limit, "offset", offset)
	data, count, err := service.userRepository.ListFollowers(user_id, searchParam, offset, limit)

	var errorResponse *httpModels.ErrorResponse = nil

	if err != nil {
		errorResponse = errorhandling.HandleDatabaseError(err, transactionId)
		return httpModels.PagedResult{}, errorResponse
	}

	return httpModels.PagedResult{Count: count, Data: data}, errorResponse
}

func (service userService) ListFollowing(user_id string, searchParam string, limit int, offset int, transactionId string) (httpModels.PagedResult, *httpModels.ErrorResponse) {
	service.logger.Infow("ListFollowing", "TransactionId", transactionId, "user_id", user_id, "searchParam", searchParam, "limit", limit, "offset", offset)
	data, count, err := service.userRepository.ListFollowing(user_id, searchParam, offset, limit)

	var errorResponse *httpModels.ErrorResponse = nil

	if err != nil {
		errorResponse = errorhandling.HandleDatabaseError(err, transactionId)
		return httpModels.PagedResult{}, errorResponse
	}

	return httpModels.PagedResult{Count: count, Data: data}, errorResponse
}

func NewUserService(userRepository *userModel.UserRepository, logger *zap.SugaredLogger) UserService {

	var ret UserService = &userService{
		userRepository: userRepository,
		logger:         logger,
	}
	return ret
}

var Module = fx.Option(fx.Provide(NewUserService))
