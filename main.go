package main

import (
	"github.com/erkindilekci/events-api/db"
	"github.com/erkindilekci/events-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
