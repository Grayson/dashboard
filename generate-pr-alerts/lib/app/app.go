package app

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/Grayson/dashboard/generate-pr-alerts/lib/github"
	"github.com/Grayson/dashboard/generate-pr-alerts/lib/output"
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

	gh := github.NewClient(http.DefaultClient, config.Token)
	target := createTarget(config)

	target.Start()

	if err := processUserRepos(config.Repos, gh, target); err != nil {
		fmt.Println(err)
		return Failure
	}

	target.StartOrganizationsPhase()
	for idx, length := 0, len(config.Orgs); idx < length; idx++ {
		if err := fetchOrgPulls(config.Orgs[idx], gh, target); err != nil {
			fmt.Println(err)
			return Failure
		}
	}
	target.EndOrganizationsPhase()

	target.End()

	return Success
}

func createTarget(config *Config) output.Target {
	if config.Json == "" {
		return output.STDOUT
	}

	return output.NewMultiTarget(
		output.STDOUT,
		output.NewJsonTarget(config.Json),
	)
}

func processUserRepos(repos []string, gh github.GitHub, target output.Target) error {
	repoLength := len(repos)
	if repoLength == 0 {
		return nil
	}

	if err := target.StartReposPhase(); err != nil {
		return err
	}

	for _, repo := range repos {
		if err := fetchRepoPulls(repo, gh, target); err != nil {
			return err
		}
	}
	return target.EndReposPhase()
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

	target.VisitOrganization(&org)
	reposUrl, _ := url.Parse(org.ReposUrl)
	repos, err := gh.OrganizationRepos(reposUrl)
	if err != nil {
		return err
	}

	repoChan := make(chan *github.OrganizationRepoInfo)

	go func() {
		for idx, length := 0, len(repos); idx < length; idx++ {
			repoChan <- &repos[idx]
		}
		close(repoChan)
	}()

	const workerLimit = 4
	wg := sync.WaitGroup{}
	for worker := 0; worker < workerLimit; worker++ {
		wg.Add(1)
		go func() {
			defer func() { wg.Done() }()
			for repo := range repoChan {
				err = repoWork(repo, url, gh, target)
				if err != nil {
					return
				}
			}
		}()
	}
	wg.Wait()
	defer func() { fmt.Println() }()
	return err
}

func repoWork(repo *github.OrganizationRepoInfo, url *url.URL, gh github.GitHub, target output.Target) error {
	target.StartRepo(repo)
	if err := logRepoPulls(url, repo, gh, target); err != nil {
		return err
	}

	if err := logRepoIssues(url, repo, gh, target); err != nil {
		return err
	}
	target.EndRepo()
	return nil
}

func logRepoPulls(url *url.URL, repo *github.OrganizationRepoInfo, gh github.GitHub, target output.Target) error {
	pullUrl, err := url.Parse(github.CleanupPullsUrl(repo.PullsUrl))
	if err != nil {
		return err
	}
	if err := printPulls(gh, pullUrl, target); err != nil {
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
		target.VisitIssue(&issues[idx])
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

	return printPulls(gh, url, target)
}

func printPulls(gh github.GitHub, url *url.URL, target output.Target) error {
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
		target.VisitPull(&pulls[idx])
	}
	target.EndPulls()

	return nil
}
