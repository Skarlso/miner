package commands

import (
	"log"

	"github.com/Skarlso/miner/config"
	"github.com/Skarlso/miner/utils"
	commander "github.com/Yitsushi/go-commander"
	"github.com/fatih/color"
)

// Start command struct.
type Start struct {
}

// Execute main entry point for the command.
func (s Start) Execute(opts *commander.CommandHelper) {
	c := config.Config{}
	c.Unmarshal()
	name := opts.Arg(0)
	if len(name) == 0 {
		name = c.Name
	}
	cyan := color.New(color.FgCyan).SprintFunc()
	version := utils.GetVersion(name)
	log.Printf("Starting server with name %s and version %s.\n", cyan(name), cyan(version))
	utils.StartServer(name, version)
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
