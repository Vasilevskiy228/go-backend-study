package main

import (
	"testing"
)

func TestValidateSong(t *testing.T) {
	service := &SongService{}

	tests := []struct {
		name    string
		cases   SongRequest
		wantErr bool
	}{
		{
			name: "все правильно",
			cases: SongRequest{
				Title:  "sex",
				Artist: "sugen",
			},
			wantErr: false,
		},
		{
			name: "пустое название",
			cases: SongRequest{
				Title:  "",
				Artist: "sugen",
			},
			wantErr: true,
		},
		{
			name: "пустой исполнитель",
			cases: SongRequest{
				Title:  "sex",
				Artist: "",
			},
			wantErr: true,
		},
		{
			name: "пробелы",
			cases: SongRequest{
				Title:  "    ",
				Artist: "sugen",
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := service.ValidateSong(tc.cases)

			if !tc.wantErr {
				if err != nil {
					t.Fatalf("не ожидали ошибку, но получили")
				}
			}

			if tc.wantErr {
				if err == nil {
					t.Fatalf("ожидали ошибку, но не получили")

				}
			}
		})
	}
}
