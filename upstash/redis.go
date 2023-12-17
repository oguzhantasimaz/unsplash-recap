package upstash

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"unsplash-recap/unsplash"
)

type Client struct {
	Token   string
	BaseURL string
}

type Response struct {
	Result *string `json:"result"`
}

// NewClient creates a new unsplash api client with the given api key
func NewClient(url, token string) *Client {
	return &Client{
		Token:   token,
		BaseURL: url,
	}
}

func (c *Client) Get(username string) (*string, error) {
	url := fmt.Sprintf("%s/get/%s", c.BaseURL, username)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("user not found")
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}

	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var r *Response
	err = json.Unmarshal(body, &r)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if r.Result == nil {
		return nil, nil
	}

	return r.Result, nil
}

func (c *Client) Set(username string, value *unsplash.Recap) error {
	url := fmt.Sprintf("%s/set/%s", c.BaseURL, username)

	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(*value)
	if err != nil {
		log.Error(err)
		return err
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Error(err)
		return err
	}

	req.Header.Add("Authorization", "Bearer "+c.Token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}

	return nil
}
