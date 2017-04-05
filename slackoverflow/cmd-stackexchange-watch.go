package slackoverflow

import (
	"os"
	"os/signal"
	"time"

	"github.com/aframevr/slackoverflow/std"
	"github.com/robfig/cron"
)

var fromDate time.Time

type cmdStackExchangeWatch struct{}

func (cse *cmdStackExchangeWatch) Execute(args []string) error {
	// Refresh the session before running this command and make sure that Slack Overflow is configured
	slackoverflow.SessionRefresh(true)
	Info("Waiting for new questions!")
	// Get at first questions from past 30 min
	fromDate = time.Now().UTC().Add(-30 * time.Minute)
	cse.StartWatching()

	watch := cron.New()
	watch.AddFunc("@every 1m", cse.StartWatching)
	go watch.Start()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
	return nil
}
func (cse *cmdStackExchangeWatch) StartWatching() {

	// Check for New Questions from Stack Exchange
	searchAdvanced := slackoverflow.StackExchange.SearchAdvanced()

	// Set it here so that it is allowed to override by config
	searchAdvanced.Parameters.Set("site", slackoverflow.config.StackExchange.Site)

	// Set all parameters from config
	for param, value := range slackoverflow.config.StackExchange.SearchAdvanced {
		searchAdvanced.Parameters.Set(param, value)
	}

	searchAdvanced.Parameters.Set("fromdate", fromDate.Unix()+1)
	fetchQuestions := true
	for fetchQuestions {
		Debug("Fetching page %d", searchAdvanced.GetCurrentPageNr())
		if results, err := searchAdvanced.Get(); results {
			// Questions recieved
			for _, q := range searchAdvanced.Result.Items {
				fromDate = time.Unix(q.CreationDate, 0).UTC()
				std.Body("Question: %s", q.Title)
				std.Body("Url:      %s", q.ShareLink)
				newq := std.NewTable("Question ID", "Time", "Answers", "Comments", "Score", "Views", "Username")
				newq.AddRow(
					q.QID,
					time.Unix(q.CreationDate, 0).Local().Format("15:04:05 Mon Jan _2 2006"),
					q.AnswerCount,
					q.CommentCount,
					q.Score,
					q.ViewCount,
					q.Owner.DisplayName,
				)
				newq.Print()
				std.Hr()
			}
			if err != nil {
				fetchQuestions = false
				Error(err.Error())
			}

			if len(searchAdvanced.Result.Items) > 0 {
				Ok("Waiting for new questions!")
			}
		}

		// Done go to next page
		if searchAdvanced.HasMore() || searchAdvanced.GetCurrentPageNr() > 10 {
			searchAdvanced.NextPage()
		} else {
			fetchQuestions = false
		}
	}

	Debug(
		"Stack Exchange Quota usage (%d/%d)",
		slackoverflow.StackExchange.GetQuotaRemaining(),
		slackoverflow.StackExchange.GetQuotaMax(),
	)
}
