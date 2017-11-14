package main

import (
	"fmt"
	"net/http"

	_ "github.com/noxoin/onitama/server"
)

func main() {
	http.HandleFunc("/_ah/health", healthCheckHandler)
	fmt.Println("Hello, World!")
	http.ListenAndServe(":8080", nil)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
