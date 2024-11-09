package main

import (
	"PowerPuff_ReviewBarang/controllers/productcontroller"
	"PowerPuff_ReviewBarang/controllers/reviewcontroller"
	"PowerPuff_ReviewBarang/controllers/usercontroller"
	"PowerPuff_ReviewBarang/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default();
	models.ConnectDatabase()
	
	r.GET("/api/user", usercontroller.Index)
	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/review", reviewcontroller.Index)
	r.GET("/api/product/:id", productcontroller.Show)
	r.POST("/api/product", productcontroller.Create)
	r.PUT("/api/product/:id", productcontroller.Update)

	r.Run()
}