package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zghost10/go-best-practices/internal/infra/http/gin/dto"
	er "github.com/zghost10/go-best-practices/internal/infra/http/gin/error"
	usecase "github.com/zghost10/go-best-practices/internal/usecase/user"
)

type UserHandler struct {
	createUserUseCase *usecase.CreateUserUseCase
	getUserUseCase    *usecase.GetUserUseCase
	listUsersUseCase  *usecase.ListUsersUseCase
}

func NewUserHandler(e *gin.Engine, createUserUseCase *usecase.CreateUserUseCase, getUserUseCase *usecase.GetUserUseCase, listUsersUseCase *usecase.ListUsersUseCase) *UserHandler {
	h := &UserHandler{createUserUseCase: createUserUseCase, getUserUseCase: getUserUseCase, listUsersUseCase: listUsersUseCase}
	g := e.Group("/users")

	g.GET("/:id", h.GetUser)
	g.POST("", h.CreateUser)
	g.GET("", h.ListUsers)

	return h
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var input dto.CreateUserDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	createUserOutput, err := h.createUserUseCase.CreateUser(usecase.CreateUserInput{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  input.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, createUserOutput)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	var input usecase.GetUserInput
	input.Identifier = c.Param("id")
	getUserOutput, err := h.getUserUseCase.GetUser(input)
	if err != nil {
		if errors.Is(err, er.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, getUserOutput)
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	listUsersOutput, err := h.listUsersUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, listUsersOutput)
}
