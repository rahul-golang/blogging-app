package repository

import (
	"context"
	"fmt"

	"blogging-app/database"
	"blogging-app/log"
	"blogging-app/pkg/models"
)

//RepositoryInterface implimets all methods in AppRepository
type RepositoryInterface interface {
	CreateUser(ctx context.Context, createReq models.User) (createResp *models.User, err error)
	GetUser(ctx context.Context, id string) (getResp *models.User, err error)
	DeleteUser(ctx context.Context, id string) (deleteResp *models.DeleteUserResp, err error)
	UpdateUser(ctx context.Context, updateReq models.User) (updateResp *models.User, err error)
	AllUsers(ctx context.Context) (getAllResp []*models.User, err error)
	GetUserProfile(ctx context.Context, id string) (*models.UserProfile, error)

	//Blog Repository FUnctions
	CreateBlog(ctx context.Context, blogReq models.Blog) (*models.Blog, error)
	GetAllBlogs(ctx context.Context) ([]*models.Blog, error)
}

// AppRepository **
type AppRepository struct {
	mysqlInterface database.MySQLClientConnInterface
}

//NewAppRepository inject dependancies of DataStore
func NewAppRepository(mysqlInterface database.MySQLClientConnInterface) RepositoryInterface {
	return &AppRepository{mysqlInterface: mysqlInterface}
}

//CreateUser add new record in datastore
func (appRepository *AppRepository) CreateUser(ctx context.Context, user models.User) (*models.User, error) {

	dbConn := appRepository.mysqlInterface.NewClientConnection()

	dbConn.AutoMigrate(&models.User{})
	//createOn := time.Now().In(time.UTC)

	d := dbConn.Create(&user)
	if d.Error != nil {
		return nil, d.Error
	}

	return &user, nil
}

//GetUser serch user bu its id retun user fron database
func (appRepository *AppRepository) GetUser(ctx context.Context, id string) (*models.User, error) {

	dbConn := appRepository.mysqlInterface.NewClientConnection()
	log.Logger(ctx).Info("in Get users repository mothod ", id)
	user := models.User{}
	err := dbConn.Where("id=?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// DeleteUser delete User From Database
func (appRepository *AppRepository) DeleteUser(ctx context.Context, id string) (*models.DeleteUserResp, error) {
	dbConn := appRepository.mysqlInterface.NewClientConnection()

	err := dbConn.Where("id=?", id).Delete(&models.User{}).Error
	if err != nil {
		return nil, err
	}
	return &models.DeleteUserResp{
		Message: "Records Deleted Sucessfully",
		ID:      id,
	}, nil
}

//UpdateUser update user in database
func (appRepository *AppRepository) UpdateUser(ctx context.Context, user models.User) (*models.User, error) {

	dbConn := appRepository.mysqlInterface.NewClientConnection()
	err := dbConn.Model(&models.User{}).Where("id=?", user.ID).Update(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

//AllUsers return all Users from database
func (appRepository *AppRepository) AllUsers(ctx context.Context) (getAllResp []*models.User, err error) {

	dbConn := appRepository.mysqlInterface.NewClientConnection()
	//log.Logger(ctx).Info("in all users repository mothod ")

	users := []models.User{}
	//err = dbConn.Model(&models.User{}).Association("blogs").Find(&users).Error

	//err = dbConn.Model(model.User{}).Find(&users).Error
	err = dbConn.Debug().Preload("Blogs").Find(&users).Error
	//err = dbConn.Model(&models.User{}).Association("blogs").Find(&users).Error

	//log.Logger(ctx).Info("in all users service mothod ")
	//err = dbConn.Find(&users).Error
	fmt.Println(users)
	if err != nil {
		return nil, err
	}
	return getAllResp, nil
}

//GetUserProfile returns user profile data
func (appRepository *AppRepository) GetUserProfile(ctx context.Context, id string) (*models.UserProfile, error) {

	dbConn := appRepository.mysqlInterface.NewClientConnection()
	fmt.Println(id)
	//log.Logger(ctx).Info("in Get users repository mothod ", id)
	userProfile := models.UserProfile{}
	err := dbConn.Where("id=?", id).First(&userProfile).Error
	if err != nil {
		return nil, err
	}
	return &userProfile, nil
}

// CreateBlog create blog in database and retub Created Blog
func (appRepository *AppRepository) CreateBlog(ctx context.Context, blogReq models.Blog) (*models.Blog, error) {

	dbConn := appRepository.mysqlInterface.NewClientConnection()

	dbConn.AutoMigrate(&models.Blog{})
	//createOn := time.Now().In(time.UTC)
	fmt.Println(blogReq)

	d := dbConn.Create(&blogReq)
	if d.Error != nil {
		return nil, d.Error
	}

	return &blogReq, nil

}

//GetAllBlogs return all Users from database
func (appRepository *AppRepository) GetAllBlogs(ctx context.Context) ([]*models.Blog, error) {

	dbConn := appRepository.mysqlInterface.NewClientConnection()
	//log.Logger(ctx).Info("in all users repository mothod ")

	blogs := []*models.Blog{}
	err := dbConn.Find(&blogs).Error
	//log.Logger(ctx).Info("in all users service mothod ")
	if err != nil {
		return nil, err
	}
	return blogs, nil
}
