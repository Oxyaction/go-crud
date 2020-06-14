package jsonutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Decode(w http.ResponseWriter, r *http.Request, entity interface{}) (int, error) {
	contentType := r.Header.Get("Content-Type")

	if contentType == "" || contentType != "application/json" {
		return http.StatusBadRequest, errors.New("Incorrect content-type")
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(entity)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			return http.StatusBadRequest, fmt.Errorf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
		case errors.As(err, &unmarshalTypeError):
			return http.StatusBadRequest, fmt.Errorf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return http.StatusBadRequest, errors.New("Request body contains badly-formed JSON")
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return http.StatusBadRequest, fmt.Errorf("Request body contains unknown field %s", fieldName)
		case errors.Is(err, io.EOF):
			return http.StatusBadRequest, errors.New("Request body must not be empty")
		case err.Error() == "http: request body too large":
			return http.StatusRequestEntityTooLarge, errors.New("Request body must not be larger than 1MB")
		default:
			return http.StatusInternalServerError, fmt.Errorf("failed to decode json %v", err)
		}
	}

	if dec.More() {
		return http.StatusBadRequest, errors.New("body must contain only one JSON object")
	}

	return http.StatusOK, nil
}
