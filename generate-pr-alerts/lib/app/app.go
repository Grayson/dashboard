package app

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Grayson/dashboard/generate-pr-alerts/lib/github"
)

type RunResult int

const (
	Success RunResult = iota
	Failure
)

func Run(config *Config, target output.Target) RunResult {
	if config.Token == "" {
		fmt.Println("No GitHub token provided")
		return Failure
	}

	gh := github.NewClient(http.DefaultClient, config.Token)

	if err := processUserRepos(config.Repos, gh, target); err != nil {
		fmt.Println(err)
		return Failure
	}

	for idx, length := 0, len(config.Orgs); idx < length; idx++ {
		target.StartOrganizationsPhase(config.Orgs)
		if err := fetchOrgPulls(config.Orgs[idx], gh, target); err != nil {
			fmt.Println(err)
			return Failure
		}
		target.EndOrganizationsPhase()
	}

	return Success
}

func processUserRepos(repos []string, gh github.GitHub, target output.Target) error {
	repoLength := len(repos)
	if repoLength != 0 {
		if err := target.StartReposPhase(); err != nil {
			return err
		}
	}

	for _, repo := range repos {
		if err := fetchRepoPulls(repo, gh, target); err != nil {
			return err
		}
	}
	return nil
}

func fetchOrgPulls(orgName string, gh github.GitHub, target output.Target) error {
	url, err := github.OrganizationInfoUrl(orgName)
	if err != nil {
		return err
	}

	org, err := gh.OrganizationInfo(url)
	if err != nil {
		return err
	}

	target.VisitOrganization(org)
	reposUrl, _ := url.Parse(org.ReposUrl)
	repos, err := gh.OrganizationRepos(reposUrl)
	if err != nil {
		return err
	}

	for idx, length := 0, len(repos); idx < length; idx++ {
		repo := &repos[idx]
		target.StartRepo(repo)
		if err := logRepoPulls(url, repo, gh, target); err != nil {
			return err
		}

		if err := logRepoIssues(url, repo, gh, target); err != nil {
			return err
		}
		target.EndRepo()
	}
	fmt.Println()

	return nil
}

func logRepoPulls(url *url.URL, repo *github.OrganizationRepoInfo, gh github.GitHub, target output.Target) error {
	pullUrl, err := url.Parse(github.CleanupPullsUrl(repo.PullsUrl))
	if err != nil {
		return err
	}
	if err := printPulls(gh, pullUrl, repo, target); err != nil {
		return err
	}
	// TODO: Evaluate this Println
	fmt.Println()
	return nil
}

func logRepoIssues(url *url.URL, repo *github.OrganizationRepoInfo, gh github.GitHub, target output.Target) error {
	url, err := url.Parse(github.CleanupIssuesUrl(repo.IssuesUrl))
	if err != nil {
		return err
	}

	issues, err := gh.Issues(url)
	if err != nil {
		return err
	}

	length := len(issues)
	if length == 0 {
		return nil
	}

	target.StartIssues(issues)
	for idx := 0; idx < length; idx++ {
		target.VisitIssue(issues[idx])
	}
	target.EndIssues()

	return nil
}

func fetchRepoPulls(repoInfo string, gh github.GitHub, target output.Target) error {
	// TODO: Fetch repo information and add StartRepo call

	split := strings.Split(repoInfo, "/")
	if len(split) != 2 {
		return fmt.Errorf("unexpected repo format '%v' (expected 'username/reponame')", repoInfo)
	}

	url, err := github.PullsUrl(split[0], split[1])
	if err != nil {
		return err
	}

	return printPulls(gh, url, repoInfo, target)
}

func printPulls(gh github.GitHub, url *url.URL, repoInfo string, target output.Target) error {
	pulls, err := gh.Pulls(url)
	if err != nil {
		return err
	}

	pullLength := len(pulls)
	if pullLength == 0 {
		return nil
	}

	target.StartPulls(pulls)
	for idx := 0; idx < pullLength; idx++ {
		target.VisitPull(pulls[idx])
	}
	target.EndPulls()

	return nil
}
