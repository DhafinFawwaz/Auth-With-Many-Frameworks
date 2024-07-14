package database

import (
	"fiber-auth-template/config"
	"fmt"

	"database/sql"

	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDatabase() {
	fmt.Printf("Connecting to database...\n")

	var connStr string
	var db *sql.DB
	var err error

	if config.DB_TYPE() == "postgres" {
		connStr = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			config.DB_USERNAME(),
			config.DB_PASSWORD(),
			config.DB_HOST(),
			config.DB_PORT(),
			config.DB_NAME(),
		)

		// connStr = fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		// 	config.DB_USERNAME(),
		// 	config.DB_PASSWORD(),
		// 	config.DB_NAME(),
		// 	config.DB_HOST(),
		// 	config.DB_PORT(),
		// )

		db, err = sql.Open("postgres", connStr)
	} else { // config.DB_TYPE() == "mysql"
		connStr = fmt.Sprintf("mysql:%s:%s@tcp(%s:%s)/%s?parseTime=true",
			config.DB_USERNAME(),
			config.DB_PASSWORD(),
			config.DB_HOST(),
			config.DB_PORT(),
			config.DB_NAME(),
		)
		db, err = sql.Open("mysql", connStr)
	}

	fmt.Println("Connection String:\n" + connStr)
	if err != nil {
		fmt.Println("Error connecting to database")
		panic(err)
	}
	DB = db
	initializeStatements()

	fmt.Println("Connected to database!")
}

// Statements
var createMahasiswaTable *sql.Stmt
var selectAllMahasiswa *sql.Stmt
var insertNewMahasiswa *sql.Stmt
var selectByEmail *sql.Stmt

func initializeStatements() {
	createMahasiswaTable = prepare(`
CREATE TABLE IF NOT EXISTS mahasiswa (
	id SERIAL PRIMARY KEY, 
	username VARCHAR(255), 
	email VARCHAR(255),
	password VARCHAR(255), 
	date_joined TIMESTAMP, 
	nim VARCHAR(255)
)`)
	exec(createMahasiswaTable)

	selectAllMahasiswa = prepare("SELECT * FROM mahasiswa")
	insertNewMahasiswa = prepare("INSERT INTO mahasiswa (username, email, password, date_joined, nim) VALUES ($1, $2, $3, $4, $5)")
	selectByEmail = prepare("SELECT * FROM mahasiswa WHERE email = $1")
}

func prepare(statementStr string) *sql.Stmt {
	statement, err := DB.Prepare(statementStr)
	if err != nil {
		fmt.Println("Error preparing statement")
		panic(err)
	}
	return statement
}

func exec(statement *sql.Stmt) {
	_, err := statement.Query()
	if err != nil {
		fmt.Println("Error preparing statement")
		panic(err)
	}
}

func DisconnectDatabase() {
	selectAllMahasiswa.Close()
	insertNewMahasiswa.Close()
	selectByEmail.Close()
	createMahasiswaTable.Close()

	DB.Close()
	fmt.Println("Disconnected from database")
}
