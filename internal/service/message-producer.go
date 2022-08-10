package service

import (
	"context"
	"encoding/json"
	"github.com/raymondgitonga/company-service/internal/config"
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

	if err != nil {
		log.Printf("Error marshalling message: %s", err)
		return
	}

	err = config.Connect("company-service", message, context.Background())

	if err != nil {
		log.Fatalf("Error sending message %s", err)
	}
}
