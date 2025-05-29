package main

import (
	"errors"
	"sync"
)

type Storage interface {
	AddQuote(quote Quote) (Quote, error)
	GetAllQuotes() ([]Quote, error)
	GetQuotesByAuthor(author string) ([]Quote, error)
	DeleteQuote(id int) error
}

type InMemoryStorage struct {
	quotes []Quote
	nextID int
	mu     sync.Mutex
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		quotes: make([]Quote, 0),
		nextID: 1,
	}
}

func (s *InMemoryStorage) AddQuote(quote Quote) (Quote, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	quote.ID = s.nextID
	s.nextID++
	s.quotes = append(s.quotes, quote)

	return quote, nil
}

func (s *InMemoryStorage) GetAllQuotes() ([]Quote, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.quotes, nil
}

func (s *InMemoryStorage) GetQuotesByAuthor(author string) ([]Quote, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var result []Quote
	for _, quote := range s.quotes {
		if quote.Author == author {
			result = append(result, quote)
		}
	}

	return result, nil
}

func (s *InMemoryStorage) DeleteQuote(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, quote := range s.quotes {
		if quote.ID == id {
			s.quotes = append(s.quotes[:i], s.quotes[i+1:]...)
			return nil
		}
	}

	return errors.New("quote not found")
}
