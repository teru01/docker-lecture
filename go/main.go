package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	server := http.Server {
		Addr: "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
