package server

import (
	"html/template"
	"net/http"
)

var (
	projectId = "api-project-377888563324"
	tmpls     map[string]*template.Template
)

func init() {
	fs := http.FileServer(http.Dir("client"))
	http.Handle("/js/", fs)
	http.Handle("/css/", fs)
	http.HandleFunc("/", mainHandler)
	tmpls = make(map[string]*template.Template)
	tmpls["index"] = template.Must(template.ParseFiles(
		"client/templates/index.html", "client/templates/base.html"))
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	tmpls["index"].ExecuteTemplate(w, "base", "Hello Onitama Server")
}
