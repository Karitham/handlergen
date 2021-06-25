// Code generated by github.com/Karitham/handlergen. DO NOT EDIT.
package main

import (
	"encoding/json"
	"net/http"
)

type Main interface {
	GetBookByName(http.ResponseWriter, *http.Request, string)
	StoreBook(http.ResponseWriter, *http.Request, Book)
}

func GetBookByName(h Main) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		Name := query.Get("name")

		h.GetBookByName(
			w,
			r,
			Name,
		)
	}
}

func StoreBook(h Main) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body Book
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "invalid body", 400)
			return
		}

		h.StoreBook(
			w,
			r,
			body,
		)
	}
}
