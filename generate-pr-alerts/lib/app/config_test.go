package app

import (
	"flag"
	"reflect"
	"testing"
)

func Test_gatherFlagArgs(t *testing.T) {
	tests := []struct {
		name  string
		want  string
		want1 []string
		want2 []string
		args  []string
	}{
		{
			"Get token",
			"token",
			[]string{},
			[]string{},
			[]string{"-token", "token"},
		},
		{
			"Get one org",
			"",
			[]string{"org"},
			[]string{},
			[]string{"-org", "org"},
		},
		{
			"Get multiple orgs",
			"",
			[]string{"org", "org2"},
			[]string{},
			[]string{"-org", "org", "-org", "org2"},
		},
		{
			"Get one repo",
			"",
			[]string{},
			[]string{"repo"},
			[]string{"-repo", "repo"},
		},
		{
			"Get multiple repos",
			"",
			[]string{},
			[]string{"repo", "repo2"},
			[]string{"-repo", "repo", "-repo", "repo2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flagSet := flag.NewFlagSet("config_test", flag.ExitOnError)
			args := append([]string{"config_test"}, tt.args...)

			got, got1, got2 := gatherFlagArgs(flagSet, args)
			if got != tt.want {
				t.Errorf("gatherFlagArgs() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("gatherFlagArgs() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("gatherFlagArgs() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
