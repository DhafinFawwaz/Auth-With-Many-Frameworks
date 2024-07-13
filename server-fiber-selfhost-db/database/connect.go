package database

import (
	"fiber-auth-template/config"
	"fmt"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDatabase() {
	fmt.Println("Connecting to database...")
	connStr := fmt.Sprintf("%s:%s@tcp(database:%s)/%s?parseTime=true",
		config.DB_USERNAME(),
		config.DB_PASSWORD(),
		config.DB_PORT(),
		config.DB_NAME(),
	)
	db, err := sql.Open("mysql", connStr)
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
	id INT AUTO_INCREMENT PRIMARY KEY, 
	username VARCHAR(255), 
	email VARCHAR(255),
	password VARCHAR(255), 
	date_joined DATETIME, 
	nim VARCHAR(255)
)`)
	exec(createMahasiswaTable)

	selectAllMahasiswa = prepare("SELECT * FROM mahasiswa")
	insertNewMahasiswa = prepare("INSERT INTO mahasiswa (username, email, password, date_joined, nim) VALUES (?, ?, ?, ?, ?)")
	selectByEmail = prepare("SELECT * FROM mahasiswa WHERE email = ?")
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
