package app

import (
	"flag"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

const (
	GitHubTokenEnvVar = "GITHUB_TOKEN"
)

type Config struct {
	Token string   `yaml:"token"`
	Repos []string `yaml:"repos"`
	Orgs  []string `yaml:"orgs"`
	Json  string   `yaml:"json"`
}

type stringArray []string

func (a *stringArray) String() string {
	return strings.Join(*a, ", ")
}

func (a *stringArray) Set(s string) error {
	*a = append(*a, s)
	return nil
}

func GatherConfig() (*Config, error) {
	config, err := getFileConfig()

	if err != nil {
		return nil, err
	}

	flagToken, flagJson, flagOrgs, flagRepos := gatherFlagArgs(flag.CommandLine, os.Args)

	if flagToken != "" {
		config.Token = flagToken
	}
	if flagJson != "" {
		config.Json = flagJson
	}

	config.Orgs = append(config.Orgs, flagOrgs...)
	config.Repos = append(config.Repos, flagRepos...)

	if config.Token == "" {
		token, _ := os.LookupEnv(GitHubTokenEnvVar)
		config.Token = token
	}

	return config, nil
}

func gatherFlagArgs(flagSet *flag.FlagSet, args []string) (string, string, []string, []string) {
	var token *string
	var json *string
	flagOrgs := make(stringArray, 0)
	flagRepos := make(stringArray, 0)

	token = flagSet.String("token", "", "GitHub Personal Access Token (can also set via environment variable `GITHUB_TOKEN`)")
	json = flagSet.String("json", "", "If specified, will generate JSON output at path provided")
	flagSet.Var(&flagOrgs, "org", "Organization name to check for repos (can be specified multiple times)")
	flagSet.Var(&flagRepos, "repo", "Repositories to check for pulls in 'user/repo-name' format (can be specified mulitple times)")
	flagSet.Parse(args[1:])

	return *token, *json, flagOrgs, flagRepos
}

func getFileConfig() (*Config, error) {
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
