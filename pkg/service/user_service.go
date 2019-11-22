package service

import (
	"blogging-app/log"
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

	log.Logger(ctx).Info("CreateUser : ", user)

	resp, err := b.userRepository.CreateUser(ctx, user)
	if err != nil {
		log.Logger(ctx).Error(err)
		return nil, err
	}

	return resp, err
}

//GetAllUser retuns users
func (b *UserServiceImpl) GetAllUser(ctx context.Context) ([]models.User, error) {
	log.Logger(ctx).Error("GetAllUser : ")
	users, err := b.userRepository.GetUsers(ctx, bson.M{})
	if err != nil {
		log.Logger(ctx).Error(err)
		return nil, err
	}
	return users, nil
}

//UpdateUser update and returns user
func (b *UserServiceImpl) UpdateUser(ctx context.Context, req models.User) (*models.User, error) {

	log.Logger(ctx).Info("UpdateUser: ", req)
	//Created filter to find and update
	filter := bson.M{"_id": bson.M{"$eq": req.ID}}
	updateResp, err := b.userRepository.UpdateUser(ctx, filter, req)
	if err != nil {
		log.Logger(ctx).Error(err)
		return nil, err
	}

	return updateResp, nil
}

//DeleteUser delets an user
func (b *UserServiceImpl) DeleteUser(ctx context.Context, id string) (deleteResp *models.User, err error) {
	log.Logger(ctx).Info("DeleteUser: ", id)
	//Created filter to find and update
	filter := bson.M{"_id": bson.M{"$eq": id}}
	deleteResp, err = b.userRepository.DeleteUser(ctx, filter)
	return deleteResp, err
}

//GetUser return user
func (b *UserServiceImpl) GetUser(ctx context.Context, id string) (*models.User, error) {

	log.Logger(ctx).Info("GetUser: ", id)
	//Created filter to find and update
	filter := bson.M{"_id": bson.M{"$eq": id}}
	createResp, err := b.userRepository.GetUsers(ctx, filter)
	return &createResp[0], err
}

// GetUserProfile return user profile
func (b *UserServiceImpl) GetUserProfile(ctx context.Context, id string) (userProfile *models.UserProfile, err error) {
	// fmt.Println("id", id)
	// userProfile, err = b.userRepository.GetUserProfile(ctx, id)
	return nil, nil
}
