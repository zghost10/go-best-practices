package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/zghost10/go-best-practices/internal/app"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	mode := os.Getenv("MODE")
	switch mode {
	case "production", "staging":
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()
	app.RegisterHTTP(router)
	if mode == "production" || mode == "staging" {
		fmt.Printf("Server is running on port %s\n", port)
	}
	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
