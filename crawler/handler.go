package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const resultsLimit = 10

//QueryHandler handler the query get request
func QueryHandler(state State) func(w http.ResponseWriter, r *http.Request) {
	fmt.Println(len(state.DF))

	return func(w http.ResponseWriter, r *http.Request) {
		query := r.FormValue("query")
		if len(query) == 0 {
			http.Error(w, "Query param not provided", 400)
			return
		}

		//TODO: Sanitize data
		results := makeQuery(query, state)

		json, err := json.Marshal(results[0:resultsLimit])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
}
