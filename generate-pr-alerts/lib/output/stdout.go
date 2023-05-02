package output

import "github.com/Grayson/dashboard/generate-pr-alerts/lib/github"

type Stdout struct {
}

func (s *Stdout) Start() error {
	return nil
}

func (s *Stdout) StartReposPhase() error {
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
	return nil
}
func (s *Stdout) EndOrganizations() error {
	return nil
}

func (s *Stdout) StartRepo(repo *github.OrganizationRepoInfo, pulls []github.Pull) error {
	return nil
}
func (s *Stdout) VisitPull(pull *github.Pull) error {
	return nil
}
func (s *Stdout) VisitIssue(issue *github.IssuesInfo) error {
	return nil
}
func (s *Stdout) EndRepo() error {
	return nil
}

func (s *Stdout) End() error {
	return nil
}
