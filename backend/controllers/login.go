package controllers

import (
	"net/http"
	"strconv"

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
	db, err := models.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, username, email, password_hash FROM users WHERE username = $1 AND password_hash = $2", loginForm.Username, loginForm.Password)
	if err != nil {
		logger.Write([]byte(err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	if rows.Next() {
		var id int
		var username string
		var email string
		var passwordHash string
		if err := rows.Scan(&id, &username, &email, &passwordHash); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			logger.Write([]byte(err.Error()))
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message":  "login successful",
			"id":       strconv.Itoa(id),
			"username": username,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to login"})
	}
}
