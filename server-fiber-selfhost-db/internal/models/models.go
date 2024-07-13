package models

import (
	"time"
)

type Mahasiswa struct {
	ID         int       `json:"id"`
	Password   string    `json:"password"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	DateJoined time.Time `json:"date_joined"`
	NIM        string    `json:"nim"`
}

type AuthenticatedMahasiswa struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Mahasiswa
}
