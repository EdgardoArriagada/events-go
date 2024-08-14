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

func login(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "Could not parse request data."})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		c.JSON(401, gin.H{"error": "Could not authenticate user."})
		return
	}

	c.JSON(200, gin.H{"message": "Login Successful."})
}

func getUsers(c *gin.Context) {

	users, err := models.GetAllUsers()

	if err != nil {
		c.JSON(500, gin.H{"error": "Could not get users"})
		return
	}

	c.JSON(200, users)
}
