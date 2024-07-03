package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"ascii_art/map_rune_art"
	"ascii_art/print_ascii"
	"ascii_art/process_files"
	"ascii_art/write_output"
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
	file               *os.File // the file where the string arts are written to.
	out_file_err       error
)

// all_slashn returns true if all string arguments are new line characters.
// Otherwise it returns false.
func all_slashn(str_splitted []string) bool {
	for _, str := range str_splitted {
		if len(str) != 0 {
			return false
		}
	}
	return true
}

func main() {
	cmd := exec.Command("gofmt", "-s", "-w", ".")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	if len(args) < 2 {
		fmt.Println("No enough arguments!! Please provide the text to be printed :)")
		return
	}

	lines, read_lines_err = process_files.ProcessFiles(bytes, read_err, lines)
	if read_lines_err {
		return
	}

	// get the string form the command line arguments, based on the arguments.
	if write_output.ValidOutputFlag(args[1]) {
		str_input = args[2]
	} else {
		str_input = args[1]
	}

	if str_input == "" {
		return
	}

	// check for non-printable ascii characters.
	for _, char := range str_input {
		if char > 126 || char < 32 {
			fmt.Println("Non-ascii printable characters encountered. Please provide only ascii printable characters in your string :)")
			return
		}
	}

	// then split the final string using '\n'.
	str_splitted := strings.Split(str_input, "\\n")

	if all_slashn(str_splitted) {
		fmt.Print(strings.Repeat("\n", len(str_splitted)-1))
		return
	}

	// map all the provided ascii printable characters to their arts.
	map_rune_art.MapRuneArt(lines, char_art_map, current_ascii_char)

	// only create the file if the output flag is valid.
	if write_output.ValidOutputFlag(args[1]) {
		if file, out_file_err = os.Create(args[1][9:]); out_file_err != nil {
			log.Fatal(out_file_err)
		}
	}

	// display results based on the conditions.
	for _, str_item := range str_splitted {
		if write_output.ValidOutputFlag(args[1]) {
			write_output.WriteOutput(char_art_map, str_item, file)
		} else {
			if str_item != "" {
				print_ascii.PrintAscii(char_art_map, str_item)
			} else {
				fmt.Println()
			}
		}
	}
}
