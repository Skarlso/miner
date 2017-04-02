package commands

import commander "github.com/Yitsushi/go-commander"

// Stop command struct.
type Stop struct {
}

// Execute main entry point for the command.
func (c Stop) Execute(opts *commander.CommandHelper) {
}

// NewStop Creates a new Stop command.
func NewStop(appName string) *commander.CommandWrapper {
	return &commander.CommandWrapper{
		Handler: &Stop{},
		Help: &commander.CommandDescriptor{
			Name:             "stop",
			ShortDescription: "Stop a running server.",
			LongDescription:  `Stop a running server. Stop also issues a stop command to the running minecraft server. Thus state will be saved.`,
			Arguments:        "",
			Examples:         []string{""},
		},
	}
}
