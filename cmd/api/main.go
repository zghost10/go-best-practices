package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zghost10/go-best-practices/internal/app"
)

func main() {
	router := gin.Default()
	app.RegisterHTTP(router)
	router.Run()
}
