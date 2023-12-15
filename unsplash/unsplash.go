package unsplash

import (
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

const (
	baseURL = "https://api.unsplash.com"
)

type Client struct {
	APIKey  string
	BaseURL string
}

// NewClient creates a new unsplash api client with the given api key
func NewClient(apiKey string) *Client {
	return &Client{
		APIKey:  apiKey,
		BaseURL: baseURL,
	}
}

func (c *Client) Get(url string, queryParameters map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	req.Header.Add("Authorization", "Client-ID "+c.APIKey)

	q := req.URL.Query()
	for k, v := range queryParameters {
		q.Add(k, v)
	}

	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return body, nil
}
