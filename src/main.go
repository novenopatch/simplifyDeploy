package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

type Config struct {
	BaseDir     string   `json:"basedir"`
	Directories []string `json:"directories"`
	Commands    []struct {
		Name            string   `json:"name"`
		Command         []string `json:"command"`
		ForceProduction bool     `json:"forceProduction"`
	} `json:"commands"`
}

func directoryExists(dir string) bool {
	_, err := os.Stat(dir)
	return !os.IsNotExist(err)
}

func runCommand(wg *sync.WaitGroup, dir string, command []string, forceProduction bool, results chan<- string) {
	defer wg.Done()

	if !directoryExists(dir) {
		results <- fmt.Sprintf("Directory %s does not exist", dir)
		return
	}

	if forceProduction && os.Getenv("APP_ENV") == "production" {
		command = append(command, "--force")
	}

	cmd := exec.Command(command[0], command[1:]...)
	cmd.Dir = dir

	err := cmd.Run()
	if err != nil {
		results <- fmt.Sprintf("Failed to run '%s' in %s: %s", command[0], dir, err)
	} else {
		results <- fmt.Sprintf("'%s' succeeded in %s", command[0], dir)
	}
}

func main() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	var config Config
	if err := json.Unmarshal(file, &config); err != nil {
		fmt.Println("Error parsing config file:", err)
		return
	}

	var wg sync.WaitGroup
	results := make(chan string, len(config.Directories)*len(config.Commands))

	for _, dir := range config.Directories {
		fullDir := filepath.Join(config.BaseDir, dir)

		if !directoryExists(fullDir) {
			fmt.Printf("Directory %s does not exist, skipping...\n", fullDir)
			continue
		}

		for _, cmd := range config.Commands {
			wg.Add(1)
			go runCommand(&wg, fullDir, cmd.Command, cmd.ForceProduction, results)
		}
	}

	wg.Wait()
	close(results)

	for result := range results {
		fmt.Println(result)
	}
}
