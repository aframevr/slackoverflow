package slackoverflow

// Commands - Register SlackOverflow main commands
type Commands struct {
	// verbosity
	Verbose bool `short:"v" long:"verbose" description:"Be more verbose, This enable loglevel Info"`
	// Debug
	Debug bool `short:"d" long:"debug" description:"Be even more verbose, This enables loglevel Debug"`

	// Config slackoverflow config
	Config cmdConfig `command:"config" description:"Display SlackOverflow configuration."`

	// Credits slackoverflow credits
	Credits cmdCredits `command:"credits" description:"List of SlackOverflow contributors."`

	// Reconfigure slackoverflow reconfigure
	Reconfigure cmdReconfigure `command:"reconfigure" alias:"init" description:"Interactive configuration of stackoverflow"`

	// Restart slackoverflow restart
	Restart cmdRestart `command:"restart" description:"Restart SlackOverflow daemon."`

	// Run slackoverflow run
	Run cmdRun `command:"run" description:"Run SlackOverflow once."`

	// Slack slackoverflow slack
	Slack cmdSlack `command:"slack" description:"Slack related commands see slackoverflow slack --help for more info."`

	// Stack Exchange slackoverflow stackexchange
	StackExchange cmdStackExchange `command:"stackexchange" description:"Stack Exchange related commands see slackoverflow stackexchange --help for more info."`

	// Start slackoverflow start
	Start cmdStart `command:"start" description:"Start SlackOverflow daemon."`

	// Status slackoverflow status
	Status cmdStatus `command:"status" description:"Get SlackOverflow daemon status."`

	// Stop slackoverflow run
	Stop cmdStop `command:"stop" description:"Stop SlackOverflow daemon."`

	// Validate slackoverflow validate
	Validate cmdValidate `command:"validate" description:"validate stackoverflow configuration"`
}
