package commands

import (
	"github.com/Skarlso/miner/config"
	"github.com/Skarlso/miner/utils"
	commander "github.com/Yitsushi/go-commander"
)

// Attach to a running server to check status and execute commands.
type Attach struct {
}

// Execute main entry point for the command.
func (a Attach) Execute(opts *commander.CommandHelper) {
	c := config.Config{}
	c.Unmarshal()
	name := opts.Arg(0)
	if len(name) == 0 {
		name = c.Name
	}
	utils.AttachServer(name)
}

// NewAttach Creates a new Attach command.
func NewAttach(appName string) *commander.CommandWrapper {
	return &commander.CommandWrapper{
		Handler: &Attach{},
		Help: &commander.CommandDescriptor{
			Name:             "attach",
			ShortDescription: "Attach to running container.",
			LongDescription:  `Attach to a running container to check on the server status and execute commands.`,
			Arguments:        "",
			Examples:         []string{""},
		},
	}
}
