package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zghost10/go-best-practices/internal/infra/http/gin/handlers"
)

func main() {
	router := gin.Default()

	router.GET("/", handlers.HelloController)

	router.Run(":8080")
}
