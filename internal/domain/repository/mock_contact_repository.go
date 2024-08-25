package repository

import (
	"helloapp/internal/domain/model"
)

type MockContactRepository struct {
	AddContactFunc                        func(userID uint, contact *model.Contact) error
	GetContactsByUserIDFunc               func(userID uint) ([]model.Contact, error)
	UpdateContactFunc                     func(contact *model.Contact) error
	DeleteContactFunc                     func(contactID uint) error
	SearchContactFunc                     func(userID uint, phoneNumber string) ([]model.Contact, error)
	GetContactByPhoneFunc                 func(phone string) (*model.Contact, error)
	GetContactByUserIDAndPhoneFunc        func(userID uint, phone string) (*model.Contact, error)
	GetContactByUserIDAndPhoneAndNameFunc func(userID uint, phone string, name string) (*model.Contact, error)
}

func (m *MockContactRepository) AddContact(userID uint, contact *model.Contact) error {
	return m.AddContactFunc(userID, contact)
}

func (m *MockContactRepository) GetContactsByUserID(userID uint) ([]model.Contact, error) {
	return m.GetContactsByUserIDFunc(userID)
}

func (m *MockContactRepository) UpdateContact(contact *model.Contact) error {
	return m.UpdateContactFunc(contact)
}

func (m *MockContactRepository) DeleteContact(contactID uint) error {
	return m.DeleteContactFunc(contactID)
}

func (m *MockContactRepository) SearchContact(userID uint, phoneNumber string) ([]model.Contact, error) {
	return m.SearchContactFunc(userID, phoneNumber)
}

func (m *MockContactRepository) GetContactByPhone(phone string) (*model.Contact, error) {
	return m.GetContactByPhoneFunc(phone)
}

func (m *MockContactRepository) GetContactByUserIDAndPhone(userID uint, phone string) (*model.Contact, error) {
	return m.GetContactByUserIDAndPhoneFunc(userID, phone)
}

func (m *MockContactRepository) GetContactByUserIDAndPhoneAndName(userID uint, phone string, name string) (*model.Contact, error) {
	return m.GetContactByUserIDAndPhoneAndNameFunc(userID, phone, name)
}
