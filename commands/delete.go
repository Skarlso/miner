package commands

import (
	"github.com/Skarlso/miner/config"
	"github.com/Skarlso/miner/utils"
	commander "github.com/Yitsushi/go-commander"
)

// Delete command struct.
type Delete struct {
}

// Execute main entry point for the command.
func (s Delete) Execute(opts *commander.CommandHelper) {
	c := config.Config{}
	c.Unmarshal()
	name := opts.Arg(0)
	if len(name) == 0 {
		name = c.Name
	}
	utils.DeleteServer(name)
}

// NewDelete Creates a new Delete command.
func NewDelete(appName string) *commander.CommandWrapper {
	return &commander.CommandWrapper{
		Handler: &Delete{},
		Help: &commander.CommandDescriptor{
			Name:             "delete",
			ShortDescription: "Delete a server.",
			LongDescription:  `Delete a server will destroy the world and container and the version file.`,
			Arguments:        "name",
			Examples:         []string{"", "my_server"},
		},
	}
}
