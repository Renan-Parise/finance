package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	db   *sql.DB
	once sync.Once
)

func GetDB() *sql.DB {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_NAME")

		dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			dbUser, dbPassword, dbHost, dbPort, dbName)

		var errDb error
		db, errDb = sql.Open("mysql", dataSourceName)
		if errDb != nil {
			log.Fatalf("Could not open db: %v", errDb)
		}

		if errPing := db.Ping(); errPing != nil {
			log.Fatalf("Could not connect to db: %v", errPing)
		}
	})
	return db
}
