package app

import (
	"github.com/gin-gonic/gin"

	gqlhandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/ast"

	"github.com/zghost10/go-best-practices/internal/infra/graphql/generated"
	"github.com/zghost10/go-best-practices/internal/infra/graphql/resolver"
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

	srv := gqlhandler.New(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{
		CreateUserUseCase: createUserUseCase,
		GetUserUseCase:    getUserUseCase,
		ListUsersUseCase:  listUsersUseCase,
	}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	if gin.Mode() != gin.ReleaseMode {
		srv.Use(extension.Introspection{})
	}

	router.POST("/query", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})

	if gin.Mode() != gin.ReleaseMode {
		router.GET("/playground", func(c *gin.Context) {
			playground.ApolloSandboxHandler("GraphQL playground", "/query").ServeHTTP(c.Writer, c.Request)
		})
	}
}
