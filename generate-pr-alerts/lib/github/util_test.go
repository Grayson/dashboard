package github

import "testing"

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
