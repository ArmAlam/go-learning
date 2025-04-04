package routes

import (
	"example/rest-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	// events
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.POST("/events/:id", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	// users
	server.POST("/signup", signup)
	server.POST("/login", login)

}
