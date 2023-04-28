package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	client              *http.Client
	personalAccessToken string
}

func NewClient(client *http.Client, personalAccessToken string) *Client {
	return &Client{
		client,
		personalAccessToken,
	}
}

func (c *Client) Pulls(url *url.URL) ([]Pull, error) {
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return []Pull{}, err
	}

	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", fmt.Sprintf("token %s", c.personalAccessToken))

	resp, err := c.client.Do(req)
	if err != nil {
		return []Pull{}, err
	}

	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Pull{}, err
	}

	var errorResponse GithubErrorResponse
	json.Unmarshal(bytes, &errorResponse)
	if errorResponse.Message != "" {
		return []Pull{}, errorResponse
	}

	var pullResponse []Pull
	if err := json.Unmarshal(bytes, &pullResponse); err != nil {
		return []Pull{}, err
	}

	return pullResponse, nil
}
