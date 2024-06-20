package read_files

import (
	"fmt"
	"os"
	"strings"
)

func ReadFiles(bytes []byte, read_err error, lines []string) []string {
	if len(os.Args) == 2 || (len(os.Args) == 3 && os.Args[2] == "standard") {
		bytes, read_err = os.ReadFile("banners/standard.txt")
	} else if os.Args[2] == "shadow" {
		bytes, read_err = os.ReadFile("banners/shadow.txt")
	} else if os.Args[2] == "thinkertoy" {
		bytes, read_err = os.ReadFile("banners/thinkertoy.txt")
	} else {
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		// return
		os.Exit(1)
	}

	if read_err != nil {
		fmt.Println(read_err)
	}

	if len(os.Args) == 3 && os.Args[2] == "thinkertoy" {
		lines = strings.Split(string(bytes), "\r\n")
	} else {
		lines = strings.Split(string(bytes), "\n")
	}

	return lines
}
