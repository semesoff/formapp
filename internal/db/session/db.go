package session

import (
	"database/sql"
	"formapp/internal/db/data"
)

func createTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS sessions (
    	session_id TEXT PRIMARY KEY,
        data bytea,
        created_at timestamptz DEFAULT NOW(),
        updated_at timestamptz default now()
	)`
	_, err := db.Exec(query)
	return err
}

func main() {
	s, _ := data.ConnectDB()
	createTable(s)
}
