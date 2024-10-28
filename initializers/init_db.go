package initializers

import (
	"database/sql"
	"jtw-go/db"
	"log"
)

var DB *sql.DB

func InitDB() {
	if DbDSN == "" {
		log.Fatal("DB_DSN environment variable not set")
	}
	db_, err := db.CreateDB(DbDSN)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	DB = db_
}
