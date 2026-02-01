package main

import (
    "net/http"
    "os"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, world!"))
}

func rebootz(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Rebooting..."))
    os.Exit(42)
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", helloWorld)
    mux.HandleFunc("/rebootz", rebootz)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    _ = http.ListenAndServe(":"+port, mux)
}
