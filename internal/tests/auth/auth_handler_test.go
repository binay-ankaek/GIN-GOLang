package auth

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"strings" // Import the strings package
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"

	"encoding/json"
	"errors"
	userhttp "helloapp/internal/adapter/delivery/http" //due to name conflict we alias it as a userhttp
	"helloapp/internal/app/user"
	"helloapp/internal/domain/model"
	"helloapp/internal/domain/repository"
)

func mockGenerateToken(userID uint) (string, error) {
	return "mockToken123", nil
}

func TestRegisterUser_Success(t *testing.T) {
	mockRepo := &repository.MockUserRepository{}
	mockRepo.CreateUserFunc = func(user *model.User) error {
		return nil
	}
	mockRepo.GetUserByPhoneFunc = func(phone string) (*model.User, error) {
		return nil, nil
	}

	userService := user.NewUserService(mockRepo, nil) // No token generation for Register
	authHandler := userhttp.NewAuthHandler(userService)

	router := gin.Default()
	router.POST("/api/auth/register", authHandler.Register)

	requestBody := `{"phone": "1234567890", "password": "password123"}`
	req, _ := http.NewRequest(http.MethodPost, "/api/auth/register", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "user registered successfully")
}

func TestLoginUser_Success(t *testing.T) {
	// Create a mock repository
	mockRepo := &repository.MockUserRepository{}

	// Define the behavior of GetUserByPhone for this test
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

	// Initialize user service with the mock repository and token generator
	userService := user.NewUserService(mockRepo, mockGenerateToken)
	authHandler := userhttp.NewAuthHandler(userService)

	// Setup Gin router
	router := gin.Default()
	router.POST("/api/auth/login", authHandler.Login)

	// Create a request to the login endpoint with valid credentials
	requestBody := `{"phone": "1234567890", "password": "password123"}`
	req, _ := http.NewRequest(http.MethodPost, "/api/auth/login", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assertions
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), `"message":"login successful"`)
	// assert.Contains(t, resp.Body.String(), `"token":"mockToken123"`)
}

