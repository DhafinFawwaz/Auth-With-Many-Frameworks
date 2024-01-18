package database

import (
	"fiber-auth-template/config"
	"fmt"

	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDatabase() {
	fmt.Println("Connecting to database...")
	dbPass := config.GetEnv("DB_PASS")
	connStr := fmt.Sprintf("postgres://postgres.bmmkodhliluefecuznfk:%s@aws-0-ap-southeast-1.pooler.supabase.com:6543/postgres", dbPass)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error connecting to database\n", err)
	}
	DB = db
	InitializeStatements()
	fmt.Println("Connected to database")
}

// Statements
var selectAllMahasiswa *sql.Stmt
var insertNewMahasiswa *sql.Stmt
var selectByEmail *sql.Stmt

func InitializeStatements() {
	var err error
	selectAllMahasiswa, err = DB.Prepare("SELECT * FROM user_api_mahasiswa")
	if err != nil {
		fmt.Println("Error preparing statement\n", err)
	}
	insertNewMahasiswa, err = DB.Prepare("INSERT INTO user_api_mahasiswa (password, last_login, is_superuser, username, first_name, last_name, email, is_staff, is_active, date_joined, nim) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)")
	if err != nil {
		fmt.Println("Error preparing statement\n", err)
	}
	selectByEmail, err = DB.Prepare("SELECT * FROM user_api_mahasiswa WHERE email = $1")
	if err != nil {
		fmt.Println("Error preparing statement\n", err)
	}
}

func DisconnectDatabase() {
	selectAllMahasiswa.Close()
	insertNewMahasiswa.Close()
	selectByEmail.Close()
	DB.Close()
	fmt.Println("Disconnected from database")
}
