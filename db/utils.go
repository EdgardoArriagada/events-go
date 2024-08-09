package db

import (
	"database/sql" // <- from go standard lib
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func dbExec(query string) {
	_, err := DB.Exec(query)
	if err != nil {
		panic(err)
	}

}

func getDbConnection(driverName string, dataSourceName string) {
	db, err := sql.Open(driverName, dataSourceName)

	DB = db

	if err != nil {
		panic("Could not connect to database.")
	}
}