func TestRegisterUser_BindJSONError(t *testing.T) {
	mockRepo := &repository.MockUserRepository{}
	mockRepo.CreateUserFunc = func(user *model.User) error {
		return nil
	}
	mockRepo.GetUserByPhoneFunc = func(phone string) (*model.User, error) {
		return nil, nil
	}

	userService := user.NewUserService(mockRepo, nil)
	authHandler := userhttp.NewAuthHandler(userService)

	router := gin.Default()
	router.POST("/api/auth/register", authHandler.Register)

	requestBody := `{"phone": "1234567890", "password": "password123"` // Malformed JSON
	req, _ := http.NewRequest(http.MethodPost, "/api/auth/register", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Contains(t, resp.Body.String(), "error")
}

func TestGetProfile_Success(t *testing.T) {
	mockRepo := &repository.MockUserRepository{
		GetUserByIDFunc: func(id uint) (*model.User, error) {
			if id == 1 {
				return &model.User{
					Phone: "1234567890",
					Email: "email@gmail.com",
					Name:  "abc com",
				}, nil
			}
			return nil, errors.New("user not found")
		},
	}

	//initialize mock repository
	userService := user.NewUserService(mockRepo, nil)
	authHandler := userhttp.NewAuthHandler(userService)

	//setup gin router
	router := gin.Default()
	router.GET("/api/auth/profile", func(c *gin.Context) {
		c.Set("userID", uint(1))
		authHandler.GetProfile(c)
	})
	//create request to the profile end point
	req, _ := http.NewRequest(http.MethodGet, "/api/auth/profile", nil)
	//record the response
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	//assertion
	assert.Equal(t, http.StatusOK, resp.Code)
	expectedResponse := `{
	    "Phone":"1234567890",
		"Email":"email@gmail.com",
		"Name":"abc com",
		"Image":"",
		"ID":0,
		"Password":"",
		"CreatedAt":"0001-01-01T00:00:00Z",
		"UpdatedAt":"0001-01-01T00:00:00Z",
		"DeletedAt":null,
		"Contacts":null
	}`
	assert.JSONEq(t, expectedResponse, resp.Body.String())
}

func TestGetProfile_InternalServerError(t *testing.T) {
	mockRepo := &repository.MockUserRepository{
		GetUserByIDFunc: func(id uint) (*model.User, error) {
			return nil, errors.New("unexpected error")
		},
	}

	// Initialize the user service with the mock repository
	userService := user.NewUserService(mockRepo, nil)
	authHandler := userhttp.NewAuthHandler(userService)

	// Setup Gin router
	router := gin.Default()
	router.GET("/api/auth/profile", func(c *gin.Context) {
		c.Set("userID", uint(1))
		authHandler.GetProfile(c)
	})

	// Create a request to the profile endpoint
	req, _ := http.NewRequest(http.MethodGet, "/api/auth/profile", nil)

	// Record the response
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assertions
	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	expectedResponse := `{"error":"unexpected error"}`
	assert.JSONEq(t, expectedResponse, resp.Body.String())
}

func TestUpdateProfile_Success(t *testing.T) {
	gin.SetMode(gin.TestMode) // Set Gin to TestMode

	// Mock repository with necessary methods
	mockRepo := &repository.MockUserRepository{
		GetUserByIDFunc: func(id uint) (*model.User, error) {
			if id == 1 {
				return &model.User{
					Phone:    "1234567890",
					Email:    "oldemail@example.com",
					Name:     "Old Name",
					Password: "oldpassword",
					Image:    "oldimage.png",
				}, nil
			}
			return nil, errors.New("user not found")
		},
		UpdateUserFunc: func(user *model.User) error {
			// Simulate successful update
			return nil
		},
	}

	// Initialize mock repository and service
	userService := user.NewUserService(mockRepo, nil)
	authHandler := userhttp.NewAuthHandler(userService)

	// Setup gin router
	router := gin.Default()
	router.PUT("/api/auth/profile", func(c *gin.Context) {
		c.Set("userID", uint(1))
		authHandler.UpdateProfile(c)
	})

	// Create request to update the profile endpoint
	reqBody := `{
        "name": "New Name",
        "email": "newemail@example.com",
        "password": "",
        "phone": "0987654321",
        "image": "newimage.png"
    }`
	req := httptest.NewRequest(http.MethodPut, "/api/auth/profile", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(context.WithValue(req.Context(), "userID", "1"))

	// Record the response
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assertions
	assert.Equal(t, http.StatusOK, resp.Code)

	// Define a custom function to compare the expected and actual responses
	var actualResponse map[string]interface{}
	var expectedResponse map[string]interface{}

	err := json.Unmarshal(resp.Body.Bytes(), &actualResponse)
	assert.NoError(t, err)

	expectedJSON := `{
        "Phone": "0987654321",
        "Email": "newemail@example.com",
        "Name": "New Name",
        "Image": "newimage.png",
        "ID": 0,
        "Password": "",
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "Contacts": null
    }`
	err = json.Unmarshal([]byte(expectedJSON), &expectedResponse)
	assert.NoError(t, err)

	// Remove the Password field from both expected and actual responses for comparison
	delete(actualResponse, "Password")
	delete(expectedResponse, "Password")

	assert.Equal(t, expectedResponse, actualResponse)
}

func TestDeleteUser_Success(t *testing.T) {
	gin.SetMode(gin.TestMode) // Set Gin to TestMode

	// Mock repository with necessary methods
	mockRepo := &repository.MockUserRepository{
		GetUserByIDFunc: func(id uint) (*model.User, error) {
			if id == 1 {
				return &model.User{
					// No need to explicitly define ID, CreatedAt, UpdatedAt, DeletedAt
					Phone:    "1234567890",
					Email:    "user@example.com",
					Name:     "Test User",
					Password: "password", // Include Password if needed
				}, nil
			}
			return nil, errors.New("user not found")
		},
		DeleteUserFunc: func(id uint) error {
			if id == 1 {
				return nil // Simulate successful deletion
			}
			return errors.New("user not found")
		},
	}

	// Initialize mock repository and service
	userService := user.NewUserService(mockRepo, nil)
	authHandler := userhttp.NewAuthHandler(userService)

	// Setup gin router
	router := gin.Default()
	router.DELETE("/api/auth/user/:id", func(c *gin.Context) {
		c.Set("userID", uint(1))
		authHandler.DeleteProfile(c)
	})

	// Create request to delete the user
	req := httptest.NewRequest(http.MethodDelete, "/api/auth/user/1", nil)
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assertions
	assert.Equal(t, http.StatusOK, resp.Code)

	// Expected response body
	expectedResponse := map[string]string{
		"message": "user profile deleted successfully",
	}
	var actualResponse map[string]string
	err := json.Unmarshal(resp.Body.Bytes(), &actualResponse)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	assert.Equal(t, expectedResponse, actualResponse)
}
