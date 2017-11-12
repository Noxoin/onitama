package server

import (
	"net/http"
)

func init() {
	http.HandleFunc("/", mainHandler)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Onitama Server")
}
