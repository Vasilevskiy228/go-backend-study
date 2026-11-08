package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("POST /songs/preview", songPreviewHandler1)

}

type SongRequest struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
}

func songPreviewHandler(w http.ResponseWriter, r *http.Request) {
	var input SongRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid json body", http.StatusBadRequest)
		return
	}

	if input.Artist == "" || input.Title == "" {
		http.Error(w, "title or artist are required", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(input); err != nil {
		log.Printf("ошибка записи ответа: %v", err)
	}
}
