package app

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Grayson/dashboard/generate-pr-alerts/lib/github"
)

type RunResult int

const (
	Success RunResult = iota
	Failure
)

func Run(config *Config) RunResult {
	if config.Token == "" {
		fmt.Println("No GitHub token provided")
		return Failure
	}

	client := github.NewClient(http.DefaultClient, config.Token)
	for _, repo := range config.Repos {
		if err := fetchRepoPulls(repo, client); err != nil {
			fmt.Println(err)
			return Failure
		}
	}

	for idx, length := 0, len(config.Orgs); idx < length; idx++ {
		if err := fetchOrgPulls(config.Orgs[idx], client); err != nil {
			fmt.Println(err)
			return Failure
		}
	}

	return Success
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
		pullUrl, err := url.Parse(github.CleanupPullsUrl(repos[idx].PullsUrl))
		if err != nil {
			return err
		}

		pulls, err := gh.Pulls(pullUrl)
		if err != nil {
			return err
		}

		pullLength := len(pulls)
		if pullLength == 0 {
			continue
		}

		fmt.Printf("%v pulls for %v (%v)\n", pullLength, repos[idx].Name, repos[idx].HtmlUrl)

		for pullIdx := 0; pullIdx < pullLength; pullIdx++ {
			printPull(&pulls[pullIdx])
		}
		fmt.Println()
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

	pulls, err := gh.Pulls(url)
	if err != nil {
		return err
	}

	pullLength := len(pulls)
	fmt.Printf("%v pulls for %v\n", pullLength, repoInfo)

	for idx := 0; idx < pullLength; idx++ {
		printPull(&pulls[idx])
	}

	fmt.Println()

	return nil
}

func printPull(pull *github.Pull) {
	fmt.Printf("\"%v\" from %v created at %v\n", pull.Title, pull.User.Login, pull.CreatedAt)
}
