package router

import (
	"net/http"

	"github.com/yaedinorog/quote_api/app/internal/handlers"
)

func HTTPSetup() {
	http.HandleFunc("/quotes", handlers.GetQuotes)
	http.HandleFunc("/quotes/", handlers.GetQuotes)
	http.HandleFunc("/quotes/random", handlers.GetRandomQuote)
}
