package app

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
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

	client := github.NewClient(http.DefaultClient, config.Token)

	if err := processUserRepos(config.Repos, client, target); err != nil {
		fmt.Println(err)
		return Failure
	}

	for idx, length := 0, len(config.Orgs); idx < length; idx++ {
		if err := fetchOrgPulls(config.Orgs[idx], client); err != nil {
			fmt.Println(err)
			return Failure
		}
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
		if err := fetchRepoPulls(repo, gh); err != nil {
			return err
		}
	}
	return nil
}

func fetchOrgPulls(orgName string, gh github.GitHub) error {
	url, err := github.OrganizationInfoUrl(orgName)
	if err != nil {
		return err
	}

	org, err := gh.OrganizationInfo(url)
	if err != nil {
		return err
	}

	fmt.Printf("# %v\n", org.Login)
	reposUrl, _ := url.Parse(org.ReposUrl)
	repos, err := gh.OrganizationRepos(reposUrl)
	if err != nil {
		return err
	}

	for idx, length := 0, len(repos); idx < length; idx++ {
		repo := &repos[idx]
		if err := logRepoPulls(url, repo, gh); err != nil {
			return err
		}

		if err := logRepoIssues(url, repo, gh); err != nil {
			return err
		}
	}
	fmt.Println()

	return nil
}

func logRepoPulls(url *url.URL, repo *github.OrganizationRepoInfo, gh github.GitHub) error {
	pullUrl, err := url.Parse(github.CleanupPullsUrl(repo.PullsUrl))
	if err != nil {
		return err
	}
	if err := printPulls(gh, pullUrl, fmt.Sprintf("[%v](%v)", repo.Name, repo.HtmlUrl)); err != nil {
		return err
	}
	fmt.Println()
	return nil
}

func logRepoIssues(url *url.URL, repo *github.OrganizationRepoInfo, gh github.GitHub) error {
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

	fmt.Printf("%v issues for [%v](%v)\n", length, repo.Name, repo.HtmlUrl)
	for idx := 0; idx < length; idx++ {
		printIssue(&issues[idx])
	}
	fmt.Println()

	return nil
}

func fetchRepoPulls(repoInfo string, gh github.GitHub) error {
	split := strings.Split(repoInfo, "/")
	if len(split) != 2 {
		return fmt.Errorf("unexpected repo format '%v' (expected 'username/reponame')", repoInfo)
	}

	url, err := github.PullsUrl(split[0], split[1])
	if err != nil {
		return err
	}

	return printPulls(gh, url, repoInfo)
}

func printPulls(gh github.GitHub, url *url.URL, repoInfo string) error {
	pulls, err := gh.Pulls(url)
	if err != nil {
		return err
	}

	pullLength := len(pulls)
	if pullLength == 0 {
		return nil
	}
	fmt.Printf("%v pulls for %v\n", pullLength, repoInfo)

	for idx := 0; idx < pullLength; idx++ {
		printPull(&pulls[idx])
	}

	fmt.Println()

	return nil
}

func printPull(pull *github.Pull) {
	number := path.Base(pull.HtmlUrl)
	fmt.Printf("* Pull: \"%v\" from %v created at %v [#%v](%v)\n", pull.Title, pull.User.Login, pull.CreatedAt, number, pull.HtmlUrl)
}

func printIssue(issue *github.IssuesInfo) {
	fmt.Printf("* Issue: \"%v\" from %v created at %v [#%v](%v)\n", issue.Title, issue.User.Login, issue.CreatedAt, issue.Number, issue.HtmlUrl)
}
