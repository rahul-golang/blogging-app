package service

import (
	"context"
	"reflect"
	"testing"

	"blogging-app/pkg/models"
	"blogging-app/pkg/repository"
)

type MokeUserRepository struct {
}

func NewMokeUserRepository() repository.UserRepositoryInterface {
	return &MokeUserRepository{}
}

func (mokeUserRepository MokeUserRepository) Create(ctx context.Context, createReq models.User) (*models.User, error) {

	return &createReq, nil
}
func (mokeUserRepository MokeUserRepository) Get(ctx context.Context, id string) (getResp *models.User, err error) {
	return
}
func (mokeUserRepository MokeUserRepository) Delete(ctx context.Context, id string) (deleteResp *models.DeleteUserResp, err error) {
	return
}
func (mokeUserRepository MokeUserRepository) Update(ctx context.Context, updateReq models.User) (updateResp *models.User, err error) {
	return
}
func (mokeUserRepository MokeUserRepository) All(ctx context.Context) (getAllResp []*models.User, err error) {
	return
}

// type MokeUserRepositoryInterface interface {
// 	Create(ctx context.Context, createReq models.User) (createResp *models.User, err error)
// 	Get(ctx context.Context, id string) (getResp *models.User, err error)
// 	Delete(ctx context.Context, id string) (deleteResp *models.DeleteUserResp, err error)
// 	Update(ctx context.Context, updateReq models.User) (updateResp *models.User, err error)
// 	All(ctx context.Context) (getAllResp []*models.User, err error)
// }

func Test_basicUsersService_CreateUser(t *testing.T) {
	//usersService := NewBasicUsersService(NewMokeUserRepository())
	type args struct {
		ctx       context.Context
		createReq models.CreateUserReq
	}
	tests := []struct {
		name         string
		usersService UsersService
		args         args
		//want    *models.CreateUserResp
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"'Test 1 : ",
			NewBasicUsersService(NewMokeUserRepository()),
			args{
				context.Background(),
				models.CreateUserReq{
					models.User{},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.usersService.CreateUser(tt.args.ctx, tt.args.createReq)
			if (err != nil) != tt.wantErr {
				t.Errorf("basicUsersService.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("basicUsersService.CreateUser() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func Test_basicUsersService_GetAllUser(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		b    UsersService
		args args
		//	wantAllRecordResp []*models.User
		wantErr bool
	}{
		{
			"'Test 1 : ",
			NewBasicUsersService(NewMokeUserRepository()),
			args{
				context.Background(),
			},
			false,
		},
		{
			"'Test 1 : ",
			NewBasicUsersService(NewMokeUserRepository()),
			args{
				context.Background(),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.b.GetAllUser(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("basicUsersService.GetAllUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(gotAllRecordResp, tt.wantAllRecordResp) {
			// 	t.Errorf("basicUsersService.GetAllUser() = %v, want %v", gotAllRecordResp, tt.wantAllRecordResp)
			// }
		})
	}
}

func Test_basicUsersService_UpdateUser(t *testing.T) {
	type args struct {
		ctx       context.Context
		upadteReq models.User
	}
	tests := []struct {
		name string
		b    UsersService
		args args
		//wantUpdateResp *models.User
		wantErr bool
	}{
		{
			"'Test 1 : ",
			NewBasicUsersService(NewMokeUserRepository()),
			args{
				context.Background(),
				models.User{},
			},
			false,
		},
		{
			"'Test 1 : ",
			NewBasicUsersService(NewMokeUserRepository()),
			args{
				context.Background(),
				models.User{FirstName: "Rahul"},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.b.UpdateUser(tt.args.ctx, tt.args.upadteReq)
			if (err != nil) != tt.wantErr {
				t.Errorf("basicUsersService.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(gotUpdateResp, tt.wantUpdateResp) {
			// 	t.Errorf("basicUsersService.UpdateUser() = %v, want %v", gotUpdateResp, tt.wantUpdateResp)
			// }
		})
	}
}

func Test_basicUsersService_DeleteUser(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name string
		b    UsersService
		args args
		//wantDeleteResp *models.DeleteUserResp
		wantErr bool
	}{
		{
			"Test 1",
			NewBasicUsersService(NewMokeUserRepository()),
			args{
				context.Background(),
				"1",
			},
			false,
		},
		{
			"Test 2",
			NewBasicUsersService(NewMokeUserRepository()),
			args{
				context.Background(),
				"2",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.b.DeleteUser(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("basicUsersService.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(gotDeleteResp, tt.wantDeleteResp) {
			// 	t.Errorf("basicUsersService.DeleteUser() = %v, want %v", gotDeleteResp, tt.wantDeleteResp)
			// }
		})
	}
}

func Test_basicUsersService_GetUser(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name string
		b    UsersService
		args args
		//wantCreateResp *models.User
		wantErr bool
	}{
		{
			"Test 1",
			NewBasicUsersService(NewMokeUserRepository()),
			args{
				context.Background(),
				"1",
			},
			false,
		},
		{
			"Test 2",
			NewBasicUsersService(NewMokeUserRepository()),
			args{
				context.Background(),
				"2",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.b.GetUser(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("basicUsersService.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(gotCreateResp, tt.wantCreateResp) {
			// 	t.Errorf("basicUsersService.GetUser() = %v, want %v", gotCreateResp, tt.wantCreateResp)
			// }
		})
	}
}

func TestNewBasicUsersService(t *testing.T) {
	type args struct {
		userRepositoryInterface repository.UserRepositoryInterface
	}
	tests := []struct {
		name string
		args args
		want UsersService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBasicUsersService(tt.args.userRepositoryInterface); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBasicUsersService() = %v, want %v", got, tt.want)
			}
		})
	}
}
