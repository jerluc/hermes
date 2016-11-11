package main

import (
	"fmt"
	hermes "github.com/jerluc/hermes/lib"
	"os"
	"os/user"
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
	usr, _ := user.Current()
	dir := usr.HomeDir
	config, configErr := hermes.LoadConfig(dir + "/.hermes.yml")
	if configErr != nil {
		fmt.Println(configErr)
		os.Exit(1)
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
