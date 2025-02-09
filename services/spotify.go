package services

import (
    "context"                               // For managing request contexts
    "fmt"                                   // For formatted I/O operations
    "os"                                    // For environment variable access
    "github.com/zmb3/spotify/v2"            // Spotify API client
    spotifyauth "github.com/zmb3/spotify/v2/auth" // Spotify authentication utilities
    "golang.org/x/oauth2/clientcredentials" // OAuth2 client credentials flow
)


func FetchTrendingSongs() ([]spotify.FullTrack, error) {
    // Load Spotify client credentials from environment variables
    config := &clientcredentials.Config{
        ClientID:     os.Getenv("1c50df945d5a43449f3e170f7a3493d9"),       // Spotify API client ID
        ClientSecret: os.Getenv("2d0793e4a4114280b3cde4a9bc28beb7"),   // Spotify API client secret
        TokenURL:     spotifyauth.TokenURL,                 // Spotify OAuth2 token URL
    }

    // Create a background context for the request
    ctx := context.Background()

    // Fetch the access token using client credentials
    token, err := config.Token(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed to get token: %v", err) // Handle token retrieval error
    }

    // Create an authenticated HTTP client
    httpClient := spotifyauth.New().Client(ctx, token)
    client := spotify.New(httpClient) // Initialize Spotify API client

    // Define the Spotify playlist ID for trending songs
    playlistID := spotify.ID("37i9dQZEVXbMDoHDwVN2tF") // Example Spotify playlist ID for Top Hits

    // Fetch playlist tracks from Spotify
    tracks, err := client.GetPlaylistTracks(ctx, playlistID)
    if err != nil {
        return nil, fmt.Errorf("failed to fetch tracks: %v", err) // Handle track retrieval error
    }

    // Extract full track information from the response
    var fullTracks []spotify.FullTrack
    for _, track := range tracks.Tracks {
        fullTracks = append(fullTracks, track.Track) // Append each track to the result slice
    }

    // Return the list of tracks and no error
    return fullTracks, nil
}