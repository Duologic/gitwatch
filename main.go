package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"
)

type Config struct {
	Repository string
	Branch     string
	Dir        string
}

func GetJsonConfig(configFile string, config interface{}) error {
	jsonFile, err := os.Open(configFile)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	return json.Unmarshal(bytes, &config)
}

func main() {
	var configfile string
	var repository string
	var branch string
	var dir string
	flag.StringVar(&configfile, "config", "", "path to the config file")
	flag.StringVar(&repository, "repository", "", "git repository (http/https/ssh)")
	flag.StringVar(&branch, "branch", "master", "remote branch name to watch")
	flag.StringVar(&dir, "dir", "/repo", "remote branch name to watch")
	flag.Parse()

	var config Config
	if configfile != "" {
		if err := GetJsonConfig(configfile, &config); err != nil {
			log.Fatal(err)
		}
	}
	if repository != "" {
		config.Repository = repository
	}
	if branch != "" {
		config.Branch = branch
	}
	if dir != "" {
		config.Dir = dir
	}
	if config.Repository == "" {
		log.Fatal("Repository needs to be set")
	}

	if err := os.MkdirAll(config.Dir, 0755); err != nil {
		log.Fatal(err)
	}

	if err := runCommand(os.TempDir(), "git", "clone", "--depth=1", config.Repository, config.Dir); err != nil {
		log.Fatalf("error running git clone: %v", err)
	}

	done := make(chan struct{})

	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		select {
		case <-done:
			ticker.Stop()
		default:
			if err := runCommand(config.Dir, "git", "clean", "-xdf"); err != nil {
				log.Fatalf("error running git clean: %v", err)
			}

			if err := runCommand(config.Dir, "git", "fetch", "origin", config.Branch); err != nil {
				log.Fatalf("error running git fetch: %v", err)
			}

			if err := runCommand(config.Dir, "git", "checkout", "-f", fmt.Sprintf("origin/%s", config.Branch)); err != nil {
				log.Fatalf("error running git checkout: %v", err)
			}
		}
	}
}

func runCommand(dir, command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
