package utils

import (
	"log"

	"os"
	"github.com/Skarlso/miner/config"
	docker "github.com/fsouza/go-dockerclient"
)

// PullImage pulls an image named `image`
func PullImage(image, version string) {
	client := getClient()
	opts := docker.PullImageOptions{
		OutputStream: os.Stdout,
		Repository:   image,
		Tag:          version,
	}
	auth := docker.AuthConfiguration{}
	client.PullImage(opts, auth)
}

// StartServer starts a server
func StartServer(server, version string) {
	var mod string
	switch config.GetMod() {
	case config.CRAFTBUKKIT:
		mod = "craftbukkit"
	case config.FORGE:
		mod = "forge"
	}
	labels := map[string]string{
		"world": server,
	}
	client := getClient()
	config := docker.Config{
		AttachStderr: true,
		AttachStdin:  true,
		AttachStdout: true,
		Cmd: []string{"bash", "-c", "echo \"eula=true\" > eula.txt ; java -jar /minecraft/" + mod + ".jar nogui"},
		Labels: labels,
	}
	opts := docker.CreateContainerOptions{
		Name:   server,
		Config: &config,
	}
	hostConfig := docker.HostConfig{
		PortBindings
	}
	container, err := client.CreateContainer(opts)
	if err != nil {
		log.Fatalln("Error creating container: ", err)
	}
	client.StartContainer(container.ID, )
}

func getClient() *docker.Client {
	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		log.Println("Error while trying to connect to docker client: ", err)
	}
	return client
}
