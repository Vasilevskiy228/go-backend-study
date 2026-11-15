package main

import (
	"errors"
	"strings"
)

type SongService struct {
}

func (s *SongService) ValidateSong(r SongRequest) error {
	if strings.TrimSpace(r.Artist) == "" || strings.TrimSpace(r.Title) == "" {
		return errors.New("Пустые данные")

	}

	return nil

}
