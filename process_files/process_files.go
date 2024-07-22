package process_files

import (
	"errors"
	"log"
	"os"
	"strings"

	"ascii_art/download_file"
	"ascii_art/file_integrity"
)

// ProcessFiles reads and returns the lines of the valid banner files.
// Returns true if an error is encountered.
func ProcessFiles(bytes []byte, read_err error, lines []string, banner string) ([]string, error) {
	// args := os.Args
	var file_path string
	var original_hash string
	var invalid_file_error error
	banners_dir := "banners/"

	// create the banners directory if it does not exist.
	if _, banners_err := os.Stat(banners_dir); os.IsNotExist(banners_err) {
		create_banners_err := os.Mkdir("banners", 0o700)
		if create_banners_err != nil {
			log.Fatal(create_banners_err)
		}
	}

	// define a default file incase a banner is not selected
	if banner == "" {
		banner = "standard"
	}

	// define files and their hash values.
	if banner == "standard" {
		file_path = banners_dir + "standard.txt"
		original_hash = "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	} else if banner == "shadow" {
		file_path = banners_dir + "shadow.txt"
		original_hash = "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	} else if banner == "thinkertoy" {
		file_path = banners_dir + "thinkertoy.txt"
		original_hash = "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3"
	} else {
		// fmt.Printf("\033[31m"+"%s"+"\033[0m"+" is an unsupported banner file. Use either standard, shadow, or thinkertoy\n", banner)
		invalid_file_error = errors.New("\033[31m" + banner + "\033[0m" + " is an unsupported banner file. Use either standard, shadow, or thinkertoy")
		return nil, invalid_file_error
	}

	bytes, read_err = os.ReadFile(file_path)

	file_altered := file_integrity.FileAltered(bytes, original_hash)

	// download original file if it doesn't exist or is altered.
	if _, file_err := os.Stat(file_path); os.IsNotExist(file_err) || file_altered {
		// fmt.Printf("The file "+"\033[31m"+"%s.txt"+"\033[0m"+" is missing, or it's data has probably been altered.\nDownloading the original version...\n", banner)
		banner_error := errors.New("The file " + banner + ".txt" + " is missing, or it's data has probably been altered.\nIf you have internet connection, then the file has been downloaded successfully. Please re-run the program.")
		_ = download_file.DownloadFile(file_path)
		return nil, banner_error
	}

	if banner == "thinkertoy" {
		lines = strings.Split(string(bytes), "\r\n")
	} else {
		lines = strings.Split(string(bytes), "\n")
	}

	return lines, nil
}
