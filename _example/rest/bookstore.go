package main

import (
	"errors"
	"sync"
)

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}

type bstore struct {
	mu    sync.Mutex
	store map[string]Book
}

func New() *bstore {
	return &bstore{
		mu:    sync.Mutex{},
		store: map[string]Book{},
	}
}

func (b *bstore) Store(book Book) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.store[book.Name] = book
	return nil
}

func (b *bstore) Get(name string) (Book, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if book, ok := b.store[name]; ok {
		return book, nil
	}

	return Book{}, errors.New("unknown book")
}
