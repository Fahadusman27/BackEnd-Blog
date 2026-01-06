package repository

import (
	"BLOG/domain/config"
	. "BLOG/domain/model"
)

//Admin

func SeacrhByID(ID int) (*Users, error) {
	user := new(Users)

	query := `SELECT ID, Username, Picture, Email, RoleID, CreatedAt FROM Users WHERE id=$1`

	err :=config.DB.QueryRow(query, ID).
		Scan(&user.ID, &user.Username, &user.Email, &user.Picture, &user.Role, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

//User

func SearchByUsername(username string) (*Users, error) {
	user := new(Users)

	searchPattern := "%" + username + "%"
	
	query := `SELECT ID, Username FROM Users WHERE Username LIKE $1`

	err := config.DB.QueryRow(query, searchPattern).
		Scan(&user.ID, &user.Username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func SearchByKategori(Name string) (*Kategori, error) {
	Kategori := new(Kategori)

	searchPattern := "%" + Name + "%"

	query := `SELECT ID, Name FROM kategori WHERE Name LIKE $1`

	err := config.DB.QueryRow(query, searchPattern).
		Scan(&Kategori.ID, &Kategori.Name, &Kategori.Slug)
	if err != nil {
		return nil, err
	}

	return Kategori, nil
}

func SearchByTitle(Name string) (*Title, error) {
	title := new(Title)

	searchPattern := "%" + Name + "%"

	query := `SELECT ID, Name FROM Title WHERE Title LIKE $1`

	err := config.DB.QueryRow(query, searchPattern).
		Scan(&title.ID, &title.Name)
	if err != nil {
		return nil, err
	}

	return title, nil
}