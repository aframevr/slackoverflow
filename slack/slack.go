package slack

// Client Slack API Client resource
type Client struct {
	channel string
	token   string
	apiHost string
}

// Load slack
func Load() *Client {
	return &Client{}
}

// SetToken set Slack API Token
func (s *Client) SetToken(token string) {
	s.token = token
}

// SetHost set Slack API host
func (s *Client) SetHost(host string) {
	s.apiHost = host
}

// SetChannel where questions will be posted
func (s *Client) SetChannel(channel string) {
	s.channel = channel
}
