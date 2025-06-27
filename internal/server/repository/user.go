package repository

import "qfzack/grpc-demo/internal/server/model"

// define repository funcs through interface
type UserRepository interface {
	GetByID(ID string) (*model.User, error)
}

/*
define struct userRepository
implement the pre-defined UserRepository interface
*/
type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

// implement the interface method
func (r *userRepository) GetByID(ID string) (*model.User, error) {
	// TODO get user from datanase
	return &model.User{
		ID:      ID,
		Name:    "Test User",
		Account: "testuser",
	}, nil
}

/*
define a mock repository (case for polymorphism/多态)
implement the pre-defined UserRepository interface
*/
type mockUserRepository struct{}

func NewMockUserRepository() UserRepository {
	return &mockUserRepository{}
}

// implement the interface method
func (r *mockUserRepository) GetByID(ID string) (*model.User, error) {
	return &model.User{
		ID:      ID,
		Name:    "Mock User",
		Account: "mockuser",
	}, nil
}
