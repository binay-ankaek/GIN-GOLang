package contact

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"helloapp/internal/domain/model"
	"helloapp/internal/domain/repository"
)

type ContactService struct {
	contactRepository repository.ContactRepository
	userRepository    repository.UserRepository
}

func NewContactService(contactRepo repository.ContactRepository, userRepo repository.UserRepository) *ContactService {
	return &ContactService{
		contactRepository: contactRepo,
		userRepository:    userRepo,
	}
}

func (s *ContactService) AddContact(userID uint, phone string, name string) error {

	if phone == "" || name == "" {
		return fmt.Errorf("phone or name is empty")
	}
	existingContact, err := s.contactRepository.GetContactByUserIDAndPhoneAndName(userID, phone, name)
	if err == nil && existingContact != nil {
		return errors.New("contact with this phone number or name already exists")
	}
	contact := &model.Contact{
		Phone: phone,
		Name:  name,
	}
	return s.contactRepository.AddContact(userID, contact)
}

func (s *ContactService) GetContacts(userID uint) ([]model.Contact, error) {
	return s.contactRepository.GetContactsByUserID(userID)
}

func (s *ContactService) UpdateContact(userID, contactID uint, phone, name string) error {
	contacts, err := s.contactRepository.GetContactsByUserID(userID)
	if err != nil {
		return err
	}

	// Find the specific contact by contactID
	var contactToUpdate *model.Contact
	for _, contact := range contacts {
		if contact.ID == contactID {
			contactToUpdate = &contact
			break
		}
	}

	if contactToUpdate == nil {
		return fmt.Errorf("contact not found")
	}

	contactToUpdate.Phone = phone
	contactToUpdate.Name = name
	return s.contactRepository.UpdateContact(contactToUpdate)
}

//delete the user

func (s *ContactService) DeleteContact(userID, contactID uint) error {
	contacts, err := s.contactRepository.GetContactsByUserID(userID)
	if err != nil {
		return err
	}

	// Find the specific contact by contactID
	var contactToDelete *model.Contact
	for _, contact := range contacts {
		if contact.ID == contactID {
			contactToDelete = &contact
			break
		}
	}

	if contactToDelete == nil {
		return fmt.Errorf("contact not found")
	}

	return s.contactRepository.DeleteContact(contactToDelete.ID)
}

func (s *ContactService) SearchContact(userID uint, phone string) (*model.Contact, error) {
	return s.contactRepository.GetContactByPhone(phone)
}

func (s *ContactService) SearchContactProfile(userID uint, phone string) (*model.Contact, *model.User, error) {
	contact, err := s.contactRepository.GetContactByUserIDAndPhone(userID, phone)
	if err != nil {
		return nil, nil, err
	}

	user, err := s.userRepository.GetUserByPhone(phone)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, nil, err
	}

	return contact, user, nil
}

func (s *ContactService) GetProfileByPhone(phone string) (*model.User, *model.Contact, error) {
	user, err := s.userRepository.GetUserByPhone(phone)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, nil, err
	}

	contact, err := s.contactRepository.GetContactByPhone(phone)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, nil, err
	}

	return user, contact, nil
}

// Implement other contact-related methods like UpdateContact, DeleteContact, SearchContact, etc.
