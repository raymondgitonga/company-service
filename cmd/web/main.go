package main

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const PORT = ":8080"

func main() {
	router := mux.NewRouter()
	app := newApp(router)
	fmt.Println("Listening on port: ", PORT)
	app.Run(PORT)
}
