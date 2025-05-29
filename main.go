package main

import (
	"log"
	"net/http"
)

func main() {
	storage := NewInMemoryStorage()
	handler := NewQuoteHandler(storage)

	http.HandleFunc("/quotes", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			author := r.URL.Query().Get("author")
			if author != "" {
				handler.GetQuotesByAuthor(w, r)
			} else {
				handler.GetAllQuotes(w, r)
			}
		case http.MethodPost:
			handler.AddQuote(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/quotes/random", handler.GetRandomQuote)
	http.HandleFunc("/quotes/", handler.DeleteQuote)

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
