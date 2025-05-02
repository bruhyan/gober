package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(c *gin.Context) {
	var request CreateUserRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	username := request.Username
	email := request.Email
	password := request.Password //hash

	fmt.Println(username, email, password)

	c.Status(http.StatusOK)
}
