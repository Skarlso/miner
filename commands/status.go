package commands

import (
	"github.com/Skarlso/miner/config"
	"github.com/Skarlso/miner/utils"
	commander "github.com/Yitsushi/go-commander"
)

// Status command struct.
type Status struct {
}

// Execute main entry point for the command.
func (s Status) Execute(opts *commander.CommandHelper) {
	c := config.Config{}
	c.Unmarshal()
	name := opts.Arg(0)
	if len(name) == 0 {
		name = c.Name
	}
	utils.StatusServer(name)
}

// NewStatus Creates a new Status command.
func NewStatus(appName string) *commander.CommandWrapper {
	return &commander.CommandWrapper{
		Handler: &Status{},
		Help: &commander.CommandDescriptor{
			Name:             "status",
			ShortDescription: "Status of a server.",
			LongDescription:  `Status a server will display information about the container and if the world was created.`,
			Arguments:        "name",
			Examples:         []string{"", "my_server"},
		},
	}
}
