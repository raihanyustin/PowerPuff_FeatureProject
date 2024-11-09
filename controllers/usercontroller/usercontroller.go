package usercontroller

import (
	"net/http"
	"PowerPuff_ReviewBarang/models"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var user []models.User

	models.DB.Find(&user)
	c.JSON(http.StatusOK, gin.H{"user": user})
	
}