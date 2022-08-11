package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xvbnm48/golang-api-crud/controllers/productController"
	"github.com/xvbnm48/golang-api-crud/models"
)

func main() {
	r := gin.Default()
	models.ConnectionDatabase()

	r.GET("/api/products", productController.Index)
	r.GET("/api/products/:id", productController.Show)
	r.POST("/api/products", productController.Create)
	r.PUT("/api/products/:id", productController.Update)
	r.DELETE("/api/products/:id", productController.Delete)

	r.Run()
}
