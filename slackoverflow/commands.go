package slackoverflow

// Commands - Register Slackoverflow main commands
type Commands struct {
	// verbosity
	Verbose bool `short:"v" long:"verbose" description:"Be more verbose, This enable loglevel Info"`
	// Debug
	Debug bool `short:"d" long:"debug" description:"Be even more verbose, This enables loglevel Debug"`

	// Config slackoverflow config
	Config cmdConfig `command:"config" description:"Display Slackoverflow configuration."`

	// Credits slackoverflow credits
	Credits cmdCredits `command:"credits" description:"List of Slackoverflow contributors."`

	// Reconfigure slackoverflow reconfigure
	Reconfigure cmdReconfigure `command:"reconfigure" alias:"init" description:"Interactive configuration of stackoverflow"`

	// Restart slackoverflow restart
	Restart cmdRestart `command:"restart" description:"Restart Slackoverflow daemon."`

	// Run slackoverflow run
	Run cmdRun `command:"run" description:"Run Slackoverflow once."`

	// Start slackoverflow start
	Slack cmdSlack `command:"slack" description:"Slack related commands see slackoverflow slack --help for more info."`

	// Start slackoverflow start
	Start cmdStart `command:"start" description:"Start Slackoverflow daemon."`

	// Status slackoverflow status
	Status cmdStatus `command:"status" description:"Get Slackoverflow daemon status."`

	// Stop slackoverflow run
	Stop cmdStop `command:"stop" description:"Stop Slackoverflow daemon."`

	// Validate slackoverflow validate
	Validate cmdValidate `command:"validate" description:"validate stackoverflow configuration"`

	// Watch slackoverflow watch
	Watch cmdWatch `command:"watch" description:"Start Slackoverflow daemon in foreground."`
}
