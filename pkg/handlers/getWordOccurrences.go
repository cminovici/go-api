package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"
)

func (s *Server) GetWordOccurrences(w http.ResponseWriter, r *http.Request) {

	wordsParam := r.URL.Query().Get("words")
	if wordsParam == "" {
		http.Error(w, "Missing words query param", http.StatusBadRequest)
		return
	}

	words := strings.Split(wordsParam, ",")
	for i, word := range words {
		words[i] = strings.TrimSpace(word)
	}

	wordCounts := make(map[string]int)

	for _, word := range words {
		var count int
		query := `SELECT count FROM word_count WHERE word = ?`
		err := s.DB.QueryRow(query, word).Scan(&count)
		if err == sql.ErrNoRows {
			wordCounts[word] = 0
		} else if err != nil {
			http.Error(w, "Error querying the database", http.StatusInternalServerError)
			return
		} else {
			wordCounts[word] = count
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wordCounts)
}
