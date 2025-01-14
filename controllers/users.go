package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"myproject/models"
	"myproject/repositories"
	"myproject/request"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userRepo *repositories.UserRepository
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		userRepo: repositories.NewUserRepository(),
	}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.userRepo.GetAllUsers()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error fetching users",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")

	// Convert string to int
	userId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	user, err := h.userRepo.GetUserByID(userId)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user request.CreateUserDTO

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	err := h.userRepo.CreateUser(&models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	})
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error creating user",
		})
		return
	}

	c.JSON(200, gin.H{
		"users": "users",
	})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	// Convert string to int
	userId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Fetch the existing user
	existingUser, err := h.userRepo.GetUserByID(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var user request.CreateUserDTO

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	existingUser.FirstName = user.FirstName
	existingUser.LastName = user.LastName
	existingUser.Email = user.Email
	existingUser.Password = user.Password

	if err := h.userRepo.UpdateUser(existingUser); err != nil {
		c.JSON(500, gin.H{
			"error": "Error updating user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "user": existingUser})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	// Convert string to int

	userId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	existingUser, err := h.userRepo.GetUserByID(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := h.userRepo.DeleteUserByID(existingUser); err != nil {
		c.JSON(500, gin.H{
			"error": "Error deleting user",
		})
		return
	}

	c.JSON(200, gin.H{
		"users": "users",
	})
}
