package driven

import (
	"context"

	userv1 "github.com/nullexp/finman-auth-service/internal/adapter/driver/grpc/proto/user/v1"
	"github.com/nullexp/finman-auth-service/internal/port/model"
	"google.golang.org/grpc"
)

type UserService struct {
	client userv1.UserServiceClient
}

func NewUserService(conn *grpc.ClientConn) *UserService {
	return &UserService{
		client: userv1.NewUserServiceClient(conn),
	}
}

func (us *UserService) GetUser(ctx context.Context, username, password string) (*model.GetUserResponse, error) {
	req := &userv1.GetUserByUsernameAndPasswordRequest{
		Username: username,
		Password: password,
	}

	resp, err := us.client.GetUserByUsernameAndPassword(ctx, req)
	if err != nil {
		return nil, err
	}

	return &model.GetUserResponse{
		Id:      resp.User.Id,
		IsAdmin: resp.User.IsAdmin,
	}, nil
}
