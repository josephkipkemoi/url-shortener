package database

import (
	"database/sql"
	"log"
)

// createUrlTableMigration creates the urlshortener table with the required table columns
func createUrlTableMigration(urlDb *sql.DB) {
	createSQLUrlTable := `CREATE TABLE urlshortener (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		original_url STRING NOT NULL,
		shortened_url STRING NOT NULL
	);` // SQL statement to create urlshortener table

	log.Println("creating urlshortener table...")

	statement, err := urlDb.Prepare(createSQLUrlTable) // prepare SQL statement
	checkErr(err)

	statement.Exec() // execute SQL statement

	log.Println("urlshortener table created...")
}
