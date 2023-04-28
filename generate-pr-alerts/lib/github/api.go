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
	Assignee           User    `json:"assignee"`
	Assignees          []User  `json:"assignees"`
	AuthorAssociation  string  `json:"author_association"`
	Body               string  `json:"body"`
	ClosedAt           string  `json:"closed_at"`
	CreatedAt          string  `json:"created_at"`
	HtmlUrl            string  `json:"html_url"`
	Id                 int     `json:"id"`
	IsDraft            bool    `json:"draft"`
	Labels             []Label `json:"labels"`
	MergedAt           string  `json:"merged_at"`
	RequestedReviewers []User  `json:"requested_reviewers"`
	State              string  `json:"state"`
	Title              string  `json:"title"`
	UpdatedAt          string  `json:"updated_at"`
	Url                string  `json:"url"`
	User               User    `json:"user"`
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
