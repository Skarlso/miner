package commands

import (
	"fmt"

	commander "github.com/Yitsushi/go-commander"
	docker "github.com/fsouza/go-dockerclient"
)

// Attach to a running server to check status and execute commands.
type Attach struct {
}

// Execute main entry point for the command.
func (a Attach) Execute(opts *commander.CommandHelper) {
	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		panic(err)
	}
	imgs, err := client.ListImages(docker.ListImagesOptions{All: false})
	if err != nil {
		panic(err)
	}
	for _, img := range imgs {
		fmt.Println("ID: ", img.ID)
		fmt.Println("RepoTags: ", img.RepoTags)
		fmt.Println("Created: ", img.Created)
		fmt.Println("Size: ", img.Size)
		fmt.Println("VirtualSize: ", img.VirtualSize)
		fmt.Println("ParentId: ", img.ParentID)
	}
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
