package repository

import (
	"helloapp/internal/domain/model"
)

type MockUserRepository struct {
	CreateUserFunc          func(user *model.User) error
	GetUserByPhoneFunc      func(phone string) (*model.User, error)
	UpdateUserFunc          func(user *model.User) error
	DeleteUserFunc          func(id uint) error
	DeleteAllUsersFunc      func() error
	GetUserByIDFunc         func(id uint) (*model.User, error)
	GetUserWithContactsFunc func(userID uint) (*model.User, error)
}

func (m *MockUserRepository) CreateUser(user *model.User) error {
	return m.CreateUserFunc(user)
}

func (m *MockUserRepository) GetUserByPhone(phone string) (*model.User, error) {
	return m.GetUserByPhoneFunc(phone)
}

func (m *MockUserRepository) UpdateUser(user *model.User) error {
	return m.UpdateUserFunc(user)
}

func (m *MockUserRepository) DeleteUser(id uint) error {
	return m.DeleteUserFunc(id)
}

func (m *MockUserRepository) DeleteAllUsers() error {
	return m.DeleteAllUsersFunc()
}

func (m *MockUserRepository) GetUserByID(id uint) (*model.User, error) {
	return m.GetUserByIDFunc(id)
}

func (m *MockUserRepository) GetUserWithContacts(userID uint) (*model.User, error) {
	return m.GetUserWithContactsFunc(userID)
}
