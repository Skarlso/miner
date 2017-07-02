package commands

import (
	"log"

	"github.com/Skarlso/miner/config"
	"github.com/Skarlso/miner/utils"
	commander "github.com/Yitsushi/go-commander"
	"github.com/fatih/color"
)

// Setup command struct.
type Setup struct {
}

// Execute main entry point for the command.
func (s Setup) Execute(opts *commander.CommandHelper) {
	c := config.Config{}
	c.Unmarshal()
	name := opts.Arg(0)
	if len(name) == 0 {
		name = c.Name
	}
	red := color.New(color.FgRed).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	version := opts.Arg(1)
	if len(version) == 0 {
		log.Fatalf(red("Please provide a version number to use."))
	}
	log.Printf(`Creating world with following options:
Name: %s
Version: %s
Mod: %s
Bind Location: %s`, cyan(name), cyan(version), cyan(config.GetMod()), cyan(c.BindBase))
	utils.PullImage(c.BindBase + ":" + version)
}

// NewSetup Creates a new Setup command.
func NewSetup(appName string) *commander.CommandWrapper {
	return &commander.CommandWrapper{
		Handler: &Setup{},
		Help: &commander.CommandDescriptor{
			Name:             "setup",
			ShortDescription: "Setup a new server.",
			LongDescription:  `Setup a new server with a given version. If non is provided, latest is used.`,
			Arguments:        "name version",
			Examples:         []string{"", "my_server 1.11.1", ""},
		},
	}
}
