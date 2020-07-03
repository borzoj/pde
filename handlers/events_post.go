package handlers

import (
	"github.com/gin-gonic/gin"
	"test.com/borzoj/pde/events"
)

// EventsPostHandler hndler
type EventsPostHandler struct {
	service events.ServiceInterface
}

// NewEventsPostHandler new
func NewEventsPostHandler() *EventsPostHandler {

	return &EventsPostHandler{
		service: events.NewService(),
	}
}

// Handle handle
func (*EventsPostHandler) Handle(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
