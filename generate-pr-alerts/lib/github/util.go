package github

import (
	"net/url"
	"strings"
)

const (
	base  = "https://api.github.com"
	org   = "orgs"
	pulls = "pulls"
	repos = "repos"
)

func PullsUrl(user string, repo string) (*url.URL, error) {
	return join(base, repos, user, repo, pulls)
}

func OrganizationInfoUrl(orgName string) (*url.URL, error) {
	return join(base, org, orgName)
}

func OrganizationReposUrl(orgName string) (*url.URL, error) {
	return join(base, org, orgName, repos)
}

func CleanupPullsUrl(urlString string) string {
	return strings.Replace(urlString, "{/number}", "", 1)
}

func join(first string, rest ...string) (*url.URL, error) {
	result, err := url.JoinPath(first, rest...)
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(result)
	if err != nil {
		return nil, err
	}
	return u, nil
}
