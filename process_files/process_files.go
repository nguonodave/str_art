package process_files

import (
	"fmt"
	"log"
	"os"
	"strings"

	"ascii_art/download_file"
	"ascii_art/file_integrity"
	"ascii_art/write_output"
)

func print_usage() {
	fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
}

func print_output_usage() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
}

func ProcessFiles(bytes []byte, read_err error, lines []string) ([]string, bool) {
	args := os.Args
	var file_arg string
	var file_path string
	var original_hash string
	banners_dir := "banners/"

	if _, banners_err := os.Stat(banners_dir); os.IsNotExist(banners_err) {
		create_banners_err := os.Mkdir("banners", 0o700)
		if create_banners_err != nil {
			log.Fatal(create_banners_err)
		}
	}

	if !write_output.ValidOutputFlag(args[1]) && len(args) > 3 {
		print_usage()
		return nil, true
	}

	if write_output.ValidOutputFlag(args[1]) {
		if len(args) == 4 {
			file_arg = args[3]
		} else if len(args) == 3 {
			file_arg = "standard"
		} else {
			print_output_usage()
			return nil, true
		}
	} else {
		if len(args) == 3 {
			file_arg = args[2]
		} else if len(args) == 2 {
			file_arg = "standard"
		}
	}

	if file_arg == "standard" {
		file_path = banners_dir + "standard.txt"
		original_hash = "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	} else if file_arg == "shadow" {
		file_path = banners_dir + "shadow.txt"
		original_hash = "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	} else if file_arg == "thinkertoy" {
		file_path = banners_dir + "thinkertoy.txt"
		original_hash = "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3"
	} else {
		fmt.Println("Invalid banner file. Use either standard, shadow, or thinkertoy")
		return nil, true
	}

	bytes, read_err = os.ReadFile(file_path)

	file_altered := file_integrity.FileAltered(bytes, original_hash)

	// download original file if it doesn't exist or is altered
	if _, file_err := os.Stat(file_path); os.IsNotExist(file_err) || file_altered {
		fmt.Println("File(s) missing or data probably altered.\nDownloading the original version...")
		fmt.Println()
		download_file.DownloadFile(file_path)
		return nil, true
	}

	if file_arg == "thinkertoy" {
		lines = strings.Split(string(bytes), "\r\n")
	} else {
		lines = strings.Split(string(bytes), "\n")
	}

	return lines, false
}
