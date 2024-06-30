package driver

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/nullexp/finman-auth-service/internal/adapter/driven"
	"github.com/nullexp/finman-auth-service/internal/port/model"
	"github.com/stretchr/testify/assert"
)

func TestAuthService_CreateToken(t *testing.T) {
	secret := "test-secret"
	expireAfter := time.Hour
	tokenService := driven.NewTokenService(secret, expireAfter)

	tests := []struct {
		name             string
		setupMockUserSvc func(*driven.MockUserService)
		request          model.CreateTokenRequest
		expectedResponse *model.CreateTokenResponse
		expectAnyError   bool
	}{
		{
			name: "successful token creation",
			setupMockUserSvc: func(mock *driven.MockUserService) {
				mock.SetGetUserResponse(&model.GetUserResponse{Id: "123", IsAdmin: false}, nil)
			},
			request: model.CreateTokenRequest{
				Username: "validUser",
				Password: "validPass",
			},
			expectedResponse: &model.CreateTokenResponse{
				Token: "", // Will be checked later since it's dynamically generated
			},
			expectAnyError: false,
		},
		{
			name:             "error on validation failure",
			setupMockUserSvc: func(mock *driven.MockUserService) {},
			request: model.CreateTokenRequest{
				Username: "",
				Password: "validPass",
			},
			expectedResponse: nil,
			expectAnyError:   true,
		},
		{
			name: "error on GetUser failure",
			setupMockUserSvc: func(mock *driven.MockUserService) {
				mock.SetGetUserResponse(nil, errors.New("user service error"))
			},
			request: model.CreateTokenRequest{
				Username: "validUser",
				Password: "validPass",
			},
			expectedResponse: nil,
			expectAnyError:   true,
		},
		{
			name: "error when user not found",
			setupMockUserSvc: func(mock *driven.MockUserService) {
				mock.SetGetUserResponse(nil, nil)
			},
			request: model.CreateTokenRequest{
				Username: "invalidUser",
				Password: "invalidPass",
			},
			expectedResponse: nil,
			expectAnyError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserService := driven.NewMockUserService()

			tt.setupMockUserSvc(mockUserService)

			authService := NewAuthService(mockUserService, tokenService)

			response, err := authService.CreateToken(context.Background(), tt.request)
			if tt.expectedResponse != nil && response != nil {
				assert.NotEmpty(t, response.Token)
				assert.Nil(t, err)
			} else {
				assert.Equal(t, tt.expectedResponse, response)
				if tt.expectAnyError {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}
			}
		})
	}
}
