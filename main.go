package main

import (
	"strconv"

	"example.com/events-go/db"
	"example.com/events-go/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.GET("/events/:id", getEvent)

	server.GET("/health", healthCheck)

	err := server.Run(":8080")

	if err != nil {
		panic(err)
	}
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(500, gin.H{"error": "Could not get events"})
		return
	}

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

	err = event.Save()

	if err != nil {
		context.JSON(500, gin.H{"error": "Could not save event"})
		return
	}

	context.JSON(201, gin.H{"message": "event created", "data": event})
}

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(500, gin.H{"error": "Could not get event"})
		return
	}

	context.JSON(200, event)
}

func healthCheck(context *gin.Context) {
	context.JSON(200, gin.H{"status": "ok"})
	return
}
