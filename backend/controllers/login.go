package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/honda-pp/go-vuejs-practice/backend/models"
)

type LoginForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func Login(c *gin.Context) {
	var loginForm LoginForm
	if err := c.ShouldBind(&loginForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := models.NewDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	user, err := db.GetUser(loginForm.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = user.CheckPassword(loginForm.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to login"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "login successful",
		"id":       strconv.Itoa(user.ID),
		"username": user.Username,
	})
}
