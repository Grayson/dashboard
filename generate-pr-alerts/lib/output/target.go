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

	StartRepo(repo *github.OrganizationRepoInfo) error

	StartPulls(pulls []github.Pull) error
	VisitPull(pull *github.Pull) error
	EndPulls() error

	StartIssues(issues []github.IssuesInfo) error
	VisitIssue(issue *github.IssuesInfo) error
	EndIssues() error

	EndRepo() error

	End() error
}
