package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	s := New()

	r := chi.NewRouter()
	r.Route("/book", func(r chi.Router) {
		r.Get("/", GetBookByName(func(w http.ResponseWriter, r *http.Request, name string) {
			book, err := s.Get(name)
			if err != nil {
				http.NotFound(w, r)
				return
			}
			json.NewEncoder(w).Encode(book)
		}))

		r.Post("/", StoreBook(func(w http.ResponseWriter, r *http.Request, body Book) {
			if err := s.Store(body); err != nil {
				http.Error(w, "error processing book", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
		}))
	})

	log.Println("server running on localhost:5555")
	if err := http.ListenAndServe(":5555", r); err != nil {
		log.Fatal(err)
	}
}
