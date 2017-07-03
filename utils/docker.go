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
	"github.com/docker/docker/api/types/filters"
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

// AttachServer attach to minecraft server
func AttachServer(server string) {
	cli := getClient()
	ctx := context.Background()
	con := getDockerContainer(server)
	log.Printf("Attaching to server '%s' with container name '%s'\n.", server, con)
	conn, err := cli.ContainerAttach(ctx, con, types.ContainerAttachOptions{
		Stderr: true,
		Stdin:  true,
		Stdout: true,
		Stream: true,
		Logs:   true,
	})
	if err != nil {
		log.Fatal("Error while attaching: ", err)
	}
	defer conn.Close()
	io.Copy(os.Stdout, conn.Reader)
}

// StopServer stopping minecraft server
func StopServer(server string) {
	cli := getClient()
	ctx := context.Background()
	con := getDockerContainer(server)
	log.Printf("Stoping server '%s' with container name '%s'\n.", server, con)
	conn, err := cli.ContainerAttach(ctx, con, types.ContainerAttachOptions{
		Stderr: true,
		Stdin:  true,
		Stdout: true,
		Stream: true,
		Logs:   true,
	})
	if err != nil {
		log.Fatal("Error while attaching: ", err)
	}
	defer conn.Close()
	_, err = conn.Conn.Write([]byte("stop\r"))
	if err != nil {
		log.Fatal("Error sending shutdown signal: ", err)
	}
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
		AttachStderr: false,
		AttachStdin:  true,
		AttachStdout: true,
		Cmd:          []string{"bash", "-c", "echo \"eula=true\" > eula.txt ; java -jar /minecraft/" + mod + ".jar nogui"},
		Labels:       labels,
		WorkingDir:   "/data",
		Image:        c.RepoTag + ":" + version,
		Tty:          true,
		OpenStdin:    true,
		StdinOnce:    false,
		Volumes: map[string]struct{}{
			"/data": struct{}{},
		},
	}
	port, _ := nat.NewPort("tcp", "25565")
	portMap := nat.PortMap(map[nat.Port][]nat.PortBinding{
		port: []nat.PortBinding{{
			HostIP:   "",
			HostPort: "25565",
		}},
	})
	containerHostConfig := &container.HostConfig{
		Binds:        []string{filepath.Join(c.BindBase, server) + ":/data"},
		PortBindings: portMap,
	}
	resp, err := cli.ContainerCreate(ctx, containerConfig, containerHostConfig, nil, "")
	if err != nil {
		log.Fatal("Error while creating container: ", err)
	}
	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
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

func getDockerContainer(serverName string) string {
	cli := getClient()
	ctx := context.Background()
	fills := filters.NewArgs()
	fills.Add("label", "world="+serverName)
	fills.Add("status", "running")
	con, err := cli.ContainerList(ctx, types.ContainerListOptions{
		All:     true,
		Filters: fills,
	})
	if err != nil {
		log.Fatal("Failed to find container for server with error: ", err)
	}
	if len(con) < 1 {
		log.Fatal("No containers found running with label: ", serverName)
	}
	return con[0].Names[0][1:]
}
