package models

import (
	"time"

	"github.com/google/uuid"
)

type Mahasiswa struct {
	id           uuid.UUID
	password     string
	last_login   time.Time
	is_superuser bool
	username     string
	first_name   string
	last_name    string
	email        string
	is_staff     bool
	is_active    bool
	date_joined  time.Time
	nim          string
}
