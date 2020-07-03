package main

import (
	"github.com/gin-gonic/gin"
	"test.com/borzoj/pde/handlers"
)

func main() {
	handler := handlers.NewEventsPostHandler()
	r := gin.Default()
	r.POST("/events", func(c *gin.Context) {
		handler.Handle(c)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
