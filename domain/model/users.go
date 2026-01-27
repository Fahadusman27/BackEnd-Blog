package model

import "time"

type Users struct {
	ID           int       `json:"id"`
	Username     string    `json:"Username"`
	Picture      string    `json:"Picture"`
	Email        string    `json:"Email"`
	PasswordHash string    `json:"-"`
	Role         int       `json:"RoleID"`
	CreatedAt    time.Time `json:"CreatedAt"`
}

type Login struct {
	ID       int    `json:"id"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
	RoleName string `json:"RoleName"`
}

type Register struct {
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Role     int    `json:"RoleID"`
	CreatedAt	time.Time	`json:"CreatedAt"`
}

type Profile struct {
	ID       int    `json:"id"`
	Name	string	`json:"Name"`
	Username string `json:"Username"`
	Picture  string `json:"Picture"`
	Email    string `json:"Email"`
	UpdatedAt	time.Time	`json:"UpadatedAt"`
}
