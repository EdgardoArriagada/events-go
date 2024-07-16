package main

import (
	"example.com/events-go/db"
	"example.com/events-go/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(200, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	event.Id = 1
	event.UserId = 1

	context.JSON(201, gin.H{"message": "event created", "data": event})
}
