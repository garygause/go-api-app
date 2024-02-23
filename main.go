package main

import (
	"net/http"

	"github.com/garygause/go-api-app/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/products", getProducts)
	server.POST("/products", createProduct)

	server.Run(":8000")
}

func getProducts(context *gin.Context) {
	products := models.GetAllProducts()
	context.JSON(http.StatusOK, products)
}

func createProduct(context *gin.Context) {
	var product models.Product
	err := context.ShouldBindJSON(&product)
	
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	product.ID = 1
	product.StoreID = 1
	product.Save()
	
	context.JSON(http.StatusCreated, gin.H{"message": "Product created.", "product": product})
}