package commands

import commander "github.com/Yitsushi/go-commander"

// List command struct.
type List struct {
}

// Execute main entry point for the command.
func (c List) Execute(opts *commander.CommandHelper) {
}

// NewList Creates a new List command.
func NewList(appName string) *commander.CommandWrapper {
	return &commander.CommandWrapper{
		Handler: &List{},
		Help: &commander.CommandDescriptor{
			Name:             "list",
			ShortDescription: "List the created servers.",
			LongDescription:  `Return a list of created servers.`,
			Arguments:        "",
			Examples:         []string{""},
		},
	}
}
