package main

import (
	"example/rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")

}

func getEvents(context *gin.Context) {

	events := models.GetAllEvents()

	context.JSON(http.StatusOK, events)

}

func createEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event) // gin will populate data from the body for event

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Can't parse request body"})

		return
	}

	// DUMMY
	event.ID = 1
	event.UserId = 2

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}
