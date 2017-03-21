package stackexchange

import "net/url"

var stackexchange *Client

// Client Stack Exchange API Client resource
type Client struct {
	apiHost        string
	apiVersion     string
	quotaMax       int
	quotaRemaining int
	apiKey         string
}

// Load Stack Exchange client
func Load() *Client {
	stackexchange = &Client{}
	return stackexchange
}

// SetHost set Stack Exchange API host
func (s *Client) SetHost(host string) {
	s.apiHost = host
}

// SetAPIVersion set Stack Exchange API version
func (s *Client) SetAPIVersion(varsion string) {
	s.apiVersion = varsion
}

// SetAPIKey Set Stack Exchange API to receive a higher request quota if exists
func (s *Client) SetAPIKey(apiKey string) {
	s.apiKey = apiKey
}

// SetQuotaRemaining set remaining quota for today
func (s *Client) SetQuotaRemaining(quotaRemaining int) {
	s.quotaRemaining = quotaRemaining
}

// GetQuotaRemaining return remaining quota for today
func (s *Client) GetQuotaRemaining() int {
	return s.quotaRemaining
}

// SetQuotaMax set maximum allowed quota
func (s *Client) SetQuotaMax(quotaMax int) {
	s.quotaMax = quotaMax
}

// GetQuotaMax return maximum allowed quota
func (s *Client) GetQuotaMax() int {
	return s.quotaMax
}

// GetAPIKey return active StackExchange API Key if one is set
func (s *Client) GetAPIKey() string {
	return s.apiKey
}

// GetEndpont returns base endpoint
func (s *Client) GetEndpont(path string) (*url.URL, error) {
	endpoint, err := url.Parse(s.apiHost + "/" + s.apiVersion + "/" + path)
	return endpoint, err
}

// SearchAdvanced https://api.stackexchange.com/docs/advanced-search
func (s *Client) SearchAdvanced() *SearchAdvanced {
	searchAdvanced := &SearchAdvanced{}
	searchAdvanced.Init()
	return searchAdvanced
}

// Questions https://api.stackexchange.com/docs/questions-by-ids
func (s *Client) Questions() *Questions {
	questions := &Questions{}
	questions.Init()
	return questions
}
