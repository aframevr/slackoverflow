package std

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// HTTPGetByURL make Get request
func HTTPGetByURL(url string, collection interface{}) error {

	client := &http.Client{
		Transport: http.DefaultTransport,
	}

	response, err := client.Get(url)

	if err != nil {
		return err
	}

	err = readResponse(response, &collection)

	return err
}

// HTTPGet make Get request
func HTTPGet(host string, path string, qp map[string]string, collection interface{}) error {

	baseURL, _ := url.Parse(host)
	endpoint, _ := baseURL.Parse(baseURL.RequestURI() + "/" + path)
	query := endpoint.Query()

	for key, value := range qp {
		query.Set(key, value)
	}

	endpoint.RawQuery = query.Encode()

	client := &http.Client{
		Transport: http.DefaultTransport,
	}

	response, err := client.Get(endpoint.String())

	if err != nil {
		return err
	}

	err = readResponse(response, &collection)

	return err
}

// ReadResponse the response from http.Response
func readResponse(response *http.Response, result interface{}) error {
	// close the body when done reading
	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, result)

	if err != nil {
		return err
	}

	if response.StatusCode == 400 {
		return fmt.Errorf("Bad Request: %s", string(bytes))
	}

	return nil
}
