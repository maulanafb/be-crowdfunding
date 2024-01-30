package handler

import (
	"be_crowdfunding/helper"
	"be_crowdfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// tangkap input dari user
	// map input dari user ke struct RegisterUserInput
	// struct di atas kita passing sebagai parameter servicec
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		var errors []string
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Error())
		}

		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(newUser, "jwttoken")
	response := helper.APIResponse("Account successfully registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
