package output

import "github.com/Grayson/dashboard/generate-pr-alerts/lib/github"

type MultiTarget struct {
	targets []Target
	count   int
}

func NewMultiTarget(t ...Target) *MultiTarget {

	return &MultiTarget{t, len(t)}
}

func (t *MultiTarget) Start() error {
	for idx := 0; idx < t.count; idx++ {
		if err := t.targets[idx].Start(); err != nil {
			return err
		}
	}
	return nil
}

func (t *MultiTarget) StartReposPhase() error {
	for idx := 0; idx < t.count; idx++ {
		if err := t.targets[idx].StartReposPhase(); err != nil {
			return err
		}
	}
	return nil
}

func (t *MultiTarget) EndReposPhase() error {
	for idx := 0; idx < t.count; idx++ {
		if err := t.targets[idx].EndReposPhase(); err != nil {
			return err
		}
	}
	return nil
}

func (t *MultiTarget) StartOrganizationsPhase() error {
	for idx := 0; idx < t.count; idx++ {
		if err := t.targets[idx].StartOrganizationsPhase(); err != nil {
			return err
		}
	}
	return nil
}

func (t *MultiTarget) EndOrganizationsPhase() error {
	for idx := 0; idx < t.count; idx++ {
		if err := t.targets[idx].EndOrganizationsPhase(); err != nil {
			return err
		}
	}
	return nil
}

func (t *MultiTarget) VisitOrganization(org *github.OrganizationInfo) error {
	for idx := 0; idx < t.count; idx++ {
		if err := t.targets[idx].VisitOrganization(org); err != nil {
			return err
		}
	}
	return nil
}

func (t *MultiTarget) StartRepo(repo *github.OrganizationRepoInfo) error {
	for idx := 0; idx < t.count; idx++ {
		if err := t.targets[idx].StartRepo(repo); err != nil {
			return err
		}
	}
	return nil
}

func (t *MultiTarget) StartPulls(pulls []github.Pull) error {
	for idx := 0; idx < t.count; idx++ {
		if err := t.targets[idx].StartPulls(pulls); err != nil {
			return err
		}
	}
	return nil
}

func (t *MultiTarget) VisitPull(pull *github.Pull) error {
	for idx := 0; idx < t.count; idx++ {
		if err := t.targets[idx].VisitPull(pull); err != nil {
			return err
		}
	}
	return nil
}

func (t *MultiTarget) EndPulls() error {
	for idx := 0; idx < t.count; idx++ {
		if err := t.targets[idx].EndPulls(); err != nil {
			return err
		}
	}
	return nil
}

func (t *MultiTarget) StartIssues(issues []github.IssuesInfo) error {
	for idx := 0; idx < t.count; idx++ {
		if err := t.targets[idx].StartIssues(issues); err != nil {
			return err
		}
	}
	return nil
}

func (t *MultiTarget) VisitIssue(issue *github.IssuesInfo) error {
	for idx := 0; idx < t.count; idx++ {
		if err := t.targets[idx].VisitIssue(issue); err != nil {
			return err
		}
	}
	return nil
}

func (t *MultiTarget) EndIssues() error {
	for idx := 0; idx < t.count; idx++ {
		if err := t.targets[idx].EndIssues(); err != nil {
			return err
		}
	}
	return nil
}

func (t *MultiTarget) EndRepo() error {
	for idx := 0; idx < t.count; idx++ {
		if err := t.targets[idx].EndRepo(); err != nil {
			return err
		}
	}
	return nil
}

func (t *MultiTarget) End() error {
	for idx := 0; idx < t.count; idx++ {
		if err := t.targets[idx].End(); err != nil {
			return err
		}
	}
	return nil
}
