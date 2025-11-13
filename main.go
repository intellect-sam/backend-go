package main

import (
	"github.com/intellect-sam/backend-go/db"
	"github.com/intellect-sam/backend-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080") // localhost:8080
}
