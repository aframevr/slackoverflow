package slack

import (
	"errors"
	"net/url"
)

type channelResponseFull struct {
	Channel      Channel   `json:"channel"`
	Channels     []Channel `json:"channels"`
	Purpose      string    `json:"purpose"`
	Topic        string    `json:"topic"`
	NotInChannel bool      `json:"not_in_channel"`
	APIresponse
}

// Channel contains information about the channel
type Channel struct {
	groupConversation
	IsChannel bool `json:"is_channel"`
	IsGeneral bool `json:"is_general"`
	IsMember  bool `json:"is_member"`
}

// GetChannels retrieves all the channels
func (client *Client) GetChannels(excludeArchived bool) ([]Channel, error) {
	values := url.Values{
		"token": {client.APIToken},
	}

	if excludeArchived {
		values.Add("exclude_archived", "1")
	}
	response, err := channelRequest("channels.list", values)
	if err != nil {
		return nil, err
	}
	return response.Channels, nil
}

// Conversation is the foundation for IM and BaseGroupConversation
type conversation struct {
	ID                 string   `json:"id"`
	Created            JSONTime `json:"created"`
	IsOpen             bool     `json:"is_open"`
	LastRead           string   `json:"last_read,omitempty"`
	UnreadCount        int      `json:"unread_count,omitempty"`
	UnreadCountDisplay int      `json:"unread_count_display,omitempty"`
}

// GroupConversation is the foundation for Group and Channel
type groupConversation struct {
	conversation
	Name       string   `json:"name"`
	Creator    string   `json:"creator"`
	IsArchived bool     `json:"is_archived"`
	Members    []string `json:"members"`
	Topic      Topic    `json:"topic"`
	Purpose    Purpose  `json:"purpose"`
}

// Topic contains information about the topic
type Topic struct {
	Value   string   `json:"value"`
	Creator string   `json:"creator"`
	LastSet JSONTime `json:"last_set"`
}

// Purpose contains information about the purpose
type Purpose struct {
	Value   string   `json:"value"`
	Creator string   `json:"creator"`
	LastSet JSONTime `json:"last_set"`
}

func channelRequest(path string, values url.Values) (*channelResponseFull, error) {
	response := &channelResponseFull{}
	err := post(path, values, response)
	if err != nil {
		return nil, err
	}
	if !response.Ok {
		return nil, errors.New(response.Error)
	}
	return response, nil
}
