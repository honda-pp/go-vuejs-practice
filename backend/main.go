package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/honda-pp/go-vuejs-practice/backend/controllers"
)

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:88"}
	config.AllowCredentials = true
	router.Use(gin.Logger())
	router.Use(cors.New(config))

	api := router.Group("/api")
	api.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Your App",
		})
	})
	{
		api.POST("/login", controllers.Login)
	}
	{
		api.POST("/signup", controllers.Signup)
	}

	router.Run(":8080")
}
