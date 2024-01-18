package models

import (
	"time"

	"github.com/lib/pq"
)

type Mahasiswa struct {
	ID          int         `json:"id"`
	Password    string      `json:"password"`
	LastLogin   pq.NullTime `json:"last_login"`
	IsSuperuser bool        `json:"is_superuser"`
	Username    string      `json:"username"`
	FirstName   string      `json:"first_name"`
	LastName    string      `json:"last_name"`
	Email       string      `json:"email"`
	IsStaff     bool        `json:"is_staff"`
	IsActive    bool        `json:"is_active"`
	DateJoined  time.Time   `json:"date_joined"`
	NIM         string      `json:"nim"`
}

type AuthenticatedMahasiswa struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Mahasiswa
}
