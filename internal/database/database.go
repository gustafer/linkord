package database

import (
	"database/sql"

	"github.com/gustafer/linkord/configs"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	// Initialize the database connection once when the package is loaded
	connStr := configs.LoadDbConnStr()
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
}

func NewConn() *sql.DB {
	// reuse the existing db pointer connection
	return db
}

func Migrate() error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS game (
		id SERIAL PRIMARY KEY,
		title VARCHAR(40),
		description VARCHAR(250)
	);
	CREATE TABLE IF NOT EXISTS "user" (
		id VARCHAR(40) PRIMARY KEY,
		username VARCHAR(40)
	);
	CREATE TABLE IF NOT EXISTS user_game (
		game_id INTEGER REFERENCES game(id),
		user_id VARCHAR(40) REFERENCES "user"(id) 
	);
	`)
	if err != nil {
		return err
	}

	return nil
}
