package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    "music-playlist-generator/handlers"
    "music-playlist-generator/middleware"
)

func main() {
    // Load environment variables
    err := godotenv.Load()
    if err != nil {
        log.Printf("No .env file found: %v", err)
    }

    // Initialize router
    r := mux.NewRouter()

    // Add middleware
    r.Use(middleware.LoggingMiddleware)
    r.Use(middleware.JSONContentTypeMiddleware)

    // Serve static files
    fs := http.FileServer(http.Dir("static"))
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

    // API routes
    api := r.PathPrefix("/api").Subrouter()
    api.HandleFunc("/playlists", handlers.GetPlaylists).Methods("GET")
    api.HandleFunc("/trending", handlers.GetTrendingSongs).Methods("GET")
    api.HandleFunc("/save", handlers.SavePlaylist).Methods("POST")

    // Serve index.html for the root path
    r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path == "/" {
            http.ServeFile(w, r, "static/index.html")
            return
        }
        http.NotFound(w, r)
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "3090"  // Using your current port
    }

    fmt.Printf("Server running on port %s...\n", port)
    log.Fatal(http.ListenAndServe(":"+port, r))
}