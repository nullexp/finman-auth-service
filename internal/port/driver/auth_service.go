package driver

import (
	"context"

	"github.com/nullexp/finman-auth-service/internal/port/dto"
)

type AuthService interface {
	CreateToken(context.Context, dto.CreateTokenRequest) dto.CreateTokenResponse
}
