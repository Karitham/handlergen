package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	store := New()
	s := &server{bookStore: store}

	r := chi.NewRouter()
	r.Route("/book", func(r chi.Router) {
		r.Get("/", GetBookByName(s, Logger))
		r.Post("/", StoreBook(s, Logger))
	})

	log.Println("server running on localhost:5555")
	if err := http.ListenAndServe(":5555", r); err != nil {
		log.Fatal(err)
	}
}

type server struct {
	bookStore interface {
		Store(Book) error
		Get(string) (Book, error)
	}
}

func (s *server) GetBookByName(w http.ResponseWriter, r *http.Request, name string) {
	book, err := s.bookStore.Get(name)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func (s *server) StoreBook(w http.ResponseWriter, r *http.Request, body Book) {
	if err := s.bookStore.Store(body); err != nil {
		http.Error(w, "error processing book", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
