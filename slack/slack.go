package slack

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

var slack *Client

// HTTPClient http.Client{}
var HTTPClient = &http.Client{}

// Client Stack Exchange API Client resource
type Client struct {
	APIHost        string
	APIToken       string
	CurrentChannel string
}

// Load Stack Exchange client
func Load() *Client {
	slack = &Client{}
	return slack
}

// SetToken Set Slack token
func (sl *Client) SetToken(token string) {
	sl.APIToken = token
}

// SetAPIHost set api host
func (sl *Client) SetAPIHost(host string) {
	sl.APIHost = host
}

// SetChannel set channel
func (sl *Client) SetChannel(channel string) {
	sl.CurrentChannel = channel
}

// JSONTime int64
type JSONTime int64

// APIresponse response message
type APIresponse struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

func postForm(endpoint string, values url.Values, intf interface{}) error {
	resp, err := HTTPClient.PostForm(endpoint, values)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return parseResponseBody(resp.Body, &intf)
}

func post(path string, values url.Values, intf interface{}) error {
	return postForm("https://slack.com/api/"+path, values, intf)
}

func parseResponseBody(body io.ReadCloser, intf *interface{}) error {
	response, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(response, &intf)
	if err != nil {
		return err
	}

	return nil
}
