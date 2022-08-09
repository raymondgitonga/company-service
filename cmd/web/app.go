package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
}

func newApp(router *mux.Router) *App {
	return &App{
		Router: router,
	}
}

func (a *App) Run(addr string) {
	a.Router.HandleFunc("/healthCheck", HealthCheck)

	if err := http.ListenAndServe(addr, a.Router); err != nil {
		log.Fatal("error connecting to server:  ", err.Error())
	}
}

type healthCheck map[string]interface{}

func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	response := make([]healthCheck, 0)

	// Add more checks i.e DB checks, kafka ...
	response = append(response, healthCheck{
		"type":   "App health",
		"status": "Healthy",
		"error":  nil,
	})

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Unable to encode JSON")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsonResponse)
}
