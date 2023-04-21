package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/honda-pp/go-vuejs-practice/backend/controllers"
)

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:88"}
	config.AllowCredentials = true
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	router.Use(gin.Logger())
	router.Use(cors.New(config))

	api := router.Group("/api")
	api.POST("/login", controllers.Login)
	api.POST("/signup", controllers.Signup)
	api.POST("/follow", controllers.Follow)
	api.GET("/userList", controllers.UserList)
	api.GET("/userInfo/:id", controllers.UserInfo)
	api.GET("/userId/", controllers.GetUserID)
	router.Run(":8080")
}
