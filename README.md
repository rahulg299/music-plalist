# Music Playlist Generator

The **Music Playlist Generator** is a web application that allows users to create, save, and view music playlists. It also provides trending songs fetched from Spotify and lyrics for individual songs.

## Features
- Fetch trending songs from Spotify.
- Save and manage custom playlists.
- Fetch song lyrics using the [Lyrics.ovh API](https://lyrics.ovh/).
- Modern UI with a responsive design.

## Technologies Used
- **Backend**: Golang, Gorilla Mux, Spotify API
- **Frontend**: HTML, CSS, JavaScript
- **Middleware**: Custom Logging and Content-Type Middleware
- **Database**: In-memory data storage (can be replaced with a DB)
- **APIs**:
  - Spotify API for trending songs.
  - Lyrics.ovh API for fetching lyrics.

## Installation

### Prerequisites
- Golang installed (v1.23.5 or later)
- Spotify Developer Account for API credentials
- `.env` file with the following variables:
  ```plaintext
  SPOTIFY_CLIENT_ID=<your_client_id>
  SPOTIFY_CLIENT_SECRET=<your_client_secret>
  REDIRECT_URL=http://localhost:3090
Steps
Clone the repository:

bash
git clone https://github.com/your-username/music-playlist-generator.git
cd music-playlist-generator
Install dependencies:

bash
go mod tidy
Run the application:

bash
go run main.go
Open your browser and visit http://localhost:3090.

API Endpoints
GET /api/trending: Fetches trending songs from Spotify.
GET /api/playlists: Retrieves saved playlists.
POST /api/save: Saves a new playlist (requires JSON body).
