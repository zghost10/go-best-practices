package app

import (
	"github.com/gin-gonic/gin"
	handler "github.com/zghost10/go-best-practices/internal/infra/http/gin/handler"
	repository "github.com/zghost10/go-best-practices/internal/infra/persistence/memory/user"
	usecase "github.com/zghost10/go-best-practices/internal/usecase/user"
)

func RegisterHTTP(router *gin.Engine) {
	handler.NewHealthHandler(router)

	userRepo := repository.NewInMemoryUserRepo()
	createUserUseCase := usecase.NewCreateUserUseCase(userRepo)
	getUserUseCase := usecase.NewGetUserUseCase(userRepo)
	listUsersUseCase := usecase.NewListUsersUseCase(userRepo)
	handler.NewUserHandler(router, createUserUseCase, getUserUseCase, listUsersUseCase)
}
