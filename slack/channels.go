package slack

import (
	"fmt"

	"github.com/aframevr/slackoverflow/std"
)

// ListChannels lists available Slack channels
func (s *Slack) ListChannels() bool {

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
	Channels []SlChannel `json:"channels"`
}

// SlChannel slack channel
type SlChannel struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Created    int    `json:"created"`
	NumMembers int    `json:"num_members"`
}
