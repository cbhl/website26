package main

import (
    "net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, world!"))
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", helloWorld)

    _ = http.ListenAndServe(":8080", mux)
}
