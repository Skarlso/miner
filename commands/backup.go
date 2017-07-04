package commands

import (
	"log"

	"github.com/Skarlso/miner/config"
	"github.com/Skarlso/miner/utils"
	commander "github.com/Yitsushi/go-commander"
	"github.com/fatih/color"
)

// Backup a selected world.
type Backup struct {
}

// Execute main entry point for the command.
func (b Backup) Execute(opts *commander.CommandHelper) {
	c := config.Config{}
	c.Unmarshal()
	name := opts.Arg(0)
	if len(name) == 0 {
		name = c.Name
	}
	cyan := color.New(color.FgCyan).SprintFunc()
	log.Printf("Backing up server with name %s.\n", cyan(name))
	utils.Backup(name)
}

// NewBackup Creates a new Backup command.
func NewBackup(appName string) *commander.CommandWrapper {
	return &commander.CommandWrapper{
		Handler: &Backup{},
		Help: &commander.CommandDescriptor{
			Name:             "backup",
			ShortDescription: "Backup a world.",
			LongDescription:  `Backup compresses a choosen world and uploads that archive to a configured S3 bucket.`,
			Arguments:        "name",
			Examples:         []string{"myserver"},
		},
	}
}
