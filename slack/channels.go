package slack

import (
	"fmt"

	"github.com/aframevr/slackoverflow/std"
)

// ListChannels lists available Slack channels
func (s *Client) ListChannels() bool {

	qp := make(std.QueryParams)
	qp.Add("token", s.token)
	qp.Add("exclude_archived", "true")

	channels := &ChannelList{}

	std.HTTPGet(s.apiHost, "channels.list", qp, &channels)

	if len(channels.Channels) > 0 {
		channelList := std.NewTable("ID", "Name", "Created", "Members")
		for _, channel := range channels.Channels {
			channelList.AddRow(
				channel.ID,
				channel.Name,
				fmt.Sprintf("%d", channel.Created),
				fmt.Sprintf("%d", channel.NumMembers),
			)
		}
		channelList.Print()
	} else {
		std.Msg("Unable to fetch any channels")
		return false
	}
	return true
}

// ChannelList list all active channels
type ChannelList struct {
	Channels []ChannelLimited `json:"channels"`
}

// ChannelLimited slack channel
// https://api.slack.com/methods/channels.list
type ChannelLimited struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Created    int    `json:"created"`
	NumMembers int    `json:"num_members"`
}
