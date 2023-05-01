package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Grayson/dashboard/generate-pr-alerts/lib/app"
	"github.com/Grayson/dashboard/generate-pr-alerts/lib/github"
)

func main() {
	config, err := app.GatherConfig()

	if err != nil {
		panic(err)
	}

	if config.Token == "" {
		fmt.Println("No GitHub token provided")
		os.Exit(-1)
	}

	client := github.NewClient(http.DefaultClient, config.Token)
	// url, _ := url.Parse("https://api.github.com/repos/Grayson/dashboard/pulls")
	// pulls, err := client.Pulls(url)
	// fmt.Println(pulls)
	// fmt.Println(err)

	// orgUrl, _ := url.Parse("https://api.github.com/orgs/objectiveceo")
	// org, orgErr := client.OrganizationInfo(orgUrl)
	// fmt.Println(org)
	// fmt.Println(orgErr)

	orgReposUrl, _ := github.OrganizationReposUrl("objectiveceo")
	orgRepos, orgReposErr := client.OrganizationRepos(orgReposUrl)
	fmt.Printf("error: %v\n", orgReposErr)
	for _, x := range orgRepos {
		fmt.Printf("%#v", x)
		fmt.Println("-")
	}
}
