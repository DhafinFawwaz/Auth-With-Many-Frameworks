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
		if err := rows.Scan(&mahasiswa.ID, &mahasiswa.Username, &mahasiswa.Email, &mahasiswa.Password, &mahasiswa.DateJoined, &mahasiswa.NIM); err != nil {
			fmt.Println(err)
		}
		mahasiswas = append(mahasiswas, mahasiswa)
	}
	return mahasiswas, nil
}

func InsertNewMahasiswa(mahasiswa models.Mahasiswa) error {
	var err error

	_, err = insertNewMahasiswa.Exec(mahasiswa.Username, mahasiswa.Email, mahasiswa.Password, time.Now(), mahasiswa.NIM)
	return err
}

func SelectByEmail(email string) (*models.Mahasiswa, error) {
	var mahasiswa models.Mahasiswa

	row := selectByEmail.QueryRow(email)
	if row.Err() != nil {
		return nil, row.Err()
	}

	err := row.Scan(&mahasiswa.ID, &mahasiswa.Username, &mahasiswa.Email, &mahasiswa.Password, &mahasiswa.DateJoined, &mahasiswa.NIM)

	if err != nil {
		return nil, err
	}

	return &mahasiswa, nil
}

// Debug
func SQL(statementStr *string) ([][]string, error) {
	statement, err := DB.Prepare(*statementStr)
	if err != nil {
		return nil, err
	}

	defer statement.Close()
	rows, err := statement.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	rawResult := make([][]byte, len(cols))
	result := make([][]string, 0)

	dest := make([]interface{}, len(cols))
	for i := range rawResult {
		dest[i] = &rawResult[i]
	}

	i := 0
	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			fmt.Println("Failed to scan row", err)
			return nil, err
		}

		result = append(result, make([]string, len(rawResult)))
		for j, raw := range rawResult {
			if raw == nil {
				result[i][j] = "nil"
			} else {
				result[i][j] = string(raw)
			}
		}
		i++
	}

	return result, nil

}
