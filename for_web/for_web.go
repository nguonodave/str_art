package for_web

import (
	"fmt"
	// "html/template"
	"net/http"

	"ascii_art/art_server"
	"ascii_art/tools"
	"ascii_art/vars"
)

func ForWeb() {
	// // asigning the template variables to their respective template files
	// vars.Home_template, _ = template.ParseFiles("templates/home.html", "templates/nav.html")
	// vars.Art_template, _ = template.ParseFiles("templates/ascii-art.html", "templates/nav.html")
	// vars.Template_404, _ = template.ParseFiles("templates/404.html", "templates/nav.html")
	vars.All_templates, _ = vars.All_templates.ParseGlob("templates/*.html")

	http.HandleFunc("/", tools.HomeOr404Page)
	http.HandleFunc("/ascii-art", art_server.Art)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	fmt.Println("Listening on :8001...")

	// starts the HTTP server and listens for incoming requests from port 8001
	http.ListenAndServe(":8001", nil)
}
