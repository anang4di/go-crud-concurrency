package handler

import (
	"go-crud-concurrency/helper"
	"go-crud-concurrency/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service user.Service
}

func NewUserHandler(service user.Service) *userHandler {
	return &userHandler{service}
}

func (h *userHandler) GetAllUsers(c *gin.Context) {
	user, err := h.service.GetAllUsers()
	if err != nil {
		response := helper.APIResponse("Get all users failed", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.APIResponse("Get all users success", http.StatusOK, "success", user)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.UserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationErrors(err)
		errorMessages := gin.H{"errors": errors}

		response := helper.APIResponse("Register user failed", http.StatusUnprocessableEntity, "error", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := h.service.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register user failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Register user success", http.StatusCreated, "success", user)
	c.JSON(http.StatusCreated, response)
}
