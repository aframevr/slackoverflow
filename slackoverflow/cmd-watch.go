package slackoverflow

// slackoverflow watch
// SlackOverflow .
type cmdWatch struct{}

func (a *cmdWatch) Execute(args []string) error {

	// Refresh the session before running this command and make sure that Slack Overflow is configured
	slackoverflow.SessionRefresh()

	return nil
}
