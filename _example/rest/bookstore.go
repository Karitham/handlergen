package main

import (
	"errors"
	"sync"
)

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}

type BookStore interface {
	Store(Book) error
	Get(string) (Book, error)
}

type bstore struct {
	m     *sync.Mutex
	store map[string]Book
}

func New() BookStore {
	return &bstore{
		m:     &sync.Mutex{},
		store: map[string]Book{},
	}
}

func (b *bstore) Store(book Book) error {
	b.m.Lock()
	defer b.m.Unlock()
	b.store[book.Name] = book
	return nil
}

func (b *bstore) Get(name string) (Book, error) {
	b.m.Lock()
	defer b.m.Unlock()
	if book, ok := b.store[name]; ok {
		return book, nil
	}

	return Book{}, errors.New("unknown book")
}
