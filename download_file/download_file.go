package download_file

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// DownloadFile downloads the banner file contents and store them in their respective files.
func DownloadFile(file_path string) {
	url := ""
	switch file_path {
	case "banners/standard.txt":
		url = "https://learn.zone01kisumu.ke/git/root/public/raw/branch/master/subjects/ascii-art/standard.txt"
	case "banners/shadow.txt":
		url = "https://learn.zone01kisumu.ke/git/root/public/raw/branch/master/subjects/ascii-art/shadow.txt"
	case "banners/thinkertoy.txt":
		url = "https://learn.zone01kisumu.ke/git/root/public/raw/branch/master/subjects/ascii-art/thinkertoy.txt"
	default:
		log.Fatalf("%s is an invalid file", file_path)
	}

	// download the body (contents)
	resp, resp_err := http.Get(url)
	if resp_err != nil {
		log.Fatalf("Check your connection. Error getting content from:\n%s", url)
	}
	defer resp.Body.Close()

	// read the contents
	body, body_err := io.ReadAll(resp.Body)
	if body_err != nil {
		log.Fatal("Error reading response body")
	}

	// write the contents to a file.
	write_err := os.WriteFile(file_path, body, 0o777)
	if write_err != nil {
		log.Fatal(write_err)
	}

	fmt.Println("File downloaded successfully. Please re-run the program.")
}
