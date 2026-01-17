package main

import (
	"github.com/gin-gonic/gin"

	userdbinmemory "github.com/zghost10/go-best-practices/internal/infra/db/in_memory/repository"
	"github.com/zghost10/go-best-practices/internal/infra/http/gin/app"
	userinfra "github.com/zghost10/go-best-practices/internal/infra/http/gin/user"
	userusecase "github.com/zghost10/go-best-practices/internal/usecase/user"
)

func main() {
	router := gin.Default()

	app.NewHealthHandler(router)

	userRepo := userdbinmemory.NewInMemoryUserRepo()
	createUserUseCase := userusecase.NewCreateUserUseCase(userRepo)
	getUserUseCase := userusecase.NewGetUserUseCase(userRepo)
	listUsersUseCase := userusecase.NewListUsersUseCase(userRepo)
	userinfra.NewUserHandler(router, createUserUseCase, getUserUseCase, listUsersUseCase)

	router.Run(":8080")
}
