package config

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func CreateDBConnection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := initialiseDB(
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func initialiseDB(userName string, password string, dbName string) (*sql.DB, error) {
	connectionString := fmt.Sprintf("port=5432 user=%s password=%s dbname=%s sslmode=disable", userName, password, dbName)
	DB, err := sql.Open("postgres", connectionString)

	err = DB.Ping()
	if err != nil {
		panic(err)
	}
	return DB, err
}
