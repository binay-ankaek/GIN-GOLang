package http

import (
	"github.com/gin-gonic/gin"
	"helloapp/internal/app/user"
	"net/http"
)

type UserHandler struct {
	userService *user.UserService
}

func NewUserHandler(userService *user.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) DeleteAllUsers(c *gin.Context) {
	err := h.userService.DeleteAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "All users deleted successfully"})
}
