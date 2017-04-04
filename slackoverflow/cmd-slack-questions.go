package slackoverflow

import (
	"fmt"

	"github.com/aframevr/slackoverflow/slack"
	"github.com/aframevr/slackoverflow/sqlite3"
)

const (
	msgNotAnswered = "#B7E0ED"
	msgIsAnswewed  = "#30AC1F"
	thumbUp        = ":+1:"
	thumbDown      = ":-1:"
)

type cmdSlackQuestions struct {
	PostNewQuestions bool `long:"post-new" description:"Post new questions origin from configured Stack Exchange Site"`
	UpdateQuestions  bool `long:"update" description:"Update information about questions already posted to slack"`
	All              bool `long:"all" description:"Get new questions and update information about existing questions"`
}

// Execute
func (sl *cmdSlackQuestions) Execute(args []string) error {

	// Refresh the session before running this command and make sure that Slack Overflow is configured
	slackoverflow.SessionRefresh(true)

	if !slackoverflow.config.Slack.Enabled {
		Notice("Running Slack related commands is disabled in configuration file. Skipping.")
		return nil
	}

	// Psot new questions from configured Stack Exchange Site
	if sl.PostNewQuestions {
		sl.postNewQuestions()
	}

	// Update questions from configured Stack Exchange Site
	if sl.UpdateQuestions {
		sl.updateQuestions()
	}

	// Get new questions from configured Stack Exchange Site
	if sl.All {
		sl.postNewQuestions()
		sl.updateQuestions()
	}
	return nil
}

func (sl *cmdSlackQuestions) postNewQuestions() {
	Info("Slack: Posting new questions.")

	tracked, questionCount := sqlite3.StackExchangeQuestion{}.Tracked(
		slackoverflow.config.StackExchange.QuestionsToWatch)
	if questionCount == 0 {
		Notice("Slack: There are no questions in database,")
		return
	}

	// Process questions
	for _, question := range tracked {
		slackQuestion := sqlite3.SlackQuestion{}.Find(question.QID)
		if slackQuestion.QID == 0 {
			user := sqlite3.StackExchangeUser{}.Find(question.UID)

			params := slack.NewPostMessageParameters()
			params.Parse = "full"
			params.LinkNames = 1
			params.UnfurlLinks = true
			params.UnfurlMedia = false
			params.Username = fmt.Sprintf("%s asked on %s:",
				user.DisplayName,
				question.Site,
			)
			params.AsUser = false
			params.IconURL = user.ProfileImage
			params.Markdown = true
			params.EscapeText = true

			color := msgNotAnswered
			if question.IsAnswered {
				color = msgIsAnswewed
			}
			thumb := thumbUp
			if question.Score < 0 {
				thumb = thumbDown
			}
			attachment := slack.Attachment{
				Fallback:  question.Title,
				Title:     question.Title,
				TitleLink: question.ShareLink,
				Color:     color,
				Text: fmt.Sprintf(":pencil: %d :speech_balloon: %d %s %d :eye: %d",
					question.AnswerCount,
					question.CommentCount,
					thumb,
					question.Score,
					question.ViewCount,
				),
				Footer:     "slackoverflow",
				FooterIcon: "https://aframe.io/images/aframe-logo-192.png",
			}
			params.Attachments = []slack.Attachment{attachment}
			channelID, timestamp, err := slackoverflow.Slack.PostMessage(slackoverflow.config.Slack.Channel, "", params)
			if err != nil {
				Error(err.Error())
				return
			}
			// Store the link
			slackQuestion.QID = question.QID
			slackQuestion.Channel = channelID
			slackQuestion.TS = timestamp
			msg, err := slackQuestion.Create()
			if err != nil {
				Error("Slack channel (%s): %s", channelID, msg)
			} else {
				Ok("Slack channel (%s): %s and question posted", channelID, msg)
			}
		} else {
			Debug("Slack: Question %d already exists", question.QID)
		}
	}
	Notice("No more new questions to post")
}

func (sl *cmdSlackQuestions) updateQuestions() {
	Info("Slack: Updating questions.")

	links, count := sqlite3.SlackQuestion{}.GetAll()

	if count == 0 {
		Notice("No questions to update.")
		return
	}
	track := 0
	for _, ql := range links {
		stackQuestion := sqlite3.StackExchangeQuestion{}.Find(ql.QID)
		if stackQuestion.QID == 0 {
			Warning("Could not find question with ID: %d.", ql.QID)
			continue
		}
		track++
		if track <= slackoverflow.config.StackExchange.QuestionsToWatch {

			params := slack.UpdateMessageParameters{}
			params.AsUser = false
			params.Timestamp = ql.TS
			color := msgNotAnswered
			if stackQuestion.IsAnswered {
				color = msgIsAnswewed
			}
			thumb := thumbUp
			if stackQuestion.Score < 0 {
				thumb = thumbDown
			}
			attachment := slack.Attachment{
				Fallback:  stackQuestion.Title,
				Title:     stackQuestion.Title,
				TitleLink: stackQuestion.ShareLink,
				Color:     color,
				Text: fmt.Sprintf(":pencil: %d :speech_balloon: %d %s %d :eye: %d",
					stackQuestion.AnswerCount,
					stackQuestion.CommentCount,
					thumb,
					stackQuestion.Score,
					stackQuestion.ViewCount,
				),
				Footer:     "slackoverflow",
				FooterIcon: "https://aframe.io/images/aframe-logo-192.png",
			}

			params.Attachments = []slack.Attachment{attachment}
			channelID, _, _, err := slackoverflow.Slack.UpdateMessageWithAttachments(slackoverflow.config.Slack.Channel, params)
			if err != nil {
				Error("Slack channel (%s): %s", channelID, err.Error())
			} else {
				Ok("Slack channel (%s) updated: %s", channelID, stackQuestion.Title)
			}

		} else {
			params := slack.UpdateMessageParameters{}
			params.AsUser = false
			params.Timestamp = ql.TS
			color := msgNotAnswered
			if stackQuestion.IsAnswered {
				color = msgIsAnswewed
			}
			attachment := slack.Attachment{
				Fallback:  stackQuestion.Title,
				Title:     stackQuestion.Title,
				TitleLink: stackQuestion.ShareLink,
				Color:     color,
			}
			params.Attachments = []slack.Attachment{attachment}
			channelID, _, _, err := slackoverflow.Slack.UpdateMessageWithAttachments(slackoverflow.config.Slack.Channel, params)
			stackQuestion.Delete()
			ql.Delete()
			if err != nil {
				Error("Slack channel (%s): %s", channelID, err.Error())
			} else {
				Ok("Quesstion: not tracking anymore. %s", stackQuestion.Title)
			}
		}
	}
}
