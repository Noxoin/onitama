package main

import (
	"fmt"
	"net/http"

	_ "github.com/noxoin/onitama/server"
	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/_ah/health", healthCheckHandler)
	fmt.Println("Hello, World!")
	appengine.Main()
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
