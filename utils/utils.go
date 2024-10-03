package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, statusCode int, data any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}

func WriteERROR(w http.ResponseWriter, err any) error {
	error := map[string]any{"error": err}
	return WriteJSON(w, http.StatusBadRequest, error)
}
