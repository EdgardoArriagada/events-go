package db

func InitDB() {
	getDbConnection("sqlite3", "data/db/api.db")

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	dbExec(`
  CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    date_time DATETIME NOT NULL,
    user_id INTEGER
  )`)
}
