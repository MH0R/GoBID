package jasonutils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MH0R/GoBID/internal/validator"
)

func EncodeJson[T any](w http.ResponseWriter, r *http.Request, statusCode int, data T) error {
	w.Header().Set("Content-type", "Application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		return fmt.Errorf("failed to encode json %w", err)
	}

	return nil
}

func decodeValidJson[T validator.Validator](r *http.Request) (T, map[string]string, error) {
	var data T

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return data, nil, fmt.Errorf("error to decode Json: %w", err)
	}

	if problems := data.Valid(r.Context()); len(problems) > 0 {
		return data, problems, fmt.Errorf("invalid %T: %d problems", data, len(problems))
	}
	return data, nil, nil
}

func DecodeJson[T any](r *http.Request) (T, error) {
	var data T

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return data, fmt.Errorf("error to decode Json: %w", err)
	}

	return data, nil
}
