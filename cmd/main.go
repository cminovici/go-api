package main

import (
	"database/sql"
	"github.com/cminovici/go-api/pkg/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("Failed to connect to the SQLite database: %v", err)
	}

	// Create a sample table to store articles
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS word_count (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		word TEXT UNIQUE,
		count NUMERIC
	);
	`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Failed to create SQLite tables: %v", err)
	}

	return db
}

func handleRequests() {
	db := initDB()

	server := &handlers.Server{
		DB: db,
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/get", server.GetWordOccurrences).Methods(http.MethodGet)
	router.HandleFunc("/save", server.SaveWordOccurrences).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	handleRequests()
}
