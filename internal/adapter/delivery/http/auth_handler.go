package http

import (
	"helloapp/internal/app/user"
	"helloapp/internal/domain/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService *user.UserService
}

func NewAuthHandler(userService *user.UserService) *AuthHandler {
	return &AuthHandler{userService: userService}
}

// Register godoc
// @Summary Register a new account
// @Description Register a new account
// @Tags Users
// @Accept  json
// @Produce  json
// @Param body body RegisterRequest true "User registration details"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	if err := h.userService.RegisterUser(req.Phone, req.Password); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "user registered successfully"})
}

// Login godoc
// @Summary Login to the account
// @Description Authenticate a user and return a JWT token
// @Tags Users
// @Accept  json
// @Produce  json
// @Param body body LoginRequest true "User login details"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	token, err := h.userService.LoginUser(req.Phone, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Invalid phone number or password"})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{
		Message: "login successful",
		Token:   token,
	})
}

// GetProfile godoc
// @Summary Get user profile
// @Description Retrieve the profile of the currently authenticated user
// @Tags Users
// @Accept  json
// @Produce  json
// @Security BearerToken
// @Success 200 {object} GetProfileResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/auth/profile [get]
func (h *AuthHandler) GetProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "userID not found in context"})
		return
	}

	user, err := h.userService.GetUserProfile(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, GetProfileResponse{
		Phone: user.Phone,
		Email: user.Email,
		Image: user.Image,
		Name:  user.Name,
	})
}

// func (h *AuthHandler) UpdateProfile(c *gin.Context) {
// 	var req struct {
// 		Phone    *string `json:"phone"`
// 		Password *string `json:"password"`
// 		Email    *string `json:"email"`
// 		Image    *string `json:"image"`
// 		Name     *string `json:"name"`
// 	}

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	userID, exists := c.Get("userID")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
// 		return
// 	}

// 	updatedUser := &model.User{}
// 	if req.Phone != nil {
// 		updatedUser.Phone = *req.Phone
// 	}
// 	if req.Password != nil {
// 		updatedUser.Password = *req.Password
// 	}
// 	if req.Email != nil {
// 		updatedUser.Email = *req.Email
// 	}
// 	if req.Image != nil {
// 		updatedUser.Image = *req.Image
// 	}
// 	if req.Name != nil {
// 		updatedUser.Name = *req.Name
// 	}

// 	userProfile, err := h.userService.UpdateUserProfile(userID.(uint), updatedUser)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, userProfile)
// }

// UpdateProfile godoc
// @Summary Update user profile
// @Description Update user profile information
// @Tags Users
// @Accept  json
// @Produce  json
// @Param body body UpdateProfileRequest true "User profile update details"
// @Security BearerToken
// @Success 200 {object} UpdateProfileResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/auth/profile [put]
func (h *AuthHandler) UpdateProfile(c *gin.Context) {
	var req UpdateProfileRequest

	// Bind the incoming JSON to the UpdateProfileRequest struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	// Retrieve userID from the context (assumes you have middleware that sets this)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Unauthorized"})
		return
	}

	// Prepare the updated user profile data
	updatedUser := &model.User{}
	if req.Phone != nil {
		updatedUser.Phone = *req.Phone
	}
	if req.Password != nil {
		updatedUser.Password = *req.Password
	}
	if req.Email != nil {
		updatedUser.Email = *req.Email
	}
	if req.Image != nil {
		updatedUser.Image = *req.Image
	}
	if req.Name != nil {
		updatedUser.Name = *req.Name
	}

	// Call the service to update the user's profile
	userProfile, err := h.userService.UpdateUserProfile(userID.(uint), updatedUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	// Return the updated profile data
	c.JSON(http.StatusOK, UpdateProfileResponse{
		Phone: userProfile.Phone,
		Email: userProfile.Email,
		Image: userProfile.Image,
		Name:  userProfile.Name,
	})
}

// DeleteProfile godoc
// @Summary Delete user profile
// @Description Delete the user profile associated with the provided token
// @Tags Users
// @Accept  json
// @Produce  json
// @Security BearerToken
// @Success 200 {object} SuccessResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/auth/profile [delete]
func (h *AuthHandler) DeleteProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "unauthorized"})
		return
	}

	err := h.userService.DeleteUserProfile(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "user profile deleted successfully"})
}

// func (h *AuthHandler) DeleteProfile(c *gin.Context) {
// 	userID, exists := c.Get("userID")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
// 		return
// 	}

// 	err := h.userService.DeleteUserProfile(userID.(uint))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "user profile deleted successfully"})
// }
