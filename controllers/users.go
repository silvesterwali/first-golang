package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"myproject/models"
	"myproject/repositories"
	"myproject/request"
	"myproject/utils"

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
		c.JSON(500, utils.FormatDefaultError(err, "Error fetching users"))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseData(users))
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
		c.JSON(http.StatusNotFound, utils.FormatDefaultError(err, "User not found"))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseData(user))
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user request.CreateUserDTO

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.FormatValidationError(err))
		return
	}

	passwordHash, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.FormatDefaultError(err, "Error hashing password"))
		return
	}

	err = h.userRepo.CreateUser(&models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  passwordHash,
	})
	if err != nil {
		c.JSON(500, utils.FormatDefaultError(err, "Error creating user"))
		return
	}

	c.JSON(http.StatusCreated, utils.ResponseData(map[string]interface{}{
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"email":     user.Email,
	}))
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
		c.JSON(http.StatusNotFound, utils.FormatDefaultError(err, "User not found"))
		return
	}

	var user request.UpdateUserDTO

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.FormatValidationError(err))
		return
	}

	existingUser.FirstName = user.FirstName
	existingUser.LastName = user.LastName
	existingUser.Email = user.Email

	if err := h.userRepo.UpdateUser(existingUser); err != nil {
		c.JSON(500, utils.FormatDefaultError(err, "Error updating user"))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseData(map[string]interface{}{
		"firstName": existingUser.FirstName,
		"lastName":  existingUser.LastName,
		"email":     existingUser.Email,
	}))
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
		c.JSON(http.StatusNotFound, utils.FormatDefaultError(err, "User not found"))
		return
	}

	if err := h.userRepo.DeleteUserByID(existingUser); err != nil {
		c.JSON(500, utils.FormatDefaultError(err, "Error deleting user"))
		return
	}

	c.JSON(200, utils.ResponseData(utils.RemoveField(existingUser, []string{"password", "Password"})))
}
