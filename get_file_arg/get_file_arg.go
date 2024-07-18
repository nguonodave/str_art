package get_file_arg

import (
	"strings"

	"ascii_art/usage"
	"ascii_art/write_output"
)

func GetFileArg(args []string) (string, bool) {
	file_arg := ""

	// get the file based on the accepted provided arguments.
	if write_output.ValidOutputFlag(args[1]) {
		if len(args) == 4 {
			file_arg = strings.ToLower(args[3])
		} else if len(args) == 3 {
			file_arg = "standard"
		} else {
			usage.PrintOutputUsage()
			return "", true
		}
	} else {
		if len(args) == 3 {
			file_arg = strings.ToLower(args[2])
		} else if len(args) == 2 {
			file_arg = "standard"
		}
	}

	return file_arg, false
}
