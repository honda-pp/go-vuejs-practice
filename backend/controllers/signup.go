package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/honda-pp/go-vuejs-practice/backend/models"
	"github.com/lib/pq"
)

func Signup(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "The password field is empty."})
		return
	}
	user.HashPassword()
	db, err := models.NewDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Internal Server Error."})
		return
	}
	defer db.Close()

	err = db.AddUser(user)
	if err != nil {
		var status int
		var message string
		if pgerr, ok := err.(*pq.Error); ok && pgerr.Code == "23505" {
			// Check if the error code is 23505, which indicates a violation of a unique constraint
			status = http.StatusConflict
			message = "A user with this username already exists"
		} else {
			status = http.StatusInternalServerError
			message = "Internal server error"
		}
		c.JSON(status, gin.H{"message": message, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success to create user.",
	})
}
