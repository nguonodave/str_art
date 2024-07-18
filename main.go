package main

import (
	"fmt"
	"log"
	"os/exec"

	"ascii_art/for_terminal"
	"ascii_art/vars"
)

func main() {
	cmd := exec.Command("gofmt", "-s", "-w", ".")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	if len(vars.Args) < 2 {
		// later, this is where the web functionality will run
		fmt.Println("No enough arguments!! Please provide the text to be printed :)")
		return
	} else {
		for_terminal.ForTerminal()
	}
}
