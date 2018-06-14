package main

import (
	"github.com/gin-gonic/gin"
)

var CASE_EVENT_STORE = make(map[string]string)

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
	r := createRouter()

	r.Run(":3000")
}
