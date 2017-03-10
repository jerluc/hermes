package main

import (
	"fmt"
	hermes "github.com/jerluc/hermes/lib"
	"os"
	"os/user"
)

func instantiateNotifier() (hermes.Notifier, error) {
	dir := "."
	if _, err := os.Stat(".hermes.yml"); os.IsNotExist(err) {
		usr, _ := user.Current()
		dir = usr.HomeDir
	}
	config, configErr := hermes.LoadConfig(dir + "/.hermes.yml")
	if configErr != nil {
		return nil, configErr
	}
	return hermes.GetNotifier(config), nil
}

func main() {
	notifier, err := instantiateNotifier()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// TODO: Check for number of args or print usage

	cmd := hermes.NewCommand(os.Args[1:])

	exitCode := cmd.Run(notifier)
	os.Exit(exitCode)
}
