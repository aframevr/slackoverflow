package slackoverflow

import (
	"strconv"

	"github.com/aframevr/slackoverflow/std"
)

// slackoverflow config
// Display configuration information
type cmdConfig struct{}

func (a *cmdConfig) Execute(args []string) error {

	selfConfig := std.NewTable("Slackoverflow Configuration", " ")
	selfConfig.AddRow("Log Level", stackoverflow.config.Slackoverflow.LogLevel)
	selfConfig.AddRow("Number of Questions to watch", strconv.Itoa(stackoverflow.config.Slackoverflow.Watch))
	selfConfig.Print()

	si := std.NewTable("Session Info", " ")
	si.AddRow("Project path", stackoverflow.projectPath)
	si.AddRow("Config file", stackoverflow.configFile)
	si.AddRow("Database file", stackoverflow.databaseFile)
	si.AddRow("Current working dir", stackoverflow.cwd)
	si.AddRow("Hostname", stackoverflow.hostname)
	si.AddRow("Userame", stackoverflow.user.Username)
	si.AddRow("Name", stackoverflow.user.Name)
	si.AddRow("User ID", stackoverflow.user.Uid)
	si.AddRow("Group", stackoverflow.group.Name)
	si.AddRow("Group ID", stackoverflow.user.Gid)
	si.Print()

	slack := std.NewTable("Slack Configuration", " ")
	slack.AddRow("API host", stackoverflow.config.Slack.APIHost)
	slack.AddRow("Token", stackoverflow.config.Slack.Token)
	slack.AddRow("Channel", stackoverflow.config.Slack.Channel)
	slack.Print()

	stackexchange := std.NewTable("StackExchange Configuration", " ")

	stackexchange.AddRow("API Host", stackoverflow.config.StackExchange.APIHost)
	stackexchange.AddRow("API Version", stackoverflow.config.StackExchange.APIVersion)

	stackexchange.AddRow("Key", stackoverflow.config.StackExchange.Key)

	stackexchange.AddRow("Site", stackoverflow.config.StackExchange.Parameters.Site)
	stackexchange.AddRow("Tagged", stackoverflow.config.StackExchange.Parameters.Tagged)
	stackexchange.Print()

	return nil
}
