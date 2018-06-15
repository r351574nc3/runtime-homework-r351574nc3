package main

import (
	"github.com/gin-gonic/gin"
	"github.com/heroku/runtime-homework-r351574nc3/events"
)

const CASE_EVENT_STORE = events.LoadEvents()

func createRouter() *gin.Engine {
	r := gin.Default()

	// Get total hours for cases value
	r.GET("/api/cases/hours", func(c *gin.Context) {
		state := c.Params.ByName("state")
		_, ok := CASE_EVENT_STORE[state]
		if ok {
			c.JSON(200, gin.H{"case_id": 0, "hours": 0})
		} else {
			c.String(404, "No events matching your criteria")
		}
	})
	return r
}

func main() {
	events.Replay(CASE_EVENT_STORE)
	r := createRouter()

	r.Run(":3000")
}
