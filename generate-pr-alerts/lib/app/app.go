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

	return Success
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
	fmt.Printf("\"%v\" from %v created at %v", pull.Title, pull.User.Login, pull.CreatedAt)
}
