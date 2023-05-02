package github

type OrganizationRepoInfo struct {
	Description string `json:"description,omitempty"`
	FullName    string `json:"full_name,omitempty"`
	GitUrl      string `json:"git_url,omitempty"`
	HtmlUrl     string `json:"html_url,omitempty"`
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	PullsUrl    string `json:"pulls_url,omitempty"`
	Url         string `json:"url,omitempty"`
}
