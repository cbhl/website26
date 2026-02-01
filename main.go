package main

import (
    "html/template"
    "net/http"
    "net/http/pprof"
    "os"
    "time"
)

var homeTemplate = template.Must(template.ParseFiles("templates/home.html"))

func home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    if err := homeTemplate.Execute(w, nil); err != nil {
        http.Error(w, "Failed to render page", http.StatusInternalServerError)
        return
    }
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

func firstExistingDir(paths ...string) string {
    for _, path := range paths {
        if info, err := os.Stat(path); err == nil && info.IsDir() {
            return path
        }
    }
    return ""
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/rebootz", rebootz)
    registerPprof(mux)
    staticRoot := firstExistingDir("static", "/home/public")
    staticHandler := http.NotFoundHandler()
    if staticRoot != "" {
        staticHandler = http.FileServer(http.Dir(staticRoot))
    }
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path == "/" || r.URL.Path == "/index.html" {
            home(w, r)
            return
        }
        staticHandler.ServeHTTP(w, r)
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    _ = http.ListenAndServe(":"+port, mux)
}
