package routes

import (
	"example.com/events-go/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	auth := server.Group("/")
	auth.Use(middlewares.Auth)

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	auth.POST("/events", createEvent)
	auth.PUT("/events/:id", updateEvent)
	auth.DELETE("/events/:id", deleteEvent)

	auth.POST("/events/:id/register", registerForEvent)
	auth.DELETE("/events/:id/register")

	server.GET("/users", getUsers)
	server.POST("/signup", signup)
	server.POST("/login", login)

	server.GET("/health", healthCheck)
}
