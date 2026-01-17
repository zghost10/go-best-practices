package app

import "github.com/gin-gonic/gin"

func NewHealthHandler(e *gin.Engine) {
	e.GET("/", HelloController)
}

func HelloController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "The API is running!",
	})
}
