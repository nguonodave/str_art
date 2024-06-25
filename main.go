package main

import (
	"fmt"
	"os"
	"strings"

	"ascii_art/map_rune_art"
	"ascii_art/print_ascii"
	"ascii_art/process_files"
)

var (
	bytes              []byte
	read_err           error
	read_lines_err     bool
	lines              []string
	char_art_map       = make(map[rune][]string)
	current_ascii_char = 32
)

func allSlashN(str_splitted []string) bool {
	for _, str := range str_splitted {
		if len(str) != 0 {
			return false
		}
	}
	return true
}

func main() {
	// providing an error message if string is not provided (no enough arguments)
	if len(os.Args) < 2 {
		fmt.Println("No enough arguments!! Please provide the text to be printed :)")
		return
	}

	lines, read_lines_err = process_files.ProcessFiles(bytes, read_err, lines)
	if read_lines_err {
		return
	}

	// get the string form the command line arguments
	str_input := os.Args[1]

	if str_input == "" {
		return
	}

	for _, char := range str_input {
		if char > 126 || char < 32 {
			fmt.Println("It's likely your string has non-ascii printable characters. Please provide only ascii printable characters in your string :)")
			return
		}
	}

	// then split the final string using '\n'
	str_splitted := strings.Split(str_input, "\\n")

	if allSlashN(str_splitted) {
		fmt.Print(strings.Repeat("\n", len(str_splitted)-1))
		return
	}

	map_rune_art.MapRuneArt(lines, char_art_map, current_ascii_char)

	for _, str_item := range str_splitted {
		// if str_item == "" && i < len(str_splitted)-1 {
		// 	fmt.Println()
		// }
		if str_item != "" {
			print_ascii.PrintAscii(char_art_map, str_item)
		} else {
			fmt.Println()
		}
	}
}
