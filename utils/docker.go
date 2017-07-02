package utils

import (
	"log"

	"os"

	docker "github.com/fsouza/go-dockerclient"
)

// PullImage pulls an image named `image`
func PullImage(image string) {
	client := getClient()
	opts := docker.PullImageOptions{
		OutputStream: os.Stdout,
		Repository:   "skarlso/minecraft",
		Tag:          "1.12",
	}
	auth := docker.AuthConfiguration{}
	client.PullImage(opts, auth)
}

func getClient() *docker.Client {
	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		log.Println("Error while trying to connect to docker client: ", err)
	}
	return client
}
