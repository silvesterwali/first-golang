package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"myproject/repositories"
	"myproject/request"
	"myproject/utils"

	"github.com/gin-gonic/gin"
)

type UserAuthHandler struct {
	userRepo *repositories.UserRepository
}

func NewAuthHandler() *UserAuthHandler {
	return &UserAuthHandler{
		userRepo: repositories.NewUserRepository(),
	}
}

func (h *UserAuthHandler) Login(c *gin.Context) {
	var login request.LoginDTO

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, utils.FormatValidationError(err))
		return
	}

	user, err := h.userRepo.GetUserByEmail(login.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.FormatDefaultError(err, "Invalid credentials"))
		return
	}

	isValidPassword := utils.CheckPasswordHash(login.Password, user.Password)

	if !isValidPassword {
		c.JSON(http.StatusUnauthorized, utils.FormatDefaultError(err, "Invalid credentials"))
		return
	}

	token, err := utils.CreateJWTToken(strconv.Itoa(user.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.FormatDefaultError(err, "Error creating token"))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseData(map[string]interface{}{
		"token": token,
		"user": map[string]interface{}{
			"ID":        user.ID,
			"firstName": user.FirstName,
			"lastName":  user.LastName,
			"email":     user.Email,
		},
	}))
}

func (h *UserAuthHandler) Profile(c *gin.Context) {
	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusUnauthorized, utils.FormatDefaultError(nil, "Unauthorized"))
		return
	}

	// Convert string to int
	userIdInt, err := strconv.Atoi(userId.(string))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	user, err := h.userRepo.GetUserByID(userIdInt)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.FormatDefaultError(err, "User not found"))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseData(map[string]interface{}{
		"ID":        user.ID,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"email":     user.Email,
	}))
}
