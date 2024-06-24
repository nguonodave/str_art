package read_files

import (
	"fmt"
	"log"
	"os"
	"strings"

	"ascii_art/file_integrity"
)

func ReadFiles(bytes []byte, read_err error, lines []string) ([]string, bool) {
	var file_path string
	var original_hash string
	banners_dir := "banners/"

	if len(os.Args) == 2 || (len(os.Args) == 3 && os.Args[2] == "standard") {
		file_path = banners_dir + "standard.txt"
		original_hash = "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	} else if os.Args[2] == "shadow" {
		file_path = banners_dir + "shadow.txt"
		original_hash = "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	} else if os.Args[2] == "thinkertoy" {
		file_path = banners_dir + "thinkertoy.txt"
		original_hash = "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3"
	} else {
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		return nil, true
	}

	if chmod_err := os.Chmod(file_path, 0o400); chmod_err != nil {
		log.Fatal("Error changing file permissions", chmod_err)
	}

	bytes, read_err = os.ReadFile(file_path)
	if read_err != nil {
		fmt.Println("Error reading file", read_err)
		return nil, true
	}

	file_altered := file_integrity.FileAltered(bytes, original_hash)

	// print error message if a file is altered
	if file_altered {
		fmt.Println("File(s) data probably altered. Please download them from the following link\nhttps://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art")
		return nil, true
	}

	if len(os.Args) == 3 && os.Args[2] == "thinkertoy" {
		lines = strings.Split(string(bytes), "\r\n")
	} else {
		lines = strings.Split(string(bytes), "\n")
	}

	return lines, false
}
