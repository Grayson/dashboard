package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"sync"
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
		pageLimit   = 32
		workerLimit = 4
	)

	urls := make(chan *url.URL, 4)
	ch := make(chan workerResult)
	page := 0
	wg := sync.WaitGroup{}

	for worker := 0; worker < workerLimit; worker++ {
		urls <- makeUrl(*u, page)
		page++
		wg.Add(1)
		go fetchUrlWork(&wg, urls, c, ch)
	}
	go func() {
		wg.Wait()
		close(urls)
		close(ch)
	}()

	return waitForAndMapResults(ch, urls, u, page)
}

func waitForAndMapResults(ch chan workerResult, urls chan *url.URL, u *url.URL, page int) ([]OrganizationRepoInfo, error) {
	const defaultPageSize = 30

	info := make([]OrganizationRepoInfo, 0)
	shouldContinue := true

	for x := range ch {
		if x.err != nil {
			return nil, x.err
		} else {
			info = append(info, x.more...)
		}

		shouldContinue = shouldContinue && x.err != nil && (len(x.more) == defaultPageSize)
		if shouldContinue {
			urls <- makeUrl(*u, page)
			page++
		} else {
			urls <- nil
		}
	}
	return info, nil
}

func makeUrl(url url.URL, page int) *url.URL {
	query := url.Query()
	query.Set("page", strconv.Itoa(page))
	url.RawQuery = query.Encode()
	return &url
}

type workerResult struct {
	more []OrganizationRepoInfo
	err  error
}

func fetchUrlWork(wg *sync.WaitGroup, urls chan *url.URL, c *Client, ch chan workerResult) {
	defer func() { wg.Done() }()
	w := wrapper[[]OrganizationRepoInfo]{}
	for url := range urls {
		if url == nil {
			return
		}
		m, e := w.get(url, c.personalAccessToken, *c.client)
		ch <- workerResult{m, e}
	}
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
