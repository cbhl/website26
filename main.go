package main

import (
    "net/http"
    "net/http/pprof"
    "os"
    "time"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, world!"))
}

func rebootz(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Rebooting..."))
    go func() {
        time.Sleep(time.Second)
        os.Exit(42)
    }()
}

func registerPprof(mux *http.ServeMux) {
    mux.HandleFunc("/debug/pprof/", pprof.Index)
    mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
    mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
    mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
    mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", helloWorld)
    mux.HandleFunc("/rebootz", rebootz)
    registerPprof(mux)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    _ = http.ListenAndServe(":"+port, mux)
}
