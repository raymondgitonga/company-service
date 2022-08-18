package service

import (
	"context"
	"encoding/json"
	"github.com/raymondgitonga/company-service/internal/config"
	"github.com/segmentio/kafka-go"
	"log"
)

type Produce struct {
	Event     string `json:"event"`
	CompanyId string `json:"company_id"`
}

type Producer interface {
	SendMutationMessage()
}

func (p Produce) SendMutationMessage() {
	message, err := json.Marshal(p)
	ctx := context.Background()

	if err != nil {
		log.Printf("Error marshalling message: %s", err)
		return
	}

	writer := config.CreateKafkaConnection("company-service")

	err = writer.WriteMessages(
		ctx,
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: message,
		},
	)

	if err != nil {

	}
	log.Print("message sent: ", string(message))

	if err != nil {
		log.Printf("Error sending message %s", err)
	}
}
