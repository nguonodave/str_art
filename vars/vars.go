package vars

import (
	"html/template"
	"os"
)

var (
	Bytes              []byte // bytes are the contents of the banner file.
	Read_err           error
	Read_lines_err     bool     // returns true if there was an error while processing files.
	Lines              []string // holds every line of the banner files.
	Char_art_map       = make(map[rune][]string)
	Current_ascii_char = 32
	Str_input          string
	Args               = os.Args
	Out_file           *os.File // the file where the string arts are written to.
	Out_file_err       error
	// Home_template, Art_template, Template_404 *template.Template
	All_templates *template.Template
	Pages         = []string{"ascii-art"}
)
