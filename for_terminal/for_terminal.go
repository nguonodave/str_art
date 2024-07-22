package for_terminal

import (
	"fmt"
	"strings"

	"ascii_art/ascii_art_result"
	"ascii_art/get_file_arg"
	"ascii_art/map_rune_art"
	"ascii_art/process_files"
	"ascii_art/tools"
	"ascii_art/vars"
)

func ForTerminal() {
	// call func for getting file, then pass it in the ProcessFiles function
	banner, get_file_err := get_file_arg.GetFileArg(vars.Args)
	if get_file_err {
		return
	}

	vars.Lines, vars.Read_lines_err = process_files.ProcessFiles(vars.Bytes, vars.Read_err, vars.Lines, banner)
	if vars.Read_lines_err != nil {
		fmt.Println(vars.Read_lines_err)
		return
	}

	// call func for getting string
	vars.Str_input = tools.GetStrInput()

	if vars.Str_input == "" {
		return
	}

	vars.Str_input = tools.EscapeNewline(vars.Str_input)

	tools.CheckNonPrintableChars(vars.Str_input)

	// then split the final string using '\n'.
	str_splitted := strings.Split(vars.Str_input, "\\n")

	if tools.AllSlashn(str_splitted) {
		fmt.Print(strings.Repeat("\n", len(str_splitted)-1))
		return
	}

	// map all the provided ascii printable characters to their arts.
	map_rune_art.MapRuneArt(vars.Lines, vars.Char_art_map, vars.Current_ascii_char)

	ascii_art_result.AsciiArtResult(vars.Out_file, vars.Out_file_err, str_splitted, vars.Char_art_map)
}
