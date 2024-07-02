package main

import (
	"fmt"
	"os/exec"
	"sync"
)

func gitPull(wg *sync.WaitGroup, dir string, results chan<- string) {
	defer wg.Done()

	cmd := exec.Command("git", "pull")
	cmd.Dir = dir

	err := cmd.Run()
	if err != nil {
		results <- fmt.Sprintf("Failed to run 'git pull' in %s: %s", dir, err)
	} else {
		results <- fmt.Sprintf("'git pull' succeeded in %s", dir)
	}
}

func main() {
	dirs := []string{
		"/home/u882585955/domains/sigrecette.com/bassar1",
		"/home/u882585955/domains/sigrecette.com/blitta1",
		"/home/u882585955/domains/sigrecette.com/doufelgounou1",
		"/home/u882585955/domains/sigrecette.com/oti1",
		"/home/u882585955/domains/sigrecette.com/otisud1",
		"/home/u882585955/domains/sigrecette.com/dankpen1",
		"/home/u882585955/domains/sigrecette.com/mo2",
		"/home/u882585955/domains/sigrecette.com/sotouboua2",
	}

	var wg sync.WaitGroup
	results := make(chan string, len(dirs))

	for _, dir := range dirs {
		wg.Add(1)
		go gitPull(&wg, dir, results)
	}

	wg.Wait()
	close(results)

	for result := range results {
		fmt.Println(result)
	}
}
