package tools

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

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

func CheckNonPrintableChars(s string) error {
	// check for non-printable ascii characters.
	for _, char := range s {
		if char > 126 || char < 32 {
			// fmt.Println("Non-ascii printable characters encountered. Please provide only ascii printable characters in your string :)")
			non_print_error := errors.New("Found non-ascii printable characters in your string. Please provide only ascii printable characters.")
			return non_print_error
		}
	}
	return nil
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

func TemplateNotFound(temp string) bool {
	_, err := os.Stat(temp)
	return os.IsNotExist(err)
}

func LinkedTemplateNotFound() bool {
	for _, temp := range vars.Linked_templates {
		_, err := os.Stat(vars.Template_dir + temp)
		if os.IsNotExist(err) {
			return true
		}
	}

	return false
}

func TmpltNotFoundMsg(w http.ResponseWriter, temp string) {
	log.Println("The file " + temp + " or some template in it is missing")
	fmt.Fprint(w, "404: Error displaying this page, please try again later.")
}

func HomeOr404Page(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

	if TemplateNotFound(vars.Template_dir+"home.html") || LinkedTemplateNotFound() {
		w.WriteHeader(http.StatusNotFound)
		TmpltNotFoundMsg(w, "home.html")
	} else if len(path) > 0 && PageNotFound(path) {
		info_404 := "Oopsie! That page is not available"
		w.WriteHeader(http.StatusNotFound)

		if TemplateNotFound(vars.Template_dir + "404.html") {
			TmpltNotFoundMsg(w, "404.html")
		} else {
			vars.All_templates.ExecuteTemplate(w, "404.html", info_404)
		}

	} else {
		vars.All_templates.ExecuteTemplate(w, "home.html", nil)
	}
}
