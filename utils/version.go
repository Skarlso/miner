package utils

import (
	"log"
	"path/filepath"

	"io/ioutil"
	"os"

	"github.com/Skarlso/miner/config"
	"github.com/fatih/color"
)

// SetVersion set version for server
func SetVersion(version, server string) {
	filename := filepath.Join(config.Path(), server+".txt")
	c := config.Config{}
	c.Unmarshal()
	cyan := color.New(color.FgCyan).SprintFunc()
	data := []byte(version)
	if err := ioutil.WriteFile(filename, data, os.ModePerm); err != nil {
		log.Fatalln("Error creating version file for server: ", err)
	}
	log.Printf("Setting version %s for server %s", cyan(version), cyan(server))
}

// GetVersion get a version for a server
func GetVersion(server string) string {
	red := color.New(color.FgRed).SprintFunc()
	filename := filepath.Join(config.Path(), server+".txt")
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("File for server doesn't exists. Filename was: ", red(filename))
		// log.Fatal("Error opening config file: ", err)
	}
	return string(data)
}
