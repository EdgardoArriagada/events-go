package main

import (
	"example.com/events-go/db"
	"example.com/events-go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	err := server.Run(":8080")

	if err != nil {
		panic(err)
	}
}
