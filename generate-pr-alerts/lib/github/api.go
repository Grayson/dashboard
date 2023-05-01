package github

import (
	"fmt"
	"net/url"
)

type GitHub interface {
	OrganizationInfo(url *url.URL) (OrganizationInfo, error)
	OrganizationRepos(url *url.URL) ([]OrganizationRepoInfo, error)
	Pulls(url *url.URL) ([]Pull, error)
}

func (e GithubErrorResponse) Error() string {
	return fmt.Sprintf("Github error: %s [%s]", e.Message, e.DocumentationURL)
}
