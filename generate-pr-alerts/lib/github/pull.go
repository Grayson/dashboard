package github

type Pull struct {
	Assignee           User    `json:"assignee"`
	Assignees          []User  `json:"assignees"`
	AuthorAssociation  string  `json:"author_association"`
	Body               string  `json:"body"`
	ClosedAt           string  `json:"closed_at"`
	CreatedAt          string  `json:"created_at"`
	HtmlUrl            string  `json:"html_url"`
	Id                 int     `json:"id"`
	IsDraft            bool    `json:"draft"`
	Labels             []Label `json:"labels"`
	MergedAt           string  `json:"merged_at"`
	RequestedReviewers []User  `json:"requested_reviewers"`
	State              string  `json:"state"`
	Title              string  `json:"title"`
	UpdatedAt          string  `json:"updated_at"`
	Url                string  `json:"url"`
	User               User    `json:"user"`
}
