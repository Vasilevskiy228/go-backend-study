package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("POST /songs/preview", songPreviewHandler)

	log.Println("Сервер запущен на порту:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))

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

	input.Artist = strings.TrimSpace(input.Artist)
	input.Title = strings.TrimSpace(input.Title)

	if input.Artist == "" || input.Title == "" {
		http.Error(w, "title or artist are required", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(input); err != nil {
		log.Printf("ошибка записи ответа: %v", err)
	}
}
