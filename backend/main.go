package main

import (
	"net/http"

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
	router.Use(cors.New(config))

	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{Domain: "localhost", Path: "/api/", MaxAge: 86400})
	router.Use(sessions.Sessions("mysession", store))

	router.Use(gin.Logger())

	router.Use(restrictAccessMiddleware())

	api := router.Group("/api")
	api.POST("/login", controllers.Login)
	api.POST("/signup", controllers.Signup)
	api.POST("/follow", controllers.Follow)
	api.GET("/isLogin", controllers.IsLogin)
	api.GET("/logout", controllers.Logout)
	api.GET("/userList", controllers.UserList)
	api.GET("/userInfo/:id", controllers.UserInfo)
	api.GET("/userId", controllers.GetUserID)

	router.Run(":8080")
}

func restrictAccessMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("userId")
		if userID == nil && c.Request.URL.Path != "/api/isLogin" && c.Request.URL.Path != "/api/login" && c.Request.URL.Path != "/api/signup" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}
