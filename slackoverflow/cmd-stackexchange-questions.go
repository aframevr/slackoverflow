package slackoverflow

import (
	"time"

	"github.com/aframevr/slackoverflow/sqlite3"
	"github.com/aframevr/slackoverflow/stackexchange"
	"github.com/aframevr/slackoverflow/std"
)

type cmdStackExchangeQuestions struct {
	GetNewQuestions bool `long:"get" description:"Get new questions from configured Stack Exchange Site"`
	UpdateQuestions bool `long:"update" description:"Update information about existing questions"`
	All             bool `long:"all" description:"Get new questions and update information about existing questions"`
}

func (cse *cmdStackExchangeQuestions) Execute(args []string) error {

	// Refresh the session before running this command and make sure that Slack Overflow is configured
	slackoverflow.SessionRefresh(true)

	if !slackoverflow.config.StackExchange.Enabled {
		Notice("Running Stack Exchange related commands is disabled in configuration file. Skipping.")
		return nil
	}

	// Get new questions from configured Stack Exchange Site
	if cse.GetNewQuestions {
		cse.getNewQuestions()
	}

	// Get new questions from configured Stack Exchange Site
	if cse.UpdateQuestions {
		cse.updateQuestions()
	}

	// Get new questions from configured Stack Exchange Site
	if cse.All {
		cse.getNewQuestions()
		cse.updateQuestions()
	}

	return nil
}

// Get new Questions
func (cse *cmdStackExchangeQuestions) getNewQuestions() {
	Info("Stack Exchange: Checking for new questions.")

	question := sqlite3.StackExchangeQuestion{}.Latest()

	// Check do we already have some questions do obtain from_date for next request
	if question.QID == 0 {
		Notice("There are no questions in database,")
		Info("That is ok if current execution is first time you run slackoverflow")
		Info("Or there are no questions tagged with %s on site %s",
			slackoverflow.config.StackExchange.SearchAdvanced["tagged"],
			slackoverflow.config.StackExchange.Site,
		)
		question.CreationDate = time.Now().UTC().Add(-24 * time.Hour)
	}

	now := time.Now().UTC().Unix()
	diff := question.CreationDate.Unix() - now
	// Allow maximum 48 h old question as from date otherwise we may exhaust
	// rate limit if last tracked questions fromdata is to far past
	// and slackoverflow has not been running for a while.
	if diff > (48 * 60) {
		question.CreationDate = time.Now().UTC().Add(-24 * time.Hour)
	}

	Debug("Checking new questions since %s", question.CreationDate.String())

	// Check for New Questions from Stack Exchange
	searchAdvanced := slackoverflow.StackExchange.SearchAdvanced()

	// Set it here so that it is allowed to override by config
	searchAdvanced.Parameters.Set("site", slackoverflow.config.StackExchange.Site)

	// Set all parameters from config
	for param, value := range slackoverflow.config.StackExchange.SearchAdvanced {
		searchAdvanced.Parameters.Set(param, value)
	}

	searchAdvanced.Parameters.Set("fromdate", question.CreationDate.Unix()+1)

	// Output query as table
	if slackoverflow.Debugging() {
		searchAdvanced.DrawQuery()
	}

	fetchQuestions := true

	for fetchQuestions {
		Info("Fetching page %d", searchAdvanced.GetCurrentPageNr())
		std.Hr()
		if results, err := searchAdvanced.Get(); results {
			// Questions recieved
			for _, q := range searchAdvanced.Result.Items {
				std.Body("Question: %s", q.Title)
				std.Body("Url:      %s", q.ShareLink)
				if slackoverflow.Debugging() {
					newq := std.NewTable("Question ID", "Time", "Answers", "Comments", "Score", "Views", "Username")
					newq.AddRow(
						q.QID,
						time.Unix(q.CreationDate, 0).UTC().Format("15:04:05 Mon Jan _2 2006"),
						q.AnswerCount,
						q.CommentCount,
						q.Score,
						q.ViewCount,
						q.Owner.DisplayName,
					)
					newq.Print()
				}

				cse.syncQuestion(q)
				std.Hr()
			}
			if err != nil {
				fetchQuestions = false
				Error(err.Error())
			}
			std.Hr()
		}

		// Done go to next page
		if searchAdvanced.HasMore() || searchAdvanced.GetCurrentPageNr() > 10 {
			searchAdvanced.NextPage()
		} else {
			fetchQuestions = false
			Ok("There are no more new questions.")
		}
	}
	Info(
		"Stack Exchange Quota usage (%d/%d)",
		slackoverflow.StackExchange.GetQuotaRemaining(),
		slackoverflow.StackExchange.GetQuotaMax(),
	)
}

