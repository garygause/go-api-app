package routes

import (
	"net/http"
	"strconv"

	"github.com/garygause/go-api-app/models"
	"github.com/gin-gonic/gin"
)

func getProductById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Product id is invalid."})
		return
	}
	product, err := models.GetProductById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve Product."})
		return
	}
	context.JSON(http.StatusOK, product)
}

func getProducts(context *gin.Context) {
	products, err := models.GetAllProducts()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get Products."})
		return
	}
	context.JSON(http.StatusOK, products)
}

func createProduct(context *gin.Context) {
	var product models.Product
	err := context.ShouldBindJSON(&product)
	
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = product.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save Product."})
		return
	}
	
	context.JSON(http.StatusCreated, gin.H{"message": "Product created.", "product": product})
}