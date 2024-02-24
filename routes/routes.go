package routes

import (
	"github.com/garygause/go-api-app/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	
	server.GET("/products", getProducts)
	server.GET("/products/:id", getProductById)
	authenticated.POST("/products", createProduct)
	authenticated.PUT("/products/:id", updateProduct)
	authenticated.DELETE("/products/:id",deleteProduct)

	server.GET("/users", getUsers)
	server.GET("/users/:id", getUserById)
	authenticated.POST("/users", createUser)
	authenticated.PUT("/users/:id", updateUser)
	authenticated.DELETE("/users/:id",deleteUser)

	server.GET("/stores", getStores)
	server.GET("/stores/:id", getStoreById)
	authenticated.POST("/stores", createStore)
	authenticated.PUT("/stores/:id", updateStore)
	authenticated.DELETE("/stores/:id",deleteStore)

	server.POST("/signup", createUser)
	server.POST("/login", loginUser)
}