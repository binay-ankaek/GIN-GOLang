package user

import (
	"errors"
	"helloapp/internal/domain/model"
	"helloapp/internal/domain/repository"

	"golang.org/x/crypto/bcrypt"
)

// TokenGenerator defines a function type for generating tokens
type TokenGenerator func(userID uint) (string, error)

type UserService struct {
	userRepository repository.UserRepository
	generateToken  TokenGenerator
}

func NewUserService(userRepository repository.UserRepository, generateToken TokenGenerator) *UserService {
	return &UserService{
		userRepository: userRepository,
		generateToken:  generateToken,
	}
}

func (s *UserService) RegisterUser(phone, password string) error {

	//check all field are entered or not
	if phone == "" || password == "" {
		return errors.New("all fields are required")
	}

	// Check if the phone number is already registered
	existingUser, err := s.userRepository.GetUserByPhone(phone)
	if err == nil && existingUser != nil {
		return errors.New("phone number already registered")
	}

	// Hash the password
	// Hash the password using utility function
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	// Create a new user
	newUser := &model.User{
		Phone:    phone,
		Password: string(hashedPassword),
	}

	// return s.userRepository.CreateUser(newUser)
	err = s.userRepository.CreateUser(newUser)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) LoginUser(phone, password string) (string, error) {
	// Check if the phone number is already registered
	existingUser, err := s.userRepository.GetUserByPhone(phone)
	if err != nil || existingUser == nil {
		return "", errors.New("phone number not registered")
	}

	// Compare hashed password with the provided password
	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(password)); err != nil {
		return "", errors.New("password incorrect")
	}

	// Generate JWT token
	token, err := generateToken(existingUser.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserService) GetUserProfile(id uint) (*model.User, error) {
	user, err := s.userRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//user update

func (s *UserService) UpdateUserProfile(id uint, updatedUser *model.User) (*model.User, error) {
	user, err := s.userRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	if updatedUser.Phone != "" {
		user.Phone = updatedUser.Phone
	}

	if updatedUser.Password != "" {
		hashedPassword, err := HashPassword(updatedUser.Password)
		if err != nil {
			return nil, err
		}
		user.Password = hashedPassword
	}
	if updatedUser.Email != "" {
		user.Email = updatedUser.Email
	}
	if updatedUser.Image != "" {
		user.Image = updatedUser.Image
	}

	if updatedUser.Name != "" {
		user.Name = updatedUser.Name
	}

	err = s.userRepository.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUserWithContacts(userID uint) (*model.User, error) {
	return s.userRepository.GetUserWithContacts(userID)
}

func (s *UserService) DeleteUserProfile(userID uint) error {
	// Check if the user exists
	_, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		return err
	}

	// Delete user profile
	return s.userRepository.DeleteUser(userID)
}

func (s *UserService) DeleteAllUsers() error {
	// Implement logic to delete all users
	// For example, using repository method to delete all users
	return s.userRepository.DeleteAllUsers()
}

// Implement other user-related methods like UpdateProfile, DeleteUser, etc.
