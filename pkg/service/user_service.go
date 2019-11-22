package service

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"blogging-app/pkg/models"
	"blogging-app/pkg/repository"
)

// UserService describes the service.
type UserService interface {
	// User Service Functions
	CreateUser(context.Context, models.User) (*models.User, error)
	GetAllUser(context.Context) ([]models.User, error)
	UpdateUser(context.Context, models.User) (*models.User, error)
	DeleteUser(context.Context, string) (*models.User, error)
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
func (b *UserServiceImpl) CreateUser(ctx context.Context, user models.User) (*models.User, error) {
	resp, err := b.userRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return resp, err
}

//GetAllUser retuns users
func (b *UserServiceImpl) GetAllUser(ctx context.Context) ([]models.User, error) {

	users, err := b.userRepository.GetUsers(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	return users, nil
}

//UpdateUser update and returns user
func (b *UserServiceImpl) UpdateUser(ctx context.Context, req models.User) (*models.User, error) {

	filter := bson.M{"_id": bson.M{"$eq": req.ID}}
	updateResp, err := b.userRepository.UpdateUser(ctx, filter, req)
	return updateResp, err
}

//DeleteUser delets an user
func (b *UserServiceImpl) DeleteUser(ctx context.Context, id string) (deleteResp *models.User, err error) {

	deleteResp, err = b.userRepository.DeleteUser(ctx, bson.M{})
	return deleteResp, err
}

//GetUser return user
func (b *UserServiceImpl) GetUser(ctx context.Context, id string) (*models.User, error) {

	createResp, err := b.userRepository.GetUsers(ctx, bson.M{})
	return &createResp[0], err
}

// GetUserProfile return user profile
func (b *UserServiceImpl) GetUserProfile(ctx context.Context, id string) (userProfile *models.UserProfile, err error) {
	// fmt.Println("id", id)
	// userProfile, err = b.userRepository.GetUserProfile(ctx, id)
	return nil, nil
}
