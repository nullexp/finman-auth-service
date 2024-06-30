package driven

import (
	"context"

	"github.com/nullexp/finman-auth-service/internal/port/model"
)

type UserService interface {
	GetUser(ctx context.Context, username, password string) (*model.GetUserResponse, error)
}
