package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	service := &SongService{}

	mux.HandleFunc("POST /songs/preview", songPreviewHandler(service))

	log.Println("Сервер запущен на порту:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))

}

type SongRequest struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
}

func songPreviewHandler(service *SongService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var input SongRequest

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "invalid json body", http.StatusBadRequest)
			return
		}

		if err := service.ValidateSong(input); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(input); err != nil {
			log.Printf("ошибка записи ответа: %v", err)
		}
	}
}
