package output

import (
	"encoding/json"
	"os"

	"github.com/Grayson/dashboard/generate-pr-alerts/lib/github"
)

func CreateJsonTarget(path string) *Json {
	return &Json{
		path: path,
		r:    root{},
	}
}

type phase int

const (
	rootPhase phase = iota
	orgPhase
)

type Json struct {
	path        string
	r           root
	phase       phase
	currentRepo *repo
	currentOrg  *org
}

type root struct {
	Repos []repo `json:"repos,omitempty"`
	Orgs  []org  `json:"orgs,omitempty"`
}

type org struct {
	Name  string `json:"name,omitempty"`
	Repos []repo `json:"repos,omitempty"`
	Url   string `json:"url,omitempty"`
}

type repo struct {
	Name   string  `json:"name,omitempty"`
	Url    string  `json:"url,omitempty"`
	Issues []issue `json:"issues,omitempty"`
	PRs    []pull  `json:"prs,omitempty"`
}

type pull struct {
	Title           string `json:"title,omitempty"`
	Username        string `json:"username,omitempty"`
	CreationDateUTC string `json:"creation_date_utc,omitempty"`
	Url             string `json:"url,omitempty"`
}

type issue struct {
	Title           string `json:"title,omitempty"`
	Username        string `json:"username,omitempty"`
	CreationDateUTC string `json:"creation_date_utc,omitempty"`
	Url             string `json:"url,omitempty"`
}

func mapOrg(o *github.OrganizationInfo) org {
	return org{
		Name:  o.Login,
		Repos: make([]repo, 0),
		Url:   o.HtmlUrl,
	}
}

func mapRepo(r *github.OrganizationRepoInfo) repo {
	return repo{
		Name:   r.Name,
		Issues: make([]issue, 0),
		PRs:    make([]pull, 0),
		Url:    r.HtmlUrl,
	}
}

func mapPull(p *github.Pull) pull {
	return pull{
		Title:           p.Title,
		Username:        p.User.Login,
		CreationDateUTC: p.CreatedAt,
		Url:             p.HtmlUrl,
	}
}

func mapIssue(i *github.IssuesInfo) issue {
	return issue{
		Title:           i.Title,
		Username:        i.User.Login,
		CreationDateUTC: i.CreatedAt,
		Url:             i.HtmlUrl,
	}
}

func (j *Json) Start() error {
	return nil
}

func (j *Json) StartReposPhase() error {
	j.phase = rootPhase
	j.r.Repos = make([]repo, 0)
	return nil
}

func (j *Json) EndReposPhase() error {
	return nil
}

func (j *Json) StartOrganizationsPhase() error {
	j.phase = orgPhase
	j.r.Orgs = make([]org, 0)
	return nil
}

func (j *Json) EndOrganizationsPhase() error {
	return nil
}

func (j *Json) VisitOrganization(org *github.OrganizationInfo) error {
	j.r.Orgs = append(j.r.Orgs, mapOrg(org))
	j.currentOrg = &j.r.Orgs[len(j.r.Orgs)-1]
	return nil
}

func (j *Json) StartRepo(repo *github.OrganizationRepoInfo) error {
	if j.phase == rootPhase {
		j.r.Repos = append(j.r.Repos, mapRepo(repo))
		j.currentRepo = &j.r.Repos[len(j.r.Repos)-1]
		return nil
	} else {
		j.currentOrg.Repos = append(j.currentOrg.Repos, mapRepo(repo))
		j.currentRepo = &j.currentOrg.Repos[len(j.currentOrg.Repos)-1]
	}
	return nil
}

func (j *Json) StartPulls(pulls []github.Pull) error {
	return nil
}

func (j *Json) VisitPull(pull *github.Pull) error {
	j.currentRepo.PRs = append(j.currentRepo.PRs, mapPull(pull))
	return nil
}

func (j *Json) EndPulls() error {
	return nil
}

func (j *Json) StartIssues(issues []github.IssuesInfo) error {
	return nil
}

func (j *Json) VisitIssue(issue *github.IssuesInfo) error {
	j.currentRepo.Issues = append(j.currentRepo.Issues, mapIssue(issue))
	return nil
}

func (j *Json) EndIssues() error {
	return nil
}

func (j *Json) EndRepo() error {
	return nil
}

func (j *Json) End() error {
	bytes, err := json.Marshal(j.r)
	if err != nil {
		return err
	}
	return os.WriteFile(j.path, bytes, 0755)
}
