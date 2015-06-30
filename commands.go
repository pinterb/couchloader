package main

import (
	"github.com/mitchellh/cli"
	"github.com/pinterb/couchloader/command"
)

func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{

		"load": func() (cli.Command, error) {
			return &command.LoadCommand{
				Meta: *meta,
			}, nil
		},

		"list": func() (cli.Command, error) {
			return &command.ListCommand{
				Meta: *meta,
			}, nil
		},

		"unload": func() (cli.Command, error) {
			return &command.UnloadCommand{
				Meta: *meta,
			}, nil
		},
	}
}
