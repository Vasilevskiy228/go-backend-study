package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSongPreviewHandlerSuccess(t *testing.T) {
	body := `{"title":"Night Drive","artist":"Ansel"}`

	req := httptest.NewRequest(http.MethodPost, "/songs/preview", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	songPreviewHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("ожидали статус 200, получили: %d", rec.Code)
	}

	if rec.Header().Get("Content-Type") != "application/json" {
		t.Errorf("ожидали `application/json`,получили: %s", rec.Header().Get("Content-Type"))
	}

	gotBody := strings.TrimSpace(rec.Body.String())

	if gotBody != body {
		t.Errorf("ожидали тело %q, получили %q", body, gotBody)
	}

}

func TestSongPreviewHandlerInvalidJSON(t *testing.T) {
	body := `{"title":"`

	req := httptest.NewRequest(http.MethodPost, "/songs/preview", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	songPreviewHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("ожидали статус 400, получили: %d", rec.Code)
	}
}

func TestSongPreviewHandlerEmptyArtist(t *testing.T) {
	body := `{"title":"Night Drive","artist":""}`

	req := httptest.NewRequest(http.MethodPost, "/songs/preview", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	songPreviewHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("ожидали статус 400, получили: %d", rec.Code)
	}
}

func TestSongPreviewHandlerBadRequests(t *testing.T) {
	tests := []struct {
		name string
		body string
	}{{
		name: "некорректный JSON",
		body: `{"title":`,
	},
		{
			name: "пустой исполнитель",
			body: `{"title":"Night Drive","artist":""}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/songs/preview", strings.NewReader(tc.body))
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			songPreviewHandler(rec, req)

			if rec.Code != http.StatusBadRequest {
				t.Fatalf("ожидали статус 400, получили:%d", rec.Code)
			}
		})
	}

}
