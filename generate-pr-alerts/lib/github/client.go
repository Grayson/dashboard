package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"sync/atomic"
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

func (c *Client) Issues(url *url.URL) ([]IssuesInfo, error) {
	// TODO: Should this be paged?
	w := wrapper[[]IssuesInfo]{}
	return w.get(url, c.personalAccessToken, *c.client)
}

func (c *Client) Pulls(url *url.URL) ([]Pull, error) {
	w := wrapper[[]Pull]{}
	return w.get(url, c.personalAccessToken, *c.client)
}

func (c *Client) OrganizationInfo(url *url.URL) (OrganizationInfo, error) {
	w := wrapper[OrganizationInfo]{}
	return w.get(url, c.personalAccessToken, *c.client)
}

func (c *Client) OrganizationRepos(u *url.URL) ([]OrganizationRepoInfo, error) {
	const (
		pageLimit       = 32
		defaultPageSize = 30
	)

	urls := make(chan *url.URL, 1)

	// Feed urls into the pipeline
	go func() {
		for page := 0; page < pageLimit; page++ {
			url := *u
			query := url.Query()
			query.Set("page", strconv.Itoa(page))
			url.RawQuery = query.Encode()
			fmt.Printf("Sending %v\n", url.String())
			urls <- &url
		}
		fmt.Println("Closing urls")
		close(urls)
	}()

	type x struct {
		more []OrganizationRepoInfo
		err  error
	}
	ch := make(chan x)

	w := wrapper[[]OrganizationRepoInfo]{}

	wg := sync.WaitGroup{}
	stop := atomic.Bool{}
	for worker := 0; worker < 4; worker++ {
		wg.Add(1)
		go func(worker int) {
			defer func() {
				fmt.Printf("Done %v\n", worker)
				wg.Done()
			}()
			for url := range urls {
				fmt.Printf("Loading %v on worker %v\n", url, worker)
				if stop.Load() {
					fmt.Printf("Stopping worker %v\n", worker)
					return
				}
				m, e := w.get(url, c.personalAccessToken, *c.client)
				fmt.Println("Sending")
				ch <- x{m, e}
			}
		}(worker)
	}
	go func() {
		fmt.Println("Waiting")
		wg.Wait()
		fmt.Println("Closing")
		close(ch)
	}()

	info := make([]OrganizationRepoInfo, 0)
	var err error
	for x := range ch {
		if x.err != nil {
			err = x.err
		} else {
			info = append(info, x.more...)
		}

		if err != nil || (len(x.more) < defaultPageSize) {
			stop.Store(true)
		}
	}
	fmt.Println("Returning")
	if err != nil {
		return nil, err
	}

	return info, nil
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
