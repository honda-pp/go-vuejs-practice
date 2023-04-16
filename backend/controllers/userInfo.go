package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/honda-pp/go-vuejs-practice/backend/models"
)

func UserInfo(c *gin.Context) {
	db, err := models.NewDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Internal Server Error."})
		return
	}
	defer db.Close()
	userInfo, err := db.GetUserFromID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Internal Server Error."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"userInfo": userInfo})

}
