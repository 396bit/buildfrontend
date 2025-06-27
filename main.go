package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	log.SetPrefix("buildfrontend: ")
	log.SetFlags(0)

	_, err := os.Stat("./frontend/package.json")
	if err != nil {
		log.Fatalf("no 'frontend' folder found (%s)", err)
	}

	abs, _ := filepath.Abs("./frontend")
	log.Printf("cd %s", abs)

	if err := os.Chdir("./frontend"); err != nil {
		log.Fatalf("unable to cd into ./frontend: %s", err)
	}

	cmd := exec.Command(
		"bun",
		"run",
		"build",
	)

	log.Printf("execute: %s", strings.Join(cmd.Args, " "))

	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("builder failed: %s", err)
	}
	for l := range strings.Lines(string(output)) {
		log.Print(l)
	}

	if !cmd.ProcessState.Success() {
		log.Fatalf("builder exited unexpectedly (%d)", cmd.ProcessState.ExitCode())
	}

	log.Print("successfully built")
}
