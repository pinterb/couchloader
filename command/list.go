package command

import (
	"github.com/mitchellh/cli"
	"strings"
)

type ListCommand struct {
	Ui cli.Ui
}

func (c *ListCommand) Run(args []string) int {
	// Write your code here
	return 0
}

func (c *ListCommand) Synopsis() string {
	return ""
}

func (c *ListCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
