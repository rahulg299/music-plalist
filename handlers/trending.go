package handlers



import (
    "net/http"
    "music-playlist-generator/services"
    "music-playlist-generator/utils"
)

func GetTrendingSongs(w http.ResponseWriter, r *http.Request) {
    songs, err := services.FetchTrendingSongs()
    if err != nil {
        utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
        return
    }
    utils.WriteJSONResponse(w, http.StatusOK, songs)
}