package service

import (
	. "BLOG/domain/model"
	"BLOG/domain/repository"
	"errors"
	"strings"
)

func IsDirty(content string) bool {
	badWords := []string{"anjing", "babi", "bangke", "goblok", "tolol", "bego",
	"sinting", "gila", "kurang ajar", "sialan", "brengsek",
	"kampret", "bedebah", "kontol", "memek", "ngentot",
	"tai", "bangsat", "perek", "lonte", "monyet", "kera",
	"asu", "jancok", "bajingan", "cok", "ajig", "kehed", "bacot",}

	for _,word := range badWords {
		if strings.Contains(strings.ToLower(content), word) {
			return true
		}
	}

	return false
}

func CommentService(UserID int, Post_ID int, content string) (*Komen, error) {
	if IsDirty(content) {
		return nil, errors.New("Komentar Mengandung Kata Tidak Pantas")
	}

	res, err := repository.CreateCommet(Post_ID)

	if err != nil {
		return nil, err
	}

	return res, nil
}