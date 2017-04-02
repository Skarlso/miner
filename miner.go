package main

import (
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/Skarlso/miner/commands"
	cmd "github.com/Yitsushi/go-commander"
)

func init() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(usr.HomeDir, ".config", "miner")); err != nil {
		if os.IsNotExist(err) {
			log.Fatalln("Please create a 'miner' folder under .config.")
		}
	}
}

func main() {
	registry := cmd.NewCommandRegistry()
	registry.Register(commands.NewCreate)
	registry.Register(commands.NewList)
	registry.Register(commands.NewStart)
	registry.Register(commands.NewStop)
	registry.Register(commands.NewAttach)
	registry.Register(commands.NewBackup)
	registry.Execute()
}
