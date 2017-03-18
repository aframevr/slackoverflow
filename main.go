// Web hook that posts tagged Stack Overflow questions to Slack, updated using reaction emojis.
// See README.md for usage.
package main

import (
	"time"

	"github.com/aframevr/slackoverflow/slackoverflow"
)

var (
	startTime = time.Now()
	buildDate string
)

// Core config object

func main() {
	slackoverflow := slackoverflow.Start()
	slackoverflow.SetStartTime(startTime)
	slackoverflow.Info.SetBuildDate(buildDate)

	slackoverflow.Run()

	slackoverflow.Close(0)
}
