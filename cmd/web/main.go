package main

import (
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
)

const PORT = ":8080"

func main() {
	router := mux.NewRouter()
	app := newApp(router)
	log.Println("Listening on port: ", PORT)
	app.Run(PORT)
}
