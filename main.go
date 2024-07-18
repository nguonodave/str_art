package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"ascii_art/for_terminal"
)

var (
	bytes              []byte // bytes are the contents of the banner file.
	read_err           error
	read_lines_err     bool     // returns true if there was an error while processing files.
	lines              []string // holds every line of the banner files.
	char_art_map       = make(map[rune][]string)
	current_ascii_char = 32
	str_input          string
	args               = os.Args
	out_file           *os.File // the file where the string arts are written to.
	out_file_err       error
)

func main() {
	cmd := exec.Command("gofmt", "-s", "-w", ".")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	if len(args) < 2 {
		// later, this is where the web functionality will run
		fmt.Println("No enough arguments!! Please provide the text to be printed :)")
		return
	} else {
		for_terminal.ForTerminal()
	}
}
