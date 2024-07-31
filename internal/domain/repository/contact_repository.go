// internal/domain/repository/contact_repository.go

package repository

import (
	"gorm.io/gorm"
	"helloapp/internal/domain/model"
)

type ContactRepository interface {
	AddContact(userID uint, contact *model.Contact) error
	GetContactsByUserID(userID uint) ([]model.Contact, error)
	UpdateContact(contact *model.Contact) error
	DeleteContact(contactID uint) error
	SearchContact(userID uint, phoneNumber string) ([]model.Contact, error)
	GetContactByPhone(phone string) (*model.Contact, error)
	GetContactByUserIDAndPhone(userID uint, phone string) (*model.Contact, error)
	GetContactByUserIDAndPhoneAndName(userID uint, phone string, name string) (*model.Contact, error)
}

type contactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) ContactRepository {
	return &contactRepository{db}
}

func (r *contactRepository) AddContact(userID uint, contact *model.Contact) error {
	contact.UserID = userID
	return r.db.Create(contact).Error
}

func (r *contactRepository) GetContactsByUserID(userID uint) ([]model.Contact, error) {
	var contacts []model.Contact
	err := r.db.Where("user_id = ?", userID).Find(&contacts).Error
	return contacts, err
}

func (r *contactRepository) UpdateContact(contact *model.Contact) error {
	return r.db.Save(contact).Error
}

func (r *contactRepository) DeleteContact(contactID uint) error {
	return r.db.Delete(&model.Contact{}, contactID).Error
}

func (r *contactRepository) SearchContact(userID uint, phoneNumber string) ([]model.Contact, error) {
	var contacts []model.Contact
	if err := r.db.Where("user_id = ? AND phone LIKE ?", userID, "%"+phoneNumber+"%").Find(&contacts).Error; err != nil {
		return nil, err
	}
	return contacts, nil
}

func (r *contactRepository) GetContactByPhone(phone string) (*model.Contact, error) {
	var contact model.Contact
	if err := r.db.Where("phone = ?", phone).First(&contact).Error; err != nil {
		return nil, err
	}
	return &contact, nil
}

func (r *contactRepository) GetContactByUserIDAndPhone(userID uint, phone string) (*model.Contact, error) {
	var contact model.Contact
	if err := r.db.Where("user_id = ? AND phone = ?", userID, phone).First(&contact).Error; err != nil {
		return nil, err
	}
	return &contact, nil
}

func (r *contactRepository) GetContactByUserIDAndPhoneAndName(userID uint, phone string, name string) (*model.Contact, error) {
	var contact model.Contact
	if err := r.db.Where("user_id = ? AND (phone = ? OR name = ?)", userID, phone, name).First(&contact).Error; err != nil {
		return nil, err
	}
	return &contact, nil
}
