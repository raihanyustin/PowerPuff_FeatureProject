package reviewcontroller

import (
	"net/http"
	"PowerPuff_ReviewBarang/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var review []models.Review

	err := models.DB.Preload("User").Preload("Product").Find(&review).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reviews": review})
	
}