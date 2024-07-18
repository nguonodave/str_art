package tools

import (
	"fmt"
	"net/http"

	"ascii_art/vars"
	"ascii_art/write_output"
)

// all_slashn returns true if all string arguments are new line characters.
// Otherwise it returns false.
func AllSlashn(str_splitted []string) bool {
	for _, str := range str_splitted {
		if len(str) != 0 {
			return false
		}
	}
	return true
}

func EscapeNewline(s string) string {
	new_str := ""

	for _, char := range s {
		if char == '\n' {
			new_str += "\\n"
		} else {
			new_str += string(char)
		}
	}

	return new_str
}

func CheckNonPrintableChars(s string) {
	// check for non-printable ascii characters.
	for _, char := range s {
		if char > 126 || char < 32 {
			fmt.Println("Non-ascii printable characters encountered. Please provide only ascii printable characters in your string :)")
			return
		}
	}
}

func GetStrInput() string {
	// get the string form the command line arguments, based on the arguments.
	if write_output.ValidOutputFlag(vars.Args[1]) {
		vars.Str_input = vars.Args[2]
	} else {
		vars.Str_input = vars.Args[1]
	}

	return vars.Str_input
}

func PageNotFound(page string) bool {
	for _, v := range vars.Pages {
		if v == page {
			return false
		}
	}

	return true
}

func HomeOr404Page(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

	if len(path) > 0 && PageNotFound(path) {
		vars.Template_404.Execute(w, nil)
	} else {
		vars.Home_template.ExecuteTemplate(w, "home.html", nil)
	}
}
