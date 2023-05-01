package github

type OrganizationRepoInfo struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	FullName    string `json:"full_name,omitempty"`
	GitUrl      string `json:"git_url,omitempty"`
	HtmlUrl     string `json:"html_url,omitempty"`
	Description string `json:"description,omitempty"`
	Url         string `json:"url,omitempty"`
	PullsUrl    string `json:"pulls_url,omitempty"`
}
