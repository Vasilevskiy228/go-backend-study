package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSongPreviewHandler(t *testing.T) {
	body := `{"title":"sex","artist":"sugen"}`

	req := httptest.NewRequest(http.MethodPost, "/songs/preview", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	songPreviewHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Ожидали статус 200, получили: %d", rec.Code)

	}

	if rec.Header().Get("Content-Type") != "application/json" {
		t.Errorf("Ожидали 'application/json', получили %s", rec.Header().Get("Content-Type"))
	}

	var input SongRequest
	output := SongRequest{
		Title:  "sex",
		Artist: "sugen",
	}

	err := json.NewDecoder(rec.Body).Decode(&input)
	if err != nil {
		t.Fatalf("Ошибка декодирования")
	}

	if input.Artist != output.Artist || input.Title != output.Title {
		t.Errorf("Ожидали %q, получили %q", output, input)
	}

}

func TestSongPreviewHandlerBadRequests(t *testing.T) {
	tests := []struct {
		name string
		body string
	}{
		{
			name: "invalid JSON",
			body: `{"title":`,
		},
		{
			name: "empty title",
			body: `{"title":"","artist":"sugen"}`,
		},
		{
			name: "empty artist",
			body: `{"title":"sex","artist":""}`,
		},
		{
			name: "whitespace title",
			body: `{"title":"	","artist":"sugen"}`,
		},
		{
			name: "whitespace artist",
			body: `{"title":"sex","artist":"	"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/songs/preview", strings.NewReader(tc.body))
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			songPreviewHandler(rec, req)

			if rec.Code != http.StatusBadRequest {
				t.Fatalf("ожидали 400, получили %d", rec.Code)
			}
		})
	}
}
