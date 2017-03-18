package stackexchange

// Client Stack Exchange API Client resource
type Client struct {
	apiHost    string
	apiVersion string
}

// Load Stack Exchange client
func Load() *Client {
	return &Client{}
}

// SetHost set Stack Exchange API host
func (s *Client) SetHost(host string) {
	s.apiHost = host
}

// SetAPIVersion set Stack Exchange API version
func (s *Client) SetAPIVersion(varsion string) {
	s.apiVersion = varsion
}
