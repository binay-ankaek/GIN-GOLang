package http

import (
	"helloapp/internal/app/contact"
	// "helloapp/internal/app/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

type ContactHandler struct {
	contactService *contact.ContactService
}

func NewContactHandler(contactService *contact.ContactService) *ContactHandler {
	return &ContactHandler{
		contactService: contactService,
	}
}

// AddContact godoc
// @Summary Add a new contact
// @Description Add a new contact for the authenticated user
// @Tags Contacts
// @Accept  json
// @Produce  json
// @Security BearerToken
// @Param body body AddContactRequest true "Contact details"
// @Success 200 {object} AddContactResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/contact/add [post]
func (h *ContactHandler) AddContact(c *gin.Context) {
	var req AddContactRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "unauthorized"})
		return
	}

	if err := h.contactService.AddContact(userID.(uint), req.Phone, req.Name); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, AddContactResponse{Message: "contact added successfully"})
}

// func (h *ContactHandler) AddContact(c *gin.Context) {
// 	var req struct {
// 		Phone string `json:"phone"`
// 		Name  string `json:"name"`
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

// 	if err := h.contactService.AddContact(userID.(uint), req.Phone, req.Name); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "contact added successfully"})
// }

// GetContacts godoc
// @Summary Get all contacts
// @Description Retrieve all contacts for the authenticated user
// @Tags Contacts
// @Security BearerToken
// @Produce json
// @Success 200 {array} ContactResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/contact/ [get]
func (h *ContactHandler) GetContacts(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "unauthorized"})
		return
	}

	contacts, err := h.contactService.GetContacts(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, contacts)
}

// func (h *ContactHandler) GetContacts(c *gin.Context) {
// 	userID, exists := c.Get("userID")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
// 		return
// 	}

// 	contacts, err := h.contactService.GetContacts(userID.(uint))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, contacts)
// }

////

// UpdateContact updates the details of an existing contact
// @Summary Update a contact
// @Description Update the details of an existing contact
// @Tags Contacts
// @Accept  json
// @Produce  json
// @Param id path int true "Contact ID"
// @Param body body UpdateContactRequest true "Contact update details"
// @Security BearerToken
// @Success 200 {object} UpdateContactResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/contact/{id} [put]
func (h *ContactHandler) UpdateContact(c *gin.Context) {
	var req UpdateContactRequest

	// Bind the incoming JSON request to the request struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	// Get userID from context (assumed to be set by authentication middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "unauthorized"})
		return
	}

	// Parse the contact ID from the URL path parameter
	contactID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid contact id"})
		return
	}

	// Call the service to update the contact details
	if err := h.contactService.UpdateContact(userID.(uint), uint(contactID), req.Phone, req.Name); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	// Return success response
	response := UpdateContactResponse{
		Message: "contact updated successfully",
	}

	c.JSON(http.StatusOK, response)
}

// func (h *ContactHandler) UpdateContact(c *gin.Context) {
// 	var req struct {
// 		Phone string `json:"phone"`
// 		Name  string `json:"name"`
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

// 	contactID, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid contact id"})
// 		return
// 	}

// 	if err := h.contactService.UpdateContact(userID.(uint), uint(contactID), req.Phone, req.Name); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "contact updated successfully"})
// }

// DeleteContact deletes a contact by ID
// @Summary Delete a contact
// @Description Deletes a contact by its ID, if the contact belongs to the authenticated user.
// @Tags Contacts
// @Accept  json
// @Produce  json
// @Param id path int true "Contact ID"
// @Security BearerToken
// @Success 200 {object} DeleteContactResponse "Success response indicating contact deletion"
// @Failure 400 {object} ErrorResponse "Bad request, usually due to invalid contact ID"
// @Failure 401 {object} ErrorResponse "Unauthorized, when the user is not logged in"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/contact/{id} [delete]
func (h *ContactHandler) DeleteContact(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "unauthorized"})
		return
	}

	contactID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid contact id"})
		return
	}

	if err := h.contactService.DeleteContact(userID.(uint), uint(contactID)); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, DeleteContactResponse{Message: "contact deleted successfully"})
}

// func (h *ContactHandler) DeleteContact(c *gin.Context) {
// 	userID, exists := c.Get("userID")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
// 		return
// 	}

// 	contactID, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid contact id"})
// 		return
// 	}

// 	if err := h.contactService.DeleteContact(userID.(uint), uint(contactID)); err != nil {

// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "contact deleted successfully"})
// }

// SearchContact searches for a contact by phone number
// @Summary Search for a contact
// @Description Search for a contact by phone number. Returns the user's profile if the phone is registered, or the contact details if it exists in the user's contact list.
// @Tags Contacts
// @Accept  json
// @Produce  json
// @Param phone query string true "Phone number to search for"
// @Security BearerToken
// @Success 200 {object} UserResponse "Success response with user profile"
// @Success 200 {object} ContactResponse "Success response with contact details"
// @Failure 400 {object} ErrorResponse "Bad request, usually due to missing query parameters"
// @Failure 401 {object} ErrorResponse "Unauthorized, when the user is not logged in"
// @Failure 404 {object} ErrorResponse "Not found, when the phone number is not registered or in contacts"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/contact/search [get]
func (h *ContactHandler) SearchContact(c *gin.Context) {
	// Bind the incoming query parameters to the request struct
	var req SearchContactRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "phone query parameter is required"})
		return
	}

	// Check for user authentication
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "unauthorized"})
		return
	}

	// Fetch user and contact details
	user, contact, err := h.contactService.GetProfileByPhone(req.Phone)
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	// Log the search for debugging or audit purposes
	log.Printf("User ID %d is searching for phone %s", userID, req.Phone)

	// Return the appropriate response
	if user != nil {
		c.JSON(http.StatusOK, UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Phone:     user.Phone,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		})
		return
	}

	if contact != nil {
		c.JSON(http.StatusOK, ContactResponse{
			Name:  contact.Name,
			Phone: contact.Phone,
		})
		return
	}

	c.JSON(http.StatusNotFound, ErrorResponse{Error: "user not registered yet"})
}

// func (h *ContactHandler) SearchContact(c *gin.Context) {

// 	userID, exists := c.Get("userID")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
// 		return
// 	}
// 	phone := c.Query("phone")
// 	if phone == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "phone query parameter is required"})
// 		return
// 	}

// 	user, contact, err := h.contactService.GetProfileByPhone(phone)
// 	if err != nil && err != gorm.ErrRecordNotFound {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Log the userID for debugging or audit purposes
// 	log.Printf("User ID %d is searching for phone %s", userID, phone)

// 	if user != nil {
// 		c.JSON(http.StatusOK, gin.H{"user": user})
// 		return
// 	}

// 	if contact != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"name":  contact.Name,
// 			"phone": contact.Phone,
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusNotFound, gin.H{"error": "user not registered yet"})
// }

func (h *ContactHandler) GetProfileByPhone(c *gin.Context) {
	phone := c.Query("phone")
	contact, user, err := h.contactService.GetProfileByPhone(phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"contact": contact,
		"user":    user,
	})
}

// Implement other contact-related handlers like UpdateContact, DeleteContact, SearchContact, etc.
