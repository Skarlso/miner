package commands

import (
	"github.com/Skarlso/miner/config"
	"github.com/Skarlso/miner/utils"
	commander "github.com/Yitsushi/go-commander"
)

// Stop command struct.
type Stop struct {
}

// Execute main entry point for the command.
func (s Stop) Execute(opts *commander.CommandHelper) {
	c := config.Config{}
	c.Unmarshal()
	name := opts.Arg(0)
	if len(name) == 0 {
		name = c.Name
	}
	utils.StopServer(name)
}

// NewStop Creates a new Stop command.
func NewStop(appName string) *commander.CommandWrapper {
	return &commander.CommandWrapper{
		Handler: &Stop{},
		Help: &commander.CommandDescriptor{
			Name:             "stop",
			ShortDescription: "Stop a running server.",
			LongDescription:  `Stop a running server. Stop also issues a stop command to the running minecraft server. Thus state will be saved.`,
			Arguments:        "name",
			Examples:         []string{"", "my_server"},
		},
	}
}
