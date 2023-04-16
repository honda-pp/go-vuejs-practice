package main

import (
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
	api.POST("/login", controllers.Login)
	api.POST("/signup", controllers.Signup)
	api.GET("/userList", controllers.UserList)
	api.GET("/userInfo/:id", controllers.UserInfo)

	router.Run(":8080")
}
