package main

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

const port = ":8080"

//InitServer starts the http server for serving the query requests
func InitServer(state State) {
	fmt.Println("[SERVER] Starting server")
	fmt.Println("[SERVER] Magic is happening in port", port)
	mux := http.NewServeMux()
	mux.HandleFunc("/", QueryHandler(state))

	//CORS accept *
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(port, handler)
}
