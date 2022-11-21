package web

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	gsp "github.com/fluxynet/go-scratch-prod"
)

const (
	// ContentTypeJSON is the content type for JSON
	ContentTypeJSON = "application/json"

	// ContentTypeText is the content type for plain text
	ContentTypeText = "text/plain"
)

// Print sends data to the browser
func Print(w http.ResponseWriter, status int, ctype string, content []byte) {
	if ctype != "" {
		w.Header().Set("Content-Type", ctype)
	}

	w.WriteHeader(status)
	w.Write(content)
}

// Json to the browser
func Json(w http.ResponseWriter, status int, r interface{}) {
	var b, err = json.Marshal(r)
	if err == nil {
		Print(w, status, ContentTypeJSON, b)
	} else {
		JsonError(w, http.StatusInternalServerError, err)
	}
}

// JsonError to the browser in json format
func JsonError(w http.ResponseWriter, status int, err error) {
	var m string

	if err != nil {
		m = strings.ReplaceAll(err.Error(), `"`, `\"`)
	}

	Print(w, status, ContentTypeJSON, []byte(`{"error":"`+m+`"}`))
}

// ReadBody from an http.Request
func ReadBody(r *http.Request) ([]byte, error) {
	if r == nil {
		return nil, gsp.ErrInvalidRequest
	}

	switch r.Method {
	case http.MethodPost, http.MethodPut, http.MethodPatch:
		break
	default:
		return nil, gsp.ErrInvalidRequest
	}

	if r.Body == nil {
		return nil, nil
	}

	defer r.Body.Close()
	var b, err = io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// ReadJsonBodyInto reads json body into a target structure
func ReadJsonBodyInto(w http.ResponseWriter, r *http.Request, target interface{}) error {
	var b, err = ReadBody(r)

	if err != nil {
		return err
	}
	err = json.Unmarshal(b, target)
	if err != nil {
		JsonError(w, http.StatusBadRequest, err)
	}

	return err
}

// ChiIDGetter for chi mux library
func ChiIDGetter(r *http.Request) (string, error) {
	var id = chi.URLParam(r, "id")

	if id == "" {
		return "", gsp.ErrInvalidRequest
	}

	if _, err := uuid.Parse(id); err != nil {
		return "", gsp.ErrInvalidRequest
	}

	return id, nil
}

// TimestampHandler is a simple http handler that returns local time. Useful for status checking
func TimestampHandler(w http.ResponseWriter, r *http.Request) {
	Print(w, http.StatusOK, ContentTypeText, []byte(time.Now().Local().Format(time.RFC3339Nano)))
}
