package model

import (
	"github.com/golang-jwt/jwt/v4"
)

type Task struct {
	Id      string `json:"id,omitempty"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment,omitempty"`
	Repeat  string `json:"repeat"`
}

type Auth struct {
	Password string `json:"password"`
}
type Claims struct {
	PasswordHash string `json:"password_hash"`
	jwt.StandardClaims
}

const (
	TimeFormat = "20060102"
)

type TasksResp struct {
	Tasks []Task `json:"tasks"`
}
