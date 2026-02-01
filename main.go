package main

import (
    "net/http"
    "os"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, world!"))
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", helloWorld)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    _ = http.ListenAndServe(":"+port, mux)
}
