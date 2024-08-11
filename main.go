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

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		c.JSON(500, gin.H{"error": "Could not get events"})
		return
	}

	c.JSON(200, events)
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	event.Id = 1
	event.UserId = 1

	err = event.Save()

	if err != nil {
		c.JSON(500, gin.H{"error": "Could not save event"})
		return
	}

	c.JSON(201, gin.H{"message": "event created", "data": event})
}

func getEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not get event"})
		return
	}

	c.JSON(200, event)
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}
