package contact

import (
	"bytes"
	"errors"
	userhttp "helloapp/internal/adapter/delivery/http"
	"helloapp/internal/app/contact"
	"helloapp/internal/app/middleware"
	"helloapp/internal/app/user"
	"helloapp/internal/domain/model"
	"helloapp/internal/domain/repository"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAddContact_Success(t *testing.T) {
	mockRepo := &repository.MockContactRepository{
		GetContactByUserIDAndPhoneAndNameFunc: func(userID uint, phone, name string) (*model.Contact, error) {
			return nil, nil
		},
		AddContactFunc: func(userID uint, contact *model.Contact) error {
			return nil
		},
	}

	contactService := contact.NewContactService(mockRepo, nil)
	contactHandler := userhttp.NewContactHandler(contactService)
	//get secret key
	secret_key := os.Getenv("SECRET")

	router := gin.Default()
	router.Use(middleware.AuthMiddleware(secret_key))
	router.POST("/api/contact/add", contactHandler.AddContact)

	token, err := user.GenerateUserToken(1)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	requestBody := `{"phone": "1234567890", "name": "John Doe"}`
	req, _ := http.NewRequest(http.MethodPost, "/api/contact/add", bytes.NewReader([]byte(requestBody)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "contact added successfully")
}

func TestAddContact_UserAlreadyExist(t *testing.T) {
	mockRepo := &repository.MockContactRepository{
		GetContactByUserIDAndPhoneAndNameFunc: func(userID uint, phone, name string) (*model.Contact, error) {
			// Simulate that the contact already exists
			return &model.Contact{Phone: phone, Name: name}, nil
		},
		AddContactFunc: func(userID uint, contact *model.Contact) error {
			return nil
		},
	}

	contactService := contact.NewContactService(mockRepo, nil)
	contactHandler := userhttp.NewContactHandler(contactService)
	//get secret key
	secret_key := os.Getenv("SECRET")

	router := gin.Default()
	router.Use(middleware.AuthMiddleware(secret_key))
	router.POST("/api/contact/add", contactHandler.AddContact)

	token, err := user.GenerateUserToken(1)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	requestBody := `{"phone": "1234567890", "name": "John Doe"}`
	req, _ := http.NewRequest(http.MethodPost, "/api/contact/add", bytes.NewReader([]byte(requestBody)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Contains(t, resp.Body.String(), "contact with this phone number or name already exists")
}

func TestAddContact_Unauthorized(t *testing.T) {
	mockRepo := &repository.MockContactRepository{
		GetContactByUserIDAndPhoneAndNameFunc: func(userID uint, phone, name string) (*model.Contact, error) {
			// Simulate that the contact already exists
			return nil, nil
		},
		AddContactFunc: func(userID uint, contact *model.Contact) error {
			return nil
		},
	}

	contactService := contact.NewContactService(mockRepo, nil)
	contactHandler := userhttp.NewContactHandler(contactService)
	//get secret key
	secret_key := os.Getenv("SECRET")

	router := gin.Default()
	router.Use(middleware.AuthMiddleware(secret_key))
	router.POST("/api/contact/add", contactHandler.AddContact)

	token, err := user.GenerateUserToken(1)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	requestBody := `{"phone": "1234567890", "name": "John Doe"}`
	req, _ := http.NewRequest(http.MethodPost, "/api/contact/add", bytes.NewReader([]byte(requestBody)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token+"invalid")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnauthorized, resp.Code)
	assert.Contains(t, resp.Body.String(), "invalid token")
}

//get contact test case

func TestGetContact_Success(t *testing.T) {
	mockRepo := &repository.MockContactRepository{
		GetContactsByUserIDFunc: func(userID uint) ([]model.Contact, error) {
			return []model.Contact{{Phone: "1234567890", Name: "baby boy"}}, nil
		},
	}
	//assing repository for mock
	contactService := contact.NewContactService(mockRepo, nil)
	contactHandler := userhttp.NewContactHandler(contactService)
	secret := os.Getenv("SECRET")
	router := gin.Default()
	router.Use(middleware.AuthMiddleware(secret))
	router.GET("/api/contact/get", contactHandler.GetContacts)

	//generate token
	token, err := user.GenerateUserToken(1)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}
	//set header
	req, _ := http.NewRequest(http.MethodGet, "/api/contact/get", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	//set up response
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	// assertain and contains
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "1234567890")

}

func TestGetContact_Unauthorized(t *testing.T) {
	mockRepo := &repository.MockContactRepository{
		GetContactsByUserIDFunc: func(userID uint) ([]model.Contact, error) {
			return nil, nil
		},
	}
	//assing repository for mock
	contactService := contact.NewContactService(mockRepo, nil)
	contactHandler := userhttp.NewContactHandler(contactService)
	secret := os.Getenv("SECRET")
	router := gin.Default()
	router.Use(middleware.AuthMiddleware(secret))
	router.GET("/api/contact/get", contactHandler.GetContacts)

	//generate token
	token, err := user.GenerateUserToken(1)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}
	//set header
	req, _ := http.NewRequest(http.MethodGet, "/api/contact/get", nil)
	req.Header.Set("Authorization", "Bearer "+token+"token")
	//set up response
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	// assertain and contains
	assert.Equal(t, http.StatusUnauthorized, resp.Code)
	assert.Contains(t, resp.Body.String(), "invalid token")

}

// update contact test
func TestUpdateContact_Success(t *testing.T) {
	mockrepo := &repository.MockContactRepository{
		GetContactsByUserIDFunc: func(userID uint) ([]model.Contact, error) {
			return []model.Contact{
				{
					Model: gorm.Model{
						ID: 1, // Setting the ID within the gorm.Model struct
					},
					UserID: 1,
					Phone:  "9876543210",
					Name:   "baby boy",
				},
			}, nil
		},
		UpdateContactFunc: func(contact *model.Contact) error {
			// Assert that the contact is updated with the correct data
			assert.Equal(t, uint(1), contact.ID)
			assert.Equal(t, "9876543210", contact.Phone)
			assert.Equal(t, "baby boy put", contact.Name)
			return nil
		},
	}

	//initialized the service
	contactService := contact.NewContactService(mockrepo, nil)
	contactHandler := userhttp.NewContactHandler(contactService)

	//router
	router := gin.Default()
	router.Use(middleware.AuthMiddleware(os.Getenv("SECRET")))
	router.PUT("/api/contact/:id", contactHandler.UpdateContact)

	//token
	token, err := user.GenerateUserToken(1)

	if err != nil {
		t.Fatalf("failed to generate token: %v", err)

	}

	//req
	reqbody := `{"Phone":"9876543210","Name":"baby boy put"}`
	req, _ := http.NewRequest(http.MethodPut, "/api/contact/1", bytes.NewReader([]byte(reqbody)))
	req.Header.Set("Authorization", "Bearer "+token)
	//response
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	//assert the response
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "contact updated successfully")

}

func TestUpdateContact_ContactNotFound(t *testing.T) {
	mockrepo := &repository.MockContactRepository{
		GetContactsByUserIDFunc: func(userID uint) ([]model.Contact, error) {
			return []model.Contact{}, nil
		},
		UpdateContactFunc: func(contact *model.Contact) error {
			// Assert that the contact is updated with the correct data

			return nil
		},
	}

	//initialized the service
	contactService := contact.NewContactService(mockrepo, nil)
	contactHandler := userhttp.NewContactHandler(contactService)

	//router
	router := gin.Default()
	router.Use(middleware.AuthMiddleware(os.Getenv("SECRET")))
	router.PUT("/api/contact/:id", contactHandler.UpdateContact)

	//token
	token, err := user.GenerateUserToken(1)

	if err != nil {
		t.Fatalf("failed to generate token: %v", err)

	}

	//req
	reqbody := `{"Phone":"9876543210","Name":"baby boy put"}`
	req, _ := http.NewRequest(http.MethodPut, "/api/contact/1", bytes.NewReader([]byte(reqbody)))
	req.Header.Set("Authorization", "Bearer "+token)
	//response
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	//assert the response
	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	assert.Contains(t, resp.Body.String(), "contact not found")

}

func TestUpdateContact_InvalidID(t *testing.T) {
	mockrepo := &repository.MockContactRepository{}

	//initialized the service
	contactService := contact.NewContactService(mockrepo, nil)
	contactHandler := userhttp.NewContactHandler(contactService)

	//router
	router := gin.Default()
	router.Use(middleware.AuthMiddleware(os.Getenv("SECRET")))
	router.PUT("/api/contact/:id", contactHandler.UpdateContact)

	//token
	token, err := user.GenerateUserToken(1)

	if err != nil {
		t.Fatalf("failed to generate token: %v", err)

	}

	//req
	reqbody := `{"Phone":"9876543210","Name":"baby boy put"}`
	req, _ := http.NewRequest(http.MethodPut, "/api/contact/abc", bytes.NewReader([]byte(reqbody)))
	req.Header.Set("Authorization", "Bearer "+token)
	//response
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	//assert the response
	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Contains(t, resp.Body.String(), "invalid contact id")

}

func TestUpdateContact_Unauthorized(t *testing.T) {
	mockrepo := &repository.MockContactRepository{}

	//initialized the service
	contactService := contact.NewContactService(mockrepo, nil)
	contactHandler := userhttp.NewContactHandler(contactService)

	//router
	router := gin.Default()
	router.Use(middleware.AuthMiddleware(os.Getenv("SECRET")))
	router.PUT("/api/contact/:id", contactHandler.UpdateContact)

	//token
	token, err := user.GenerateUserToken(1)

	if err != nil {
		t.Fatalf("failed to generate token: %v", err)

	}

	//req
	reqbody := `{"Phone":"9876543210","Name":"baby boy put"}`
	req, _ := http.NewRequest(http.MethodPut, "/api/contact/abc", bytes.NewReader([]byte(reqbody)))
	req.Header.Set("Authorization", "Bearer "+token+"token")
	//response
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	//assert the response
	assert.Equal(t, http.StatusUnauthorized, resp.Code)
	assert.Contains(t, resp.Body.String(), "invalid token")

}

//## Delete the contact

// func TestDeleteContact_Success(t *testing.T) {
// 	mockrepo := &repository.MockContactRepository{
// 		//Mock the DeleteContact method to return nil error
// 		GetContactsByUserIDFunc: func(userID uint) ([]model.Contact, error) {
// 			return []model.Contact{{Phone: "1234567890", Name: "baby boy"}}, nil
// 		},
// 		DeleteContactFunc: func(contactID uint) error {
// 			if contactID == 1 {
// 				return nil
// 			}
// 			return nil
// 		},
// 	}

// 	//initialized the service
// 	contactService := contact.NewContactService(mockrepo, nil)
// 	contactHandler := userhttp.NewContactHandler(contactService)

// 	//router
// 	router := gin.Default()
// 	router.Use(middleware.AuthMiddleware(os.Getenv("SECRET")))
// 	router.DELETE("/api/contact/:id", contactHandler.DeleteContact)

// 	//token
// 	token, err := user.GenerateUserToken(1)

// 	if err != nil {
// 		t.Fatalf("failed to generate token: %v", err)

// 	}

// 	//req

// 	req, _ := http.NewRequest(http.MethodDelete, "/api/contact/1", nil)
// 	req.Header.Set("Authorization", "Bearer "+token)
// 	//response
// 	resp := httptest.NewRecorder()
// 	router.ServeHTTP(resp, req)
// 	//assert the response
// 	assert.Equal(t, http.StatusOK, resp.Code)
// 	assert.Contains(t, resp.Body.String(), "contact not found")

// }

func TestDeleteContact_Success(t *testing.T) {
	mockRepo := &repository.MockContactRepository{
		// Mock the GetContactsByUserID method
		GetContactsByUserIDFunc: func(userID uint) ([]model.Contact, error) {
			return []model.Contact{{
				Model: gorm.Model{
					ID: 1, // Setting the ID within the gorm.Model struct
				},
				UserID: 1,
				Phone:  "1234567890",
				Name:   "baby boy",
			},
			}, nil
		},
		// Mock the DeleteContact method to return nil error for contactID 1
		DeleteContactFunc: func(contactID uint) error {
			if contactID == 1 {
				return nil // Successful deletion
			}
			return errors.New("No such contact") // Simulate not found error
		},
	}

	// Initialize the service
	contactService := contact.NewContactService(mockRepo, nil)
	contactHandler := userhttp.NewContactHandler(contactService)

	// Set up the router
	router := gin.Default()
	router.Use(middleware.AuthMiddleware(os.Getenv("SECRET")))
	router.DELETE("/api/contact/:id", contactHandler.DeleteContact)

	// Generate a valid token
	token, err := user.GenerateUserToken(1)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	// Create a DELETE request without a body
	req, _ := http.NewRequest(http.MethodDelete, "/api/contact/1", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	// Record the response
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "contact deleted successfully")
}

func TestDeleteContact_Unauthorized(t *testing.T) {
	mockRepo := &repository.MockContactRepository{
		// Mock the GetContactsByUserID method
		GetContactsByUserIDFunc: func(userID uint) ([]model.Contact, error) {
			return []model.Contact{}, nil
		},
		// Mock the DeleteContact method to return nil error for contactID 1
		DeleteContactFunc: func(contactID uint) error {
			return errors.New("No such contact") // Simulate not found error
		},
	}

	// Initialize the service
	contactService := contact.NewContactService(mockRepo, nil)
	contactHandler := userhttp.NewContactHandler(contactService)

	// Set up the router
	router := gin.Default()
	router.Use(middleware.AuthMiddleware(os.Getenv("SECRET")))
	router.DELETE("/api/contact/:id", contactHandler.DeleteContact)

	// Generate a valid token
	token, err := user.GenerateUserToken(1)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	// Create a DELETE request without a body
	req, _ := http.NewRequest(http.MethodDelete, "/api/contact/1", nil)
	req.Header.Set("Authorization", "Bearer "+token+"token")

	// Record the response
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert the response
	assert.Equal(t, http.StatusUnauthorized, resp.Code)
	assert.Contains(t, resp.Body.String(), "invalid token")
}

func TestDeleteContact_InvalidID(t *testing.T) {
	mockRepo := &repository.MockContactRepository{
		// Mock the GetContactsByUserID method
		GetContactsByUserIDFunc: func(userID uint) ([]model.Contact, error) {
			return []model.Contact{}, nil
		},
		// Mock the DeleteContact method to return nil error for contactID 1
		DeleteContactFunc: func(contactID uint) error {
			return errors.New("No such contact") // Simulate not found error
		},
	}

	// Initialize the service
	contactService := contact.NewContactService(mockRepo, nil)
	contactHandler := userhttp.NewContactHandler(contactService)

	// Set up the router
	router := gin.Default()
	router.Use(middleware.AuthMiddleware(os.Getenv("SECRET")))
	router.DELETE("/api/contact/:id", contactHandler.DeleteContact)

	// Generate a valid token
	token, err := user.GenerateUserToken(1)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	// Create a DELETE request without a body
	req, _ := http.NewRequest(http.MethodDelete, "/api/contact/abc", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	// Record the response
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert the response
	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Contains(t, resp.Body.String(), "invalid contact id")
}

func TestDeleteContact_NotFound(t *testing.T) {
	mockRepo := &repository.MockContactRepository{
		// Mock the GetContactsByUserID method
		GetContactsByUserIDFunc: func(userID uint) ([]model.Contact, error) {
			return []model.Contact{}, nil
		},
		// Mock the DeleteContact method to return nil error for contactID 1
		DeleteContactFunc: func(contactID uint) error {
			if contactID == 1 {
				return errors.New("No such contact")
			}
			return nil // Simulate not found error
		},
	}

	// Initialize the service
	contactService := contact.NewContactService(mockRepo, nil)
	contactHandler := userhttp.NewContactHandler(contactService)

	// Set up the router
	router := gin.Default()
	router.Use(middleware.AuthMiddleware(os.Getenv("SECRET")))
	router.DELETE("/api/contact/:id", contactHandler.DeleteContact)

	// Generate a valid token
	token, err := user.GenerateUserToken(1)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	// Create a DELETE request without a body
	req, _ := http.NewRequest(http.MethodDelete, "/api/contact/1", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	// Record the response
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert the response
	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	assert.Contains(t, resp.Body.String(), "contact not found")
}

// search contact
func TestSearchContact_Success(t *testing.T) {
	mockRepo := &repository.MockContactRepository{
		SearchContactFunc: func(userID uint, phoneNumber string) ([]model.Contact, error) {
			if phoneNumber == "1234567890" {
				return []model.Contact{{Phone: "1234567890", Name: "baby boy"}}, nil
			}
			return nil, nil
		},
		// Ensure that other mock methods are defined if used in the service
	}

	contactService := contact.NewContactService(mockRepo, nil)
	contactHandler := userhttp.NewContactHandler(contactService)

	router := gin.Default()
	router.Use(middleware.AuthMiddleware(os.Getenv("SECRET")))
	router.GET("/api/contact/search", contactHandler.SearchContact)

	token, err := user.GenerateUserToken(1)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	req, _ := http.NewRequest(http.MethodGet, "/api/contact/search?phone=1234567890", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Check status code
	assert.Equal(t, http.StatusOK, resp.Code)

	// Check response body
	expectedBody := `{"name":"baby boy","phone":"1234567890"}`
	assert.JSONEq(t, expectedBody, resp.Body.String())
}
