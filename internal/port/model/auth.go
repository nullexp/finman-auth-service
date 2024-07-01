package model

import (
	"context"

	validator "github.com/go-playground/validator/v10"
)

type CreateTokenRequest struct {
	Username string `json:"username" validate:"required,gte=1"`
	Password string `json:"password" validate:"required,gte=1"`
}

func (dto CreateTokenRequest) Validate(ctx context.Context) error {
	validate := validator.New()
	return validate.StructCtx(ctx, dto)
}

type CreateTokenResponse struct {
	Token string `json:"token"`
}
