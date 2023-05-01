package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/Grayson/dashboard/generate-pr-alerts/lib/github"
	"gopkg.in/yaml.v2"
)

const (
	GitHubTokenEnvVar = "GITHUB_TOKEN"
)

type Config struct {
	Token string `yaml:"token"`
}

func main() {
	token, _ := os.LookupEnv(GitHubTokenEnvVar)
	config, err := getConfig()

	if err != nil {
		panic(err)
	}

	if config.Token != "" {
		token = config.Token
	}

	actualToken := flag.String("token", token, "GitHub Personal Access Token (can also set via environment variable `GITHUB_TOKEN`)")
	flag.Parse()

	if *actualToken == "" {
		fmt.Println("No GitHub token provided")
		os.Exit(-1)
	}

	client := github.NewClient(http.DefaultClient, *actualToken)
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

func getConfig() (*Config, error) {
	filename := "config.yaml"

	if _, err := os.Stat(filename); err != nil {
		return &Config{}, nil
	}

	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var c Config
	if err := yaml.Unmarshal(bytes, &c); err != nil {
		return nil, err
	}

	return &c, nil
}
