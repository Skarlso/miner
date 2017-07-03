package utils

import (
	"context"
	"io"
	"log"

	"os"

	"path/filepath"

	"github.com/Skarlso/miner/config"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/fatih/color"
)

// PullImage pulls an image named `image`
func PullImage(image, version string) {
	cli := getClient()
	ctx := context.Background()
	log.Println("Pulling image: ", image+":"+version)
	out, err := cli.ImagePull(ctx, image+":"+version, types.ImagePullOptions{})
	if err != nil {
		log.Fatal("Error while pulling image: ", err)
	}
	defer out.Close()
	io.Copy(os.Stdout, out)
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
	c := config.Config{}
	c.Unmarshal()
	labels := map[string]string{
		"world": server,
	}
	cli := getClient()
	ctx := context.Background()
	containerConfig := &container.Config{
		AttachStderr: true,
		AttachStdin:  true,
		AttachStdout: true,
		Cmd:          []string{"bash", "-c", "echo \"eula=true\" > eula.txt ; java -jar /minecraft/" + mod + ".jar nogui"},
		Labels:       labels,
		WorkingDir:   "/data",
		Image:        c.RepoTag + ":" + version,
		Tty:          true,
		Volumes: map[string]struct{}{
			"/data": struct{}{},
		},
	}
	port, _ := nat.NewPort("tcp", "25565")
	bindings := []nat.PortBinding{
		HostIP:   "",
		HostPort: "25565",
	}
	portMap := map[nat.Port][]nat.PortBinding{
		port: bindings,
	}
	containerHostConfig := &container.HostConfig{
		Binds:        []string{filepath.Join(c.BindBase, server) + ":/data"},
		PortBindings: portMap,
	}
	resp, err := cli.ContainerCreate(ctx, containerConfig, nil, nil, "")
	if err != nil {
		log.Fatal("Error while creating container: ", err)
	}
	log.Println("Container started with ID: ", resp.ID)
	containerStartOpts := types.ContainerStartOptions{}
	container, err := client.CreateContainer(opts)
	if err != nil {
		log.Fatalln("Error creating container: ", err)
	}
	err = client.StartContainer(container.ID, &hostConfig)
	if err != nil {
		red := color.New(color.FgRed).SprintFunc()
		log.Fatal("Failed to start container with error: ", red(err.Error()))
	}
	green := color.New(color.FgGreen).SprintFunc()
	log.Println(green("Server successfully launched."))
}

func getClient() *client.Client {
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatal("Error while creating client: ", err)
	}
	return cli
}
