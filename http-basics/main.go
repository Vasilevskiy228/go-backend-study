package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", healthHandler)
	mux.HandleFunc("POST /songs/preview", songPreviewHandler)

	mux.HandleFunc("GET /songs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"songs":[]}`))
	})

	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))

}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func songPreviewHandler(w http.ResponseWriter, r *http.Request) {
	var input SongRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	if input.Title == "" || input.Artist == "" {
		http.Error(w, "title or artist are required", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(input); err != nil {
		log.Printf("ошибка записи ответа: %v", err)
	}
}

type SongRequest struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
}
