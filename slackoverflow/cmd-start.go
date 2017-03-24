package slackoverflow

// slackoverflow start
// Start SlackOverflow daemon.
type cmdStart struct{}

func (a *cmdStart) Execute(args []string) error {

	// Refresh the session before running this command and make sure that Slack Overflow is configured
	slackoverflow.SessionRefresh(true)

	return nil
}
