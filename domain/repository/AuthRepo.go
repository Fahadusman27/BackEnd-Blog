package repository

import (
	"BLOG/domain/model"
	"database/sql"
	"time"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db}
}

func (r *AuthRepository) CreateUsers(user *model.Register) error {
	query := `
		INSERT INTO users (Email, Username, Password, RoleID)
		VALUES ($1, $2, $3, 2)
	`
	_, err := r.db.Exec(query, user.Email, user.Username, user.Password)
	return err
}

func (r *AuthRepository) FindByEmail(Email string) (model.Login, error) {
	query := `
		SELECT u.ID, u.Email, u.Password, r.Name 
		FROM users u
		JOIN role r ON u.RoleID = r.ID
		WHERE u.Email = $1`

	var login model.Login
	err := r.db.QueryRow(query, Email).
		Scan(&login.ID, &login.Email, &login.Password, &login.RoleName)
	
	if err != nil {
		return model.Login{}, err
	}

	return login, nil
}

func (r *AuthRepository) UpadateProfile(ID int, profile *model.Profile) error {
	query := `UPDATE users SET Username = $1, Picture = $2, Name = $3, UpdatedAt = $4 WHERE id = $5`

	_, err := r.db.Exec(query, profile.Username, profile.Picture, profile.Name, time.Now(), ID)

	return err
}

func (r *AuthRepository) GetProfile(ID int) (model.Profile, error) {
	query :=`SELECT ID, Email, Picture, Username FROM users WHERE ID=$1`
	
	var Profile model.Profile
	err := r.db.QueryRow(query, ID).
	Scan(&Profile.ID, &Profile.Email, &Profile.Picture, &Profile.Username)
	if err != nil {
		return model.Profile{}, err
	}
	
	return Profile, nil
}

//fitur admin

func (r *AuthRepository) RoleSelected(ID int) (model.Role, error) {
	query :=`SELECT ID, Name FROM role WHERE ID=$1`

	var role model.Role
	err := r.db.QueryRow(query, ID).
	Scan(&role.ID, &role.Name)
	if err != nil {
		return model.Role{}, err
	}

	return role, nil
}