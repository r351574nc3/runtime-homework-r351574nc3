package main

import (
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/heroku/runtime-homework-r351574nc3/common"
	"github.com/heroku/runtime-homework-r351574nc3/events"
	"github.com/heroku/runtime-homework-r351574nc3/events/store"
	"github.com/heroku/runtime-homework-r351574nc3/events/types"
)

func createRouter() *gin.Engine {
	r := gin.Default()

	// Get total hours for cases value
	r.POST("/cases", func(c *gin.Context) {
		var event_arr []types.Event

		team := c.DefaultQuery("team", "Runtime")
		state := c.DefaultQuery("state", "open")

		if err := c.ShouldBindJSON(&event_arr); err == nil {
			sort.Sort(events.ByTimestamp(event_arr))
			events.Replay(event_arr)
		}

		summaries, ok := store.DurationIndex.Filter(func(e types.StateEventSummary) bool {
			if state != "" && team != "" {
				return e.Status == state && e.Team == team
			}

			if state != "" {
				return e.Status == state
			}

			if team != "" {
				return e.Team == team
			}

			return true
		})

		if ok {
			response := make(common.HoursResponse, 0)
			for _, summary := range summaries {
				response = response.Increment(common.DurationSummary{CaseId: summary.Id, Duration: summary.Duration.Hours()})
			}

			c.JSON(200, response)
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
