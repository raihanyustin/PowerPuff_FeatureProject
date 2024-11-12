package main

import (
	"PowerPuff_ReviewBarang/controllers/reviewcontroller"

	"PowerPuff_ReviewBarang/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default();
	models.ConnectDatabase()
	
	r.POST("/api/review", reviewcontroller.Create)
	r.GET("/api/review", reviewcontroller.Index)
	r.GET("/api/review/:ProductName", reviewcontroller.Show)
	r.PUT("/api/review/:ID", reviewcontroller.Update)

	r.Run()
}