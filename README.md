# Music Playlist Generator

A web application that allows users to generate and manage music playlists using trending songs from Spotify.

## Features
- Fetches trending songs from Spotify
- Create and save custom playlists
- View saved playlists
- Fetch song lyrics (via external API)
- Simple UI for managing playlists

## Installation

1. **Clone the repository**  
   ```sh
   git clone https://github.com/yourusername/music-playlist-generator.git
   cd music-playlist-generator
Set up environment variables
Create a .env file and add your credentials:

SPOTIFY_CLIENT_ID=your-client-id
SPOTIFY_CLIENT_SECRET=your-client-secret
REDIRECT_URL=http://localhost:3090

Install dependencies
go mod tidy

Run the server
go run main.go

Access the web app
Open http://localhost:3090 in your browser.

API Endpoints
GET /api/playlists - Get saved playlists
GET /api/trending - Fetch trending songs
POST /api/save - Save a playlist
Technologies Used
Go (Gorilla Mux for routing)
Spotify API (for fetching trending songs)
JavaScript, HTML, CSS (Frontend)
Lyrics API (for song lyrics)
