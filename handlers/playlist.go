package handlers

import (
    "encoding/json"                         // For encoding and decoding JSON
    "net/http"                              // For HTTP handling
    "music-playlist-generator/models"       // Importing custom models for playlists and songs
    "music-playlist-generator/utils"        // Utility functions for handling responses
)

// GetPlaylists -  HTTP GET request to retrieve a list of playlists
func GetPlaylists(w http.ResponseWriter, r *http.Request) {

    // Create a sample list of playlists with songs
    playlists := []models.Playlist{
        {
            Name: "Top Hits", // Playlist name
            Songs: []models.Song{
                {Title: "Song 1", Artist: "Artist 1", Album: "Album 1"},
                {Title: "Song 2", Artist: "Artist 2", Album: "Album 2"},
            },
        },
    }
    //  HTTP status 200 (OK)
    utils.WriteJSONResponse(w, http.StatusOK, playlists)
}

// SavePlaylist controll HTTP POST request to save a new playlist
func SavePlaylist(w http.ResponseWriter, r *http.Request) {
    var playlist models.Playlist

    // Decode the JSON request body into the playlist object
    if err := json.NewDecoder(r.Body).Decode(&playlist); err != nil {
        
        //  fails, respond with a 400 (Bad Request) error
        utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
        return
    }

    // Respond  success message
    utils.WriteJSONResponse(w, http.StatusOK, map[string]string{"message": "Playlist saved successfully!"})
}