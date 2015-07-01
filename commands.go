package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/mitchellh/cli"
	"github.com/pinterb/couchloader/command"
)

// Commands is the mapping of all the available couchloader commands.
var Commands map[string]cli.CommandFactory

func init() {
	ui := &cli.BasicUi{Writer: os.Stdout}

	Commands = map[string]cli.CommandFactory{
		"load": func() (cli.Command, error) {
			return &command.LoadCommand{
				Ui: ui,
			}, nil
		},

		"list": func() (cli.Command, error) {
			return &command.ListCommand{
				Ui: ui,
			}, nil
		},

		"unload": func() (cli.Command, error) {
			return &command.UnloadCommand{
				Ui: ui,
			}, nil
		},

		"version": func() (cli.Command, error) {
			ver := Version
			rel := VersionPrerelease
			if GitDescribe != "" {
				ver = GitDescribe
			}
			if GitDescribe == "" && rel == "" {
				rel = "dev"
			}

			return &command.VersionCommand{
				Revision:          GitCommit,
				Version:           ver,
				VersionPrerelease: rel,
				Ui:                ui,
			}, nil
		},
	}
}

// makeShutdownCh returns a channel that can be used for shutdown
// notifications for commands. This channel will send a message for every
// interrupt or SIGTERM received.
func makeShutdownCh() <-chan struct{} {
	resultCh := make(chan struct{})

	signalCh := make(chan os.Signal, 4)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		for {
			<-signalCh
			resultCh <- struct{}{}
		}
	}()

	return resultCh
}
