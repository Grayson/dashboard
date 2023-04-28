package github

type Label struct {
	Color       string `json:"color"`
	Description string `json:"description"`
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Url         string `json:"url"`
}
