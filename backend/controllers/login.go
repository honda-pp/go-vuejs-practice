package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/honda-pp/go-vuejs-practice/backend/models"
)

func Login(c *gin.Context) {
	logger := gin.DefaultWriter
	var loginForm models.LoginForm
	if err := c.ShouldBind(&loginForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logger.Write([]byte(loginForm.Username + " "))
	logger.Write([]byte(loginForm.Password + "\n"))
	c.JSON(http.StatusOK, gin.H{"message": "login form received"})
}
