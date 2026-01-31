package main

import (
	"net/http"
	"os"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorld)

	_ = http.ListenAndServe(":"+port, mux)
}
