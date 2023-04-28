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
	w := wrapper[[]Pull]{}
	return w.get(url, c.personalAccessToken, *c.client)
}

func (c *Client) OrganizationInfo(url *url.URL) (OrganizationInfo, error) {
	w := wrapper[OrganizationInfo]{}
	return w.get(url, c.personalAccessToken, *c.client)
}

type wrapper[T any] struct {
	empty T
}

func (w *wrapper[T]) get(url *url.URL, personalAccessToken string, client http.Client) (T, error) {
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return w.empty, err
	}

	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", fmt.Sprintf("token %s", personalAccessToken))

	resp, err := client.Do(req)
	if err != nil {
		return w.empty, err
	}

	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return w.empty, err
	}

	var errorResponse GithubErrorResponse
	json.Unmarshal(bytes, &errorResponse)
	if errorResponse.Message != "" {
		return w.empty, errorResponse
	}

	var response T
	if err := json.Unmarshal(bytes, &response); err != nil {
		return w.empty, err
	}

	return response, nil
}
