package github

// https://docs.github.com/en/rest/issues/issues?apiVersion=2022-11-28#list-repository-issues

type IssuesInfo struct {
	Id          int     `json:"id,omitempty"`
	Url         string  `json:"url,omitempty"`
	HtmlUrl     string  `json:"html_url,omitempty"`
	Number      int     `json:"number,omitempty"`
	State       string  `json:"state,omitempty"`
	Title       string  `json:"title,omitempty"`
	Body        string  `json:"body,omitempty"`
	User        User    `json:"user,omitempty"`
	Labels      []Label `json:"labels,omitempty"`
	Assignee    User    `json:"assignee,omitempty"`
	Assigness   []User  `json:"assigness,omitempty"`
	CreatedAt   string  `json:"created_at,omitempty"`
	UpdatedAt   string  `json:"updated_at,omitempty"`
	ClosedAt    string  `json:"closed_at,omitempty"`
	StateReason string  `json:"state_reason,omitempty"`
}
