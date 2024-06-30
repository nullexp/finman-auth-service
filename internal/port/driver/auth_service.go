package driver

import (
	"context"

	"github.com/nullexp/finman-auth-service/internal/port/model"
)

type AuthService interface {
	CreateToken(context.Context, model.CreateTokenRequest) (*model.CreateTokenResponse, error)
}
