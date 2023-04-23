package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/honda-pp/go-vuejs-practice/backend/models"
)

type FollowRequest struct {
	FolloweeID int `json:"followee_id" binding:"required"`
}

func Follow(c *gin.Context) {
	var req FollowRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := sessions.Default(c)
	followerID, ok := session.Get("userId").(int)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	db, err := models.NewDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Internal Server Error."})
		return
	}
	defer db.Close()

	err = db.Follow(followerID, req.FolloweeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Error occurred while attempting to follow. Please try again later."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}
