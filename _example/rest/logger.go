package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

func Logger(w http.ResponseWriter, r *http.Request, err error) {
	type log struct {
		Time   time.Time `json:"time,omitempty"`
		Err    string    `json:"err,omitempty"`
		Method string    `json:"method,omitempty"`
		Path   string    `json:"path,omitempty"`
	}

	l := log{
		Time:   time.Now(),
		Err:    err.Error(),
		Method: r.Method,
		Path:   r.URL.Path,
	}

	json.NewEncoder(os.Stderr).Encode(l)
}
