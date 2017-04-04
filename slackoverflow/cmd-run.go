package slackoverflow

// slackoverflow run
// Execute slackoverflow once
type cmdRun struct{}

// Execute
func (a *cmdRun) Execute(args []string) error {

	// Refresh the session before running this command
	slackoverflow.SessionRefresh(true)
	stackechangeSync := cmdStackExchangeQuestions{}
	stackechangeSync.Sync = true
	stackechangeSync.Execute(args)

	slackSync := cmdSlackQuestions{}
	slackSync.All = true
	slackSync.Execute(args)

	return nil
}
