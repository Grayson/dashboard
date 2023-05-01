package github

import (
	"net/url"
	"reflect"
	"testing"
)

func TestCleanupPullsUrl(t *testing.T) {
	type args struct {
		urlString string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Clean up {/number} syntax",
			args{"https://api.github.com/repos/objectiveceo/friendly-timestamp/pulls{/number}"},
			"https://api.github.com/repos/objectiveceo/friendly-timestamp/pulls",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CleanupPullsUrl(tt.args.urlString); got != tt.want {
				t.Errorf("CleanupPullsUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrganizationInfoUrl(t *testing.T) {
	basicUrl, _ := url.Parse("https://api.github.com/orgs/test")

	type args struct {
		orgName string
	}
	tests := []struct {
		name    string
		args    args
		want    *url.URL
		wantErr bool
	}{
		{
			"Basic url generation",
			args{"test"},
			basicUrl,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OrganizationInfoUrl(tt.args.orgName)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrganizationInfoUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrganizationInfoUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrganizationReposUrl(t *testing.T) {
	basicUrl, _ := url.Parse("https://api.github.com/orgs/test/repos")

	type args struct {
		orgName string
	}
	tests := []struct {
		name    string
		args    args
		want    *url.URL
		wantErr bool
	}{
		{
			"Basic url generation",
			args{"test"},
			basicUrl,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OrganizationReposUrl(tt.args.orgName)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrganizationReposUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrganizationReposUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPullsUrl(t *testing.T) {
	basicUrl, _ := url.Parse("https://api.github.com/name/repo/pulls")

	type args struct {
		user string
		repo string
	}
	tests := []struct {
		name    string
		args    args
		want    *url.URL
		wantErr bool
	}{
		{
			"Basic url generation",
			args{"name", "repo"},
			basicUrl,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PullsUrl(tt.args.user, tt.args.repo)
			if (err != nil) != tt.wantErr {
				t.Errorf("PullsUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PullsUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
