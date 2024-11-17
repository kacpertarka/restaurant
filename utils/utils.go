package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, statusCode int, data any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}

func WriteERROR(w http.ResponseWriter, err error) error {
	error := map[string]any{"error": err.Error()}
	return WriteJSON(w, http.StatusBadRequest, error)
}

func ParseJSON(r *http.Request, payload any) error {
	// check if r.Body contains any data
	defer r.Body.Close()
	if r.ContentLength == 0 {
		return fmt.Errorf("missing request data")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}
