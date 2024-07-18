package ascii_art_result

import (
	"fmt"
	"log"
	"os"

	"ascii_art/print_ascii"
	"ascii_art/write_output"
)

var args = os.Args

func AsciiArtResult(out_file *os.File, out_file_err error, str_splitted []string, char_art_map map[rune][]string) {
	// only create the file if the output flag is valid.
	if write_output.ValidOutputFlag(args[1]) {
		if out_file, out_file_err = os.Create(args[1][9:]); out_file_err != nil {
			log.Fatal(out_file_err)
		}
	}

	// display results based on the conditions.
	for _, str_item := range str_splitted {
		if write_output.ValidOutputFlag(args[1]) {
			write_output.WriteOutput(char_art_map, str_item, out_file)
		} else {
			if str_item != "" {
				print_ascii.PrintAscii(char_art_map, str_item)
			} else {
				fmt.Println()
			}
		}
	}
}
