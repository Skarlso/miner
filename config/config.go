package config

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
)

// REGION to operate in.
var REGION string

var configPath string

// WAITFREQUENCY global wait frequency default.
var WAITFREQUENCY = 1

// SPINNER is the index of which spinner to use. Defaults to 7.
var SPINNER int

// Path retrieves the main configuration path.
func Path() string {
	// Get configuration path
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
	}
	return filepath.Join(usr.HomeDir, ".config", "go-furnace")
}

func init() {
	configPath = Path()
	REGION = os.Getenv("FURNACE_REGION")
	spinner := os.Getenv("FURNACE_SPINNER")
	if len(spinner) < 1 {
		SPINNER = 7
	} else {
		SPINNER, _ = strconv.Atoi(spinner)
	}
	if len(REGION) < 1 {
		log.Fatal("Please define a region to operate in with FURNACE_REGION exp: eu-central-1.")
	}
}
