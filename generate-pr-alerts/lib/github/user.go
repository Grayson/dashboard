package github

type User struct {
	AvatarUrl  string `json:"avatar_url"`
	GravatarId string `json:"gravatar_id"`
	HtmlUrl    string `json:"html_url"`
	Login      string `json:"login"`
	Url        string `json:"url"`
}
