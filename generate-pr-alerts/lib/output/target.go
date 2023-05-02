package output

import (
	"github.com/Grayson/dashboard/generate-pr-alerts/lib/github"
)

type Target interface {
	Start() error

	StartReposPhase() error
	EndReposPhase() error
	StartOrganizationsPhase() error
	EndOrganizationsPhase() error

	StartOrganizations(orgs []github.OrganizationInfo) error
	VisitOrganization(org *github.OrganizationInfo) error
	EndOrganizations() error

	StartRepo(repo *github.OrganizationRepoInfo, pulls []github.Pull) error
	VisitPull(pull *github.Pull) error
	VisitIssue(issue *github.IssuesInfo) error
	EndRepo() error

	End() error
}
