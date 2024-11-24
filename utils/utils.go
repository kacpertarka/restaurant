package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
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

func GetEnvVariable(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	return fallback
}

func GetEnvVariableAsInt(key string, fallback int) int {
	val := GetEnvVariable(key, strconv.Itoa(fallback))
	valInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	return valInt
}
