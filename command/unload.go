package command

import (
	"github.com/mitchellh/cli"
	"strings"
)

type UnloadCommand struct {
	Ui cli.Ui
}

func (c *UnloadCommand) Run(args []string) int {
	// Write your code here
	return 0
}

func (c *UnloadCommand) Synopsis() string {
	return ""
}

func (c *UnloadCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
