package art_server

import (
	"net/http"
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

	tools.CheckNonPrintableChars(str)

	vars.Lines, vars.Read_lines_err = process_files.ProcessFiles(vars.Bytes, vars.Read_err, vars.Lines, banner)
	if vars.Read_lines_err {
		art = "The file " + banner + ".txt" + " is missing, or it's data has probably been altered.\nA newer version has been downloaded successfully. Please re-run the program"
		vars.All_templates.ExecuteTemplate(w, "ascii-art.html", art)
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
