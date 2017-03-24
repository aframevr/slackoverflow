package slackoverflow

import (
	"strconv"

	"github.com/aframevr/slackoverflow/std"
)

// slackoverflow config
// Display configuration information
type cmdConfig struct{}

func (a *cmdConfig) Execute(args []string) error {

	// Refresh the session before running this command and make sure that Slack Overflow is configured
	slackoverflow.SessionRefresh(false)

	selfConfig := std.NewTable("SlackOverflow Configuration", " ")
	selfConfig.AddRow("Log Level", slackoverflow.config.SlackOverflow.LogLevel)
	selfConfig.AddRow("Number of Questions to watch", strconv.Itoa(slackoverflow.config.SlackOverflow.Watch))
	selfConfig.Print()

	si := std.NewTable("Session Info", " ")
	si.AddRow("Project path", slackoverflow.projectPath)
	si.AddRow("Config file", slackoverflow.configFile)
	si.AddRow("Database file", slackoverflow.databaseFile)
	si.AddRow("Current working dir", slackoverflow.cwd)
	si.AddRow("Hostname", slackoverflow.hostname)
	si.AddRow("Userame", slackoverflow.user.Username)
	si.AddRow("Name", slackoverflow.user.Name)
	si.AddRow("User ID", slackoverflow.user.Uid)
	si.AddRow("Group", slackoverflow.group.Name)
	si.AddRow("Group ID", slackoverflow.user.Gid)
	si.Print()

	slack := std.NewTable("Slack Configuration", " ")
	slack.AddRow("API host", slackoverflow.config.Slack.APIHost)
	slack.AddRow("Token", slackoverflow.config.Slack.Token)
	slack.AddRow("Channel", slackoverflow.config.Slack.Channel)
	slack.Print()

	stackexchange := std.NewTable("StackExchange Configuration", " ")

	stackexchange.AddRow("API Host", slackoverflow.config.StackExchange.APIHost)
	stackexchange.AddRow("API Version", slackoverflow.config.StackExchange.APIVersion)

	stackexchange.AddRow("Key", slackoverflow.config.StackExchange.Key)

	stackexchange.AddRow("Site", slackoverflow.config.StackExchange.Site)
	stackexchange.AddRow("Tagged", slackoverflow.config.StackExchange.SearchAdvanced["tagged"])
	stackexchange.Print()

	return nil
}
