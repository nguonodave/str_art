package main

import (
	"log"
	"os/exec"

	"ascii_art/for_terminal"
	"ascii_art/for_web"
	"ascii_art/vars"
)

func main() {
	cmd := exec.Command("gofmt", "-s", "-w", ".")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	if len(vars.Args) < 2 {
		for_web.ForWeb()
	} else {
		for_terminal.ForTerminal()
	}
}
