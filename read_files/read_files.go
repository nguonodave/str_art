package read_files

import (
	"fmt"
	"log"
	"os"
	"strings"

	"ascii_art/file_integrity"
)

func ReadFiles(bytes []byte, read_err error, lines []string) ([]string, bool) {
	std_chmod_error := os.Chmod("banners/standard.txt", 0o400)
	shdw_chmod_error := os.Chmod("banners/shadow.txt", 0o400)
	thnky_chmod_error := os.Chmod("banners/thinkertoy.txt", 0o400)

	if std_chmod_error != nil {
		log.Fatal(std_chmod_error)
	} else if shdw_chmod_error != nil {
		log.Fatal(shdw_chmod_error)
	} else if thnky_chmod_error != nil {
		log.Fatal(thnky_chmod_error)
	}

	var file_altered bool

	if len(os.Args) == 2 || (len(os.Args) == 3 && os.Args[2] == "standard") {
		bytes, read_err = os.ReadFile("banners/standard.txt")
		file_altered = file_integrity.FileAltered(bytes, "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf")
	} else if os.Args[2] == "shadow" {
		bytes, read_err = os.ReadFile("banners/shadow.txt")
		file_altered = file_integrity.FileAltered(bytes, "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73")
	} else if os.Args[2] == "thinkertoy" {
		bytes, read_err = os.ReadFile("banners/thinkertoy.txt")
		file_altered = file_integrity.FileAltered(bytes, "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3")
	} else {
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		return nil, true
	}

	if read_err != nil {
		fmt.Println(read_err)
	}

	// print error message if a file is altered
	if file_altered {
		fmt.Println("Seems like you altered the data file(s). Please download them from here")
		return nil, true
	}

	if len(os.Args) == 3 && os.Args[2] == "thinkertoy" {
		lines = strings.Split(string(bytes), "\r\n")
	} else {
		lines = strings.Split(string(bytes), "\n")
	}

	return lines, false
}
