package config

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
)

type Received struct {
	Message string
	Offset  int64
}

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

func ConnectKafka(topic string, ctx context.Context, msgChan chan Received) {
	defer close(msgChan)

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	err := reader.SetOffset(1)

	defer func() {
		err := reader.Close()
		if err != nil {
			fmt.Println("Error closing consumer: ", err)
			return
		}
		fmt.Println("Consumer closed")
	}()

	if err != nil {
		log.Printf("Failed to set offset %s", err)
	}

	for {
		message, err := reader.ReadMessage(ctx)
		if err != nil {
			if err == context.Canceled {
				fmt.Println("Signal interrupt error ", err)
				break
			}
			fmt.Println("Error reading message ", err)
			break
		}

		msgChan <- Received{
			Message: string(message.Value),
			Offset:  message.Offset,
		}
	}
	if err := reader.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
