package slackoverflow

// slackoverflow stackexchange
type cmdService struct {

	// Install slackoverflow service install
	Install cmdServiceInstall `command:"install" description:"Install SlackOverflow service."`

	// Remove slackoverflow service install
	Remove cmdServiceRemove `command:"remove" description:"Remove SlackOverflow service."`

	// Restart slackoverflow service restart
	Restart cmdServiceRestart `command:"restart" description:"Restart SlackOverflow service."`

	// Start slackoverflow service start
	Start cmdServiceStart `command:"start" description:"Start SlackOverflow service."`

	// Status slackoverflow service status
	Status cmdServiceStatus `command:"status" description:"Get SlackOverflow service status."`

	// Stop slackoverflow service stop
	Stop cmdServiceStop `command:"stop" description:"Stop SlackOverflow service."`
}

// slackoverflow install
// Install SlackOverflow daemon.
type cmdServiceInstall struct{}

func (a *cmdServiceInstall) Execute(args []string) error {

	// Refresh the session before running this command and make sure that Slack Overflow is configured
	slackoverflow.SessionRefresh(true)
	status, err := slackoverflow.Service.Install()
	if err != nil {
		Notice(status)
		Error("%v", err)
		slackoverflow.Close(1)
	} else {
		Notice(status)
	}
	return nil
}

// slackoverflow remove
// Remove SlackOverflow daemon.
type cmdServiceRemove struct{}

func (a *cmdServiceRemove) Execute(args []string) error {

	// Refresh the session before running this command and make sure that Slack Overflow is configured
	slackoverflow.SessionRefresh(true)
	status, err := slackoverflow.Service.Remove()
	if err != nil {
		Notice(status)
		Error("%v", err)
		slackoverflow.Close(1)
	} else {
		Notice(status)
	}
	return nil
}

// slackoverflow restart
// Restart SlackOverflow daemon.
type cmdServiceRestart struct{}

func (a *cmdServiceRestart) Execute(args []string) error {

	// Refresh the session before running this command and make sure that Slack Overflow is configured
	stop := cmdServiceStop{}
	stop.IgnoreErrors = true
	start := cmdServiceStart{}
	start.IgnoreErrors = true
	stop.Execute(args)
	start.Execute(args)
	return nil
}

// slackoverflow start
// Start SlackOverflow daemon.
type cmdServiceStart struct {
	IgnoreErrors bool `long:"ignore" description:"Ignore errors" hidden:"true"`
}

func (a *cmdServiceStart) Execute(args []string) error {

	// Refresh the session before running this command and make sure that Slack Overflow is configured
	slackoverflow.SessionRefresh(true)
	status, err := slackoverflow.Service.Start()
	if err != nil && !a.IgnoreErrors {
		Notice(status)
		Error("%v", err)
		slackoverflow.Close(1)
	} else {
		Notice(status)
	}
	return nil
}

// slackoverflow status
// Get SlackOverflow daemon status.
type cmdServiceStatus struct{}

func (a *cmdServiceStatus) Execute(args []string) error {

	// Refresh the session before running this command and make sure that Slack Overflow is configured
	slackoverflow.SessionRefresh(true)
	status, err := slackoverflow.Service.Status()
	if err != nil {
		Notice(status)
		Error("%v", err)
		slackoverflow.Close(1)
	} else {
		Notice(status)
	}
	return err
}

// slackoverflow stop
// Stop SlackOverflow daemon.
type cmdServiceStop struct {
	IgnoreErrors bool `long:"ignore" description:"Ignore errors" hidden:"true"`
}

func (a *cmdServiceStop) Execute(args []string) error {

	// Refresh the session before running this command and make sure that Slack Overflow is configured
	slackoverflow.SessionRefresh(true)
	status, err := slackoverflow.Service.Stop()
	if err != nil && !a.IgnoreErrors {
		Notice(status)
		Error("%v", err)
		slackoverflow.Close(1)
	} else {
		Notice(status)
	}
	return nil
}
