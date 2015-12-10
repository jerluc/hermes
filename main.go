package main

import (
	"os"

	hermes "github.com/jerluc/hermes/lib"
)

var notifier hermes.Notifier

func handleFailure(cmd *hermes.Command, err error) {
	notifier.Failure(cmd, err)
	// TODO: Get the platform-independent exit code
	os.Exit(1)
}

func handleSuccess(cmd *hermes.Command) {
	notifier.Success(cmd)
	os.Exit(0)
}

func handleStdout(cmd *hermes.Command, buffer string) {
	
}

func handleStderr(cmd *hermes.Command, buffer string) {
	
}

func installNotifier() {
	// TODO: Load config from the filesystem
	config := hermes.Config{
		"notifier": hermes.Config{
			"type": "slack",
			"webHookUrl": "https://hooks.slack.com/services/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",
			"recipient": "@jerluc",
		},
	}
	notifier = hermes.GetNotifier(config)
}

func main() {
	installNotifier()

	// TODO: Check for number of args or print usage

	cmd := hermes.NewCommand(os.Args[1:])

	runErr := cmd.Run()
	if runErr != nil {
		handleFailure(cmd, runErr)
	}

	handleSuccess(cmd)
}