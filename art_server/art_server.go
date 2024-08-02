package art_server

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"ascii_art/art_string"
	"ascii_art/map_rune_art"
	"ascii_art/process_files"
	"ascii_art/tools"
	"ascii_art/vars"
)

func Art(w http.ResponseWriter, r *http.Request) {
	str := r.FormValue("str")
	banner := strings.ToLower(r.FormValue("banner"))
	var art string

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		vars.All_templates.ExecuteTemplate(w, "file-err.html", "Method not allowed. Submit your request from the provided from")
		return
	}

	if non_print_error := tools.CheckNonPrintableChars(str); non_print_error != nil {
		w.WriteHeader(http.StatusBadRequest)
		vars.All_templates.ExecuteTemplate(w, "file-err.html", non_print_error)
		return
	}

	vars.Lines, vars.Read_lines_err = process_files.ProcessFiles(vars.Bytes, vars.Read_err, vars.Lines, banner)
	if vars.Read_lines_err != nil {
		if _, file_err := os.Stat("banners/" + banner + ".txt"); os.IsNotExist(file_err) {
			fmt.Println("not found")
			w.WriteHeader(http.StatusNotFound)
		} else {
			fmt.Println("found")
			w.WriteHeader(http.StatusInternalServerError)
		}
		file_err := vars.Read_lines_err
		vars.All_templates.ExecuteTemplate(w, "file-err.html", file_err)
		return
	}

	// map all the provided ascii printable characters to their arts.
	map_rune_art.MapRuneArt(vars.Lines, vars.Char_art_map, vars.Current_ascii_char)

	str_splitted := strings.Split(str, "\\n")

	if tools.AllSlashn(str_splitted) {
		art = strings.Repeat("\n", len(str_splitted)-1)
		vars.All_templates.ExecuteTemplate(w, "ascii-art.html", art)
		return
	}

	// display results based on the conditions.
	for _, str_item := range str_splitted {
		art += art_string.ArtString(vars.Char_art_map, str_item)
	}

	vars.All_templates.ExecuteTemplate(w, "ascii-art.html", art)
}
