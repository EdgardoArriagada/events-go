package db

import (
	"database/sql" // <- from go standard lib
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func getDbConnection(driverName string, dataSourceName string) *sql.DB {
	db, err := sql.Open(driverName, dataSourceName)

	if err != nil {
		panic("Could not connect to database.")
	}

	return db
}

func InitDB() {
	DB = getDbConnection("sqlite3", "data/db/api.db")

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	_, err := DB.Exec(`
  CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    date_time DATETIME NOT NULL,
    user_id INTEGER
  )`,
	)

	if err != nil {
		panic("Could not create table events.")
	}
}
