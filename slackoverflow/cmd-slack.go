package slackoverflow

// slackoverflow slack
type cmdSlack struct {
	Channels  cmdSlackChannels  `command:"channels" description:"This method returns a list of all Slack channels in the team."`
	Questions cmdSlackQuestions `command:"questions" description:"Post new or update tracked Stack Exchange questions on Slack channel"`
}
