package commands

import commander "github.com/Yitsushi/go-commander"

// Create command struct.
type Create struct {
}

// Execute main entry point for the command.
func (c Create) Execute(opts *commander.CommandHelper) {
}

// NewCreate Creates a new Create command.
func NewCreate(appName string) *commander.CommandWrapper {
	return &commander.CommandWrapper{
		Handler: &Create{},
		Help: &commander.CommandDescriptor{
			Name:             "create",
			ShortDescription: "Create a new server.",
			LongDescription:  `Create a new server with a given version. If non is provided, latest is used.`,
			Arguments:        "version",
			Examples:         []string{"11.1", "1.9", ""},
		},
	}
}
