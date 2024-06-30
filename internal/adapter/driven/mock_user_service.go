package driven

import (
	"context"
	"sync"

	"github.com/nullexp/finman-auth-service/internal/port/model"
)

type MockUserService struct {
	mu       sync.Mutex
	response *model.GetUserResponse
	err      error
}

func NewMockUserService() *MockUserService {
	return &MockUserService{}
}

func (m *MockUserService) GetUser(ctx context.Context, username, password string) (*model.GetUserResponse, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.response, m.err
}

func (m *MockUserService) SetGetUserResponse(response *model.GetUserResponse, err error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.response = response
}
