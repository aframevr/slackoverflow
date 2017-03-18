package slack

// Slack packages
type Slack struct {
	channel string
	token   string
	apiHost string
}

// Load slack
func Load() *Slack {
	return &Slack{}
}

// SetToken set Slack API Token
func (s *Slack) SetToken(token string) {
	s.token = token
}

// SetHost set Slack API host
func (s *Slack) SetHost(host string) {
	s.apiHost = host
}

// SetChannel where questions will be posted
func (s *Slack) SetChannel(channel string) {
	s.channel = channel
}
