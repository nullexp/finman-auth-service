package grpc

import (
	"context"
	"log"

	authv1 "github.com/nullexp/finman-auth-service/internal/adapter/driver/grpc/proto/auth/v1"
	"github.com/nullexp/finman-auth-service/internal/port/driver"
	"github.com/nullexp/finman-auth-service/internal/port/model"
)

type AuthService struct {
	authv1.UnimplementedAuthServiceServer
	service driver.AuthService
}

func NewAuthService(as driver.AuthService) *AuthService {
	return &AuthService{service: as}
}

func (as AuthService) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	log.Println("CALL: Login")
	result, err := as.service.CreateToken(ctx, model.CreateTokenRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &authv1.LoginResponse{Token: result.Token}, nil
}
