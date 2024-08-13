package routes

import (
	"example.com/events-go/models"
	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "Could not parse request data."})
		return
	}

	err = user.Save()
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not save user."})
		return
	}

	c.JSON(201, gin.H{"message": "User created successfully."})
}
