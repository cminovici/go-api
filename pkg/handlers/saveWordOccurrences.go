package handlers

import (
	"encoding/json"
	"github.com/cminovici/go-api/pkg/models"
	"net/http"
	"strings"
)

func (s *Server) SaveWordOccurrences(w http.ResponseWriter, r *http.Request) {

	var requestBody models.RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(requestBody.Body) == "" {
		http.Error(w, "Text body cannot be empty", http.StatusBadRequest)
		return
	}

	lowercasedText := strings.ToLower(requestBody.Body)
	words := strings.Fields(lowercasedText)

	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	txn, err := s.DB.Begin()
	if err != nil {
		http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
		return
	}

	insertSQL := `INSERT INTO word_count (word, count) VALUES (?, ?) ON CONFLICT(word) DO UPDATE SET count = count + excluded.count`
	stmt, err := txn.Prepare(insertSQL)
	if err != nil {
		http.Error(w, "Failed to prepare statement", http.StatusInternalServerError)
		txn.Rollback()
		return
	}
	defer stmt.Close()

	for word, count := range wordCount {
		_, err := stmt.Exec(word, count)
		if err != nil {
			http.Error(w, "Failed to insert or update word in database", http.StatusInternalServerError)
			txn.Rollback()
			return
		}
	}

	err = txn.Commit()
	if err != nil {
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Word counts saved successfully"}`))
}
