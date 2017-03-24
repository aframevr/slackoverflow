package slackoverflow

import (
	"fmt"

	"github.com/aframevr/slackoverflow/std"
)

type cmdSlackChannels struct{}

// Execute
func (slack *cmdSlackChannels) Execute(args []string) error {

	// Refresh the session before running this command and make sure that Slack Overflow is configured
	slackoverflow.SessionRefresh(true)

	channels, err := slackoverflow.Slack.GetChannels(true)
	if err != nil {
		std.Msg("Unable to fetch any channels")
		return nil
	}

	if len(channels) > 0 {
		channelList := std.NewTable("ID", "Name", "Created")
		for _, channel := range channels {
			channelList.AddRow(
				channel.ID,
				channel.Name,
				fmt.Sprintf("%d", channel.Created),
			)
		}
		channelList.Print()
	}
	return nil
}
