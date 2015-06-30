package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestUnloadCommand_implement(t *testing.T) {
	var _ cli.Command = &UnloadCommand{}
}
