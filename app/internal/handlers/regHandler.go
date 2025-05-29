package handlers

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"encoding/json"
)

type Quote struct {
	Author string `json:"author"`
	Text   string `json:"quote"`
}

type QuoteList struct {
	List []Quote
}

var quotes QuoteList = QuoteList{}

func GetQuotes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if author := r.URL.Query().Get("author"); author == "" {

			res_json, err := json.Marshal(quotes.List)

			if err != nil {
				log.Panic("Error while marshaling structure")
			}

			w.Write(res_json)
			return
		} else {
			var filtered []Quote
			for _, quote := range quotes.List {
				if quote.Author == author {
					filtered = append(filtered, quote)
				}
			}

			res_json, err := json.Marshal(filtered)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}

			w.Write(res_json)
		}
	case "POST":
		var quote Quote

		if err := json.NewDecoder(r.Body).Decode(&quote); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		quotes.List = append(quotes.List, quote)

		res, err := json.Marshal(quote)

		if err != nil {
			log.Printf("Error while marshalling Quote from POST request %v", err)
		}

		w.Write(res)
	case "DELETE":
		path := r.URL.Path
		id_for_deletion, err := strconv.Atoi(strings.Split(path, "/")[2])

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		quotes.List = append(quotes.List[:id_for_deletion-1], quotes.List[id_for_deletion:]...)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	res_json, err := json.Marshal(quotes.List[(rand.Int() % len(quotes.List))])

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(res_json)
}
