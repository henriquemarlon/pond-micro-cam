package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/henriquemarlon/pond-micro-cam/backend/internal/domain/dto"
	"github.com/henriquemarlon/pond-micro-cam/backend/internal/usecase"
)

type UserHandlers struct {
	UserUseCase *usecase.UserUseCase
}

func NewUserHandlers(userUseCase *usecase.UserUseCase) *UserHandlers {
	return &UserHandlers{
		UserUseCase: userUseCase,
	}
}

// CreateUser
// @Summary Create a new User entity
// @Description Create a new User entity
// @Tags Users
// @Accept json
// @Produce json
// @Param input body dto.CreateUserInputDTO true "User entity to create"
// @Success 201 {object} dto.CreateUserOutputDTO
// @Router /user [post]
func (h *UserHandlers) CreateUser(c *gin.Context) {
	var input dto.CreateUserInputDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	output, err := h.UserUseCase.CreateUser(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, output)
}

// FindAllUsersHandler
// @Summary Retrieve all User entities
// @Description Retrieve all User entities
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} dto.FindUserOutputDTO
// @Router /user [get]
func (h *UserHandlers) FindAllUsersHandler(c *gin.Context) {
	output, err := h.UserUseCase.FindAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

// FindUserByIdHandler
// @Summary Retrieve a User entity by ID
// @Description Retrieve a User entity by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} dto.FindUserOutputDTO
// @Router /user/{id} [get]
func (h *UserHandlers) FindUserByIdHandler(c *gin.Context) {
	var input dto.FindUserByIdInputDTO
	input.ID = c.Param("id")
	output, err := h.UserUseCase.FindUserById(input.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

// LoginUserHandler
// @Summary Login a User
// @Description Login a User
// @Tags Login
// @Accept json
// @Produce json
// @Param input body dto.LoginInputDTO true "User entity"
// @Success 200 {object} dto.LoginOutputDTO
// @Router /login [post]
func (h *UserHandlers) LoginUserHandler(c *gin.Context) {
	var input dto.LoginInputDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	output, err := h.UserUseCase.LoginUser(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}
