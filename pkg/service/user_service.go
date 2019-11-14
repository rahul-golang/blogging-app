package service

import (
	"context"
	"fmt"

	"blogging-app/pkg/models"
	"blogging-app/pkg/repository"
)

// UserService describes the service.
type UserService interface {
	// User Service Functions
	CreateUser(context.Context, models.User) (*models.User, error)
	GetAllUser(context.Context) ([]*models.User, error)
	UpdateUser(context.Context, models.User) (*models.User, error)
	DeleteUser(context.Context, string) (*models.DeleteUserResp, error)
	GetUser(context.Context, string) (*models.User, error)
	GetUserProfile(context.Context, string) (*models.UserProfile, error)
}

//UserServiceImpl user services
type UserServiceImpl struct {
	userRepository repository.UserRepository
}

// NewUserServiceImpl returns a naive, stateless implementation of AppService.
func NewUserServiceImpl(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{userRepository: userRepository}
}

//CreateUser return created user response
func (b *UserServiceImpl) CreateUser(ctx context.Context, createReq models.CreateUserReq) (*models.CreateUserResp, error) {
	user, err := b.userRepository.CreateUser(ctx, createReq.User)
	if err != nil {
		return nil, err
	}

	return &models.CreateUserResp{
		Message: "record created sucessfully",
		User:    user,
	}, err
}

//GetAllUser retuns users
func (b *UserServiceImpl) GetAllUser(ctx context.Context) (allRecordResp []*models.User, err error) {

	allRecordResp, err = b.userRepository.AllUsers(ctx)
	fmt.Println(allRecordResp)
	return allRecordResp, err
}

//UpdateUser update and returns user
func (b *UserServiceImpl) UpdateUser(ctx context.Context, upadteReq models.User) (updateResp *models.User, err error) {
	updateResp, err = b.userRepository.UpdateUser(ctx, upadteReq)
	return updateResp, err
}

//DeleteUser delets an user
func (b *UserServiceImpl) DeleteUser(ctx context.Context, id string) (deleteResp *models.DeleteUserResp, err error) {
	deleteResp, err = b.userRepository.DeleteUser(ctx, id)
	return deleteResp, err
}

//GetUser return user
func (b *UserServiceImpl) GetUser(ctx context.Context, id string) (createResp *models.User, err error) {
	fmt.Println("id", id)
	createResp, err = b.userRepository.GetUser(ctx, id)
	return createResp, err
}

// GetUserProfile return user profile
func (b *UserServiceImpl) GetUserProfile(ctx context.Context, id string) (userProfile *models.UserProfile, err error) {
	fmt.Println("id", id)
	userProfile, err = b.userRepository.GetUserProfile(ctx, id)
	return userProfile, err
}
