package commands

import (
	"bytes"
	"log"

	"github.com/Skarlso/miner/config"
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
	containerName := opts.Arg(0)
	if len(containerName) < 1 {
		c := config.Config{}
		c.Unmarshal()
		containerName = c.Name
	}
	log.Println("Attaching to:", containerName)
	var buf bytes.Buffer
	err = client.AttachToContainer(docker.AttachToContainerOptions{
		Container:    containerName,
		OutputStream: &buf,
		Logs:         true,
		Stdout:       true,
		Stderr:       true,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(buf.String())
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
