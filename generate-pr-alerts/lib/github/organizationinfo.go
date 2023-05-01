package github

type OrganizationInfo struct {
	AvatarUrl   string `json:"avatar_url"`
	Company     string `json:"company"`
	Description string `json:"description"`
	HtmlUrl     string `json:"html_url"`
	Id          int    `json:"id"`
	Login       string `json:"login"`
	Name        string `json:"name"`
	ReposUrl    string `json:"repos_url"`
	Url         string `json:"url"`
}
