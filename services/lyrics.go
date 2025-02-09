
package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LyricsResponse struct {
	Lyrics string `json:"lyrics"`
}

func FetchLyrics(artist, title string) (string, error) {
	// Build API URL
	url := fmt.Sprintf("https://api.lyrics.ovh/v1/%s/%s", artist, title)

	// Make HTTP request
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch lyrics: %v", err)
	}
	defer resp.Body.Close()

	// Decode response
	var lyricsResp LyricsResponse
	if err := json.NewDecoder(resp.Body).Decode(&lyricsResp); err != nil {
		return "", fmt.Errorf("failed to decode lyrics: %v", err)
	}

	return lyricsResp.Lyrics, nil
}