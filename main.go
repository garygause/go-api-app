package main

import (
	"github.com/garygause/go-api-app/db"
	"github.com/garygause/go-api-app/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8000")
}
