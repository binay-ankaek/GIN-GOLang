package http

import (
// "net/http"
)

// RegisterRequest represents the request payload for user registration.
type RegisterRequest struct {
	Phone    string `json:"phone" example:"1234567890" format:"string" description:"User's phone number"`
	Password string `json:"password" example:"password123" format:"string" description:"User's password"`
}

// LoginRequest represents the request payload for user login.
type LoginRequest struct {
	Phone    string `json:"phone" example:"1234567890" format:"string" description:"User's phone number"`
	Password string `json:"password" example:"password123" format:"string" description:"User's password"`
}

// LoginResponse represents the response payload for user login.
type LoginResponse struct {
	Message string `json:"message" example:"login successful" description:"Success message"`
	Token   string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." description:"JWT token"`
}

// GetProfileResponse represents the response payload for user profile.
type GetProfileResponse struct {
	Phone string `json:"phone" example:"1234567890" format:"string" description:"User's phone number"`
	Email string `json:"email" example:"user@example.com" format:"string" description:"User's email address"`
	Image string `json:"image" example:"http://example.com/image.jpg" format:"string" description:"User's profile image URL"`
	Name  string `json:"name" example:"John Doe" format:"string" description:"User's name"`
}

// UpdateProfileRequest represents the request payload for updating the user profile.
type UpdateProfileRequest struct {
	Phone    *string `json:"phone" example:"1234567890" format:"string" description:"User's phone number"`
	Password *string `json:"password" example:"newpassword123" format:"string" description:"User's new password"`
	Email    *string `json:"email" example:"newuser@example.com" format:"string" description:"User's new email address"`
	Image    *string `json:"image" example:"http://example.com/newimage.jpg" format:"string" description:"User's new profile image URL"`
	Name     *string `json:"name" example:"Jane Doe" format:"string" description:"User's new name"`
}

// UpdateProfileResponse represents the response payload for updating the user profile.
type UpdateProfileResponse struct {
	Phone string `json:"phone" example:"1234567890" format:"string" description:"User's updated phone number"`
	Email string `json:"email" example:"newuser@example.com" format:"string" description:"User's updated email address"`
	Image string `json:"image" example:"http://example.com/newimage.jpg" format:"string" description:"User's updated profile image URL"`
	Name  string `json:"name" example:"Jane Doe" format:"string" description:"User's updated name"`
}

// SuccessResponse represents a successful response message.
type SuccessResponse struct {
	Message string `json:"message" example:"operation successful" description:"Success message"`
}

// ErrorResponse represents an error response message.
type ErrorResponse struct {
	Error string `json:"error" example:"Invalid input" description:"Error message"`
}

// AddContactRequest represents the request payload for adding a contact.
type AddContactRequest struct {
	Phone string `json:"phone" example:"1234567890" description:"Contact's phone number"`
	Name  string `json:"name" example:"John Doe" description:"Contact's name"`
}

// AddContactResponse represents the response message for adding a contact.
type AddContactResponse struct {
	Message string `json:"message" example:"contact added successfully" description:"Success message"`
}

// ContactResponse represents the contact details in the response.
type ContactResponse struct {
	ID    uint   `json:"id" example:"1" description:"Contact ID"`
	Phone string `json:"phone" example:"1234567890" description:"Contact's phone number"`
	Name  string `json:"name" example:"John Doe" description:"Contact's name"`
}
