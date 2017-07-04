package config

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"bytes"

	"gopkg.in/yaml.v2"
)

// Path retrieves the main configuration path.
func Path() string {
	// Get configuration path
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
	}
	return filepath.Join(usr.HomeDir, ".config", "miner")
}

const (
	// CRAFTBUKKIT use craftbukkit mod
	CRAFTBUKKIT = iota
	// FORGE use forge mod
	FORGE
)

// Config is a struct of default configuration values
type Config struct {
	Spinner    int    `yaml:"spinner"`
	Bucket     string `yaml:"bucket"`
	Name       string `yaml:"name"`
	RepoTag    string `yaml:"repoTag"`
	BindBase   string `yaml:"bindBase"`
	AwsProfile string `yaml:"awsProfile"`
}

// Unmarshal config values
func (c *Config) Unmarshal() {
	data, err := ioutil.ReadFile(getConfigFilePath())
	if err != nil {
		log.Fatal("Error opening config file: ", err)
	}
	err = yaml.Unmarshal(data, c)
	if err != nil {
		log.Fatal("Error unmarshalling config file: ", err)
	}
}

// GetMod returns the modding option to use with a new server
func GetMod() int {
	mod := os.Getenv("MINER_FORGE_MOD")
	if len(mod) == 0 {
		return CRAFTBUKKIT
	}
	return FORGE
}

func getConfigFilePath() string {
	configPath := Path()
	return filepath.Join(configPath, "miner_config.yaml")
}

// Init initialize the base configuration entry
func Init() {
	configPath := Path()
	if _, err := os.Stat(configPath); err != nil {
		if os.IsNotExist(err) {
			log.Println("Detecting first run. Creating default config and user folder under: ", configPath)
			if err := os.MkdirAll(configPath, os.ModeDir|os.ModePerm); err != nil {
				log.Fatalf("Failed to create base folder at '%s' with error: '%s'", configPath, err.Error())
			}
			bindBase := configPath
			defaultConfig := []byte(`spinner: 7
bucket: my-minecraft-backup-bucket
name: miner_server
repoTag: skarlso/minecraft
bindBase: <replace>
awsProfile: default`)
			defaultConfig = bytes.Replace(defaultConfig, []byte("<replace>"), []byte(bindBase), -1)
			ioutil.WriteFile(filepath.Join(configPath, "miner_config.yaml"), defaultConfig, os.ModePerm)
		}
	}
}
