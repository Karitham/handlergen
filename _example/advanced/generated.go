// Code generated by github.com/Karitham/handlergen. DO NOT EDIT.
package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Karitham/handlergen/gen"
)

func example1(handler func(http.ResponseWriter, *http.Request, uint, gen.Template)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		query_page := query.Get("page")
		page64, err := strconv.ParseUint(query_page, 10, 64)
		if err != nil {
			http.Error(w, "invalid query", 400)
			return
		}
		page := uint(page64)

		var body gen.Template
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "invalid body", 400)
			return
		}

		handler(
			w,
			r,
			page,
			body,
		)
	}
}

func example2(handler func(http.ResponseWriter, *http.Request, string, uint, int, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		user := query.Get("user")

		query_user_id := query.Get("user_id")
		user_id64, err := strconv.ParseUint(query_user_id, 10, 64)
		if err != nil {
			http.Error(w, "invalid query", 400)
			return
		}
		user_id := uint(user_id64)

		query_page := query.Get("page")
		page, err := strconv.Atoi(query_page)
		if err != nil {
			http.Error(w, "invalid query", 400)
			return
		}

		query_per_page := query.Get("per_page")
		per_page, err := strconv.Atoi(query_per_page)
		if err != nil {
			http.Error(w, "invalid query", 400)
			return
		}

		handler(
			w,
			r,
			user,
			user_id,
			page,
			per_page,
		)
	}
}
