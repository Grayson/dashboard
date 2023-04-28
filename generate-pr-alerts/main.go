package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/Grayson/dashboard/generate-pr-alerts/lib/github"
)

func main() {
	fmt.Println("generate-pr-alerts")

	client := github.NewClient(http.DefaultClient, "redacted")
	url, _ := url.Parse("https://api.github.com/repos/Grayson/dashboard/pulls")
	pulls, err := client.Pulls(url)
	print(pulls)
	print(err)
}
