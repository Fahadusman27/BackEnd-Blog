package service

import (
	. "BLOG/domain/repository"
	"errors"
	"fmt"
)

func Search(keyword string) (map[string]interface{}, error) {
	if keyword == "" {
		return nil, errors.New("Kata Kunci tidak ada")
	}

	user, err := SearchByUsername(keyword)
	if err != nil {
		fmt.Println("User Tidak Ditemukan: %v\n", err)
	}

	kategori, err := SearchByKategori(keyword)
	if err != nil {
		fmt.Println("Kategori Tidak Ditemukan: %v\n", err)
	}

	title, err := SearchByTitle(keyword)
	if err != nil {
		fmt.Println("Judul Tidak Ditemukan: %v\n", err)
	}

	result := map[string]interface{} {
		"found_user":	user,
		"found_category":	kategori,
		"found,title":	title,
		"search_keyword":	keyword,
	}

	return result, nil
}