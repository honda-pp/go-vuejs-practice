package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/honda-pp/go-vuejs-practice/backend/models"
)

func UserList(c *gin.Context) {
	db, err := models.NewDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Internal Server Error."})
		return
	}
	defer db.Close()
	userList, err := db.GetUserList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Internal Server Error."})
		return
	}
	jsonBytesUserList, err := json.Marshal(userList)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Internal Server Error."})
		return
	}
	jsonStringUserList := string(jsonBytesUserList)
	c.JSON(http.StatusOK, gin.H{"userList": jsonStringUserList})

}
