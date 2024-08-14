package routes

import (
	"strconv"

	"example.com/events-go/models"
	"github.com/gin-gonic/gin"
)

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

func updateEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	_, err = models.GetEventById(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not get event"})
		return
	}

	var updatedEvent models.Event

	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	updatedEvent.Id = id

	err = updatedEvent.Update()
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not update event"})
		return
	}

	c.JSON(200, gin.H{"message": "Event updated", "data": updatedEvent})
}

func deleteEvent(c *gin.Context) {
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

	err = event.Delete()
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not delete event"})
		return
	}

	c.JSON(200, gin.H{"message": "Event deleted", "data": event})
}
