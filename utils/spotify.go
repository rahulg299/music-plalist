// utils/spotify.go
package utils

import (
    "context"
    "fmt"
    "os"
    // "net/http"
    "golang.org/x/oauth2"
    "github.com/zmb3/spotify/v2"
    spotifyauth "github.com/zmb3/spotify/v2/auth"
)

func GetSpotifyClient() (*spotify.Client, error) {
    clientID := os.Getenv("SPOTIFY_CLIENT_ID")
    clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")

    if clientID == "" || clientSecret == "" {
        return nil, fmt.Errorf("missing Spotify credentials in environment variables")
    }

    // Create Spotify authenticator (scopes might not be needed for client credentials)
    auth := spotifyauth.New(
        spotifyauth.WithClientID(clientID),
        spotifyauth.WithClientSecret(clientSecret),
        spotifyauth.WithScopes(
            spotifyauth.ScopePlaylistReadPrivate,
            spotifyauth.ScopePlaylistModifyPublic,
            spotifyauth.ScopePlaylistModifyPrivate,
            spotifyauth.ScopeUserReadPrivate,
        ),
    )

    ctx := context.Background()
    
    // Use the authenticator's config to handle token exchange
    token, err := auth.Exchange(ctx, "", oauth2.SetAuthURLParam("grant_type", "client_credentials"))
    if err != nil {
        return nil, fmt.Errorf("couldn't get token: %v", err)
    }

    // Create HTTP client with the obtained token
    httpClient := auth.Client(ctx, token)

    // Create Spotify client with the authenticated HTTP client
    client := spotify.New(httpClient)

    return client, nil
}