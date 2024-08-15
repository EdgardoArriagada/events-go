package routes

import (
	"strconv"

	"example.com/events-go/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not get event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not register for event"})
		return
	}

	c.JSON(201, gin.H{"message": "Registered for event", "data": event})
}

func cancelRegistration(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not get event"})
		return
	}

	err = event.CancelRegistration(userId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not cancel registration for event"})
		return
	}

	c.JSON(200, gin.H{"message": "Registration canceled", "data": event})
}
