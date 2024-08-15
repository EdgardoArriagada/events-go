package middlewares

import (
	"example.com/events-go/utils"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	c.Set("userId", userId)

	c.Next()
}
