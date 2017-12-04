package server

import (
	"html/template"
	"net/http"
)

var (
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
	if r.URL.Path != "/" {
		gameHandler(w, r)
		return
	}
	if r.Method == http.MethodPost {
		createGameHandler(w, r)
		return
	}
	tmpls["index"].ExecuteTemplate(w, "base", nil)
}
