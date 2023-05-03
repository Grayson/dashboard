package output

import (
	"fmt"
	"path"

	"github.com/Grayson/dashboard/generate-pr-alerts/lib/github"
)

var STDOUT = &Stdout{}

type Stdout struct {
	repo *github.OrganizationRepoInfo
}

func (s *Stdout) Start() error {
	return nil
}

func (s *Stdout) StartReposPhase() error {
	fmt.Println("# Repos")
	return nil
}

func (s *Stdout) EndReposPhase() error {
	return nil
}

func (s *Stdout) StartOrganizationsPhase() error {
	return nil
}

func (s *Stdout) EndOrganizationsPhase() error {
	return nil
}

func (s *Stdout) StartOrganizations(orgs []github.OrganizationInfo) error {
	return nil
}

func (s *Stdout) VisitOrganization(org *github.OrganizationInfo) error {
	fmt.Printf("# %v\n", org.Login)
	return nil
}

func (s *Stdout) EndOrganizations() error {
	return nil
}

func (s *Stdout) StartRepo(repo *github.OrganizationRepoInfo) error {
	s.repo = repo
	return nil
}

func (s *Stdout) StartPulls(pulls []github.Pull) error {
	fmt.Printf("%v pulls for [%v](%v)\n", len(pulls), s.repo.Name, s.repo.HtmlUrl)
	return nil
}

func (s *Stdout) VisitPull(pull *github.Pull) error {
	number := path.Base(pull.HtmlUrl)
	fmt.Printf("* Pull: \"%v\" from %v created at %v [#%v](%v)\n", pull.Title, pull.User.Login, pull.CreatedAt, number, pull.HtmlUrl)
	return nil
}

func (s *Stdout) EndPulls() error {
	fmt.Println()
	return nil
}

func (s *Stdout) StartIssues(issues []github.IssuesInfo) error {
	fmt.Printf("%v issues for [%v](%v)\n", len(issues), s.repo.Name, s.repo.HtmlUrl)
	return nil
}

func (s *Stdout) VisitIssue(issue *github.IssuesInfo) error {
	fmt.Printf("* Issue: \"%v\" from %v created at %v [#%v](%v)\n", issue.Title, issue.User.Login, issue.CreatedAt, issue.Number, issue.HtmlUrl)
	return nil
}

func (s *Stdout) EndIssues() error {
	return nil
}

func (s *Stdout) EndRepo() error {
	return nil
}

func (s *Stdout) End() error {
	return nil
}
