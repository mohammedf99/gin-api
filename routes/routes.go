package routes

import (
	"example/goUdemyRest/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents) // could be: GET, PUT, PATCH, DELETE
	server.GET("/events/:id", getEvent)

	protected := server.Group("/")
	protected.Use(middlewares.Authenticate)
	protected.POST("/events", createEvent)
	protected.PUT("/events/:id", updateEvent)
	protected.DELETE("/events/:id", deleteEvent)
	protected.POST("/events/:id/register", registerForEvent)
	protected.DELETE("/events/:id/cancel", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
	// server.POST("/events", middlewares.Authenticate, createEvent) // First way
}

// There are two ways to use the middleware. First way is in createEvent route
