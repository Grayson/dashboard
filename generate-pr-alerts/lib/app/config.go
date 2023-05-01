package app

import (
	"flag"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	GitHubTokenEnvVar = "GITHUB_TOKEN"
)

type Config struct {
	Token string   `yaml:"token"`
	Repos []string `yaml:"repos"`
	Orgs  []string `yaml:"orgs"`
}

func GatherConfig() (*Config, error) {
	token, _ := os.LookupEnv(GitHubTokenEnvVar)
	config, err := getFileConfig()

	if err != nil {
		return nil, err
	}

	if config.Token != "" {
		token = config.Token
	}

	config.Token = *flag.String("token", token, "GitHub Personal Access Token (can also set via environment variable `GITHUB_TOKEN`)")
	flag.Parse()
	return config, nil
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
