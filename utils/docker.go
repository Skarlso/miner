package utils

import (
	"log"

	"os"

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
	client := getClient()
	opts := docker.CreateContainerOptions{
		Name: server,
	}
	container := client.CreateContainer
}

func getClient() *docker.Client {
	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		log.Println("Error while trying to connect to docker client: ", err)
	}
	return client
}
