package main

import (
	"github.com/gorilla/mux"
	"github.com/raymondgitonga/company-service/internal/handler"
	"github.com/raymondgitonga/company-service/internal/middleware"
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
	a.Router.HandleFunc("/authorize", middleware.GenerateJWT).Methods(http.MethodGet)
	a.Router.HandleFunc("/companies", handler.GetCompanies).Methods(http.MethodGet)
	a.Router.HandleFunc("/company", handler.GetCompany).Methods(http.MethodGet)

	if err := http.ListenAndServe(addr, a.Router); err != nil {
		log.Fatal("error connecting to server:  ", err.Error())
	}
}
