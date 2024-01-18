package database

import (
	"fiber-auth-template/internal/models"
	"fmt"
	"time"
)

func SelectAllMahasiswa() ([]models.Mahasiswa, error) {
	var err error

	rows, err := selectAllMahasiswa.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mahasiswas []models.Mahasiswa
	for rows.Next() {
		var mahasiswa models.Mahasiswa
		if err := rows.Scan(&mahasiswa.ID, &mahasiswa.Password, &mahasiswa.LastLogin, &mahasiswa.IsSuperuser, &mahasiswa.Username, &mahasiswa.FirstName, &mahasiswa.LastName, &mahasiswa.Email, &mahasiswa.IsStaff, &mahasiswa.IsActive, &mahasiswa.DateJoined, &mahasiswa.NIM); err != nil {
			fmt.Println(err)
		}
		mahasiswas = append(mahasiswas, mahasiswa)
	}
	return mahasiswas, nil
}

func InsertNewMahasiswa(mahasiswa models.Mahasiswa) error {
	var err error

	_, err = insertNewMahasiswa.Exec(mahasiswa.Password, time.Now(), mahasiswa.IsSuperuser, mahasiswa.Username, mahasiswa.FirstName, mahasiswa.LastName, mahasiswa.Email, mahasiswa.IsStaff, mahasiswa.IsActive, time.Now(), mahasiswa.NIM)
	return err
}

func SelectByEmail(email string) (models.Mahasiswa, error) {
	var err error
	var mahasiswa models.Mahasiswa

	err = selectByEmail.QueryRow(email).Scan(&mahasiswa.ID, &mahasiswa.Password, &mahasiswa.LastLogin, &mahasiswa.IsSuperuser, &mahasiswa.Username, &mahasiswa.FirstName, &mahasiswa.LastName, &mahasiswa.Email, &mahasiswa.IsStaff, &mahasiswa.IsActive, &mahasiswa.DateJoined, &mahasiswa.NIM)

	fmt.Println("err:")
	fmt.Println(err)
	return mahasiswa, err
}
