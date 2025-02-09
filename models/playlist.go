package models

// Song represents a song with its details
type Song struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Album  string `json:"album"`
}

// Playlist represents a playlist with a list of songs
type Playlist struct {
	Name  string `json:"name"`
	Songs []Song `json:"songs"`
}