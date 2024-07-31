// internal/domain/repository/user_repository.go

package repository

import (
	"gorm.io/gorm"
	"helloapp/internal/domain/model"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserByPhone(phone string) (*model.User, error)
	GetUserByID(id uint) (*model.User, error) // add this method
	GetUserWithContacts(userID uint) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error
	DeleteAllUsers() error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetUserByPhone(phone string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("phone = ?", phone).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) DeleteUser(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}

func (r *userRepository) GetUserByID(id uint) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) DeleteAllUsers() error {
	// Delete all users from the 'users' table
	if err := r.db.Delete(&model.User{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetUserWithContacts(userID uint) (*model.User, error) {
	var user model.User
	if err := r.db.Preload("Contacts").First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