// Get new Questions
func (cse *cmdStackExchangeQuestions) updateQuestions() {
	Info("Stack Exchange: Updating existing questions.")

	questionIds, questionIdsCount := sqlite3.StackExchangeQuestion{}.TrackedIds(
		slackoverflow.config.StackExchange.QuestionsToWatch)

	// Check do we already have some questions do obtain from_date for next request
	if questionIdsCount == 0 {
		Notice("There are no questions in database,")
		Info("That is ok if current execution is first time you run slackoverflow")
		Info("and this case run 'slackoverflow stackechange guestions' --get first")
	}

	Debug("Checking updates for %d questions. Max to be tracked (%d)",
		questionIdsCount, slackoverflow.config.StackExchange.QuestionsToWatch)

	// Check for New Questions from Stack Exchange
	updateQuestions := slackoverflow.StackExchange.Questions()

	// Set it here so that it is allowed to override by config
	updateQuestions.Parameters.Set("site", slackoverflow.config.StackExchange.Site)

	for param, value := range slackoverflow.config.StackExchange.Questions {
		updateQuestions.Parameters.Set(param, value)
	}

	// Output query as table
	if slackoverflow.Debugging() {
		updateQuestions.DrawQuery(questionIds)
	}

	fetchQuestions := true

	for fetchQuestions {
		Info("Fetching page %d", updateQuestions.GetCurrentPageNr())
		std.Hr()
		if results, err := updateQuestions.Get(questionIds); results {
			// Questions recieved
			for _, q := range updateQuestions.Result.Items {
				std.Body("Question: %s", q.Title)
				std.Body("Url:      %s", q.ShareLink)
				if slackoverflow.Debugging() {
					newq := std.NewTable("Question ID", "Time", "Answers", "Comments", "Score", "Views", "Username")
					newq.AddRow(
						q.QID,
						time.Unix(q.CreationDate, 0).UTC().Format("15:04:05 Mon Jan _2 2006"),
						q.AnswerCount,
						q.CommentCount,
						q.Score,
						q.ViewCount,
						q.Owner.DisplayName,
					)
					newq.Print()
				}
				cse.syncQuestion(q)
				std.Hr()
			}
			if err != nil {
				fetchQuestions = false
				Error(err.Error())
			}
			std.Hr()
		}

		// Done go to next page
		if updateQuestions.HasMore() {
			updateQuestions.NextPage()
		} else {
			fetchQuestions = false
			Ok("There are no more questions to update.")
		}
	}
	Info(
		"Stack Exchange Quota usage (%d/%d)",
		slackoverflow.StackExchange.GetQuotaRemaining(),
		slackoverflow.StackExchange.GetQuotaMax(),
	)
}

// Store new questions
func (cse *cmdStackExchangeQuestions) syncQuestion(q stackexchange.QuestionObj) {
	var ok string
	var err error
	// Create or Update user
	ok, err = sqlite3.StackExchangeUser{}.SyncShallowUser(q.Owner)
	if err != nil {
		Error(err.Error())
	} else {
		Ok(ok)
	}
	// Create or Update user
	ok, err = sqlite3.StackExchangeQuestion{}.SyncFromSite(q, slackoverflow.config.StackExchange.Site)
	if err != nil {
		Error(err.Error())
	} else {
		Ok(ok)
	}
}
