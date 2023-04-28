package github

import (
	"fmt"
	"net/url"
)

type User struct {
	AvatarUrl  string `json:"avatar_url"`
	GravatarId string `json:"gravatar_id"`
	HtmlUrl    string `json:"html_url"`
	Login      string `json:"login"`
	Url        string `json:"url"`
}

type Label struct {
	Color       string `json:"color"`
	Description string `json:"description"`
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Url         string `json:"url"`
}

type Pull struct {
	Url                string  `json:"url"`
	Id                 int     `json:"id"`
	HtmlUrl            string  `json:"html_url"`
	State              string  `json:"state"`
	Title              string  `json:"title"`
	User               User    `json:"user"`
	Body               string  `json:"body"`
	Labels             []Label `json:"labels"`
	CreatedAt          string  `json:"created_at"`
	UpdatedAt          string  `json:"updated_at"`
	ClosedAt           string  `json:"closed_at"`
	MergedAt           string  `json:"merged_at"`
	Assignee           User    `json:"assignee"`
	Assignees          []User  `json:"assignees"`
	RequestedReviewers []User  `json:"requested_reviewers"`
	AuthorAssociation  string  `json:"author_association"`
	IsDraft            bool    `json:"draft"`
}

type GithubErrorResponse struct {
	Message          string `json:"message"`
	DocumentationURL string `json:"documentation_url"`
}

type GitHub interface {
	Pulls(url *url.URL) ([]Pull, error)
}

func (e GithubErrorResponse) Error() string {
	return fmt.Sprintf("Github error: %s [%s]", e.Message, e.DocumentationURL)
}
