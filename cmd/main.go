package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cminovici/go-api/pkg/handlers"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Article REST API!")
	fmt.Println("Article REST API")
}

func handleRequests() {
	// create a new instance of a mux router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/save", handlers.SaveText).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	handleRequests()
}
