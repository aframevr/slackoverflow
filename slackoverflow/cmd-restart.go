package slackoverflow

// slackoverflow restart
// Restart SlackOverflow daemon.
type cmdRestart struct{}

func (a *cmdRestart) Execute(args []string) error {

	// Refresh the session before running this command and make sure that Slack Overflow is configured
	slackoverflow.SessionRefresh(false)

	return nil
}
