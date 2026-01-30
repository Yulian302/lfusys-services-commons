package mocks

import (
	"context"

	authtypes "github.com/Yulian302/lfusys-services-gateway/auth/types"
	"github.com/stretchr/testify/mock"
)

type MockDynamoDbStore struct {
	mock.Mock
}

func (m *MockDynamoDbStore) GetByEmail(
	ctx context.Context,
	email string,
) (*authtypes.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(*authtypes.User), args.Error(1)
}

func (m *MockDynamoDbStore) Create(
	ctx context.Context,
	user authtypes.User,
) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockDynamoDbStore) FindExisting(
	ctx context.Context,
	email string,
) (bool, error) {
	args := m.Called(ctx, email)
	return args.Bool(0), args.Error(1)
}

func (m *MockDynamoDbStore) IsReady(ctx context.Context) error {
	return nil
}

func (m *MockDynamoDbStore) Name() string {
	return "Mock"
}

func (m *MockDynamoDbStore) ResetMock() {
	m.ExpectedCalls = nil
	m.Calls = nil
}
