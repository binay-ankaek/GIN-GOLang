// package auth

// import (
// 	"bytes"
// 	"encoding/json"
// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	"helloapp/internal/adapter/delivery/http"
// 	"helloapp/internal/domain/model"
// 	"net/http/httptest"
// 	"testing"
// )

// type MockUserService struct {
// 	mock.Mock
// }

// func (m *MockUserService) RegisterUser(phone, password string) error {
// 	args := m.Called(phone, password)
// 	return args.Error(0)
// }

// func (m *MockUserService) LoginUser(phone, password string) (string, error) {
// 	args := m.Called(phone, password)
// 	return args.String(0), args.Error(1)
// }

// func (m *MockUserService) GetUserProfile(UserID uint) (*model.User, error) {
// 	args := m.Called(UserID)
// 	return args.Get(0).(*model.User), args.Error(1)
// }

// func (m *MockUserService) UpdateUserProfile(userID uint, updatedUser *model.User) (*model.User, error) {
// 	args := m.Called(userID, updatedUser)
// 	return args.Get(0).(*model.User), args.Error(1)
// }

// func (m *MockUserService) DeleteUserProfile(userID uint) error {
// 	args := m.Called(userID)
// 	return args.Error(0)
// }

// func TestRegister(t *testing.T) {
// 	mockUserService := new(MockUserService)
// 	handler := http.NewAuthHandler(mockUserService)

// 	router := gin.Default()
// 	router.POST("/register", handler.Register)

// 	mockUserService.On("RegisterUser", "1234567890", "password").Return(nil)

// 	body := map[string]string{
// 		"phone":    "1234567890",
// 		"password": "password",
// 	}
// 	jsonValue, _ := json.Marshal(body)
// 	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonValue))
// 	resp := httptest.NewRecorder()
// 	router.ServeHTTP(resp, req)

// 	assert.Equal(t, http.StatusOK, resp.Code)
// 	mockUserService.AssertExpectations(t)
// }

// internal/adapter/delivery/http/auth_test.go

// internal/tests/auth/auth_handler_test.go

package auth

import (
	// "net/http"
	// "net/http/httptest"
	// "strings"
	// "testing"

	// userhttp "helloapp/internal/adapter/delivery/http"
	// "helloapp/internal/app/user"
	// "helloapp/internal/domain/model"
	// "helloapp/internal/domain/repository"

	// "github.com/gin-gonic/gin"
	// "github.com/stretchr/testify/assert"
	// "golang.org/x/crypto/bcrypt"

	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	userhttp "helloapp/internal/adapter/delivery/http"
	"helloapp/internal/app/user"
	"helloapp/internal/domain/model"
	"helloapp/internal/domain/repository"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

// Mock token generator function
func mockGenerateToken(_ uint) (string, error) {
	return "mockToken123", nil
}

func TestRegisterUser_Success(t *testing.T) {
	// Create a new instance of the mock repository
	mockRepo := &repository.MockUserRepository{}

	// Set up mock behavior
	mockRepo.CreateUserFunc = func(user *model.User) error {
		return nil // Simulate successful user creation
	}
	mockRepo.GetUserByPhoneFunc = func(phone string) (*model.User, error) {
		return nil, nil // Simulate that no user exists with this phone number
	}

	// Create a new instance of UserService with the mock repository
	userService := user.NewUserService(mockRepo)
	authHandler := userhttp.NewAuthHandler(userService)

	// Set up Gin router and test recorder
	router := gin.Default()
	router.POST("/api/auth/register", authHandler.Register)

	// Create a request with the proper route and body
	requestBody := `{"phone": "1234567890", "password": "password123"}`
	req, _ := http.NewRequest(http.MethodPost, "/api/auth/register", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert results
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "user registered successfully")
}

// func TestRegisterUser_UserAlreadyExists(t *testing.T) {
// 	// Create a new instance of the mock repository
// 	mockRepo := &repository.MockUserRepository{}

// 	// Set up mock behavior for GetUserByPhone to simulate user already exists
// 	mockRepo.GetUserByPhoneFunc = func(phone string) (*model.User, error) {
// 		return &model.User{}, nil // Simulate that a user exists with this phone number
// 	}
// 	mockRepo.CreateUserFunc = func(user *model.User) error {
// 		return nil // This should not be called
// 	}

// 	userService := user.NewUserService(mockRepo)
// 	authHandler := userhttp.NewAuthHandler(userService)

// 	// Set up Gin router and test recorder
// 	router := gin.Default()
// 	router.POST("/api/auth/login", authHandler.Register)

// 	// Create a request with the proper route and body
// 	requestBody := `{"phone": "1234567890", "password": "password123"}`
// 	req, _ := http.NewRequest(http.MethodPost, "/api/auth/login", strings.NewReader(requestBody))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp := httptest.NewRecorder()
// 	router.ServeHTTP(resp, req)

// 	// Assert results
// 	assert.Equal(t, http.StatusBadRequest, resp.Code)
// 	assert.Contains(t, resp.Body.String(), "user already exists")
// }

func TestRegisterUser_BindJSONError(t *testing.T) {
	// Create a new instance of the mock repository
	mockRepo := &repository.MockUserRepository{}

	// Set up mock behavior
	mockRepo.CreateUserFunc = func(user *model.User) error {
		return nil // Simulate successful user creation
	}
	mockRepo.GetUserByPhoneFunc = func(phone string) (*model.User, error) {
		return nil, nil // Simulate that no user exists with this phone number
	}

	userService := user.NewUserService(mockRepo)
	authHandler := userhttp.NewAuthHandler(userService)

	// Set up Gin router and test recorder
	router := gin.Default()
	router.POST("/api/auth/register", authHandler.Register)

	// Create a request with malformed JSON
	requestBody := `{"phone": "1234567890", "password": "password123"` // Missing closing brace
	req, _ := http.NewRequest(http.MethodPost, "/api/auth/register", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert results
	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Contains(t, resp.Body.String(), "error")
}

func TestLoginUser_Success(t *testing.T) {

	// Create a new instance of the mock repository
	mockRepo := &repository.MockUserRepository{}

	// Set up mock behavior for a successful login
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	mockRepo.GetUserByPhoneFunc = func(phone string) (*model.User, error) {
		if phone == "1234567890" {
			return &model.User{
				Phone:    phone,
				Password: string(hashedPassword),
			}, nil
		}
		return nil, nil
	}

	// Create a new instance of UserService with the mock repository
	userService := user.NewUserService(mockRepo)
	authHandler := userhttp.NewAuthHandler(userService)

	// Set up Gin router and test recorder
	router := gin.Default()
	router.POST("/api/auth/login", authHandler.Login)

	// Create a request with the proper route and body
	requestBody := `{"phone": "1234567890", "password": "password123"}`
	req, _ := http.NewRequest(http.MethodPost, "/api/auth/login", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert results
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "login successful")
	assert.Contains(t, resp.Body.String(), "mockToken123")
}
