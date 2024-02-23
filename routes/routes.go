package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/products", getProducts)
	server.GET("/products/:id", getProductById)
	server.POST("/products", createProduct)
	server.PUT("/products/:id", updateProduct)
}