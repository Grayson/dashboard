package main

import (
	"os"

	"github.com/Grayson/dashboard/generate-pr-alerts/lib/app"
	"github.com/Grayson/dashboard/generate-pr-alerts/lib/output"
)

func main() {
	config, err := app.GatherConfig()

	if err != nil {
		panic(err)
	}

	if exitCode := app.Run(config, output.STDOUT); exitCode != app.Success {
		os.Exit(int(exitCode))
	}
}
