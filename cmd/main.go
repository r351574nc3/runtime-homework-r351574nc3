package main

import (
	"github.com/gin-gonic/gin"
	"github.com/heroku/runtime-homework-r351574nc3/common"
	"github.com/heroku/runtime-homework-r351574nc3/events"
	"github.com/heroku/runtime-homework-r351574nc3/events/types"
)

const CASE_EVENT_STORE = events.LoadEvents()

func createRouter() *gin.Engine {
	r := gin.Default()

	// Get total hours for cases value
	r.GET("/api/cases/hours", func(c *gin.Context) {
		state := c.Params.ByName("state")
		team := c.Params.ByName("team")
		summaries, ok := events.DurationIndex.Filter(func(e *types.StateEventSummary) {
			return e.Status == state && e.Team == team
		})

		if ok {
			response HoursResponse = make(HoursResponse)
			for summary := range summaries {
				response = append(response, DurationSummary{ CaseId: summary.Id, Duration: summary.Duration })
			}

			c.JSON(200, gin.H{ response })
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
