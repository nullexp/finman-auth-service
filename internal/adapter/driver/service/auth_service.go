package driver

import (
	"context"

	"github.com/nullexp/finman-auth-service/internal/domain"
	"github.com/nullexp/finman-auth-service/internal/port/driven"
	"github.com/nullexp/finman-auth-service/internal/port/model"
)

type AuthService struct {
	userService  driven.UserService
	tokenService driven.TokenService
}

func NewAuthService(userService driven.UserService, tokenService driven.TokenService) *AuthService {
	return &AuthService{userService: userService, tokenService: tokenService}
}

func (as AuthService) CreateToken(ctx context.Context, dto model.CreateTokenRequest) (*model.CreateTokenResponse, error) {
	if err := dto.Validate(ctx); err != nil {
		return nil, err
	}

	user, err := as.userService.GetUser(ctx, dto.Username, dto.Password)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, domain.ErrInvalidAuth
	}

	token, err := as.tokenService.CreateToken(model.Subject{UserId: user.Id, IsAdmin: user.IsAdmin})

	if err != nil {
		return nil, err
	}

	return &model.CreateTokenResponse{
		Token: token,
	}, nil
}
