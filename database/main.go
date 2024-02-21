package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

func init() {
	ConnentDB()
}

var DB *sql.DB

// ConnectDb function runs the createDatabase for our application
func ConnentDB() {
	createDatabase()
}

// createDatabase function creates the database for our application
func createDatabase() {
	var fp string = "./database/files/urlshortener.db"
	ok := checkFileExists(fp)

	if !ok {
		file, err := os.Create("./database/files/urlshortener.db") // create SQLite file
		checkErr(err)
		file.Close()
		log.Println("urlshortener.db created...")
	}

	urlDb, _ := sql.Open("sqlite3", "./database/files/urlshortener.db") // open the created pinacle db file

	if !ok {
		createUrlTableMigration(urlDb)
	}

	log.Println("urlshortener database running...")

	DB = urlDb
}
