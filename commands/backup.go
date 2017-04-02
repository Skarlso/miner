package commands

import commander "github.com/Yitsushi/go-commander"

// Backup a selected world.
type Backup struct {
}

// Execute main entry point for the command.
func (b Backup) Execute(opts *commander.CommandHelper) {
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
