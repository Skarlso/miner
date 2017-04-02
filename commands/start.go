package commands

import commander "github.com/Yitsushi/go-commander"

// Start command struct.
type Start struct {
}

// Execute main entry point for the command.
func (c Start) Execute(opts *commander.CommandHelper) {
}

// NewStart Creates a new Start command.
func NewStart(appName string) *commander.CommandWrapper {
	return &commander.CommandWrapper{
		Handler: &Start{},
		Help: &commander.CommandDescriptor{
			Name:             "start",
			ShortDescription: "Start a given server.",
			LongDescription:  `Start a server with a given name. If non is provided, a default name will be used.`,
			Arguments:        "name",
			Examples:         []string{"", "myserver", "zombie_apo_map"},
		},
	}
}
