package main

import (
	"github.com/Skarlso/miner/commands"
	"github.com/Skarlso/miner/config"
	cmd "github.com/Yitsushi/go-commander"
)

func init() {
	config.Init()
}

func main() {
	registry := cmd.NewCommandRegistry()
	registry.Register(commands.NewSetup)
	registry.Register(commands.NewStart)
	registry.Register(commands.NewStop)
	registry.Register(commands.NewAttach)
	registry.Register(commands.NewBackup)
	registry.Register(commands.NewStatus)
	registry.Register(commands.NewDelete)
	registry.Execute()
}
