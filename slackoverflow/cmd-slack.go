package slackoverflow

// slackoverflow slack
type cmdSlack struct {
	Channels cmdSlackChannels `command:"channels" description:"This method returns a list of all Slack channels in the team."`
}

type cmdSlackChannels struct{}

// Execute
func (slack *cmdSlackChannels) Execute(args []string) error {
	// Refresh the session before running this command
	slackoverflow.Slack.ListChannels()
	return nil
}
