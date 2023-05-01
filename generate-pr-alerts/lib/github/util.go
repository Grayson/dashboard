package github

import "net/url"

const (
	base  = "https://api.github.com"
	org   = "orgs"
	repos = "repos"
)

func OrganizationInfoUrl(orgName string) (*url.URL, error) {
	return join(base, org, orgName)
}

func OrganizationReposUrl(orgName string) (*url.URL, error) {
	return join(base, org, orgName, repos)
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
