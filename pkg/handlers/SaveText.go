package handlers

import (
	"fmt"
	"net/http"
)

func SaveText(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Post Only", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err := fmt.Fprintf(w, `{[]}`)
	if err != nil {
		return
	}
}
