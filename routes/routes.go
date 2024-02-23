package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/products", getProducts)
	server.GET("/products/:id", getProductById)
	server.POST("/products", createProduct)
	server.PUT("/products/:id", updateProduct)
	server.DELETE("/products/:id",deleteProduct)

	server.GET("/users", getUsers)
	server.GET("/users/:id", getUserById)
	server.POST("/users", createUser)
	server.PUT("/users/:id", updateUser)
	server.DELETE("/users/:id",deleteUser)

	server.GET("/stores", getStores)
	server.GET("/stores/:id", getStoreById)
	server.POST("/stores", createStore)
	server.PUT("/stores/:id", updateStore)
	server.DELETE("/stores/:id",deleteStore)

	server.POST("/signup", createUser)
}