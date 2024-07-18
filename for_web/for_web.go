package for_web

import (
	"ascii_art/art_server"
	"ascii_art/tools"
	"ascii_art/vars"
	"fmt"
	"html/template"
	"net/http"
)

func ForWeb() {
	// asigning the template variables to their respective template files
	vars.Home_template, _ = template.ParseFiles("templates/home.html")
	vars.Art_template, _ = template.ParseFiles("templates/ascii-art.html")
	vars.Template_404, _ = template.ParseFiles("templates/404.html")

	http.HandleFunc("/", tools.HomeOr404Page)
	http.HandleFunc("/ascii-art", art_server.Art)

	fmt.Println("Listening on :8001...")

	// starts the HTTP server and listens for incoming requests from port 8001
	http.ListenAndServe(":8001", nil)
}
