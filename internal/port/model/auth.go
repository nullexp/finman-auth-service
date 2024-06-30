package model

import (
	"context"

	validator "github.com/go-playground/validator/v10"
)

type CreateTokenRequest struct {
	Username string `validate:"required,gte=1"`
	Password string `validate:"required,gte=1"`
}

func (dto CreateTokenRequest) Validate(ctx context.Context) error {
	validate := validator.New()
	return validate.StructCtx(ctx, dto)
}

type CreateTokenResponse struct {
	Token string
}
